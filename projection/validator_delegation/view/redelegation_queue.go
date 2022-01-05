package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type RedelegationQueue interface {
	Upsert(row RedelegationQueueRow) error
	FindBy(completionTime utctime.UTCTime) (RedelegationQueueRow, bool, error)
	DequeueAllMatureRedelegationQueue(blockTime utctime.UTCTime) ([]DVVTriplet, error)
}

type RedelegationQueueView struct {
	rdb *rdb.Handle
}

func NewRedelegationQueueView(handle *rdb.Handle) RedelegationQueue {
	return &RedelegationQueueView{
		handle,
	}
}

func (view *RedelegationQueueView) Upsert(row RedelegationQueueRow) error {

	return nil
}

func (view *RedelegationQueueView) FindBy(completionTime utctime.UTCTime) (RedelegationQueueRow, bool, error) {

	// TODO handle the row not found

	return RedelegationQueueRow{}, true, nil
}

func (view *RedelegationQueueView) DequeueAllMatureRedelegationQueue(blockTime utctime.UTCTime) ([]DVVTriplet, error) {

	// TODO find all mature RedelegationQueueRow, then concate their DVVTriplets

	// Optional TODO: de-duplicate, a same DVVTriplet could appear multiple times, we should avoid that

	// TODO Delete the mature rows

	return nil, nil
}

// TODO UNIQUE CompletionTime
type RedelegationQueueRow struct {
	CompletionTime utctime.UTCTime `json:"completionTime"`
	DVVTriplets    []DVVTriplet    `json:"dvvTriplets"`
}

type DVVTriplet struct {
	DelegatorAddress    string `json:"delegatorAddress"`
	ValidatorSrcAddress string `json:"validatorSrcAddress"`
	ValidatorDstAddress string `json:"validatorDstAddress"`
}
