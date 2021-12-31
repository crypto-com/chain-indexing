package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// Keep track of validator's total number of delegations in each height
type DelegationTotal interface {
	Increment(identity string, total int64) error
}

type DelegationTotalView struct {
	*view.Total
}

func NewDelegationTotalView(rdbHandle *rdb.Handle) DelegationTotal {
	return &DelegationTotalView{
		view.NewTotal(rdbHandle, "view_vd_delegation_total"),
	}
}
