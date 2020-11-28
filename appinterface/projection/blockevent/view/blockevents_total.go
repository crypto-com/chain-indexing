package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type BlockEventsTotal struct {
	*view.Total
}

func NewBlockEventsTotal(rdbHandle *rdb.Handle) *BlockEventsTotal {
	return &BlockEventsTotal{
		view.NewTotal(rdbHandle, "view_block_events_total"),
	}
}
