package main

import (
	"fmt"
	"time"

	"github.com/crypto-com/chainindex/usecase/parser"

	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type SyncManager struct {
	client  *tendermint.HTTPClient
	logger  applogger.Logger
	Subject *chainfeed.BlockSubject

	moduleAccounts *parser.ModuleAccounts

	currentSyncHeight int64

	pollingInterval time.Duration
}

// NewSyncManager creates a new feed with polling for latest block starts at a specific height
func NewSyncManager(
	logger applogger.Logger,
	tendermintRPCUrl string,
	startAtHeight int64,
	moduleAccounts *parser.ModuleAccounts,
) *SyncManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &SyncManager{
		client: tendermintClient,
		logger: logger.WithFields(applogger.LogFields{
			"module": "SyncManager",
		}),

		moduleAccounts: moduleAccounts,

		currentSyncHeight: startAtHeight,

		pollingInterval: DEFAULT_POLLING_INTERVAL,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (manager *SyncManager) SyncBlocks(latestHeight int64) error {
	for manager.currentSyncHeight < latestHeight {
		// Request tendermint RPC
		block, rawBlock, err := manager.client.Block(manager.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", manager.currentSyncHeight, err)
		}

		blockResults, err := manager.client.BlockResults(manager.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block_results at %d: %v", manager.currentSyncHeight, err)
		}

		// Create new block notification and notify subscribers
		notif := notification.NewBlockNotification(
			manager.currentSyncHeight, block, rawBlock, blockResults,
		)
		manager.Subject.Notify(notif)

		// Log and sync next
		manager.logger.Infof("Synced and produce event: %d", block.Height)
		manager.currentSyncHeight++
	}
	return nil
}

// InitSubject creates subject and attach subscribers
func (manager *SyncManager) InitSubject() *chainfeed.BlockSubject {
	// Currently only the chain processor subscriber
	// add more subscriber base on the need
	chainProcessor := chainfeed.NewBlockSubscriber(manager.moduleAccounts)

	blockSubject := chainfeed.NewBlockSubject()
	blockSubject.Attach(chainProcessor)

	return blockSubject
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (manager *SyncManager) Run() error {
	manager.Subject = manager.InitSubject()

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
