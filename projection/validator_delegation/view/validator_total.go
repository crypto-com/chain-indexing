package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// Keep track of the number of total validators in each height
type ValidatorTotal interface {
	Increment(identity string, total int64) error
}

type ValidatorTotalView struct {
	*view.Total
}

func NewValidatorTotalView(rdbHandle *rdb.Handle) ValidatorTotal {
	return &ValidatorTotalView{
		view.NewTotal(rdbHandle, "view_vd_validator_total"),
	}
}
