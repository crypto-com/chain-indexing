package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type ValidatorBlockCommitmentsTotal struct {
	*view.Total
}

func NewValidatorBlockCommitmentsTotal(rdbHandle *rdb.Handle) *ValidatorBlockCommitmentsTotal {
	return &ValidatorBlockCommitmentsTotal{
		view.NewTotal(rdbHandle, "view_validator_block_commitments_total"),
	}
}
