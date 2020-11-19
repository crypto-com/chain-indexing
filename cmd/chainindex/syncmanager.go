package main

import (
	"errors"
	"fmt"
	"time"

	event_interface "github.com/crypto-com/chainindex/appinterface/event"
	projection_interface "github.com/crypto-com/chainindex/appinterface/projection"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/appinterface/rdbstatusstore"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/projection"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
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

	txDecoder         *parser.TxDecoder
	statusStore       *rdbstatusstore.RDbStatusStoreImpl
	eventStore        *event_interface.RDbStore
	projectionManager *projection.Manager
	subject           *chainfeed.BlockSubject
	eventRegistry     *event.Registry
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	logger applogger.Logger,
	rdbConn rdb.Conn,

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

		txDecoder:       txDecoder,
		pollingInterval: DEFAULT_POLLING_INTERVAL,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64) error {
	maybeLastIndexedHeight, err := manager.statusStore.GetLastIndexedBlockHeight()
	if err != nil {
		return fmt.Errorf("error running GetLastIndexedBlockHeight %v", err)
	}

	// if none of the block has been indexed before, start with 0
	lastIndexedHeight := int64(0)
	if maybeLastIndexedHeight != nil {
		lastIndexedHeight = *maybeLastIndexedHeight
	}

	// Sync next height to avoid duplication
	currentIndexingHeight := lastIndexedHeight + 1

	manager.logger.Infof("going to synchronized blocks from %d to %d", currentIndexingHeight, latestHeight)
	for currentIndexingHeight < latestHeight {
		// Request tendermint RPC
		block, rawBlock, err := manager.client.Block(currentIndexingHeight)
		if err != nil {
			return fmt.Errorf("error requesting chain block at height %d: %v", currentIndexingHeight, err)
		}

		blockResults, err := manager.client.BlockResults(currentIndexingHeight)
		if err != nil {
			return fmt.Errorf("error requesting chain block_results at height %d: %v", currentIndexingHeight, err)
		}

		// Create new block notification and notify subscribers
		tx, err := manager.rdbConn.Begin()
		txHandle := tx.ToHandle()
		if err != nil {
			return fmt.Errorf("error beginning transaction: %v", err)
		}
		eventStore := event_interface.NewRDbStore(txHandle, manager.eventRegistry)
		notif := notification.NewBlockNotification(
			currentIndexingHeight, block, rawBlock, blockResults,
		)
		manager.subject.Notify(notif, eventStore)

		// Current block indexing done, update db and sync next height
		manager.logger.WithFields(applogger.LogFields{
			"blockHeight": block.Height,
		}).Info("block synced and events produced")

		if err := manager.statusStore.UpdateLastIndexedBlockHeight(txHandle, currentIndexingHeight); err != nil {
			return fmt.Errorf("error updating last indexed block height to %d: %v", currentIndexingHeight, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("error committing block synchronization outcomes: %v", err)
		}

		// If there is any error before, short-circuit return in the error handling
		// while the local currentIndexingHeight won't be incremented and will be retried later
		currentIndexingHeight += 1
	}
	return nil
}

func (manager *SyncManager) InitSubject() *chainfeed.BlockSubject {
	// Currently only the chain processor subscriber
	// add more subscriber base on the need
	chainProcessor := chainfeed.NewBlockSubscriber(manager.txDecoder)

	blockSubject := chainfeed.NewBlockSubject(manager.logger)

	// add more subscribers here
	blockSubject.Attach(chainProcessor)

	return blockSubject
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
		return nil, fmt.Errorf("error registering projection to manager %v", err)
	}

	return projectionManager, nil
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (manager *SyncManager) Run() error {
	var err error
	manager.subject = manager.InitSubject()
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
	for {
		latestHeight := tracker.GetLatestBlockHeight()
		if latestHeight == nil {
			<-time.After(manager.pollingInterval)
			continue
		}
		if err := manager.SyncBlocks(*latestHeight); err != nil {
			manager.logger.Errorf("error synchronizing blocks to latest height %d: %v", *latestHeight, err)
		}

		<-time.After(manager.pollingInterval)
	}
}
