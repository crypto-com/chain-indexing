package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
)

func (projection *ValidatorDelegation) insertRedelegationQueue(
	rdbTxHandle *rdb.Handle,
	red view.RedelegationRow,
	completionTime utctime.UTCTime,
) error {
	dvvTriplet := view.DVVTriplet{
		DelegatorAddress:    red.DelegatorAddress,
		ValidatorSrcAddress: red.ValidatorSrcAddress,
		ValidatorDstAddress: red.ValidatorDstAddress,
	}

	redelegationQueueView := NewRedelegationQueue(rdbTxHandle)

	row, found, err := redelegationQueueView.FindBy(completionTime)
	if err != nil {
		return fmt.Errorf("error redelegationQueueView.FindBy(): %v", err)
	}
	if !found {
		row = view.RedelegationQueueRow{
			CompletionTime: completionTime,
			DVVTriplets:    []view.DVVTriplet{dvvTriplet},
		}
	} else {
		row.DVVTriplets = append(row.DVVTriplets, dvvTriplet)
	}

	if err := redelegationQueueView.Upsert(row); err != nil {
		return fmt.Errorf("error redelegationQueueView.Upsert(): %v", err)
	}

	return nil
}
