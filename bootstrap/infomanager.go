package bootstrap

import (
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
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
	tendermintRPCUrl string,
	insecureTendermintClient bool,
	strictGenesisParsing bool,
) *InfoManager {
	var tendermintClient *tendermint.HTTPClient
	if insecureTendermintClient {
		tendermintClient = tendermint.NewInsecureHTTPClient(
			tendermintRPCUrl,
			strictGenesisParsing,
		)
	} else {
		tendermintClient = tendermint.NewHTTPClient(
			tendermintRPCUrl,
			strictGenesisParsing,
		)
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
				manager.logger.Errorf("error querying TendermintApp status: %v", err)
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
