package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
)

type UnbondingValidators interface {
	Insert(operatorAddress string, UnbondingTime utctime.UTCTime) error
	RemoveIfExist(operatorAddress string) error
	GetMatureEntries(blockTime utctime.UTCTime) ([]UnbondingValidatorRow, error)
	DeleteMatureEntries(blockTime utctime.UTCTime) error
}

type UnbondingValidatorsView struct {
	rdb *rdb.Handle
}

func NewUnbondingValidatorsView(handle *rdb.Handle) UnbondingValidators {
	return &UnbondingValidatorsView{
		handle,
	}
}

func (view *UnbondingValidatorsView) Insert(operatorAddress string, UnbondingTime utctime.UTCTime) error {

	return nil
}

func (view *UnbondingValidatorsView) RemoveIfExist(operatorAddress string) error {
	// First to check if the row exist, then remove it

	return nil
}

func (view *UnbondingValidatorsView) GetMatureEntries(blockTime utctime.UTCTime) ([]UnbondingValidatorRow, error) {

	return nil, nil
}

func (view *UnbondingValidatorsView) DeleteMatureEntries(blockTime utctime.UTCTime) error {

	return nil
}

// TODO: UNIQUE OperatorAddress
type UnbondingValidatorRow struct {
	OperatorAddress string          `json:"operatorAddress"`
	UnbondingTime   utctime.UTCTime `json:"UnbondingTime"`
}
