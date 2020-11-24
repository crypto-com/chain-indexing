package syncstrategy

import (
	"github.com/crypto-com/chainindex/entity/command"
)

type Strategy interface {
	Sync(currentHeight int64, latestHeight int64, worker SyncBlockWorker) ([][]command.Command, SyncedHeight, error)
}

type SyncBlockWorker = func(blockHeight int64) ([]command.Command, error)

type SyncedHeight = int64
