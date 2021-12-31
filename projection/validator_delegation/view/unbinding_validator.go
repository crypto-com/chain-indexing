package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UnbindingValidators interface {
	Insert(operatorAddress string, unbindingTime utctime.UTCTime) error
	RemoveIfExist(operatorAddress string) error
	GetMatureEntries(blockTime utctime.UTCTime) ([]UnbindingValidatorRow, error)
	DeleteMatureEntries(blockTime utctime.UTCTime) error
}

type UnbindingValidatorsView struct {
	rdb *rdb.Handle
}

func NewUnbindingValidatorsView(handle *rdb.Handle) UnbindingValidators {
	return &UnbindingValidatorsView{
		handle,
	}
}

func (view *UnbindingValidatorsView) Insert(operatorAddress string, unbindingTime utctime.UTCTime) error {

	return nil
}

func (view *UnbindingValidatorsView) RemoveIfExist(operatorAddress string) error {

	return nil
}

func (view *UnbindingValidatorsView) GetMatureEntries(blockTime utctime.UTCTime) ([]UnbindingValidatorRow, error) {

	return nil, nil
}

func (view *UnbindingValidatorsView) DeleteMatureEntries(blockTime utctime.UTCTime) error {

	return nil
}

type UnbindingValidatorRow struct {
	OperatorAddress string          `json:"operatorAddress"`
	UnbindingTime   utctime.UTCTime `json:"unbindingTime"`
}
