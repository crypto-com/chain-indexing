package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountRawEventsTotalView struct {
	*view.Total
}

func NewAccountRawEventsTotalView(rdbHandle *rdb.Handle) *AccountRawEventsTotalView {
	return &AccountRawEventsTotalView{
		view.NewTotal(rdbHandle, "view_account_raw_events_total"),
	}
}
