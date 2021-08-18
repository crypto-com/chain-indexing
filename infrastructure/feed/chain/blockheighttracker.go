package chain

import (
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"

	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/appinterface/tendermint"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

const DEFAULT_POLLING_INTERVAL = 5 * time.Second
const DEFAULT_MAX_RETRY_INTERVAL = 15 * time.Minute
const DEFAULT_MAX_RETRY_TIME = 0 // 0 means always retry

type BlockHeightTracker struct {
	logger applogger.Logger
	client tendermint.Client

	pollingInterval  time.Duration
	maxRetryInterval time.Duration
	maxRetryTime     time.Duration

	subscriptions []chan<- int64

	latestBlockHeight *int64
	rwMutex           sync.RWMutex
}

func NewBlockHeightTracker(logger applogger.Logger, client tendermint.Client) *BlockHeightTracker {
	tracker := &BlockHeightTracker{
		logger: logger.WithFields(applogger.LogFields{
			"module": "BlockHeightTracker",
		}),
		client: client,

		pollingInterval:  DEFAULT_POLLING_INTERVAL,
		maxRetryInterval: DEFAULT_MAX_RETRY_INTERVAL,
		maxRetryTime:     DEFAULT_MAX_RETRY_TIME,

		subscriptions: make([]chan<- int64, 0),

		latestBlockHeight: primptr.Int64Nil(),
	}

	go tracker.Run()

	return tracker
}

func (tracker *BlockHeightTracker) Run() {
	for {
		operation := func() error {
			height, err := tracker.client.LatestBlockHeight()
			if err != nil {
				return fmt.Errorf("error getting chain latest block height: %v", err)
			}

			for _, subscription := range tracker.subscriptions {
				select {
				case subscription <- height:
				default:
					tracker.logger.Info("block subscription channel is blocked, maybe busy?")
				}
			}

			tracker.rwMutex.Lock()
			tracker.latestBlockHeight = &height
			tracker.rwMutex.Unlock()

			tracker.logger.Infof("updated chain latest block height: %d", height)
			<-time.After(tracker.pollingInterval)
			return nil
		}
		notifyFn := func(opErr error, backoffDuration time.Duration) {
			tracker.logger.Errorf("retrying in %s: %v", backoffDuration.String(), opErr)
		}
		neverStopExponentialBackoff := backoff.NewExponentialBackOff()
		neverStopExponentialBackoff.MaxElapsedTime = tracker.maxRetryTime
		neverStopExponentialBackoff.MaxInterval = tracker.maxRetryInterval
		if err := backoff.RetryNotify(operation, neverStopExponentialBackoff, notifyFn); err != nil {
			tracker.logger.Errorf("stopping retry after too many errors: %s: %v", err)
		}
	}
}

func (tracker *BlockHeightTracker) Subscribe(ch chan<- int64) {
	tracker.subscriptions = append(tracker.subscriptions, ch)
}

func (tracker *BlockHeightTracker) GetLatestBlockHeight() *int64 {
	tracker.rwMutex.RLock()
	defer tracker.rwMutex.RUnlock()

	return tracker.latestBlockHeight
}
