package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/crypto-com/chainindex/usecase/executor"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/syncstrategy"

	event_interface "github.com/crypto-com/chainindex/appinterface/event"
	projection_interface "github.com/crypto-com/chainindex/appinterface/projection"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/appinterface/rdbstatusstore"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/projection"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/parser"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type SyncManager struct {
	rdbConn         rdb.Conn
	client          *tendermint.HTTPClient
	logger          applogger.Logger
	pollingInterval time.Duration
	windowSize      int

	latestBlockHeight *int64
	shouldSyncCh      chan bool

	txDecoder          *parser.TxDecoder
	windowSyncStrategy *syncstrategy.Window
	statusStore        *rdbstatusstore.RDbStatusStoreImpl
	eventStore         *event_interface.RDbStore
	projectionManager  *projection.Manager
	eventRegistry      *event.Registry
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	windowSize int,
	tendermintRPCUrl string,
	txDecoder *parser.TxDecoder,
) *SyncManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &SyncManager{
		rdbConn: rdbConn,
		client:  tendermintClient,
		logger: logger.WithFields(applogger.LogFields{
			"module": "SyncManager",
		}),
		pollingInterval: DEFAULT_POLLING_INTERVAL,
		windowSize:      windowSize,

		shouldSyncCh: make(chan bool),

		txDecoder:          txDecoder,
		windowSyncStrategy: syncstrategy.NewWindow(logger, windowSize),
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64) error {
	maybeLastIndexedHeight, err := manager.statusStore.GetLastIndexedBlockHeight()
	if err != nil {
		return fmt.Errorf("error running GetLastIndexedBlockHeight %v", err)
	}

	// if none of the block has been indexed before, start with 0
	currentIndexingHeight := int64(0)
	if maybeLastIndexedHeight != nil {
		currentIndexingHeight = *maybeLastIndexedHeight + 1
	}

	manager.logger.Infof("going to synchronized blocks from %d to %d", currentIndexingHeight, latestHeight)
	for currentIndexingHeight < latestHeight {
		blocksCommands, syncedHeight, err := manager.windowSyncStrategy.Sync(
			currentIndexingHeight, latestHeight, manager.syncBlockWorker,
		)
		if err != nil {
			manager.logger.Errorf("error when synchronizing block with window strategy: %v", err)
			<-time.After(5 * time.Second)
		}

		manager.logger.Debug("start persisting blocks commands")
		tx, err := manager.rdbConn.Begin()
		txHandle := tx.ToHandle()
		if err != nil {
			return fmt.Errorf("error beginning transaction: %v", err)
		}
		eventStore := event_interface.NewRDbStore(txHandle, manager.eventRegistry)
		for i, commands := range blocksCommands {
			blockHeight := currentIndexingHeight + int64(i)
			executor := executor.NewBlockExecutor(manager.logger, blockHeight)
			executor.AddAllCommands(commands)

			// generate all events, make them persistent
			if err := executor.ExecAllCommands(); err != nil {
				return fmt.Errorf("error generating all events%v", err)
			}
			if err := executor.StoreAllEvents(eventStore); err != nil {
				return fmt.Errorf("error storing all events%v", err)
			}
		}
		if err := manager.statusStore.UpdateLastIndexedBlockHeight(txHandle, syncedHeight); err != nil {
			return fmt.Errorf("error updating last indexed block height to %d: %v", currentIndexingHeight, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("error committing block synchronization outcomes: %v", err)
		}

		// If there is any error before, short-circuit return in the error handling
		// while the local currentIndexingHeight won't be incremented and will be retried later
		manager.logger.Infof("successfully synced to block height %d", syncedHeight)
		currentIndexingHeight = syncedHeight + 1
	}
	return nil
}

func (manager *SyncManager) syncBlockWorker(blockHeight int64) ([]command.Command, error) {
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

func (manager *SyncManager) InitRegistry() *event.Registry {
	registry := event.NewRegistry()
	event_usecase.RegisterEvents(registry)
	return registry
}

func (manager *SyncManager) InitStatusStore() *rdbstatusstore.RDbStatusStoreImpl {
	StatusStore := rdbstatusstore.NewRDbStatusStoreImpl(
		manager.logger,
		manager.rdbConn.ToHandle(),
	)
	return StatusStore
}

func (manager *SyncManager) InitEventStore() (*event_interface.RDbStore, error) {
	if manager.eventRegistry == nil {
		return nil, errors.New("cannot init event store without an Register instance")
	}
	EventStore := event_interface.NewRDbStore(manager.rdbConn.ToHandle(), manager.eventRegistry)
	return EventStore, nil
}

func (manager *SyncManager) InitProjectionManager() (*projection.Manager, error) {
	if manager.eventStore == nil {
		return nil, errors.New("cannot init projection manager without an eventStore instance")
	}
	projectionManager := projection.NewManager(manager.logger, manager.eventStore)

	// register more projections here
	blockProjection := projection_interface.NewBlock(manager.logger, manager.rdbConn)
	if err := projectionManager.RegisterProjection(blockProjection); err != nil {
		return nil, fmt.Errorf("error registering block projection to manager %v", err)
	}

	transactionProjection := projection_interface.NewTransaction(manager.logger, manager.rdbConn)
	if err := projectionManager.RegisterProjection(transactionProjection); err != nil {
		return nil, fmt.Errorf("error registering transaction projection to manager %v", err)
	}

	blockEventProjection := projection_interface.NewBlockEvent(manager.logger, manager.rdbConn)
	if err := projectionManager.RegisterProjection(blockEventProjection); err != nil {
		return nil, fmt.Errorf("error registering block event projection to manager %v", err)
	}

	return projectionManager, nil
}

// Run starts the polling service for blocks
func (manager *SyncManager) Run() error {
	var err error
	manager.eventRegistry = manager.InitRegistry()
	manager.statusStore = manager.InitStatusStore()

	if manager.eventStore, err = manager.InitEventStore(); err != nil {
		return fmt.Errorf("error init event store %v", err)
	}

	if manager.projectionManager, err = manager.InitProjectionManager(); err != nil {
		return fmt.Errorf("error init projection manager %v", err)
	}

	manager.projectionManager.Run()

	tracker := chainfeed.NewBlockHeightTracker(manager.logger, manager.client)
	manager.latestBlockHeight = tracker.GetLatestBlockHeight()
	blockHeightCh := make(chan int64)
	go func() {
		latestBlockHeight := <-blockHeightCh
		manager.latestBlockHeight = &latestBlockHeight
		manager.logger.Infof("updated latest block height to %d", *manager.latestBlockHeight)
		manager.drainShouldSyncCh()
		manager.shouldSyncCh <- true
	}()
	tracker.Subscribe(blockHeightCh)

	for {
		if manager.latestBlockHeight == nil {
			manager.logger.Info("the chain has no block yet")
		} else {
			if err := manager.SyncBlocks(*manager.latestBlockHeight); err != nil {
				manager.logger.Errorf("error synchronizing blocks to latest height %d: %v", *manager.latestBlockHeight, err)
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
