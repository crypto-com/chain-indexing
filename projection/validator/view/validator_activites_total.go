package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const VALIDATOR_ACTIVITIES_TABLE_NAME = "view_validator_activities_total"

type ValidatorActivitiesTotal interface {
	Set(string, int64) error
	Increment(string, int64) error
	IncrementAll([]string, int64) error
	DecrementAll([]string, int64) error
	FindBy(string) (int64, error)
	SumBy([]string) (int64, error)
}

type ValidatorActivitiesTotalView struct {
	*view.Total
}

func NewValidatorActivitiesTotalView(rdbHandle *rdb.Handle) ValidatorActivitiesTotal {
	return &ValidatorActivitiesTotalView{
		view.NewTotal(rdbHandle, VALIDATOR_ACTIVITIES_TABLE_NAME),
	}
}
