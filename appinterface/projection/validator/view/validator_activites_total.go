package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type ValidatorActivitiesTotal struct {
	*view.Total
}

func NewValidatorActivitiesTotal(rdbHandle *rdb.Handle) *ValidatorActivitiesTotal {
	return &ValidatorActivitiesTotal{
		view.NewTotal(rdbHandle, "view_validator_activities_total"),
	}
}
