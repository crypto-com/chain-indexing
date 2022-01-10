package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

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
		projection.config.validatorAddressPrefix,
		projection.config.accountAddressPrefix,
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
