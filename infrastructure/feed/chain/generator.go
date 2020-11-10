package chain

import (
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"time"
)

const POLLING_INTERVAL = 5

func LatestBlockHeightGenerator(client tendermint.HTTPClient, logger applogger.Logger) chan int64 {
	latestBlockHeightChan := make(chan int64)

	go func() {
		for {
			// get latest block height of the chain
			height, err := client.LatestBlockHeight()
			if err != nil {
				logger.Errorf("error getting chain's latest block height", err)
			}
			latestBlockHeightChan <- height
			logger.Infof("current chain's latest block height: %d", height)

			<-time.After(POLLING_INTERVAL * time.Second)
		}
	}()

	return latestBlockHeightChan
}
