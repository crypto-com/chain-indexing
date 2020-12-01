package main

import (
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

const INFO_DEFAULT_POLLING_INTERVAL = 1 * time.Second

type InfoManager struct {
	rdbConn         rdb.Conn
	client          *tendermint.HTTPClient
	pollingInterval time.Duration
	viewStatus      *polling.Status
	logger 			applogger.Logger
}

func NewInfoManager(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	tendermintRPCUrl string,
) *InfoManager {
	tendermintClient := tendermint.NewHTTPClient(tendermintRPCUrl)

	viewStatus := polling.NewStatus(rdbConn.ToHandle())
	return &InfoManager{
		logger: logger,
		rdbConn:    rdbConn,
		client:     tendermintClient,
		viewStatus: viewStatus,
		pollingInterval: INFO_DEFAULT_POLLING_INTERVAL,
	}

}

func (manager *InfoManager) Run() {
	manager.logger.Infof("infomanager started")
	go func() {
		for {
			status, _ := manager.client.Status()
			result := (*status)["result"]
			syncInfo := result.(map[string]interface{})["sync_info"]
			latestHeight := syncInfo.(map[string]interface{})["latest_block_height"].(string)
			// upsert
			_= manager.viewStatus.Insert("LatestHeight", latestHeight)
			time.Sleep(manager.pollingInterval)
		}
	}()
}
