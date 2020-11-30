package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type TransactionsTotal struct {
	*view.Total
}

func NewTransactionsTotal(rdbHandle *rdb.Handle) *TransactionsTotal {
	return &TransactionsTotal{
		view.NewTotal(rdbHandle, "view_transactions_total"),
	}
}
