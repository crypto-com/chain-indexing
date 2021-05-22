package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DENOMS_TOTAL_TABLE_NAME = "view_nft_denoms_total"

type DenomsTotal struct {
	*view.Total
}

func NewDenomsTotal(rdbHandle *rdb.Handle) *DenomsTotal {
	return &DenomsTotal{
		view.NewTotal(rdbHandle, DENOMS_TOTAL_TABLE_NAME),
	}
}
