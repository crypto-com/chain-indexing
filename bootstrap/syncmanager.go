package bootstrap

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	eventhandler_interface "github.com/crypto-com/chain-indexing/appinterface/eventhandler"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	"github.com/crypto-com/chain-indexing/usecase/model"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	command_entity "github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	chainfeed "github.com/crypto-com/chain-indexing/infrastructure/feed/chain"
	"github.com/crypto-com/chain-indexing/infrastructure/metric/prometheus"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"github.com/crypto-com/chain-indexing/usecase/syncstrategy"
)

const MAX_RETRY_TIME_ALWAYS_RETRY = 0
const DEFAULT_POLLING_INTERVAL = 5 * time.Second
const DEFAULT_MAX_RETRY_INTERVAL = 15 * time.Minute
const DEFAULT_MAX_RETRY_TIME = MAX_RETRY_TIME_ALWAYS_RETRY

type SyncManager struct {
	rdbConn              rdb.Conn
	tendermintClient     *tendermint.HTTPClient
	cosmosClient         cosmosapp_interface.Client
	logger               applogger.Logger
	pollingInterval      time.Duration
	maxRetryInterval     time.Duration
	maxRetryTime         time.Duration
	strictGenesisParsing bool

	accountAddressPrefix string
	stakingDenom         string

	windowSyncStrategy *syncstrategy.Window

	eventHandler eventhandler_interface.Handler

	// SyncManager state
	latestBlockHeight *int64
	shouldSyncCh      chan bool

	parserManager *utils.CosmosParserManager

	startingBlockHeight int64
}

type SyncManagerParams struct {
	Logger    applogger.Logger
	RDbConn   rdb.Conn
	TxDecoder *utils.TxDecoder

	Config SyncManagerConfig
}

type SyncManagerConfig struct {
	WindowSize               int
	TendermintRPCUrl         string
	CosmosAppHTTPRPCURL      string
	InsecureTendermintClient bool
	InsecureCosmosAppClient  bool
	StrictGenesisParsing     bool
	AccountAddressPrefix     string
	StakingDenom             string
	StartingBlockHeight      int64
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	params SyncManagerParams,
	pm *utils.CosmosParserManager,
	eventHandler eventhandler_interface.Handler,
) *SyncManager {
	var tendermintClient *tendermint.HTTPClient
	if params.Config.InsecureTendermintClient {
		tendermintClient = tendermint.NewInsecureHTTPClient(
			params.Config.TendermintRPCUrl,
			params.Config.StrictGenesisParsing,
		)
	} else {
		tendermintClient = tendermint.NewHTTPClient(
			params.Config.TendermintRPCUrl,
			params.Config.StrictGenesisParsing,
		)
	}

	var cosmosClient cosmosapp_interface.Client
	if params.Config.InsecureCosmosAppClient {
		cosmosClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
			params.Config.CosmosAppHTTPRPCURL, params.Config.StakingDenom,
		)
	} else {
		cosmosClient = cosmosapp_infrastructure.NewHTTPClient(
			params.Config.CosmosAppHTTPRPCURL, params.Config.StakingDenom,
		)
	}

	return &SyncManager{
		rdbConn:          params.RDbConn,
		tendermintClient: tendermintClient,
		cosmosClient:     cosmosClient,
		logger: params.Logger.WithFields(applogger.LogFields{
			"module": "SyncManager",
		}),
		pollingInterval:      DEFAULT_POLLING_INTERVAL,
		maxRetryInterval:     DEFAULT_MAX_RETRY_INTERVAL,
		maxRetryTime:         DEFAULT_MAX_RETRY_TIME,
		strictGenesisParsing: params.Config.StrictGenesisParsing,

		accountAddressPrefix: params.Config.AccountAddressPrefix,
		stakingDenom:         params.Config.StakingDenom,

		shouldSyncCh: make(chan bool, 1),

		windowSyncStrategy: syncstrategy.NewWindow(params.Logger, params.Config.WindowSize),

		eventHandler: eventHandler,

		parserManager: pm,

		startingBlockHeight: params.Config.StartingBlockHeight,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64, isRetry bool) error {
	maybeLastIndexedHeight, err := manager.eventHandler.GetLastHandledEventHeight()
	if err != nil {
		return fmt.Errorf("error running GetLastIndexedBlockHeight %v", err)
	}

	// if none of the block has been indexed before, start with 0
	currentIndexingHeight := int64(0)
	if maybeLastIndexedHeight != nil {
		currentIndexingHeight = *maybeLastIndexedHeight + 1
	}

	targetHeight := latestHeight
	if isRetry {
		// Reduce the block size to be synced when retrying to avoid spamming and wasting resource
		if latestHeight > currentIndexingHeight {
			targetHeight = currentIndexingHeight
			if targetHeight < manager.startingBlockHeight {
				targetHeight = manager.startingBlockHeight
			}
		}
	}

	// Skip for latest height < statring block height
	if currentIndexingHeight != 0 && latestHeight < manager.startingBlockHeight {
		return nil
	}
	manager.logger.Infof("going to synchronized blocks from %d to %d", currentIndexingHeight, targetHeight)
	for currentIndexingHeight <= targetHeight {
		startTime := time.Now()
		endHeight := targetHeight
		if currentIndexingHeight == 0 {
			// Genesis Block as an individual window, size = 1
			endHeight = 0
		} else if currentIndexingHeight < manager.startingBlockHeight {
			currentIndexingHeight = manager.startingBlockHeight
		}

		blocksCommands, syncedHeight, err := manager.windowSyncStrategy.Sync(
			currentIndexingHeight, endHeight, manager.syncBlockWorker,
		)
		if err != nil {
			return fmt.Errorf("error when synchronizing block with window strategy: %v", err)
		}

		if err != nil {
			return fmt.Errorf("error beginning transaction: %v", err)
		}
		for i, commands := range blocksCommands {
			blockHeight := currentIndexingHeight + int64(i)

			events := make([]event.Event, 0, len(commands))
			for _, command := range commands {
				event, err := command.Exec()
				if err != nil {

					return fmt.Errorf(
						"error executing command %sV%d to produce events: %v",
						command.Name(), command.Version(), err,
					)
				}
				events = append(events, event)
			}

			err := manager.eventHandler.HandleEvents(blockHeight, events)
			if err != nil {
				return fmt.Errorf("error handling events: %v", err)
			}
			prometheus.RecordProjectionExecTime(manager.eventHandler.Id(), time.Since(startTime).Milliseconds())
		}

		// If there is any error before, short-circuit return in the error handling
		// while the local currentIndexingHeight won't be incremented and will be retried later
		manager.logger.Infof("successfully synced to block height %d", syncedHeight)
		prometheus.RecordProjectionLatestHeight(manager.eventHandler.Id(), syncedHeight)

		currentIndexingHeight = syncedHeight + 1
	}
	return nil
}

func (manager *SyncManager) syncBlockWorker(blockHeight int64) ([]command_entity.Command, error) {
	logger := manager.logger.WithFields(applogger.LogFields{
		"submodule":   "SyncBlockWorker",
		"blockHeight": blockHeight,
	})

	logger.Info("synchronizing block")

	if blockHeight == int64(0) {
		genesis, err := manager.tendermintClient.Genesis()
		if err != nil {
			return nil, fmt.Errorf("error requesting chain genesis: %v", err)
		}

		return parser.ParseGenesisCommands(genesis, manager.accountAddressPrefix)
	}

	// Request tendermint RPC
	block, rawBlock, err := manager.tendermintClient.Block(blockHeight)
	if err != nil {
		return nil, fmt.Errorf("error requesting chain block at height %d: %v", blockHeight, err)
	}

	blockResults, err := manager.tendermintClient.BlockResults(blockHeight)
	if err != nil {
		return nil, fmt.Errorf("error requesting chain block_results at height %d: %v", blockHeight, err)
	}

	txs := make([]model.Tx, 0)
	for _, txHex := range block.Txs {
		var tx *model.Tx
		tx, err = manager.cosmosClient.Tx(parser.TxHash(txHex))
		if err != nil {
			return nil, fmt.Errorf("error requesting chain txs (%s) at height %d: %v", txHex, blockHeight, err)
		}
		txs = append(txs, *tx)
	}

	parseBlockToCommandsLogger := manager.logger.WithFields(applogger.LogFields{
		"submodule":   "ParseBlockToCommands",
		"blockHeight": blockHeight,
	})

	commands, err := parser.ParseBlockToCommands(
		parseBlockToCommandsLogger,
		manager.parserManager,
		manager.cosmosClient,
		block,
		rawBlock,
		blockResults,
		txs,
		manager.accountAddressPrefix,
		manager.stakingDenom,
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing block data to commands %v", err)
	}

	return commands, nil
}

// Run starts the polling service for blocks
func (manager *SyncManager) Run() error {
	tracker := chainfeed.NewBlockHeightTracker(manager.logger, manager.tendermintClient)
	manager.latestBlockHeight = tracker.GetLatestBlockHeight()
	blockHeightCh := make(chan int64, 1)
	go func() {
		for {
			latestBlockHeight := <-blockHeightCh
			manager.latestBlockHeight = &latestBlockHeight
			manager.drainShouldSyncCh()
			manager.shouldSyncCh <- true
		}
	}()
	tracker.Subscribe(blockHeightCh)

	parser.InitParsers(manager.parserManager)
	parser.RegisterBreakingVersionParsers(manager.parserManager)

	for {
		isRetry := false
		operation := func() error {
			if manager.latestBlockHeight == nil {
				manager.logger.Info("the chain has no block yet")
			} else {
				if syncErr := manager.SyncBlocks(*manager.latestBlockHeight, isRetry); syncErr != nil {
					return fmt.Errorf(
						"error synchronizing blocks to latest height %d: %v", *manager.latestBlockHeight, syncErr,
					)
				}
			}

			select {
			case <-manager.shouldSyncCh:
			case <-time.After(manager.pollingInterval):
			}
			return nil
		}
		notifyFn := func(opErr error, backoffDuration time.Duration) {
			isRetry = true
			manager.logger.Errorf(
				"retrying in %s: %v", backoffDuration.String(), opErr,
			)
		}

		neverStopExponentialBackoff := backoff.NewExponentialBackOff()
		neverStopExponentialBackoff.MaxElapsedTime = manager.maxRetryTime
		neverStopExponentialBackoff.MaxInterval = manager.maxRetryInterval
		if err := backoff.RetryNotify(
			operation,
			backoff.NewExponentialBackOff(),
			notifyFn,
		); err != nil {
			manager.logger.Errorf("stopping retry after too many errors: %s: %v", err)
		}
	}
}

func (manager *SyncManager) drainShouldSyncCh() {
	select {
	case <-manager.shouldSyncCh:
	default:
	}
}
