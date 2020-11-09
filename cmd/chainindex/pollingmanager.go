package main

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type PollingManager struct {
	client            tendermint.HTTPClient
	currentSyncHeight int64
	logger            applogger.Logger
	Subject           feed.Subject
}

// NewBlockPollingFeed creates a new feed with polling for latest block starts at a specific height
func NewPollingManager(logger applogger.Logger, tendermintRPCUrl string, startAtHeight int64) *PollingManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &PollingManager{
		client:            *tendermintClient,
		currentSyncHeight: startAtHeight,
		logger:            logger,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (pm *PollingManager) SyncBlocks(latestHeight int64) error {
	for pm.currentSyncHeight < latestHeight {
		// Get all subscriber ready for new block
		newBlockNotif := notification.NewNotifNewBlock(pm.currentSyncHeight)
		pm.Subject.Notify(newBlockNotif)

		// Request tendermint RPC
		block, rawBlock, err := pm.client.Block(pm.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", pm.currentSyncHeight, err)
		}

		// Prepare notification payload and dispatch
		blockNotif := notification.NewNotifBlockReceived(block, rawBlock)
		pm.Subject.Notify(blockNotif)

		doneNotif := notification.NewNotifCurrBlockDone()
		pm.Subject.Notify(doneNotif)

		// Log and sync next
		pm.logger.Infof("Synced and produce event: %d", block.Height)
		pm.currentSyncHeight++
	}
	return nil
}

// InitSubject creates subject and attach subscribers
func (pm *PollingManager) InitSubject() feed.Subject {
	// Currently only the chain processor subscriber
	// add more subscriber base on the need
	chainProcessor := chainfeed.NewBlockSubscriber(0)

	blockSubject := chainfeed.NewBlockSubject()
	blockSubject.Attach(chainProcessor)

	return blockSubject
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (pm *PollingManager) Run() error {
	pm.Subject = pm.InitSubject()

	for latestHeight := range chainfeed.LatestBlockHeightGenerator(pm.client) {
		// Notify all the listener
		fmt.Println("current height:", latestHeight)
		if err := pm.SyncBlocks(latestHeight); err != nil {
			pm.logger.Error("Error syncing blocks")
		}
	}

	return nil
}
