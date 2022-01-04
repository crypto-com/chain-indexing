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
	if err := projection.AddValidatorRecord(rdbTxHandle, validatorRow); err != nil {
		return fmt.Errorf("error in projection.AddValidatorRecord(): %v", err)
	}

	if err := projection.handleDelegate(
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
	validator, err := validatorsView.FindByOperatorAddr(
		validatorAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
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

func (projection *ValidatorDelegation) handleDelegate(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	amount coin.Int,
) error {

	validatorsView := NewValidators(rdbTxHandle)
	delegationsView := NewDelegations(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByOperatorAddr(
		validatorAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}

	// Create/Get Delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		delegation = view.DelegationRow{
			Height:           height,
			ValidatorAddress: validatorAddress,
			DelegatorAddress: delegatorAddress,
			Shares:           coin.ZeroDec(),
		}

		if errAddDelRecord := projection.AddDelegationRecord(rdbTxHandle, delegation); errAddDelRecord != nil {
			return fmt.Errorf("error projection.AddDelegationRecord(): %v", errAddDelRecord)
		}
	}

	// Calculate new shares
	newShares, err := utils.ValidatorIssueNewShares(validator, amount)
	if err != nil {
		return fmt.Errorf("error IssueNewShares(): %v", err)
	}

	// Update Delegation
	delegation.Shares = delegation.Shares.Add(newShares)
	if err := delegationsView.Update(delegation); err != nil {
		return fmt.Errorf("error delegationsView.Update(): %v", err)
	}

	// Update Validator
	validator.Tokens = validator.Tokens.Add(amount)
	validator.Shares = validator.Shares.Add(newShares)
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleUndelegate(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	amount coin.Int,
	completionTime utctime.UTCTime,
) error {

	shares, err := projection.UnbondAmountToShares(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	)
	if err != nil {
		return fmt.Errorf("error projection.UnbondAmountToShares(): %v", err)
	}

	returnAmount, err := projection.Unbond(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		shares,
	)
	if err != nil {
		return fmt.Errorf("error projection.Unbond(): %v", err)
	}

	ubd, err := projection.SetUnbondingDelegationEntry(
		rdbTxHandle,
		delegatorAddress,
		validatorAddress,
		height,
		completionTime,
		returnAmount,
	)
	if err != nil {
		return fmt.Errorf("error projection.SetUnbondingDelegationEntry(): %v", err)
	}

	if err := projection.InsertUBDQueue(rdbTxHandle, ubd, completionTime); err != nil {
		return fmt.Errorf("error projection.InsertUBDQueue(): %v", err)
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

	if err := JailValidator(rdbTxHandle, validator); err != nil {
		return fmt.Errorf("error in JailValidator(): %v", err)
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
	validator, err := validatorsView.FindByOperatorAddr(
		validatorAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
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
		validator, err := validatorsView.FindByOperatorAddr(
			entry.OperatorAddress,
			height,
		)
		if err != nil {
			return fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
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

	// Get all mature DVPair from UBDQueue
	matureEntries, err := ubdQueueView.GetMatureEntries(blockTime)
	if err != nil {
		return fmt.Errorf("error ubdQueueView.GetMatureEntries(): %v", err)
	}

	for _, entry := range matureEntries {

		if err := projection.CompleteUnbonding(
			rdbTxHandle,
			height,
			blockTime,
			entry.DelegatorAddress,
			entry.ValidatorAddress,
		); err != nil {
			// Not 100% sure why the error is ignored here.
			// This is probably due to a same DVPair could show up in matureEntries for multiple time.
			// reference: https://github.com/cosmos/cosmos-sdk/blob/59793e68fd7227c0ea31cf9456822dd566db4da7/x/staking/keeper/val_state_change.go#L46
			continue
		}

	}

	// Remove mature UBDQueue entries
	if err := ubdQueueView.DeleteMatureEntries(blockTime); err != nil {
		return fmt.Errorf("error ubdQueueView.DeleteMatureEntries(): %v", err)
	}

	return nil
}
