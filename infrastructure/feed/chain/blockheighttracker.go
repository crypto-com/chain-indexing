package chain

import (
	"sync"
	"time"

	"github.com/crypto-com/chainindex/internal/primptr"

	"github.com/crypto-com/chainindex/appinterface/tendermint"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second

type BlockHeightTracker struct {
	logger applogger.Logger
	client tendermint.Client

	pollingInterval time.Duration

	latestBlockHeight *int64
	rwMutex           sync.RWMutex
}

func NewBlockHeightTracker(logger applogger.Logger, client tendermint.Client) *BlockHeightTracker {
	tracker := &BlockHeightTracker{
		logger: logger.WithFields(applogger.LogFields{
			"module": "BlockHeightTracker",
		}),
		client: client,

		pollingInterval: DEFAULT_POLLING_INTERVAL,

		latestBlockHeight: primptr.Int64Nil(),
	}

	go tracker.Run()

	return tracker
}

func (tracker *BlockHeightTracker) Run() {
	for {
		height, err := tracker.client.LatestBlockHeight()
		if err != nil {
			tracker.logger.Errorf("error getting chain latest block height", err)
			<-time.After(1 * time.Second)
			continue
		}

		tracker.rwMutex.Lock()
		tracker.latestBlockHeight = &height
		tracker.rwMutex.Unlock()

		tracker.logger.Infof("updated chain latest block height: %d", height)
		<-time.After(tracker.pollingInterval)
	}

}

func (tracker *BlockHeightTracker) GetLatestBlockHeight() *int64 {
	tracker.rwMutex.RLock()
	defer tracker.rwMutex.RUnlock()

	return tracker.latestBlockHeight
}
