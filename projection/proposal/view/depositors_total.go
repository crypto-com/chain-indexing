package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DEPOSITORS_TOTAL_TABLE_NAME = "view_proposal_depositors_total"

type DepositorsTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type DepositorsTotalView struct {
	*view.Total
}

func NewDepositorsTotalView(rdbHandle *rdb.Handle) DepositorsTotal {
	return &DepositorsTotalView{
		view.NewTotal(rdbHandle, DEPOSITORS_TOTAL_TABLE_NAME),
	}
}
