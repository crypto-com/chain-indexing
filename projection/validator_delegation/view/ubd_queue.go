package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UBDQueue interface {
	Insert(row UBDQueueRow) error
	Update(row UBDQueueRow) error
	FindBy(completionTime utctime.UTCTime) (UBDQueueRow, bool, error)
	GetMatureEntries(blockTime utctime.UTCTime) ([]DVPair, error)
	DeleteMatureEntries(blockTime utctime.UTCTime) error
}

type UBDQueueView struct {
	rdb *rdb.Handle
}

func NewUBDQueueView(handle *rdb.Handle) UBDQueue {
	return &UBDQueueView{
		handle,
	}
}

func (view *UBDQueueView) Insert(row UBDQueueRow) error {

	return nil
}

func (view *UBDQueueView) Update(row UBDQueueRow) error {

	return nil
}

func (view *UBDQueueView) FindBy(completionTime utctime.UTCTime) (UBDQueueRow, bool, error) {
	var row UBDQueueRow

	// TODO handle the row not found

	return row, true, nil
}

func (view *UBDQueueView) GetMatureEntries(blockTime utctime.UTCTime) ([]DVPair, error) {

	// TODO find all mature UDBQueueRow, then concate their DVPairs

	// TODO de-duplicate, a same DVPair could appear multiple times, we should avoid that

	return nil, nil
}

func (view *UBDQueueView) DeleteMatureEntries(blockTime utctime.UTCTime) error {

	return nil
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
