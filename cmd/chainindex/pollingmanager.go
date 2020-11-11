package main

import (
	"fmt"
	"github.com/crypto-com/chainindex/infrastructure/feed"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"time"
)

const SYNC_INTERVAL = 5 * time.Second

type PollingManager struct {
	client            tendermint.HTTPClient
	currentSyncHeight int64
	logger            applogger.Logger
	Subject           *chainfeed.BlockSubject

	interval time.Duration
}

// NewBlockPollingFeed creates a new feed with polling for latest block starts at a specific height
func NewPollingManager(logger applogger.Logger, tendermintRPCUrl string, startAtHeight int64) *PollingManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &PollingManager{
		client:            *tendermintClient,
		currentSyncHeight: startAtHeight,
		logger:            logger,

		interval: SYNC_INTERVAL,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (pm *PollingManager) SyncBlocks(latestHeight int64) error {
	for pm.currentSyncHeight < latestHeight {
		// Request tendermint RPC
		block, rawBlock, err := pm.client.Block(pm.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", pm.currentSyncHeight, err)
		}

		// Create new block notification and notify subscribers
		notif := notification.NewBlockNotification(pm.currentSyncHeight, block, rawBlock)
		pm.Subject.Notify(notif)

		// Log and sync next
		pm.logger.Infof("Synced and produce event: %d", block.Height)
		pm.currentSyncHeight++
	}
	return nil
}

// InitSubject creates subject and attach subscribers
func (pm *PollingManager) InitSubject() *chainfeed.BlockSubject {
	// Currently only the chain processor subscriber
	// add more subscriber base on the need
	chainProcessor := chainfeed.NewBlockSubscriber(0)

	blockSubject := chainfeed.NewBlockSubject()
	blockSubject.Attach(chainProcessor)

	return blockSubject
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (pm *PollingManager) Run() {
	pm.Subject = pm.InitSubject()

	// Kickoff the block height syncing feed
	blockHeightFeed := feed.NewBlockHeightFeed(pm.client, pm.logger)
	go blockHeightFeed.Run()

	// Constantly gets the latest height and runs sync service to catch up to the height
	for {
		latestHeight := blockHeightFeed.GetLatestBlockHeight()

		if err := pm.SyncBlocks(latestHeight); err != nil {
			pm.logger.Errorf("error when the sync service catching to block %d", latestHeight)
		}

		<-time.After(pm.interval)
	}
}
