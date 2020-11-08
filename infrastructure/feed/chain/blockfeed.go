package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/feed"
	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/infrastructure/notification"
	"github.com/crypto-com/chainindex/infrastructure/processor/chain"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	"os"
)

type BlockFeed struct {
	client            tendermint.HTTPClient
	currentSyncHeight int64
	logger            *infrastructure.ZerologLogger
	Subject           feed.Subject
}


// NewBlockPollingFeed creates a new feed with polling for latest block starts at a specific height
func NewBlockFeed(tendermintRPCUrl string, startAtHeight int64) *BlockFeed {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stderr)

	return &BlockFeed{
		client:            *tendermintClient,
		currentSyncHeight: startAtHeight,
		logger:            logger,
	}
}

// SyncBlocks makes request to tendermint, create and dispatch notifications
func (bf *BlockFeed) SyncBlocks(latestHeight int64) error {
	for bf.currentSyncHeight < latestHeight {
		// TODO: sync both Block and BlockResult
		block, rawBlock, err := bf.client.Block(bf.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", bf.currentSyncHeight, err)
		}

		// Prepare notification payload and dispatch
		blockNotif := notification.NewBlockNotif(block, rawBlock)
		bf.Subject.Notify(blockNotif)

		doneNotif := notification.NewCurrBlockDoneNotif()
		bf.Subject.Notify(doneNotif)

		// Log and sync next
		bf.logger.Infof("Synced and produce event: %d", block.Height)
		bf.currentSyncHeight++
	}
	return nil
}

// InitSubject creates subject and attach subscribers
func (bf *BlockFeed) InitSubject() feed.Subject {
	// Currently only the chain processor subscriber
	// add more subscriber base on the need
	chainProcessor := chain.NewChainProcessor(0)

	blockSubject := NewBlockSubject()
	blockSubject.Attach(chainProcessor)

	return blockSubject
}

// Run starts the polling service for blocks
// new BlockFeedSubject and add listeners
func (bf *BlockFeed) Run() error {
	bf.Subject = bf.InitSubject()

	for latestHeight := range LatestBlockHeightGenerator(bf.client) {
		// Notify all the listener
		fmt.Println("current height:", latestHeight)
		if err := bf.SyncBlocks(latestHeight); err != nil {
			bf.logger.Error("Error syncing blocks")
		}
	}

	return nil
}
