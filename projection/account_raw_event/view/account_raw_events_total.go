package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type AccountRawEventsTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
}

type AccountRawEventsTotalView struct {
	*view.Total
}

func NewAccountRawEventsTotalView(rdbHandle *rdb.Handle) *AccountRawEventsTotalView {
	return &AccountRawEventsTotalView{
		view.NewTotal(rdbHandle, "view_account_raw_events_total"),
	}
}
