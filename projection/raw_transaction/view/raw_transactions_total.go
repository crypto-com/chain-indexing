package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type RawTransactionsTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type RawTransactionsTotalView struct {
	*view.Total
}

func NewRawTransactionsTotalView(rdbHandle *rdb.Handle) RawTransactionsTotal {
	return &RawTransactionsTotalView{
		view.NewTotal(rdbHandle, "view_raw_transactions_total"),
	}
}
