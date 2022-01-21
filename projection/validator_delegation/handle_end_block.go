package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
)

func (projection *ValidatorDelegation) handlePowerChanged(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	tendermintPubKey string,
	power string,
) error {

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.config.ConNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error getting consensusNodeAddress from tendermint pub key: %v", err)
	}

	validatorsView := NewValidators(rdbTxHandle)
	unbondingValidatorsView := NewUnbondingValidators(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByConsensusNodeAddr(consensusNodeAddress, height)
	if err != nil {
		return fmt.Errorf("error finding validator by consensusNodeAddr: %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found, conNodeAddr: %v", consensusNodeAddress)
	}

	if power == "0" {
		validator.Status = types.UNBONDING

		// UnbondingHeight is the height when Unbonding start
		// UnbondingTime is the time when Unbonding is finished
		validator.UnbondingHeight = height
		validator.UnbondingTime = blockTime.Add(projection.config.UnbondingTime)

		// Insert the validator to UnbondingValidators set
		if err := unbondingValidatorsView.Insert(validator.OperatorAddress, validator.UnbondingTime); err != nil {
			return fmt.Errorf("error inserting unbonding validator entry: %v", err)
		}

	} else {
		// In this case, power should be larger than 0.

		validator.Status = types.BONDED

		// Remove the validator from UnbondingValidators set, if exist
		if err := unbondingValidatorsView.RemoveIfExist(validator.OperatorAddress); err != nil {
			return fmt.Errorf("error removing unbonding validator entry: %v", err)
		}
	}
	validator.Power = power

	// Update the validator
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error updating validator: %v", err)
	}

	return nil
}

// Update Validator.Status to `bonded`, if the validator finish the unbonding period.
func (projection *ValidatorDelegation) handleMatureUnbondingValidators(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
) error {

	validatorsView := NewValidators(rdbTxHandle)
	unbondingValidatorsView := NewUnbondingValidators(rdbTxHandle)

	// Get all UnbondingValidator entry that finish the Unbonding period
	matureEntries, err := unbondingValidatorsView.GetMatureEntries(blockTime)
	if err != nil {
		return fmt.Errorf("error getting all mature unbonding validator entries: %v", err)
	}

	// Update validators status
	for _, entry := range matureEntries {
		validator, found, err := validatorsView.FindByOperatorAddr(entry.OperatorAddress, height)
		if err != nil {
			return fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
		}
		if !found {
			return fmt.Errorf("error validator not found: %v, %v", entry.OperatorAddress, height)
		}

		validator.Status = types.UNBONDED

		if err := validatorsView.Update(validator); err != nil {
			return fmt.Errorf("error updating validator: %v", err)
		}

	}

	// Remove those mature UnbondingValidator entry
	if err := unbondingValidatorsView.DeleteMatureEntries(blockTime); err != nil {
		return fmt.Errorf("error removing all mature unbonding validator entries: %v", err)
	}

	return nil
}

// Find mature UnbindingDelegation through UnbondingDelegationQueue and then handle them one by one
func (projection *ValidatorDelegation) handleMatureUnbondingDelegationQueueEntries(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
) error {

	ubdQueueView := NewUnbondingDelegationQueue(rdbTxHandle)

	// Get all mature unbonding delegations from UnbondingDelegationQueue
	matureUnbonds, err := ubdQueueView.DequeueAllMatureUnbondingDelegationQueue(blockTime)
	if err != nil {
		return fmt.Errorf("error dequeue all mature entries from unbonding delegation queue: %v", err)
	}

	for _, dvPair := range matureUnbonds {

		if err := projection.completeUnbonding(
			rdbTxHandle,
			height,
			blockTime,
			dvPair.DelegatorAddress,
			dvPair.ValidatorAddress,
		); err != nil {
			// Not 100% sure why the error is ignored here.
			// This is probably due to a same DVPair could show up in matureUnbonds for multiple time.
			// reference: https://github.com/cosmos/cosmos-sdk/blob/59793e68fd7227c0ea31cf9456822dd566db4da7/x/staking/keeper/val_state_change.go#L46
			continue
		}

	}

	return nil
}

// Find mature Redelegation through RedelegationQueue and then handle them one by one
func (projection *ValidatorDelegation) handleMatureRedelegationQueueEntries(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
) error {

	redelegationQueueView := NewRedelegationQueue(rdbTxHandle)

	// Get all mature redelegations from RedelegationQueue
	matureRelegations, err := redelegationQueueView.DequeueAllMatureRedelegationQueue(blockTime)
	if err != nil {
		return fmt.Errorf("error dequeue all mature entries from redelegation queue: %v", err)
	}

	for _, dvvTriplet := range matureRelegations {

		if err := projection.completeRedelegation(
			rdbTxHandle,
			height,
			blockTime,
			dvvTriplet.DelegatorAddress,
			dvvTriplet.ValidatorSrcAddress,
			dvvTriplet.ValidatorDstAddress,
		); err != nil {
			// Not 100% sure why the error is ignored here.
			// This is probably due to a same DVVTriplet could show up in matureEntries for multiple time.
			// reference: https://github.com/cosmos/cosmos-sdk/blob/28541e7ca36097d6a98637c26c276a5d55f408b0/x/staking/keeper/val_state_change.go#L82
			continue
		}

	}

	return nil
}
