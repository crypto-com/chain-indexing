package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountRawTransactionsTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type AccountRawTransactionsTotalView struct {
	*view.Total
}

func NewAccountRawTransactionsTotalView(rdbHandle *rdb.Handle) AccountRawTransactionsTotal {
	return &AccountRawTransactionsTotalView{
		view.NewTotal(rdbHandle, "view_account_raw_transactions_total"),
	}
}
