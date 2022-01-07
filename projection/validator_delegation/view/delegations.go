package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Delegations interface {
	Clone(previousHeight int64, currentHeight int64) error

	Insert(row DelegationRow) error
	Update(row DelegationRow) error
	Delete(row DelegationRow) error

	FindBy(validatorAddress string, delegatorAddress string, height int64) (DelegationRow, bool, error)

	ListByValidatorAddr(
		validatorAddress string,
		height int64,
		pagination *pagination.Pagination,
	) ([]DelegationRow, *pagination.PaginationResult, error)
	ListByDelegatorAddr(
		delegatorAddress string,
		height int64,
		pagination *pagination.Pagination,
	) ([]DelegationRow, *pagination.PaginationResult, error)
}

type DelegationsView struct {
	rdb *rdb.Handle
}

func NewDelegationsView(handle *rdb.Handle) Delegations {
	return &DelegationsView{
		handle,
	}
}

func (view *DelegationsView) Clone(previousHeight, currentHeight int64) error {

	return nil
}

func (view *DelegationsView) Insert(row DelegationRow) error {

	return nil
}

func (view *DelegationsView) Update(row DelegationRow) error {

	return nil
}

func (view *DelegationsView) Delete(row DelegationRow) error {

	return nil
}

func (view *DelegationsView) FindBy(
	validatorAddress string,
	delegatorAddress string,
	height int64,
) (DelegationRow, bool, error) {

	// TODO need to handle the case when row NOT exist

	return DelegationRow{}, false, nil
}

func (view *DelegationsView) ListByValidatorAddr(
	validatorAddress string,
	height int64,
	pagination *pagination.Pagination,
) ([]DelegationRow, *pagination.PaginationResult, error) {

	return nil, nil, nil
}

func (view *DelegationsView) ListByDelegatorAddr(
	delegatorAddress string,
	height int64,
	pagination *pagination.Pagination,
) ([]DelegationRow, *pagination.PaginationResult, error) {

	return nil, nil, nil
}

// TODO:
// - UNIQUE(height, validatorAddress, delegatorAddress)
// - INDEX(height, validatorAddress)
// - INDEX(height, delegatorAddress)
type DelegationRow struct {
	Height           int64    `json:"height"`
	ValidatorAddress string   `json:"validatorAddress"`
	DelegatorAddress string   `json:"delegatorAddress"`
	Shares           coin.Dec `json:"shares"`
}
