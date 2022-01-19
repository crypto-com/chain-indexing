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
		return coin.ZeroDec(), fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
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
		return coin.ZeroDec(), fmt.Errorf("error finding delegation: %v", err)
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
			return coin.ZeroDec(), fmt.Errorf("error inserting delegation record: %v", errAddDelRecord)
		}
	}

	// Calculate new shares
	newShares, err := utils.ValidatorIssueNewShares(validator, amount)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error in validator issuing new shares: %v", err)
	}

	// Update Delegation
	delegation.Shares = delegation.Shares.Add(newShares)
	if err := delegationsView.Update(delegation); err != nil {
		return coin.ZeroDec(), fmt.Errorf("error updating delegation: %v", err)
	}

	// Update Validator
	validator.Tokens = validator.Tokens.Add(amount)
	validator.Shares = validator.Shares.Add(newShares)
	if err := validatorsView.Update(validator); err != nil {
		return coin.ZeroDec(), fmt.Errorf("error updating validator: %v", err)
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
		return fmt.Errorf("error converting the unbond amout to corresponding shares on that validator: %v", err)
	}

	returnAmount, err := projection.unbond(
		rdbTxHandle,
		height,
		validatorAddress,
		delegatorAddress,
		shares,
	)
	if err != nil {
		return fmt.Errorf("error unbonding shares from delegation and validator: %v", err)
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
		return fmt.Errorf("error setting an unbonding delegation entry: %v", err)
	}

	if err := projection.insertUnbondingDelegationQueue(rdbTxHandle, ubd, completionTime); err != nil {
		return fmt.Errorf("error inserting an entry to unbonding delegation queue: %v", err)
	}

	return nil
}

// unbondAmountToShares - convert the unbond amout to shares on the target validator
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
		return coin.ZeroDec(), fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
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
		return coin.ZeroDec(), fmt.Errorf("error finding delegation: %v", err)
	}
	if !found {
		return coin.ZeroDec(), fmt.Errorf("error delegation not found: %v, %v, %v", validatorAddress, delegatorAddress, height)
	}

	var shares coin.Dec
	shares, err = utils.ValidatorSharesFromTokens(validator, amount)
	if err != nil {
		return coin.ZeroDec(), fmt.Errorf("error converting unbond amount to shares on that validator: %v", err)
	}
	// Cap the shares at the delegation's shares.
	if shares.GT(delegation.Shares) {
		shares = delegation.Shares
	}

	return shares, nil
}

// unbond - remove shares on both the delegation and validator
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
		return coin.ZeroInt(), fmt.Errorf("error finding validator by OperatorAddress (AKA ValidatorAddress): %v", err)
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
		return coin.ZeroInt(), fmt.Errorf("error finding delegation: %v", err)
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
		projection.config.ValidatorAddressPrefix,
		projection.config.AccountAddressPrefix,
	)
	if err != nil {
		return coin.ZeroInt(), fmt.Errorf("error in checking if ValidatorAddr and DelegatorAddr are the same address: %v", err)
	}

	// Remaining tokens in the delegation
	delegationRemainingTokens := utils.ValidatorTokensFromShares(validator, delegation.Shares)

	// If the delegation belongs to operator of the validator and undelegating will decrease the validator's
	// self-delegation below their minimum, we jail the validator.
	if isOperatorDelegation && !validator.Jailed &&
		delegationRemainingTokens.TruncateInt().LT(validator.MinSelfDelegation) {

		if errJailVal := projection.jailValidator(rdbTxHandle, validator); errJailVal != nil {
			return coin.ZeroInt(), fmt.Errorf("error jailing the validator: %v", errJailVal)
		}
		validator.Jailed = true
	}

	// Update delegation
	if delegation.Shares.IsZero() {

		// Delete the DelegationRow
		if errDeleteDel := delegationsView.Delete(delegation); errDeleteDel != nil {
			return coin.ZeroInt(), fmt.Errorf("error deleting delegation: %v", errDeleteDel)
		}

	} else {

		if errUpdateDel := delegationsView.Update(delegation); errUpdateDel != nil {
			return coin.ZeroInt(), fmt.Errorf("error in updating delegation: %v", errUpdateDel)
		}

	}

	// remove the shares and tokens from the validator
	var removedValTokens coin.Int
	validator, removedValTokens, err = utils.ValidatorRemoveDelShares(validator, shares)
	if err != nil {
		return coin.ZeroInt(), fmt.Errorf("error in removing shares and tokens on validator: %v", err)
	}

	// Update validator
	if err := validatorsView.Update(validator); err != nil {
		return coin.ZeroInt(), fmt.Errorf("error updating validator: %v", err)
	}

	if validator.Shares.IsZero() && validator.IsUnbonded() {
		// if not unbonded, we must instead remove validator in EndBlocker once it finishes its unbonding period
		if err := validatorsView.Delete(validator); err != nil {
			return coin.ZeroInt(), fmt.Errorf("error deleting validator: %v", err)
		}
	}

	return removedValTokens, nil
}
