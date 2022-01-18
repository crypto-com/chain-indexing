package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UnbondingDelegationQueue interface {
	Upsert(row UnbondingDelegationQueueRow) error
	FindBy(completionTime utctime.UTCTime) (UnbondingDelegationQueueRow, bool, error)
	DequeueAllMatureUnbondingDelegationQueue(blockTime utctime.UTCTime) ([]DVPair, error)
}

type UnbondingDelegationQueueView struct {
	rdb *rdb.Handle
}

func NewUnbondingDelegationQueueView(handle *rdb.Handle) UnbondingDelegationQueue {
	return &UnbondingDelegationQueueView{
		handle,
	}
}

func (view *UnbondingDelegationQueueView) Upsert(row UnbondingDelegationQueueRow) error {

	return nil
}

func (view *UnbondingDelegationQueueView) FindBy(completionTime utctime.UTCTime) (UnbondingDelegationQueueRow, bool, error) {
	var row UnbondingDelegationQueueRow

	// TODO handle the row not found

	return row, true, nil
}

func (view *UnbondingDelegationQueueView) DequeueAllMatureUnbondingDelegationQueue(blockTime utctime.UTCTime) ([]DVPair, error) {

	// TODO find all mature UnbondingDelegationQueueRow, then concate their DVPairs

	// Optional TODO: de-duplicate, a same DVPair could appear multiple times, we should avoid that

	// TODO Delete the mature rows

	return nil, nil
}

// NOTES:
// - UNIQUE CompletionTime
type UnbondingDelegationQueueRow struct {
	CompletionTime utctime.UTCTime `json:"completionTime"`
	DVPairs        []DVPair        `json:"dvPairs"`
}

type DVPair struct {
	DelegatorAddress string `json:"delegatorAddress"`
	ValidatorAddress string `json:"validatorAddress"`
}
