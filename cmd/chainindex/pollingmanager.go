package main

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/appinterface/rdbstatusstore"
	"github.com/crypto-com/chainindex/infrastructure/feed"
	chainfeed "github.com/crypto-com/chainindex/infrastructure/feed/chain"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"

	"time"
)

const SYNC_INTERVAL = 5 * time.Second

type PollingManager struct {
	rdbConn               rdb.Conn
	client                tendermint.HTTPClient
	logger                applogger.Logger
	currentIndexingHeight int64
	Subject               *chainfeed.BlockSubject

	interval time.Duration
}

// NewBlockPollingFeed creates a new feed with polling for latest block starts at a specific height
func NewPollingManager(logger applogger.Logger, tendermintRPCUrl string, rdbConn rdb.Conn) *PollingManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	return &PollingManager{
		rdbConn: rdbConn,
		client:  *tendermintClient,
		logger:  logger,

		interval: SYNC_INTERVAL,
	}
}

func (pm *PollingManager) SetCurrentIndexingHeight() error {
	var err error
	statusStore := rdbstatusstore.NewRDbStatusStoreImpl(pm.rdbConn.ToHandle())
	pm.currentIndexingHeight, err = statusStore.GetLastIndexedBlockHeight()
	if err != nil {
		return fmt.Errorf("error running GetLastIndexedBlockHeight %v", err)
	}
	return nil
}

func (pm *PollingManager) UpdateIndexedHeight(handle *rdb.Handle, nextHeight int64) error {
	statusStore := rdbstatusstore.NewRDbStatusStoreImpl(handle)
	err := statusStore.UpdateLastIndexedBlockHeight(nextHeight)
	if err != nil {
		return fmt.Errorf("error running UpdateLastIndexedBlockHeight %v", err)
	}
	return nil
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (pm *PollingManager) SyncBlocks(latestHeight int64) error {
	for pm.currentIndexingHeight < latestHeight {
		// TODO: begin DB transaction
		// Request tendermint RPC
		block, rawBlock, err := pm.client.Block(pm.currentIndexingHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", pm.currentIndexingHeight, err)
		}

		// Create new block notification and notify subscribers
		notif := notification.NewBlockNotification(pm.currentIndexingHeight, block, rawBlock)
		pm.Subject.Notify(notif)

		// Current block indexing done, update db and sync next
		pm.UpdateIndexedHeight(pm.rdbConn.ToHandle(), pm.currentIndexingHeight + 1)
		pm.SetCurrentIndexingHeight()
		pm.logger.Infof("block height %d synced and events produced", block.Height)
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
