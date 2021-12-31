package validator_delegation

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
)

func AddDelegationRecord(
	rdbTxHandle *rdb.Handle,
	delegation view.DelegationRow,
) error {
	delegationsView := NewDelegations(rdbTxHandle)
	delegationTotalView := NewDelegationTotal(rdbTxHandle)

	// Insert an DelegationRow record
	if err := delegationsView.Insert(delegation); err != nil {
		return fmt.Errorf("error delegationsView.Insert(): %v", err)
	}

	// Increase number of total delegations of this validator at this height
	validatorIdentity := fmt.Sprintf("%s:%s", delegation.ValidatorAddress, strconv.FormatInt(delegation.Height, 10))
	if err := delegationTotalView.Increment(validatorIdentity, 1); err != nil {
		return fmt.Errorf("error in validatorTotalView.Increment() with validatorAddress: %v", err)
	}

	// Increase number of total delegations of this delegator at this height
	delegatorIdentity := fmt.Sprintf("%s:%s", delegation.DelegatorAddress, strconv.FormatInt(delegation.Height, 10))
	if err := delegationTotalView.Increment(delegatorIdentity, 1); err != nil {
		return fmt.Errorf("error in validatorTotalView.Increment() with delegatorAddress: %v", err)
	}

	return nil
}
