package view

import (
	"errors"
	"fmt"
	"time"

	pagination_appinterface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	jsoniter "github.com/json-iterator/go"
)

type UnbondingDelegations interface {
	Insert(row UnbondingDelegationRow) error
	Update(row UnbondingDelegationRow) error
	Delete(row UnbondingDelegationRow) error
	FindBy(delegatorAddress, validatorAddress string, height int64) (UnbondingDelegationRow, bool, error)
	// For internal projection logic
	ListByValidator(validatorAddress string, height int64) ([]UnbondingDelegationRow, error)
	// For HTTP API
	ListByValidatorWithPagination(
		validatorAddress string,
		height int64,
		pagination *pagination_appinterface.Pagination,
	) ([]UnbondingDelegationRow, *pagination_appinterface.PaginationResult, error)
}

type UnbondingDelegationsView struct {
	rdb *rdb.Handle
}

func NewUnbondingDelegationsView(handle *rdb.Handle) UnbondingDelegations {
	return &UnbondingDelegationsView{
		handle,
	}
}

func (view *UnbondingDelegationsView) Insert(row UnbondingDelegationRow) error {
	entriesJSON, err := jsoniter.MarshalToString(row.Entries)
	if err != nil {
		return fmt.Errorf("error JSON marshalling UnbondingDelegationRow.Entries for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
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
			fmt.Sprintf("[%v,)", row.Height),
			row.DelegatorAddress,
			row.ValidatorAddress,
			entriesJSON,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building UnbondingDelegation insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting UnbondingDelegation into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting UnbondingDelegation into the table: no row upserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *UnbondingDelegationsView) Update(row UnbondingDelegationRow) error {

	// Check if there is an record's lower bound start with this height.
	found, err := view.checkIfUnbondingDelegationRecordExistByHeightLowerBound(row)
	if err != nil {
		return fmt.Errorf("error in checking new UnbondingDelegation record existence at this height: %v", err)
	}

	if found {
		// If there is a record that height = `[row.Height,)`, then update the existed one

		entriesJSON, err := jsoniter.MarshalToString(row.Entries)
		if err != nil {
			return fmt.Errorf("error JSON marshalling UnbondingDelegationRow.Entries for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		sql, sqlArgs, err := view.rdb.StmtBuilder.
			Update(
				"view_vd_unbonding_delegations",
			).
			SetMap(map[string]interface{}{
				"entries": entriesJSON,
			}).
			Where(
				"delegator_address = ? AND validator_address = ? AND height @> ?::int8",
				row.DelegatorAddress,
				row.ValidatorAddress,
				row.Height,
			).
			ToSql()

		if err != nil {
			return fmt.Errorf("error building UnbondingDelegation update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		result, err := view.rdb.Exec(sql, sqlArgs...)
		if err != nil {
			return fmt.Errorf("error updating UnbondingDelegation into the table: %v: %w", err, rdb.ErrWrite)
		}
		if result.RowsAffected() != 1 {
			return fmt.Errorf("error updating UnbondingDelegation into the table: row updated: %v: %w", result.RowsAffected(), rdb.ErrWrite)
		}

	} else {
		// If there is not an existed record, then update the previous record's height range and insert a new record

		err := view.setUnbondingDelegationRecordHeightUpperBound(row)
		if err != nil {
			return fmt.Errorf("error updating UnbondingDelegation.Height upper bound: %v", err)
		}
		err = view.Insert(row)
		if err != nil {
			return fmt.Errorf("error inserting a new record for this UnbondingDelegation: %v", err)
		}

	}

	return nil
}

func (view *UnbondingDelegationsView) checkIfUnbondingDelegationRecordExistByHeightLowerBound(row UnbondingDelegationRow) (bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"COUNT(*)",
		).
		From("view_vd_unbonding_delegations").
		Where(
			"delegator_address = ? AND validator_address = ? AND height = ?",
			row.DelegatorAddress,
			row.ValidatorAddress,
			fmt.Sprintf("[%v,)", row.Height),
		).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("error building sql to check unbonding delegation at specific height: %v: %w", err, rdb.ErrPrepare)
	}

	var count int64
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&count,
	); err != nil {
		return false, fmt.Errorf("error scanning count: %v: %w", err, rdb.ErrQuery)
	}

	return count == 1, nil
}

func (view *UnbondingDelegationsView) Delete(row UnbondingDelegationRow) error {

	return view.setUnbondingDelegationRecordHeightUpperBound(row)
}

func (view *UnbondingDelegationsView) setUnbondingDelegationRecordHeightUpperBound(row UnbondingDelegationRow) error {

	// Set the upper bound for record height: `[<PREVIOUS-LOWER-BOUND>, row.Height)`.
	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`
		UPDATE view_vd_unbonding_delegations
		SET height = int8range(lower(height), ?, '[)')
		WHERE delegator_address = ? AND validator_address = ? AND height @> ?::int8
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building UnbondingDelegation upper(height) update sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		row.Height,
		row.DelegatorAddress,
		row.ValidatorAddress,
		row.Height,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error executing UnbondingDelegation upper(height)update sql: %v: %w", execErr, rdb.ErrWrite)
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
			"delegator_address = ? AND validator_address = ? AND height @> ?::int8",
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
			"validator_address = ? AND height @> ?::int8",
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

func (view *UnbondingDelegationsView) ListByValidatorWithPagination(
	validatorAddress string,
	height int64,
	pagination *pagination_appinterface.Pagination,
) ([]UnbondingDelegationRow, *pagination_appinterface.PaginationResult, error) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"delegator_address",
			"entries",
		).
		From(
			"view_vd_unbonding_delegations",
		).
		Where(
			"validator_address = ? AND height @> ?::int8",
			validatorAddress,
			height,
		)

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building UnbondingDelegation select by validator SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing UnbondingDelegation select by validator SQL: %v: %w", err, rdb.ErrQuery)
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
			return nil, nil, fmt.Errorf("error scanning UnbondingDelegation row: %v: %w", err, rdb.ErrQuery)
		}

		if err = jsoniter.UnmarshalFromString(entriesJSON, &unbondingDelegation.Entries); err != nil {
			return nil, nil, fmt.Errorf("error unmarshalling UnbondingDelegationRow.Entries JSON: %v: %w", err, rdb.ErrQuery)
		}

		unbondingDelegations = append(unbondingDelegations, unbondingDelegation)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return unbondingDelegations, paginationResult, nil
}

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
		Height:           creationHeight,
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
