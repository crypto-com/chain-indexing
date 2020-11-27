package main

import (
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/rdb"

	applogger "github.com/crypto-com/chainindex/internal/logger"
	"github.com/crypto-com/chainindex/usecase/parser"
)

type IndexService struct {
	logger  applogger.Logger
	rdbConn rdb.Conn

	baseDenom             string
	consNodeAddressPrefix string
	windowSize            int
	tendermintHTTPRPCURL  string
}

// NewIndexService creates a new server instance for polling and indexing
func NewIndexService(logger applogger.Logger, rdbConn rdb.Conn, config *FileConfig) *IndexService {
	return &IndexService{
		logger:  logger,
		rdbConn: rdbConn,

		baseDenom:             config.Blockchain.BaseDenom,
		consNodeAddressPrefix: config.Blockchain.ConNodeAddressPrefix,
		windowSize:            config.Sync.WindowSize,
		tendermintHTTPRPCURL:  config.Tendermint.HTTPRPCURL,
	}
}

// Run function runs the polling server to index the data from Tendermint
func (service *IndexService) Run() error {
	service.logger.Debug("TODO: should load module accounts configuration")

	txDecoder := parser.NewTxDecoder(service.baseDenom)

	syncManager := NewSyncManager(
		service.logger,
		service.rdbConn,
		service.consNodeAddressPrefix,
		service.windowSize,
		service.tendermintHTTPRPCURL,
		txDecoder,
	)
	if err := syncManager.Run(); err != nil {
		return fmt.Errorf("error running sync manager %v", err)
	}

	return nil
}
