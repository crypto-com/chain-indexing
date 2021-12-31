package validator_delegation

import (
	"fmt"
	"strconv"

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
) error {

	validatorsView := NewValidators(rdbTxHandle)
	validatorTotalView := NewValidatorTotal(rdbTxHandle)

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.conNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error in GetConsensusNodeAddress: %v", err)
	}

	// Insert an ValidatorRow.
	validatorRow := view.ValidatorRow{
		Height: height,

		OperatorAddress:      validatorAddress,
		ConsensusNodeAddress: consensusNodeAddress,

		Status: types.Unbonded,
		Jailed: false,
		Power:  "0",

		UnbindingHeight: int64(0),
		UnbindingTime:   utctime.UTCTime{},

		Tokens: coin.ZeroInt(),
		Shares: coin.ZeroDec(),
	}
	if err := validatorsView.Insert(validatorRow); err != nil {
		return fmt.Errorf("error in validatorsView.Insert(): %v", err)
	}

	// Increase number of total validators at this height
	identity := strconv.FormatInt(height, 10)
	if err := validatorTotalView.Increment(identity, 1); err != nil {
		return fmt.Errorf("error in validatorTotalView.Increment(): %v", err)
	}

	if err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	); err != nil {
		return fmt.Errorf("error in handleDelegate(): %v", err)
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
		return fmt.Errorf("error validatorsView.FindByConNodeAddress(): %v", err)
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

		if err := AddDelegationRecord(
			rdbTxHandle,
			delegation,
		); err != nil {
			return fmt.Errorf("error CreateDelegation(): %v", err)
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
		return fmt.Errorf("error validatorsView.FindByConNodeAddress(): %v", err)
	}

	// Update Valdiator
	validator.Jailed = true

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
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
	unbindingValidatorsView := NewUnbindingValidators(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByConsensusNodeAddr(
		consensusNodeAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByConNodeAddress(): %v", err)
	}

	if power == "0" {
		validator.Status = types.Unbonding

		// UnbindingHeight is the height when unbinding start
		// UnbindingTime is the time when unbinding is finished
		validator.UnbindingHeight = height
		validator.UnbindingTime = blockTime.Add(projection.unbindingTime)

		// Insert the validator to UnbindingValidators set
		if err := unbindingValidatorsView.Insert(validator.OperatorAddress, validator.UnbindingTime); err != nil {
			return fmt.Errorf("error unbindingValidatorsView.Insert(): %v", err)
		}

	} else {
		// In this case, power should be larger than 0.

		validator.Status = types.Bonded

		// Remove the validator from UnbindingValidators set, if exist
		if err := unbindingValidatorsView.RemoveIfExist(validator.OperatorAddress); err != nil {
			return fmt.Errorf("error unbindingValidatorsView.Insert(): %v", err)
		}
	}

	// Update the validator
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) unbindAllMatureUnbindingValidators(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
) error {

	validatorsView := NewValidators(rdbTxHandle)
	unbindingValidatorsView := NewUnbindingValidators(rdbTxHandle)

	// Get all UnbindingValidator entry that finish the unbinding period
	matureEntries, err := unbindingValidatorsView.GetMatureEntries(blockTime)
	if err != nil {
		return fmt.Errorf("error unbindingValidatorsView.GetMatureEntries(): %v", err)
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

		validator.Status = types.Unbonded

		if err := validatorsView.Update(validator); err != nil {
			return fmt.Errorf("error validatorsView.Update(): %v", err)
		}

	}

	// Remove those mature UnbindingValidator entry
	if err := unbindingValidatorsView.DeleteMatureEntries(blockTime); err != nil {
		return fmt.Errorf("error unbindingValidatorsView.DeleteMatureEntries(): %v", err)
	}

	return nil
}
