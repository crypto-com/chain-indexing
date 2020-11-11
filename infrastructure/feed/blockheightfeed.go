package feed

import (
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"time"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type BlockHeightFeed struct {
	height int64
	logger applogger.Logger
	client tendermint.HTTPClient

	interval time.Duration
}

func NewBlockHeightFeed(client tendermint.HTTPClient, logger applogger.Logger) *BlockHeightFeed {
	return &BlockHeightFeed{
		height: int64(0),
		logger: logger,
		client: client,

		interval: DEFAULT_POLLING_INTERVAL,
	}
}

func (feed *BlockHeightFeed) GetLatestBlockHeight() int64 {
	return feed.height
}

func (feed *BlockHeightFeed) Run() {
	latestBlockHeightChan := make(chan int64)

	go func() {
		for {
			// get latest block height of the chain
			height, err := feed.client.LatestBlockHeight()
			if err != nil {
				feed.logger.Errorf("error getting chain's latest block height", err)
			}
			latestBlockHeightChan <- height
			feed.logger.Infof("chain's current latest block height: %d", height)

			<-time.After(feed.interval)
		}
	}()

	for {
		feed.height = <-latestBlockHeightChan
	}
}
