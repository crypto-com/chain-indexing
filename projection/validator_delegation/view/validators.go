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
	Insert(row ValidatorRow) error
	Update(row ValidatorRow) error
	Delete(row ValidatorRow) error
	FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, bool, error)
	FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, bool, error)
	FindByTendermintAddr(tendermintAddress string, height int64) (ValidatorRow, error)
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

func (view *ValidatorsView) Delete(row ValidatorRow) error {

	return nil
}

func (view *ValidatorsView) FindByOperatorAddr(operatorAddress string, height int64) (ValidatorRow, bool, error) {

	// TODO handle the error when validator is NOT FOUND

	return ValidatorRow{}, true, nil
}

func (view *ValidatorsView) FindByConsensusNodeAddr(consensusNodeAddress string, height int64) (ValidatorRow, bool, error) {

	// TODO handle the error when validator is NOT FOUND

	return ValidatorRow{}, true, nil
}

func (view *ValidatorsView) FindByTendermintAddr(tendermintAddress string, height int64) (ValidatorRow, error) {

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

// Notes on `Validator.Status`:
//
// - MsgCreateValidator, CreateGenesisValidator: Validator.Status = Unbonded
// - PowerChanged event:
//   - power > 0: Validator.Status = Bonded
//   - power = 0: Validator.Status = Unbonding, UnbondingValidator entry created
// - UnbondingValidator complete Unbonding period: Validator.Status = Unbonded

// TODO:
// - UNIQUE(height, operatorAddress)
// - UNIQUE(height, consensusNodeAddress)
type ValidatorRow struct {
	Height int64 `json:"height"`

	OperatorAddress      string `json:"operatorAddress"`
	ConsensusNodeAddress string `json:"consensusNodeAddress"`
	TendermintAddress    string `json:"tendermintAddress"`

	Status types.ValidatorStatus `json:"status"`
	Jailed bool                  `json:"jailed"`
	Power  string                `json:"power"`

	// `UnbondingHeight` and `UnbondingTime` only useful when `Status` is `Unbonding`
	// The height start the Unbonding
	UnbondingHeight int64 `json:"UnbondingHeight"`
	// The time when Unbonding is finished
	UnbondingTime utctime.UTCTime `json:"UnbondingTime"`

	Tokens            coin.Int `json:"tokens"`
	Shares            coin.Dec `json:"shares"`
	MinSelfDelegation coin.Int `json:"minSelfDelegation"`
}

func (v *ValidatorRow) IsBonded() bool {
	return v.Status == types.BONDED
}

func (v *ValidatorRow) IsUnbonded() bool {
	return v.Status == types.UNBONDED
}

func (v *ValidatorRow) IsUnbonding() bool {
	return v.Status == types.UNBONDING
}
