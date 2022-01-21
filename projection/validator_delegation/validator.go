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

func (projection *ValidatorDelegation) handleGenesisCreateNewValidator(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	tendermintPubKey string,
	amount coin.Int,
	minSelfDelegationInString string,
) error {

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.config.ConNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error in GetConsensusNodeAddress: %v", err)
	}

	tendermintAddress, err := utils.GetTendermintAddress(tendermintPubKey)
	if err != nil {
		return fmt.Errorf("error in GetTendermintAddress: %v", err)
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
		TendermintAddress:    tendermintAddress,

		// For MsgCreateValidator in genesis block, we assumed they are always BONDED.
		// Here, as genesis has no block_results and power_changed event, we just set the power as `1`,
		// indicating that's a bonded validator.
		//
		// TODO: What if gen_txs contains more validators than maximum validators?
		//       In that case, may consider parsing `max_validators` from genesis.
		//       And only marked top `max_validators` validators as bonded.
		Status: types.BONDED,
		Jailed: false,
		Power:  "1",

		UnbondingHeight: int64(0),
		UnbondingTime:   utctime.UTCTime{},

		Tokens:            coin.ZeroInt(),
		Shares:            coin.ZeroDec(),
		MinSelfDelegation: minSelfDelegation,
	}
	if err := validatorsView.Insert(validatorRow); err != nil {
		return fmt.Errorf("error inserting ValidatorRow: %v", err)
	}

	if _, err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	); err != nil {
		return fmt.Errorf("error handling Delegate: %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleCreateNewValidator(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	tendermintPubKey string,
	amount coin.Int,
	minSelfDelegationInString string,
) error {

	consensusNodeAddress, err := utils.GetConsensusNodeAddress(tendermintPubKey, projection.config.ConNodeAddressPrefix)
	if err != nil {
		return fmt.Errorf("error in GetConsensusNodeAddress: %v", err)
	}

	tendermintAddress, err := utils.GetTendermintAddress(tendermintPubKey)
	if err != nil {
		return fmt.Errorf("error in GetTendermintAddress: %v", err)
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
		TendermintAddress:    tendermintAddress,

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
		return fmt.Errorf("error inserting ValidatorRow: %v", err)
	}

	if _, err := projection.handleDelegate(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		amount,
	); err != nil {
		return fmt.Errorf("error handling Delegate: %v", err)
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
		return fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	minSelfDelegation, ok := coin.NewIntFromString(minSelfDelegationInString)
	if !ok {
		return fmt.Errorf("Failed to parse minSelfDelegationInString: %v", minSelfDelegationInString)
	}
	// Update Validator
	validator.MinSelfDelegation = minSelfDelegation

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error updating validator: %v", err)
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
	validator, found, err := validatorsView.FindByConsensusNodeAddr(
		consensusNodeAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error finding validator by consensusNodeAddr: %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found, conNodeAddr: %v", consensusNodeAddress)
	}

	if err := projection.jailValidator(rdbTxHandle, validator); err != nil {
		return fmt.Errorf("error in jailing the validator: %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) jailValidator(
	rdbTxHandle *rdb.Handle,
	validator view.ValidatorRow,
) error {
	validatorsView := NewValidators(rdbTxHandle)

	// Update Validator
	validator.Jailed = true

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error updating validator: %v", err)
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
		return fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	// Update Validator
	validator.Jailed = false

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error updating validator: %v", err)
	}

	return nil
}

// removeValidatorTokens - Deduct from validator's bonded tokens and update the validator.
func (projection *ValidatorDelegation) removeValidatorTokens(
	rdbTxHandle *rdb.Handle,
	validator view.ValidatorRow,
	tokensToRemove coin.Int,
) error {
	validatorsView := NewValidators(rdbTxHandle)

	if tokensToRemove.IsNegative() {
		return fmt.Errorf("should not happen: trying to remove negative tokens %v", tokensToRemove)
	}

	if validator.Tokens.LT(tokensToRemove) {
		return fmt.Errorf("should not happen: only have %v tokens, trying to remove %v", validator.Tokens, tokensToRemove)
	}

	validator.Tokens = validator.Tokens.Sub(tokensToRemove)

	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error updating validator: %v", err)
	}

	return nil
}
