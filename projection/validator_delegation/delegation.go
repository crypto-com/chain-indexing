package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) AddDelegationRecord(
	rdbTxHandle *rdb.Handle,
	delegation view.DelegationRow,
) error {
	delegationsView := NewDelegations(rdbTxHandle)

	// Insert an DelegationRow record
	if err := delegationsView.Insert(delegation); err != nil {
		return fmt.Errorf("error delegationsView.Insert(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) RemoveDelegationRecord(
	rdbTxHandle *rdb.Handle,
	delegation view.DelegationRow,
) error {
	delegationsView := NewDelegations(rdbTxHandle)

	// Delete an DelegationRow record
	if err := delegationsView.Delete(delegation); err != nil {
		return fmt.Errorf("error delegationsView.Delete(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) UnbondAmountToShares(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	amount coin.Int,
) (coin.Dec, error) {
	var shares coin.Dec

	validatorsView := NewValidators(rdbTxHandle)
	delegationsView := NewDelegations(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByOperatorAddr(
		validatorAddress,
		height,
	)
	if err != nil {
		return shares, fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}

	// Get the delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return shares, fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		return shares, fmt.Errorf("error delegation not found: %v, %v, %v", validatorAddress, delegatorAddress, height)
	}

	shares, err = utils.ValidatorSharesFromTokens(validator, amount)
	if err != nil {
		return shares, fmt.Errorf("error utils.ValidatorSharesFromTokens(): %v", err)
	}
	// Cap the shares at the delegation's shares.
	if shares.GT(delegation.Shares) {
		shares = delegation.Shares
	}

	return shares, nil
}

func (projection *ValidatorDelegation) Unbond(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	shares coin.Dec,
) (coin.Int, error) {
	var removedValTokens coin.Int

	validatorsView := NewValidators(rdbTxHandle)
	delegationsView := NewDelegations(rdbTxHandle)

	// Get the validator
	validator, err := validatorsView.FindByOperatorAddr(
		validatorAddress,
		height,
	)
	if err != nil {
		return removedValTokens, fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}

	// Get the delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return removedValTokens, fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		return removedValTokens, fmt.Errorf("error delegation not found: %v, %v, %v", validatorAddress, delegatorAddress, height)
	}

	// Subtract shares from delegation
	delegation.Shares = delegation.Shares.Sub(shares)

	// Check if this delegation is the validator self delegation
	isOperatorDelegation, err := utils.IsValAddrEqualsDelAddr(
		validatorAddress,
		delegatorAddress,
		projection.validatorAddressPrefix,
		projection.accountAddressPrefix,
	)
	if err != nil {
		return removedValTokens, fmt.Errorf("error in IsValAddrEqualsDelAddr: %v", err)
	}

	// Remaining tokens in the delegation
	delegationRemainingTokens := utils.ValidatorTokensFromShares(validator, delegation.Shares)

	// If the delegation belongs to operator of the validator and undelegating will decrease the validator's
	// self-delegation below their minimum, we jail the validator.
	if isOperatorDelegation && !validator.Jailed &&
		delegationRemainingTokens.TruncateInt().LT(validator.MinSelfDelegation) {

		if err := JailValidator(rdbTxHandle, validator); err != nil {
			return removedValTokens, fmt.Errorf("error in JailValidator(): %v", err)
		}
		validator.Jailed = true
	}

	// Update delegation
	if delegation.Shares.IsZero() {

		if err := projection.RemoveDelegationRecord(rdbTxHandle, delegation); err != nil {
			return removedValTokens, fmt.Errorf("error in projection.RemoveDelegationRecord(): %v", err)
		}

	} else {

		if err := delegationsView.Update(delegation); err != nil {
			return removedValTokens, fmt.Errorf("error in delegationsView.Remove(): %v", err)
		}

	}

	// remove the shares and tokens from the validator
	validator, removedValTokens = utils.ValidatorRemoveDelShares(validator, shares)

	// Update validator
	if err := validatorsView.Update(validator); err != nil {
		return removedValTokens, fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	if validator.Shares.IsZero() && validator.Status == types.UNBONDED {
		// if not unbonded, we must instead remove validator in EndBlocker once it finishes its unbonding period
		if err := projection.RemoveValidatorRecord(rdbTxHandle, validator); err != nil {
			return removedValTokens, fmt.Errorf("error in projection.RemoveValidatorRecord(): %v", err)
		}
	}

	return removedValTokens, nil
}
