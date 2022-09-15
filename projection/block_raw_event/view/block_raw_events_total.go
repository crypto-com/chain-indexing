package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type BlockRawEventsTotal struct {
	*view.Total
}

func NewBlockRawEventsTotal(rdbHandle *rdb.Handle) *BlockRawEventsTotal {
	return &BlockRawEventsTotal{
		view.NewTotal(rdbHandle, "view_block_raw_events_total"),
	}
}
