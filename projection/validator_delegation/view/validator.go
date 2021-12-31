package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Validators interface {
	Clone(previousHeight int64, currentHeight int64) error
	Insert(ValidatorRow) error
	Update(ValidatorRow) error
	FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, error)
	FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, error)
	List(height int64, pagination *pagination.Pagination) ([]ValidatorRow, *pagination.PaginationResult, error)
}

type ValidatorsView struct {
	rdb *rdb.Handle
}

func NewValidatorsView(handle *rdb.Handle) Validators {
	return &ValidatorsView{
		handle,
	}
}

func (view *ValidatorsView) Clone(previousHeight, currentHeight int64) error {

	return nil
}

func (view *ValidatorsView) Insert(row ValidatorRow) error {

	return nil
}

func (view *ValidatorsView) Update(row ValidatorRow) error {

	return nil
}

func (view *ValidatorsView) FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, error) {

	return ValidatorRow{}, nil
}

func (view *ValidatorsView) FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, error) {

	return ValidatorRow{}, nil
}

func (view *ValidatorsView) List(
	height int64,
	pagination *pagination.Pagination,
) (
	[]ValidatorRow,
	*pagination.PaginationResult,
	error,
) {

	return nil, nil, nil
}

type ValidatorRow struct {
	Height int64 `json:"height"`

	OperatorAddress      string `json:"operatorAddress"`
	ConsensusNodeAddress string `json:"consensusNodeAddress"`

	Status types.ValidatorStatus `json:"status"`
	Jailed bool                  `json:"jailed"`
	Power  string                `json:"power"`

	// `UnbindingHeight` and `UnbindingTime` only useful when `Status` is `unbinding`
	// The height start the unbinding
	UnbindingHeight int64 `json:"unbindingHeight"`
	// The time when unbinding is finished
	UnbindingTime utctime.UTCTime `json:"unbindingTime"`

	Tokens coin.Int `json:"tokens"`
	Shares coin.Dec `json:"shares"`
}
