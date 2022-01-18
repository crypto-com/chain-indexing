package view

import (
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Redelegations interface {
	Clone(previousHeight int64, currentHeight int64) error

	Upsert(row RedelegationRow) error
	Delete(row RedelegationRow) error
	FindBy(delegatorAddress, validatorSrcAddress, validatorDstAddress string, height int64) (RedelegationRow, bool, error)
	ListBySrcValidator(validatorSrcAddress string, height int64) ([]RedelegationRow, error)
}

type RedelegationsView struct {
	rdb *rdb.Handle
}

func NewRedelegationsView(handle *rdb.Handle) Redelegations {
	return &RedelegationsView{
		handle,
	}
}

func (view *RedelegationsView) Clone(previousHeight, currentHeight int64) error {

	return nil
}

func (view *RedelegationsView) Upsert(row RedelegationRow) error {

	return nil
}

func (view *RedelegationsView) Delete(row RedelegationRow) error {

	return nil
}

func (view *RedelegationsView) FindBy(
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
	height int64,
) (RedelegationRow, bool, error) {

	// TODO Handle the error when row is NOT FOUND

	return RedelegationRow{}, true, nil
}

func (view *RedelegationsView) ListBySrcValidator(
	validatorSrcAddress string,
	height int64,
) ([]RedelegationRow, error) {

	return nil, nil
}

// Notes:
// - UNIQUE(delegatorAddress, validatorSrcAddress, validatorDstAddress, height)
// - Index(validatorSrcAddress, height)
type RedelegationRow struct {
	Height              int64               `json:"height"`
	DelegatorAddress    string              `json:"delegatorAddress"`
	ValidatorSrcAddress string              `json:"validatorSrcAddress"`
	ValidatorDstAddress string              `json:"validatorDstAddress"`
	Entries             []RedelegationEntry `json:"entries"`
}

func NewRedelegationRow(
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
	sharesDst coin.Dec,
) RedelegationRow {
	return RedelegationRow{
		DelegatorAddress:    delegatorAddress,
		ValidatorSrcAddress: validatorSrcAddress,
		ValidatorDstAddress: validatorSrcAddress,
		Entries: []RedelegationEntry{
			NewRedelegationEntry(creationHeight, completionTime, balance, sharesDst),
		},
	}
}

// AddEntry - append entry to the redelegation
func (red *RedelegationRow) AddEntry(
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
	sharesDst coin.Dec,
) {
	entry := NewRedelegationEntry(creationHeight, completionTime, balance, sharesDst)
	red.Entries = append(red.Entries, entry)
}

// RemoveEntry - remove entry at index i to the unbonding delegation
func (red *RedelegationRow) RemoveEntry(i int64) {
	red.Entries = append(red.Entries[:i], red.Entries[i+1:]...)
}

type RedelegationEntry struct {
	CreationHeight int64           `json:"creationHeight"`
	CompletionTime utctime.UTCTime `json:"completionTime"`
	InitialBalance coin.Int        `json:"initialBalance"`
	SharesDst      coin.Dec        `json:"sharesDst"`
}

func NewRedelegationEntry(
	creationHeight int64,
	completionTime utctime.UTCTime,
	balance coin.Int,
	sharesDst coin.Dec,
) RedelegationEntry {
	return RedelegationEntry{
		CreationHeight: creationHeight,
		CompletionTime: completionTime,
		InitialBalance: balance,
		SharesDst:      sharesDst,
	}
}

func (e *RedelegationEntry) IsMature(currentTime utctime.UTCTime) bool {
	complete := time.Unix(0, e.CompletionTime.UnixNano())
	current := time.Unix(0, currentTime.UnixNano())

	return complete.Before(current)
}
