package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) handleCreateNewValidator(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	tendermintPubKey string,
	amount coin.Int,
	minSelfDelegationInString string,
) error {

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.conNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error in GetConsensusNodeAddress: %v", err)
	}

	minSelfDelegation, ok := coin.NewIntFromString(minSelfDelegationInString)
	if !ok {
		return fmt.Errorf("Failed to parse minSelfDelegationInString: %v", minSelfDelegationInString)
	}

	// Insert an ValidatorRow.
	validatorsView := NewValidators(rdbTxHandle)
	validatorRow := view.ValidatorRow{
		Height: height,

		OperatorAddress:      validatorAddress,
		ConsensusNodeAddress: consensusNodeAddress,

		Status: types.UNBONDED,
		Jailed: false,
		Power:  "0",

		UnbondingHeight: int64(0),
		UnbondingTime:   utctime.UTCTime{},

		Tokens:            coin.ZeroInt(),
		Shares:            coin.ZeroDec(),
		MinSelfDelegation: minSelfDelegation,
	}
	if err := validatorsView.Insert(validatorRow); err != nil {
		return fmt.Errorf("error validatorsView.Insert(): %v", err)
	}

	if _, err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	); err != nil {
		return fmt.Errorf("error in projection.handleDelegate(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleEditValidator(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	minSelfDelegationInString string,
) error {

	validatorsView := NewValidators(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByOperatorAddr(validatorAddress, height)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	minSelfDelegation, ok := coin.NewIntFromString(minSelfDelegationInString)
	if !ok {
		return fmt.Errorf("Failed to parse minSelfDelegationInString: %v", minSelfDelegationInString)
	}
	// Update Valdiator
	validator.MinSelfDelegation = minSelfDelegation

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}

// handleDelegate - Handle create/update delegation. Return newShares and error.
func (projection *ValidatorDelegation) handleDelegate(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	amount coin.Int,
) (coin.Dec, error) {

	validatorsView := NewValidators(rdbTxHandle)
	delegationsView := NewDelegations(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByOperatorAddr(validatorAddress, height)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}
	if !found {
		return coin.ZeroDec(), fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	// Create/Get Delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		delegation = view.DelegationRow{
			Height:           height,
			ValidatorAddress: validatorAddress,
			DelegatorAddress: delegatorAddress,
			Shares:           coin.ZeroDec(),
		}

		// Insert an DelegationRow record
		if errAddDelRecord := delegationsView.Insert(delegation); errAddDelRecord != nil {
			return coin.ZeroDec(), fmt.Errorf("error delegationsView.Insert(): %v", errAddDelRecord)
		}
	}

	// Calculate new shares
	newShares, err := utils.ValidatorIssueNewShares(validator, amount)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error IssueNewShares(): %v", err)
	}

	// Update Delegation
	delegation.Shares = delegation.Shares.Add(newShares)
	if err := delegationsView.Update(delegation); err != nil {
		return coin.ZeroDec(), fmt.Errorf("error delegationsView.Update(): %v", err)
	}

	// Update Validator
	validator.Tokens = validator.Tokens.Add(amount)
	validator.Shares = validator.Shares.Add(newShares)
	if err := validatorsView.Update(validator); err != nil {
		return coin.ZeroDec(), fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return newShares, nil
}

func (projection *ValidatorDelegation) handleUndelegate(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	amount coin.Int,
	completionTime utctime.UTCTime,
) error {

	shares, err := projection.unbondAmountToShares(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	)
	if err != nil {
		return fmt.Errorf("error projection.unbondAmountToShares(): %v", err)
	}

	returnAmount, err := projection.unbond(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		shares,
	)
	if err != nil {
		return fmt.Errorf("error projection.unbond(): %v", err)
	}

	ubd, err := projection.setUnbondingDelegationEntry(
		rdbTxHandle,
		delegatorAddress,
		validatorAddress,
		height,
		completionTime,
		returnAmount,
	)
	if err != nil {
		return fmt.Errorf("error projection.setUnbondingDelegationEntry(): %v", err)
	}

	if err := projection.insertUBDQueue(rdbTxHandle, ubd, completionTime); err != nil {
		return fmt.Errorf("error projection.insertUBDQueue(): %v", err)
	}

	return nil
}

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
		return fmt.Errorf("error projection.unbondAmountToShares(): %v", err)
	}

	srcValReturnAmount, err := projection.unbond(
		rdbTxHandle,
		height,
		validatorSrcAddress,
		delegatorAddress,
		shares,
	)
	if err != nil {
		return fmt.Errorf("error projection.unbond(): %v", err)
	}

	dstValSharesCreated, err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorDstAddress,
		delegatorAddress,
		srcValReturnAmount,
	)
	if err != nil {
		return fmt.Errorf("error projection.handleDelegate(): %v", err)
	}

	completionTime, completeNow, err := projection.calculateRedelegationCompleteTime(
		rdbTxHandle,
		height,
		blockTime,
		validatorSrcAddress,
	)

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
		return fmt.Errorf("error projection.setRedelegationEntry(): %v", err)
	}

	if err := projection.insertRedelegationQueue(rdbTxHandle, red, completionTime); err != nil {
		return fmt.Errorf("error projection.insertRedelegationQueue(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleValidatorJailed(
	rdbTxHandle *rdb.Handle,
	height int64,
	consensusNodeAddress string,
) error {

	validatorsView := NewValidators(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByConsensusNodeAddr(
		consensusNodeAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByConsensusNodeAddr(): %v", err)
	}

	if err := projection.jailValidator(rdbTxHandle, validator); err != nil {
		return fmt.Errorf("error in projection.jailValidator(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleValidatorUnjailed(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
) error {

	validatorsView := NewValidators(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByOperatorAddr(validatorAddress, height)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	// Update Valdiator
	validator.Jailed = false

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handlePowerChanged(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	tendermintPubKey string,
	power string,
) error {

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.conNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error GetConsensusNodeAddress(): %v", err)
	}

	validatorsView := NewValidators(rdbTxHandle)
	unbondingValidatorsView := NewUnbondingValidators(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByConsensusNodeAddr(
		consensusNodeAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByConsensusNodeAddr(): %v", err)
	}

	if power == "0" {
		validator.Status = types.UNBONDING

		// UnbondingHeight is the height when Unbonding start
		// UnbondingTime is the time when Unbonding is finished
		validator.UnbondingHeight = height
		validator.UnbondingTime = blockTime.Add(projection.UnbondingTime)

		// Insert the validator to UnbondingValidators set
		if err := unbondingValidatorsView.Insert(validator.OperatorAddress, validator.UnbondingTime); err != nil {
			return fmt.Errorf("error unbondingValidatorsView.Insert(): %v", err)
		}

	} else {
		// In this case, power should be larger than 0.

		validator.Status = types.BONDED

		// Remove the validator from UnbondingValidators set, if exist
		if err := unbondingValidatorsView.RemoveIfExist(validator.OperatorAddress); err != nil {
			return fmt.Errorf("error unbondingValidatorsView.RemoveIfExist(): %v", err)
		}
	}

	// Update the validator
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
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
		return fmt.Errorf("error unbondingValidatorsView.GetMatureEntries(): %v", err)
	}

	// Update validators status
	for _, entry := range matureEntries {
		validator, found, err := validatorsView.FindByOperatorAddr(entry.OperatorAddress, height)
		if err != nil {
			return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
		}
		if !found {
			return fmt.Errorf("error validator not found: %v, %v", entry.OperatorAddress, height)
		}

		validator.Status = types.UNBONDED

		if err := validatorsView.Update(validator); err != nil {
			return fmt.Errorf("error validatorsView.Update(): %v", err)
		}

	}

	// Remove those mature UnbondingValidator entry
	if err := unbondingValidatorsView.DeleteMatureEntries(blockTime); err != nil {
		return fmt.Errorf("error unbondingValidatorsView.DeleteMatureEntries(): %v", err)
	}

	return nil
}

// Find mature UnbindingDelegation through UBDQueue and then handle them one by one
func (projection *ValidatorDelegation) handleMatureUBDQueueEntries(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
) error {

	ubdQueueView := NewUBDQueue(rdbTxHandle)

	// Get all mature unbonding delegations from UBDQueue
	matureUnbonds, err := ubdQueueView.DequeueAllMatureUBDQueue(blockTime)
	if err != nil {
		return fmt.Errorf("error ubdQueueView.DequeueAllMatureUBDQueue(): %v", err)
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
		return fmt.Errorf("error redelegationQueueView.DequeueAllMatureRedelegationQueue(): %v", err)
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
