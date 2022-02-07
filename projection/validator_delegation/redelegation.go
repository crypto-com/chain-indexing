package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) handleRedelegate(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
	amount coin.Int,
) error {

	shares, err := projection.unbondAmountToShares(
		rdbTxHandle,
		height,
		validatorSrcAddress,
		delegatorAddress,
		amount,
	)
	if err != nil {
		return fmt.Errorf("error converting the unbond amout to corresponding shares on that validator: %v", err)
	}

	srcValReturnAmount, err := projection.unbond(
		rdbTxHandle,
		height,
		validatorSrcAddress,
		delegatorAddress,
		shares,
	)
	if err != nil {
		return fmt.Errorf("error unbonding shares from delegation and validator: %v", err)
	}

	dstValSharesCreated, err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorDstAddress,
		delegatorAddress,
		srcValReturnAmount,
	)
	if err != nil {
		return fmt.Errorf("error hanlding delegate operation: %v", err)
	}

	completionTime, completeNow, err := projection.calculateRedelegationCompleteTime(
		rdbTxHandle,
		height,
		blockTime,
		validatorSrcAddress,
	)
	if err != nil {
		return fmt.Errorf("error calculating relegation completion time: %v", err)
	}

	if completeNow { // no need to create the redelegation object
		return nil
	}

	red, err := projection.setRedelegationEntry(
		rdbTxHandle,
		delegatorAddress,
		validatorSrcAddress,
		validatorDstAddress,
		height,
		completionTime,
		srcValReturnAmount,
		dstValSharesCreated,
	)
	if err != nil {
		return fmt.Errorf("error setting redelegation entry: %v", err)
	}

	if err := projection.insertRedelegationQueue(rdbTxHandle, red, completionTime); err != nil {
		return fmt.Errorf("error inserting RedelegationQueue entry: %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) calculateRedelegationCompleteTime(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	validatorSrcAddress string,
) (completionTime utctime.UTCTime, completeNow bool, err error) {

	validatorsView := NewValidators(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByOperatorAddr(validatorSrcAddress, height)
	if err != nil {
		return completionTime, false, fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
	}

	switch {
	case !found || validator.IsBonded():
		// the longest wait - just unbonding period from now
		completionTime = blockTime.Add(projection.config.UnbondingTime)

		return completionTime, false, nil

	case validator.IsUnbonded():
		return completionTime, true, nil

	case validator.IsUnbonding():
		return validator.UnbondingTime, false, nil

	default:
		return completionTime, false, fmt.Errorf("unknown validator status: %s", validator.Status)
	}
}

func (projection *ValidatorDelegation) setRedelegationEntry(
	rdbTxHandle *rdb.Handle,
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
	sharesDst coin.Dec,
) (view.RedelegationRow, error) {

	redelegationsView := NewRedelegations(rdbTxHandle)

	red, found, err := redelegationsView.FindBy(delegatorAddress, validatorSrcAddress, validatorDstAddress, creationHeight)
	if err != nil {
		return red, fmt.Errorf("error finding redelegation: %v", err)
	}
	if found {
		red.AddEntry(creationHeight, completionTime, balance, sharesDst)

		if err := redelegationsView.Update(red); err != nil {
			return red, fmt.Errorf("error updating reledegation record: %v", err)
		}
	} else {
		red = view.NewRedelegationRow(
			delegatorAddress,
			validatorSrcAddress,
			validatorDstAddress,
			creationHeight,
			completionTime,
			balance,
			sharesDst,
		)

		if err := redelegationsView.Insert(red); err != nil {
			return red, fmt.Errorf("error inserting reledegation record: %v", err)
		}
	}

	return red, nil
}

func (projection *ValidatorDelegation) completeRedelegation(
	rdbTxHandle *rdb.Handle,
	height int64,
	currentBlockTime utctime.UTCTime,
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
) error {

	redelegationsView := NewRedelegations(rdbTxHandle)

	red, found, err := redelegationsView.FindBy(delegatorAddress, validatorSrcAddress, validatorDstAddress, height)
	if err != nil {
		return fmt.Errorf("error in finding redelegation: %v", err)
	}
	if !found {
		return fmt.Errorf("Redelegation not found: %v, %v, %v, %v", delegatorAddress, validatorSrcAddress, validatorDstAddress, height)
	}

	for i := 0; i < len(red.Entries); i++ {
		entry := red.Entries[i]

		if entry.IsMature(currentBlockTime) {
			red.RemoveEntry(int64(i))
			i--
		}

		// Update the redelegation or delete it if there are no more entries
		if len(red.Entries) == 0 {

			if err := redelegationsView.Delete(red); err != nil {
				return fmt.Errorf("error deleting redelegation entry: %v", err)
			}

		} else {

			if err := redelegationsView.Update(red); err != nil {
				return fmt.Errorf("error updating reledegation record: %v", err)
			}

		}
	}

	return nil
}
