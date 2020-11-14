package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/crypto-com/chainindex/usecase/domain"

	app_event "github.com/crypto-com/chainindex/appinterface/event"
	app_projection "github.com/crypto-com/chainindex/appinterface/projection"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/appinterface/rdbstatusstore"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/projection"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"github.com/crypto-com/chainindex/usecase/parser"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type SyncManager struct {
	rdbConn         rdb.Conn
	client          *tendermint.HTTPClient
	logger          applogger.Logger
	pollingInterval time.Duration

	moduleAccounts *parser.ModuleAccounts

	StatusStore       *rdbstatusstore.RDbStatusStoreImpl
	EventStore        *app_event.RDbStore
	ProjectionManager *projection.Manager
	Subject           *chainfeed.BlockSubject
	Registry          *event.Registry
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	logger applogger.Logger,
	tendermintRPCUrl string,
	rdbConn rdb.Conn,
	moduleAccounts *parser.ModuleAccounts,
) *SyncManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &SyncManager{
		rdbConn: rdbConn,
		client:  tendermintClient,
		logger: logger.WithFields(applogger.LogFields{
			"module": "SyncManager",
		}),

		moduleAccounts:  moduleAccounts,
		pollingInterval: DEFAULT_POLLING_INTERVAL,
	}
}

func (manager *SyncManager) UpdateIndexedHeight(nextHeight int64, handle *rdb.Handle) error {
	err := manager.StatusStore.UpdateLastIndexedBlockHeight(nextHeight)
	if err != nil {
		return fmt.Errorf("error updating last indexed block height %v", err)
	}
	return nil
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64) error {
	maybeLastIndexedHeight, err := manager.StatusStore.GetLastIndexedBlockHeight()
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

	for currentIndexingHeight < latestHeight {
		// Request tendermint RPC
		block, rawBlock, err := manager.client.Block(currentIndexingHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", currentIndexingHeight, err)
		}

		blockResults, err := manager.client.BlockResults(currentIndexingHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block_results at %d: %v", currentIndexingHeight, err)
		}

		// Create new block notification and notify subscribers
		notif := notification.NewBlockNotification(
			currentIndexingHeight, block, rawBlock, blockResults,
		)
		manager.Subject.Notify(notif, manager.EventStore)

		// Current block indexing done, update db and sync next height
		manager.logger.WithFields(applogger.LogFields{
			"blockHeight": block.Height,
		}).Info("block synced and events produced")
		err = manager.UpdateIndexedHeight(currentIndexingHeight, manager.rdbConn.ToHandle())
		if err != nil {
			return fmt.Errorf("error updating last indexed height for height %d: %v", currentIndexingHeight, err)
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
	chainProcessor := chainfeed.NewBlockSubscriber(manager.moduleAccounts)

	blockSubject := chainfeed.NewBlockSubject()

	// add more subscribers here
	blockSubject.Attach(chainProcessor)

	return blockSubject
}

func (manager *SyncManager) InitRegistry() *event.Registry {
	registry := event.NewRegistry()
	domain.RegisterEvents(registry)
	return registry
}

func (manager *SyncManager) InitStatusStore() *rdbstatusstore.RDbStatusStoreImpl {
	StatusStore := rdbstatusstore.NewRDbStatusStoreImpl(
		manager.logger,
		manager.rdbConn.ToHandle(),
	)
	return StatusStore
}

func (manager *SyncManager) InitEventStore() (*app_event.RDbStore, error) {
	if manager.Registry == nil {
		return nil, errors.New("cannot init event store without an Register instance")
	}
	EventStore := app_event.NewRDbStore(manager.rdbConn.ToHandle(), manager.Registry)
	return EventStore, nil
}

func (manager *SyncManager) InitProjectionManager() (*projection.Manager, error) {
	if manager.EventStore == nil {
		return nil, errors.New("cannot init projection manager without an EventStore instance")
	}
	projectionManager := projection.NewManager(manager.logger, manager.EventStore)

	// register more projections here
	blockProjection := app_projection.NewBlock(manager.logger, manager.rdbConn)
	if err := projectionManager.RegisterProjection(blockProjection); err != nil {
		return nil, fmt.Errorf("error registering projection to manager %v", err)
	}

	return projectionManager, nil
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (manager *SyncManager) Run() error {
	var err error
	manager.Subject = manager.InitSubject()
	manager.Registry = manager.InitRegistry()
	manager.StatusStore = manager.InitStatusStore()

	if manager.EventStore, err = manager.InitEventStore(); err != nil {
		return fmt.Errorf("error init event store %v", err)
	}

	if manager.ProjectionManager, err = manager.InitProjectionManager(); err != nil {
		return fmt.Errorf("error init projection manager %v", err)
	}

	manager.ProjectionManager.Run()

	tracker := chainfeed.NewBlockHeightTracker(manager.logger, manager.client)
	for {
		latestHeight := tracker.GetLatestBlockHeight()
		if latestHeight == nil {
			<-time.After(manager.pollingInterval)
			continue
		}
		if err := manager.SyncBlocks(*latestHeight); err != nil {
			manager.logger.Errorf("Error synchronizing blocks: %v", err)
		}

		<-time.After(manager.pollingInterval)
	}
}
