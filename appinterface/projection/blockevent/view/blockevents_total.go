package view

import (
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
)

type BlockEventsTotal struct {
	*view.Total
}

func NewBlockEventsTotal(rdbHandle *rdb.Handle) *BlockEventsTotal {
	return &BlockEventsTotal{
		view.NewTotal(rdbHandle, "view_block_events_total"),
	}
}
