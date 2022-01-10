package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func (projection *ValidatorDelegation) jailValidator(
	rdbTxHandle *rdb.Handle,
	validator view.ValidatorRow,
) error {
	validatorsView := NewValidators(rdbTxHandle)

	// Update Valdiator
	validator.Jailed = true

	// Write the updated Validator to DB
	if err := validatorsView.Update(validator); err != nil {
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}

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
		return fmt.Errorf("error validatorsView.Update(): %v", err)
	}

	return nil
}
