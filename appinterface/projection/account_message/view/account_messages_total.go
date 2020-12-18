package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountMessagesTotal struct {
	*view.Total
}

func NewAccountMessagesTotal(rdbHandle *rdb.Handle) *AccountMessagesTotal {
	return &AccountMessagesTotal{
		view.NewTotal(rdbHandle, "view_account_message_ids_total"),
	}
}
