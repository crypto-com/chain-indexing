package view

import (
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type UnbondingDelegations interface {
	Clone(previousHeight int64, currentHeight int64) error

	Upsert(row UnbondingDelegationRow) error
	Delete(row UnbondingDelegationRow) error
	FindBy(delegatorAddress, validatorAddress string, height int64) (UnbondingDelegationRow, bool, error)
	ListByValidator(validatorAddress string, height int64) ([]UnbondingDelegationRow, error)
}

type UnbondingDelegationsView struct {
	rdb *rdb.Handle
}

func NewUnbondingDelegationsView(handle *rdb.Handle) UnbondingDelegations {
	return &UnbondingDelegationsView{
		handle,
	}
}

func (view *UnbondingDelegationsView) Clone(previousHeight, currentHeight int64) error {

	return nil
}

func (view *UnbondingDelegationsView) Upsert(row UnbondingDelegationRow) error {

	return nil
}

func (view *UnbondingDelegationsView) Delete(row UnbondingDelegationRow) error {

	return nil
}

func (view *UnbondingDelegationsView) FindBy(
	delegatorAddress string,
	validatorAddress string,
	height int64,
) (UnbondingDelegationRow, bool, error) {

	return UnbondingDelegationRow{}, true, nil
}

func (view *UnbondingDelegationsView) ListByValidator(
	validatorAddress string,
	height int64,
) ([]UnbondingDelegationRow, error) {

	return nil, nil
}

// NOTES:
// - UNIQUE(delegatorAddress, validatorAddress, height)
// - INDEX(validatorAddress, height)
// - INDEX(delegatorAddress, height)
type UnbondingDelegationRow struct {
	Height           int64                      `json:"height"`
	DelegatorAddress string                     `json:"delegatorAddress"`
	ValidatorAddress string                     `json:"validatorAddress"`
	Entries          []UnbondingDelegationEntry `json:"entries"`
}

func NewUnbondingDelegationRow(
	delegatorAddress string,
	validatorAddress string,
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
) UnbondingDelegationRow {
	return UnbondingDelegationRow{
		DelegatorAddress: delegatorAddress,
		ValidatorAddress: validatorAddress,
		Entries: []UnbondingDelegationEntry{
			NewUnbondingDelegationEntry(creationHeight, completionTime, balance),
		},
	}
}

// AddEntry - append entry to the unbonding delegation
func (ubd *UnbondingDelegationRow) AddEntry(
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
) {
	entry := NewUnbondingDelegationEntry(creationHeight, completionTime, balance)
	ubd.Entries = append(ubd.Entries, entry)
}

// RemoveEntry - remove entry at index i to the unbonding delegation
func (ubd *UnbondingDelegationRow) RemoveEntry(i int64) {
	ubd.Entries = append(ubd.Entries[:i], ubd.Entries[i+1:]...)
}

type UnbondingDelegationEntry struct {
	CreationHeight int64           `json:"creationHeight"`
	CompletionTime utctime.UTCTime `json:"completionTime"`
	InitialBalance coin.Int        `json:"initialBalance"`
	Balance        coin.Int        `json:"balance"`
}

func NewUnbondingDelegationEntry(
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
) UnbondingDelegationEntry {
	return UnbondingDelegationEntry{
		CreationHeight: creationHeight,
		CompletionTime: completionTime,
		InitialBalance: balance,
		Balance:        balance,
	}
}

func (e *UnbondingDelegationEntry) IsMature(currentTime utctime.UTCTime) bool {
	complete := time.Unix(0, e.CompletionTime.UnixNano())
	current := time.Unix(0, currentTime.UnixNano())

	return complete.Before(current)
}
