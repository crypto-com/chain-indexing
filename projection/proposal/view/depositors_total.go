package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DEPOSITORS_TOTAL_TABLE_NAME = "view_proposal_depositors_total"

type DepositorsTotal struct {
	*view.Total
}

func NewDepositorsTotal(rdbHandle *rdb.Handle) *DepositorsTotal {
	return &DepositorsTotal{
		view.NewTotal(rdbHandle, DEPOSITORS_TOTAL_TABLE_NAME),
	}
}
