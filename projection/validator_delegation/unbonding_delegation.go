package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) setUnbondingDelegationEntry(
	rdbTxHandle *rdb.Handle,
	delegatorAddress string,
	validatorAddress string,
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
) (view.UnbondingDelegationRow, error) {

	unbondingDelegationsView := NewUnbondingDelegations(rdbTxHandle)

	ubd, found, err := unbondingDelegationsView.FindBy(delegatorAddress, validatorAddress, creationHeight)
	if err != nil {
		return ubd, fmt.Errorf("error finding unbonding delegation record: %v", err)
	}
	if found {
		ubd.AddEntry(creationHeight, completionTime, balance)

		if err := unbondingDelegationsView.Update(ubd); err != nil {
			return ubd, fmt.Errorf("error updating unbonding delegation record: %v", err)
		}
	} else {
		ubd = view.NewUnbondingDelegationRow(delegatorAddress, validatorAddress, creationHeight, completionTime, balance)

		if err := unbondingDelegationsView.Insert(ubd); err != nil {
			return ubd, fmt.Errorf("error inserting unbonding delegation record: %v", err)
		}
	}

	return ubd, nil
}

func (projection *ValidatorDelegation) completeUnbonding(
	rdbTxHandle *rdb.Handle,
	height int64,
	currentBlockTime utctime.UTCTime,
	delegatorAddress string,
	validatorAddress string,
) error {

	unbondingDelegationsView := NewUnbondingDelegations(rdbTxHandle)

	ubd, found, err := unbondingDelegationsView.FindBy(delegatorAddress, validatorAddress, height)
	if err != nil {
		return fmt.Errorf("error finding unbonding delegation: %v", err)
	}
	if !found {
		return fmt.Errorf("UnbondingDelegation not found: %v, %v, %v", delegatorAddress, validatorAddress, height)
	}

	for i := 0; i < len(ubd.Entries); i++ {
		entry := ubd.Entries[i]

		if entry.IsMature(currentBlockTime) {
			ubd.RemoveEntry(int64(i))
			i--
		}

		// Update the unbonding delegation or delete it if there are no more entries
		if len(ubd.Entries) == 0 {

			if err := unbondingDelegationsView.Delete(ubd); err != nil {
				return fmt.Errorf("error deleting unbonding delegation: %v", err)
			}

		} else {

			if err := unbondingDelegationsView.Update(ubd); err != nil {
				return fmt.Errorf("error updating unbonding delegation: %v", err)
			}

		}
	}

	return nil
}
