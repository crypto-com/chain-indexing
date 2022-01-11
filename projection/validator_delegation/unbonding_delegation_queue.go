package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
)

func (projection *ValidatorDelegation) insertUBDQueue(
	rdbTxHandle *rdb.Handle,
	ubd view.UnbondingDelegationRow,
	completionTime utctime.UTCTime,
) error {
	dvPair := view.DVPair{
		DelegatorAddress: ubd.DelegatorAddress,
		ValidatorAddress: ubd.ValidatorAddress,
	}

	ubdQueueView := NewUBDQueue(rdbTxHandle)

	row, found, err := ubdQueueView.FindBy(completionTime)
	if err != nil {
		return fmt.Errorf("error finding unbonding delegation queue entry: %v", err)
	}
	if !found {
		row = view.UBDQueueRow{
			CompletionTime: completionTime,
			DVPairs:        []view.DVPair{dvPair},
		}
	} else {
		row.DVPairs = append(row.DVPairs, dvPair)
	}

	if err := ubdQueueView.Upsert(row); err != nil {
		return fmt.Errorf("error upserting unbonding delegation queue entry: %v", err)
	}

	return nil
}
