package bootstrap

import (
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap/config"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
)

// TODO: Move InfoManager to CronJob
const INFO_DEFAULT_POLLING_INTERVAL = 5 * time.Second

type InfoManager struct {
	rdbConn         rdb.Conn
	client          *tendermint.HTTPClient
	pollingInterval time.Duration
	viewStatus      *polling.Status
	logger          applogger.Logger
}

func NewInfoManager(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	tendermintApp config.TendermintApp,
) *InfoManager {
	var tendermintClient *tendermint.HTTPClient
	if tendermintApp.Insecure {
		tendermintClient = tendermint.NewInsecureHTTPClient(
			tendermintApp.HTTPRPCUrl,
			tendermintApp.StrictGenesisParsing,
		)
	} else {
		tendermintClient = tendermint.NewHTTPClient(
			tendermintApp.HTTPRPCUrl,
			tendermintApp.StrictGenesisParsing,
		)
	}
	if tendermintApp.MaybeAuthQueryKV != nil {
		tendermintClient.SetAuthQueryKV(tendermint.HTTPClientAuthKV{
			Key:   tendermintApp.MaybeAuthQueryKV.Key,
			Value: tendermintApp.MaybeAuthQueryKV.Value,
		})
	}

	viewStatus := polling.NewStatus(rdbConn.ToHandle())
	return &InfoManager{
		logger: logger.WithFields(applogger.LogFields{
			"module": "InfoManager",
		}),
		rdbConn:         rdbConn,
		client:          tendermintClient,
		viewStatus:      viewStatus,
		pollingInterval: INFO_DEFAULT_POLLING_INTERVAL,
	}

}

func (manager *InfoManager) Run() {
	manager.logger.Infof("InfoManager started")
	go func() {
		for {
			status, err := manager.client.Status()
			if err != nil {
				manager.logger.Errorf("error querying Tendermint status: %v", err)
				time.Sleep(manager.pollingInterval)
				continue
			}
			result := (*status)["result"]
			syncInfo := result.(map[string]interface{})["sync_info"]
			latestHeight := syncInfo.(map[string]interface{})["latest_block_height"].(string)

			err = manager.viewStatus.Upsert("LatestHeight", latestHeight)
			if err != nil {
				manager.logger.Errorf("error upserting latest height: %v", err)
				time.Sleep(manager.pollingInterval)
				continue
			}

			time.Sleep(manager.pollingInterval)
		}
	}()
}
