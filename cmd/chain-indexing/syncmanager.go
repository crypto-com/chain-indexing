package main

import (
	"fmt"
	"time"

	eventhandler_interface "github.com/crypto-com/chain-indexing/appinterface/eventhandler"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	command_entity "github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/entity/event"
	chainfeed "github.com/crypto-com/chain-indexing/infrastructure/feed/chain"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	"github.com/crypto-com/chain-indexing/usecase/syncstrategy"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type SyncManager struct {
	rdbConn         rdb.Conn
	client          *tendermint.HTTPClient
	logger          applogger.Logger
	pollingInterval time.Duration

	txDecoder          *parser.TxDecoder
	windowSyncStrategy *syncstrategy.Window

	eventHandler eventhandler_interface.Handler

	// SyncManager state
	latestBlockHeight *int64
	shouldSyncCh      chan bool
}

type SyncManagerParams struct {
	Logger    applogger.Logger
	RDbConn   rdb.Conn
	TxDecoder *parser.TxDecoder

	Config SyncManagerConfig
}

type SyncManagerConfig struct {
	WindowSize       int
	TendermintRPCUrl string
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	params SyncManagerParams,
	eventHandler eventhandler_interface.Handler,
) *SyncManager {
	tendermintClient := tendermint.NewHTTPClient(params.Config.TendermintRPCUrl)

	return &SyncManager{
		rdbConn: params.RDbConn,
		client:  tendermintClient,
		logger: params.Logger.WithFields(applogger.LogFields{
			"module": "SyncManager",
		}),
		pollingInterval: DEFAULT_POLLING_INTERVAL,

		shouldSyncCh: make(chan bool, 1),

		txDecoder:          params.TxDecoder,
		windowSyncStrategy: syncstrategy.NewWindow(params.Logger, params.Config.WindowSize),

		eventHandler: eventHandler,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64) error {
	maybeLastIndexedHeight, err := manager.eventHandler.GetLastHandledEventHeight()
	if err != nil {
		return fmt.Errorf("error running GetLastIndexedBlockHeight %v", err)
	}

	// if none of the block has been indexed before, start with 0
	currentIndexingHeight := int64(0)
	if maybeLastIndexedHeight != nil {
		currentIndexingHeight = *maybeLastIndexedHeight + 1
	}

	manager.logger.Infof("going to synchronized blocks from %d to %d", currentIndexingHeight, latestHeight)
	for currentIndexingHeight <= latestHeight {
		blocksCommands, syncedHeight, err := manager.windowSyncStrategy.Sync(
			currentIndexingHeight, latestHeight, manager.syncBlockWorker,
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
					return fmt.Errorf("error generating event: %v", err)
				}
				events = append(events, event)
			}

			err := manager.eventHandler.HandleEvents(blockHeight, events)
			if err != nil {
				return fmt.Errorf("error handling events: %v", err)
			}
		}

		// If there is any error before, short-circuit return in the error handling
		// while the local currentIndexingHeight won't be incremented and will be retried later
		manager.logger.Infof("successfully synced to block height %d", syncedHeight)
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
		genesis, err := manager.client.Genesis()
		if err != nil {
			return nil, fmt.Errorf("error requesting chain genesis: %v", err)
		}

		return parser.ParseGenesisCommands(genesis)
	}

	// Request tendermint RPC
	block, rawBlock, err := manager.client.Block(blockHeight)
	if err != nil {
		return nil, fmt.Errorf("error requesting chain block at height %d: %v", blockHeight, err)
	}

	blockResults, err := manager.client.BlockResults(blockHeight)
	if err != nil {
		return nil, fmt.Errorf("error requesting chain block_results at height %d: %v", blockHeight, err)
	}

	commands, err := parser.ParseBlockToCommands(
		manager.txDecoder,
		block,
		rawBlock,
		blockResults,
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing block data to commands %v", err)
	}

	return commands, nil
}

// Run starts the polling service for blocks
func (manager *SyncManager) Run() error {
	tracker := chainfeed.NewBlockHeightTracker(manager.logger, manager.client)
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

	for {
		if manager.latestBlockHeight == nil {
			manager.logger.Info("the chain has no block yet")
		} else {
			if err := manager.SyncBlocks(*manager.latestBlockHeight); err != nil {
				manager.logger.Errorf("error synchronizing blocks to latest height %d: %v", *manager.latestBlockHeight, err)
				<-time.After(5 * time.Second)
			}
		}

		select {
		case <-manager.shouldSyncCh:
		case <-time.After(manager.pollingInterval):
		}
	}
}

func (manager *SyncManager) drainShouldSyncCh() {
	select {
	case <-manager.shouldSyncCh:
	default:
	}
}
