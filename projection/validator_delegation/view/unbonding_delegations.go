package view

import (
	"errors"
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	jsoniter "github.com/json-iterator/go"
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

	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`INSERT INTO view_vd_unbonding_delegations(
		height,
		delegator_address,
		validator_address,
		entries
	) SELECT
		?,
		delegator_address,
		validator_address,
		entries
	FROM view_vd_unbonding_delegations WHERE height = ?
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building UnbondingDelegation clone sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		currentHeight,
		previousHeight,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error cloning UnbondingDelegation into the table: %v: %w", execErr, rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingDelegationsView) Upsert(row UnbondingDelegationRow) error {

	entriesJSON, err := jsoniter.MarshalToString(row.Entries)
	if err != nil {
		return fmt.Errorf("error JSON marshalling UnbondingDelegationRow.Entries for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_unbonding_delegations",
		).
		Columns(
			"height",
			"delegator_address",
			"validator_address",
			"entries",
		).
		Values(
			row.Height,
			row.DelegatorAddress,
			row.ValidatorAddress,
			entriesJSON,
		).
		Suffix("ON CONFLICT(delegator_address, validator_address, height) DO UPDATE SET entries = EXCLUDED.entries").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building UnbondingDelegation upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting UnbondingDelegation into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting UnbondingDelegation into the table: no row upserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingDelegationsView) Delete(row UnbondingDelegationRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_unbonding_delegations",
		).
		Where(
			"delegator_address = ? AND validator_address = ? AND height = ?",
			row.DelegatorAddress,
			row.ValidatorAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building UnbondingDelegation deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting UnbondingDelegation from the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error deleting UnbondingDelegation from the table: no row deleted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingDelegationsView) FindBy(
	delegatorAddress string,
	validatorAddress string,
	height int64,
) (UnbondingDelegationRow, bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"entries",
		).
		From("view_vd_unbonding_delegations").
		Where(
			"delegator_address = ? AND validator_address = ? AND height = ?",
			delegatorAddress,
			validatorAddress,
			height,
		).
		ToSql()
	if err != nil {
		return UnbondingDelegationRow{}, false, fmt.Errorf("error building select UnbondingDelegation sql: %v: %w", err, rdb.ErrPrepare)
	}

	var unbondingDelegation UnbondingDelegationRow
	unbondingDelegation.Height = height
	unbondingDelegation.DelegatorAddress = delegatorAddress
	unbondingDelegation.ValidatorAddress = validatorAddress

	var entriesJSON string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&entriesJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return UnbondingDelegationRow{}, false, nil
		}
		return UnbondingDelegationRow{}, false, fmt.Errorf("error scanning UnbondingDelegation: %v: %w", err, rdb.ErrQuery)
	}

	if err = jsoniter.UnmarshalFromString(entriesJSON, &unbondingDelegation.Entries); err != nil {
		return UnbondingDelegationRow{}, false, fmt.Errorf("error unmarshalling UnbondingDelegation.Entries JSON: %v: %w", err, rdb.ErrQuery)
	}

	return unbondingDelegation, true, nil
}

func (view *UnbondingDelegationsView) ListByValidator(
	validatorAddress string,
	height int64,
) ([]UnbondingDelegationRow, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"delegator_address",
			"entries",
		).
		From(
			"view_vd_unbonding_delegations",
		).
		Where(
			"validator_address = ? AND height = ?",
			validatorAddress,
			height,
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building UnbondingDelegation select by validator SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing UnbondingDelegation select by validator SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	unbondingDelegations := make([]UnbondingDelegationRow, 0)
	for rowsResult.Next() {
		var unbondingDelegation UnbondingDelegationRow
		unbondingDelegation.Height = height
		unbondingDelegation.ValidatorAddress = validatorAddress

		var entriesJSON string
		if err = rowsResult.Scan(
			&unbondingDelegation.DelegatorAddress,
			&entriesJSON,
		); err != nil {
			return nil, fmt.Errorf("error scanning UnbondingDelegation row: %v: %w", err, rdb.ErrQuery)
		}

		if err = jsoniter.UnmarshalFromString(entriesJSON, &unbondingDelegation.Entries); err != nil {
			return nil, fmt.Errorf("error unmarshalling UnbondingDelegationRow.Entries JSON: %v: %w", err, rdb.ErrQuery)
		}

		unbondingDelegations = append(unbondingDelegations, unbondingDelegation)
	}

	return unbondingDelegations, nil
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
