package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type RawBlockEventsTotal struct {
	*view.Total
}

func NewRawBlockEventsTotal(rdbHandle *rdb.Handle) *RawBlockEventsTotal {
	return &RawBlockEventsTotal{
		view.NewTotal(rdbHandle, "view_raw_block_events_total"),
	}
}
