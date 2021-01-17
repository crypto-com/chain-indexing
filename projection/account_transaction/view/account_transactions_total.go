package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountTransactionsTotal struct {
	*view.Total
}

func NewAccountTransactionsTotal(rdbHandle *rdb.Handle) *AccountTransactionsTotal {
	return &AccountTransactionsTotal{
		view.NewTotal(rdbHandle, "view_account_transactions_total"),
	}
}

func (total *AccountTransactionsTotal) Search(address string) (bool, error) {
	_, err := total.FindBy(fmt.Sprintf("%s:-", address))
	if err != nil {
		return false, err
	}
	return true, nil
}
