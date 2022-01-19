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

	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`INSERT INTO view_vd_redelegations(
		height,
		delegator_address,
		validator_src_address,
		validator_dst_address,
		entries
	) SELECT
		?,
		delegator_address,
		validator_src_address,
		validator_dst_address,
		entries
	FROM view_vd_redelegations WHERE height = ?
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building redelegation clone sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		currentHeight,
		previousHeight,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error cloning redelegation into the table: %v: %w", execErr, rdb.ErrWrite)
	}

	return nil
}

func (view *RedelegationsView) Upsert(row RedelegationRow) error {

	entriesJSON, err := jsoniter.MarshalToString(row.Entries)
	if err != nil {
		return fmt.Errorf("error JSON marshalling RedelegationRow.Entries for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_redelegations",
		).
		Columns(
			"height",
			"delegator_address",
			"validator_src_address",
			"validator_dst_address",
			"entries",
		).
		Values(
			row.Height,
			row.DelegatorAddress,
			row.ValidatorSrcAddress,
			row.ValidatorDstAddress,
			entriesJSON,
		).
		Suffix("ON CONFLICT(delegator_address, validator_src_address, validator_dst_address, height) DO UPDATE SET entries = EXCLUDED.entries").
		ToSql()

	if err != nil {
		return fmt.Errorf("error building redelegation upsertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error upserting redelegation into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error upserting redelegation into the table: no row upserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *RedelegationsView) Delete(row RedelegationRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_redelegations",
		).
		Where(
			"delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND height = ?",
			row.DelegatorAddress,
			row.ValidatorSrcAddress,
			row.ValidatorDstAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building redelegation deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting redelegation from the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error deleting redelegation from the table: no row deleted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *RedelegationsView) FindBy(
	delegatorAddress string,
	validatorSrcAddress string,
	validatorDstAddress string,
	height int64,
) (RedelegationRow, bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"entries",
		).
		From("view_vd_redelegations").
		Where(
			"delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND height = ?",
			delegatorAddress,
			validatorSrcAddress,
			validatorDstAddress,
			height,
		).
		ToSql()
	if err != nil {
		return RedelegationRow{}, false, fmt.Errorf("error building select redelegation sql: %v: %w", err, rdb.ErrPrepare)
	}

	var redelegation RedelegationRow
	redelegation.Height = height
	redelegation.DelegatorAddress = delegatorAddress
	redelegation.ValidatorSrcAddress = validatorSrcAddress
	redelegation.ValidatorDstAddress = validatorDstAddress

	var entriesJSON string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&entriesJSON,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return RedelegationRow{}, false, nil
		}
		return RedelegationRow{}, false, fmt.Errorf("error scanning redelegation: %v: %w", err, rdb.ErrQuery)
	}

	if err = jsoniter.UnmarshalFromString(entriesJSON, &redelegation.Entries); err != nil {
		return RedelegationRow{}, false, fmt.Errorf("error unmarshalling RedelegationRow.Entries JSON: %v: %w", err, rdb.ErrQuery)
	}

	return redelegation, true, nil
}

func (view *RedelegationsView) ListBySrcValidator(
	validatorSrcAddress string,
	height int64,
) ([]RedelegationRow, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"delegator_address",
			"validator_dst_address",
			"entries",
		).
		From(
			"view_vd_redelegations",
		).
		Where(
			"validator_src_address = ? AND height = ?",
			validatorSrcAddress,
			height,
		).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building redelegation select by src validator SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing redelegation select by src validator SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	redelegations := make([]RedelegationRow, 0)
	for rowsResult.Next() {
		var redelegation RedelegationRow
		redelegation.Height = height
		redelegation.ValidatorSrcAddress = validatorSrcAddress

		var entriesJSON string
		if err = rowsResult.Scan(
			&redelegation.DelegatorAddress,
			&redelegation.ValidatorSrcAddress,
			&entriesJSON,
		); err != nil {
			return nil, fmt.Errorf("error scanning redelegation row: %v: %w", err, rdb.ErrQuery)
		}

		if err = jsoniter.UnmarshalFromString(entriesJSON, &redelegation.Entries); err != nil {
			return nil, fmt.Errorf("error unmarshalling RedelegationRow.Entries JSON: %v: %w", err, rdb.ErrQuery)
		}

		redelegations = append(redelegations, redelegation)
	}

	return redelegations, nil
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
