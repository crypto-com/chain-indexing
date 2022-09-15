package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountRawEventsTotal struct {
	*view.Total
}

func NewAccountRawEventsTotal(rdbHandle *rdb.Handle) *AccountRawEventsTotal {
	return &AccountRawEventsTotal{
		view.NewTotal(rdbHandle, "view_account_raw_events_total"),
	}
}
