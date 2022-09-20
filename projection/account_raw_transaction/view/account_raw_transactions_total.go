package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountRawTransactionsTotal struct {
	*view.Total
}

func NewAccountRawTransactionsTotal(rdbHandle *rdb.Handle) *AccountRawTransactionsTotal {
	return &AccountRawTransactionsTotal{
		view.NewTotal(rdbHandle, "view_account_raw_transactions_total"),
	}
}

func (total *AccountRawTransactionsTotal) Search(address string) (bool, error) {
	numberOfRowsFound, err := total.FindBy(fmt.Sprintf("%s:-", address))
	if err != nil {
		return false, err
	}
	if numberOfRowsFound == 0 {
		return false, nil
	}
	return true, nil
}
