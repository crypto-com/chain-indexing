package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UBDQueue interface {
	Upsert(row UBDQueueRow) error
	FindBy(completionTime utctime.UTCTime) (UBDQueueRow, bool, error)
	DequeueAllMatureUBDQueue(blockTime utctime.UTCTime) ([]DVPair, error)
}

type UBDQueueView struct {
	rdb *rdb.Handle
}

func NewUBDQueueView(handle *rdb.Handle) UBDQueue {
	return &UBDQueueView{
		handle,
	}
}

func (view *UBDQueueView) Upsert(row UBDQueueRow) error {

	return nil
}

func (view *UBDQueueView) FindBy(completionTime utctime.UTCTime) (UBDQueueRow, bool, error) {
	var row UBDQueueRow

	// TODO handle the row not found

	return row, true, nil
}

func (view *UBDQueueView) DequeueAllMatureUBDQueue(blockTime utctime.UTCTime) ([]DVPair, error) {

	// TODO find all mature UDBQueueRow, then concate their DVPairs

	// Optional TODO: de-duplicate, a same DVPair could appear multiple times, we should avoid that

	// TODO Delete the mature rows

	return nil, nil
}

// TODO UNIQUE CompletionTime
type UBDQueueRow struct {
	CompletionTime utctime.UTCTime `json:"completionTime"`
	DVPairs        []DVPair        `json:"dvPairs"`
}

type DVPair struct {
	DelegatorAddress string `json:"delegatorAddress"`
	ValidatorAddress string `json:"validatorAddress"`
}
