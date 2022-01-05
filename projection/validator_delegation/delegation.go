package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) unbondAmountToShares(
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

	// Get the delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		return coin.ZeroDec(), fmt.Errorf("error delegation not found: %v, %v, %v", validatorAddress, delegatorAddress, height)
	}

	var shares coin.Dec
	shares, err = utils.ValidatorSharesFromTokens(validator, amount)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error utils.ValidatorSharesFromTokens(): %v", err)
	}
	// Cap the shares at the delegation's shares.
	if shares.GT(delegation.Shares) {
		shares = delegation.Shares
	}

	return shares, nil
}

func (projection *ValidatorDelegation) unbond(
	rdbTxHandle *rdb.Handle,
	height int64,
	validatorAddress string,
	delegatorAddress string,
	shares coin.Dec,
) (coin.Int, error) {

	validatorsView := NewValidators(rdbTxHandle)
	delegationsView := NewDelegations(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByOperatorAddr(validatorAddress, height)
	if err != nil {
		return coin.ZeroInt(), fmt.Errorf("error validatorsView.FindByOperatorAddr(): %v", err)
	}
	if !found {
		return coin.ZeroInt(), fmt.Errorf("error validator not found: %v, %v", validatorAddress, height)
	}

	// Get the delegation
	delegation, found, err := delegationsView.FindBy(
		validatorAddress,
		delegatorAddress,
		height,
	)
	if err != nil {
		return coin.ZeroInt(), fmt.Errorf("error delegationsView.FindBy(): %v", err)
	}
	if !found {
		return coin.ZeroInt(), fmt.Errorf("error delegation not found: %v, %v, %v", validatorAddress, delegatorAddress, height)
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
		return coin.ZeroInt(), fmt.Errorf("error in IsValAddrEqualsDelAddr: %v", err)
	}

	// Remaining tokens in the delegation
	delegationRemainingTokens := utils.ValidatorTokensFromShares(validator, delegation.Shares)

	// If the delegation belongs to operator of the validator and undelegating will decrease the validator's
	// self-delegation below their minimum, we jail the validator.
	if isOperatorDelegation && !validator.Jailed &&
		delegationRemainingTokens.TruncateInt().LT(validator.MinSelfDelegation) {

		if err := projection.jailValidator(rdbTxHandle, validator); err != nil {
			return coin.ZeroInt(), fmt.Errorf("error in projection.jailValidator(): %v", err)
		}
		validator.Jailed = true
	}

	// Update delegation
	if delegation.Shares.IsZero() {

		// Delete the DelegationRow
		if err := delegationsView.Delete(delegation); err != nil {
			return coin.ZeroInt(), fmt.Errorf("error delegationsView.Delete(): %v", err)
		}

	} else {

		if err := delegationsView.Update(delegation); err != nil {
			return coin.ZeroInt(), fmt.Errorf("error in delegationsView.Remove(): %v", err)
		}

	}

	// remove the shares and tokens from the validator
	var removedValTokens coin.Int
	validator, removedValTokens = utils.ValidatorRemoveDelShares(validator, shares)

	// Update validator
	if err := validatorsView.Update(validator); err != nil {
		return coin.ZeroInt(), fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	if validator.Shares.IsZero() && validator.IsUnbonded() {
		// if not unbonded, we must instead remove validator in EndBlocker once it finishes its unbonding period
		if err := validatorsView.Delete(validator); err != nil {
			return coin.ZeroInt(), fmt.Errorf("error validatorsView.Delete(): %v", err)
		}
	}

	return removedValTokens, nil
}
