package tendermint

import (
	"fmt"
	"github.com/crypto-com/chainindex/entity/executor"
	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/infrastructure/parser"
	"os"
	"time"
)

type BlockPollingService struct {
	client            HTTPClient
	currentSyncHeight int64
	logger            *infrastructure.ZerologLogger
}

// NewBlockPollingFeed creates a new feed with polling for latest block starts at a specific height
func NewBlockPollingService(tendermintRPCUrl string, startAtHeight int64) *BlockPollingService {
	tendermintClient := NewHTTPClient(tendermintRPCUrl)
	logger := infrastructure.NewZerologLoggerWithColor(os.Stderr)

	return &BlockPollingService{
		client:            *tendermintClient,
		currentSyncHeight: startAtHeight,
		logger:            logger,
	}
}

// SyncBlocks sync block data until latestHeight
// TODO: add BlockResult and producing events, add error retry
func (s *BlockPollingService) SyncBlocks(latestHeight int64) error {
	for s.currentSyncHeight < latestHeight {
		// TODO: sync both Block and BlockResult
		block, rawBlock, err := s.client.Block(s.currentSyncHeight)
		if err != nil {
			return fmt.Errorf("error getting chain's block at %d: %v", s.currentSyncHeight, err)
		}

		// height 1001 => struct BlockParser, BlockResultParser, ValidatorParser
		// TODO: produce commands with parsers, extract following logic to a 'manager'
		cmds, err := parser.ParseBlockToCommands(block, rawBlock)
		if err != nil {
			return fmt.Errorf("error parsing block data to commands %d: %v", s.currentSyncHeight, err)
		}

		// Create executor for event generation and persistent
		executor := executor.New()
		executor.AddCommands(cmds)

		// TODO: Process commands to events & store

		// Log and sync next
		s.logger.Infof("Synced and produce event: %d %v", block.Height, executor.Commands)
		s.currentSyncHeight++
	}
	return nil
}

// Run starts the polling service for blocks
func (s *BlockPollingService) Run() error {
	// Background routine for polling latestBlockHeight
	go func() {
		for {
			// get latest block height of the chain
			latestHeight, err := s.client.LatestBlockHeight()
			if err != nil {
				s.logger.Errorf("error getting chain's latest block height: %v", err)
			}

			// already sync to latest height, wait
			if s.currentSyncHeight >= latestHeight {
				continue
			}

			// sync until to latestHeight
			err = s.SyncBlocks(latestHeight)
			if err != nil {
				s.logger.Errorf("error sync block data to latest block height: %v", err)
			}

			<-time.After(5 * time.Second)
		}
	}()

	// TODO: refactor later with API server
	// Main thread waits for background thread to sync blocks
	onExitCh := make(chan bool)
	<-onExitCh

	return nil
}
