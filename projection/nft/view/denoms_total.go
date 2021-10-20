package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DENOMS_TOTAL_TABLE_NAME = "view_nft_denoms_total"

type DenomsTotal interface {
	Increment(identity string, total int64) error
	FindBy(identity string) (int64, error)
}

type DenomsTotalView struct {
	*view.Total
}

func NewDenomsTotalView(rdbHandle *rdb.Handle) DenomsTotal {
	return &DenomsTotalView{
		view.NewTotal(rdbHandle, DENOMS_TOTAL_TABLE_NAME),
	}
}
