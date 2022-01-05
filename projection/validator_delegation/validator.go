package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
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
