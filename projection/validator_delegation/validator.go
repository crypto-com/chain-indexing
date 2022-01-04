package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
)

func (projection *ValidatorDelegation) AddValidatorRecord(
	rdbTxHandle *rdb.Handle,
	validator view.ValidatorRow,
) error {
	validatorsView := NewValidators(rdbTxHandle)

	// Insert an ValidatorRow record
	if err := validatorsView.Insert(validator); err != nil {
		return fmt.Errorf("error validatorsView.Insert(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) RemoveValidatorRecord(
	rdbTxHandle *rdb.Handle,
	validator view.ValidatorRow,
) error {
	validatorsView := NewValidators(rdbTxHandle)

	// Delete an ValidatorRow record
	if err := validatorsView.Delete(validator); err != nil {
		return fmt.Errorf("error validatorsView.Delete(): %v", err)
	}

	return nil
}

func JailValidator(
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
