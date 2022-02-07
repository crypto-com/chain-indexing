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

type Redelegations interface {
	Insert(row RedelegationRow) error
	Update(row RedelegationRow) error
	Delete(row RedelegationRow) error
	FindBy(delegatorAddress, validatorSrcAddress, validatorDstAddress string, height int64) (RedelegationRow, bool, error)
	// For internal projection logic
	ListBySrcValidator(validatorSrcAddress string, height int64) ([]RedelegationRow, error)
	// For HTTP API
	ListBySrcValidatorWithPagination(
		validatorSrcAddress string,
		height int64,
		pagination *pagination_appinterface.Pagination,
	) ([]RedelegationRow, *pagination_appinterface.PaginationResult, error)
}

type RedelegationsView struct {
	rdb *rdb.Handle
}

func NewRedelegationsView(handle *rdb.Handle) Redelegations {
	return &RedelegationsView{
		handle,
	}
}

func (view *RedelegationsView) Insert(row RedelegationRow) error {

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
			fmt.Sprintf("[%v,)", row.Height),
			row.DelegatorAddress,
			row.ValidatorSrcAddress,
			row.ValidatorDstAddress,
			entriesJSON,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building redelegation insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting redelegation into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting redelegation into the table: no row upserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *RedelegationsView) Update(row RedelegationRow) error {

	// Check if there is an record's lower bound start with this height.
	found, err := view.checkIfRedelegationRecordExistByHeightLowerBound(row)
	if err != nil {
		return fmt.Errorf("error in checking new Redelegation record existence at this height: %v", err)
	}

	if found {
		// If there is a record lower(height) == row.Height, then update the existed one

		entriesJSON, err := jsoniter.MarshalToString(row.Entries)
		if err != nil {
			return fmt.Errorf("error JSON marshalling RedelegationRow.Entries for upsertion: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		sql, sqlArgs, err := view.rdb.StmtBuilder.
			Update(
				"view_vd_redelegations",
			).
			SetMap(map[string]interface{}{
				"entries": entriesJSON,
			}).
			Where(
				"delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND height @> ?::int8",
				row.DelegatorAddress,
				row.ValidatorSrcAddress,
				row.ValidatorDstAddress,
				row.Height,
			).
			ToSql()

		if err != nil {
			return fmt.Errorf("error building Redelegation update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		result, err := view.rdb.Exec(sql, sqlArgs...)
		if err != nil {
			return fmt.Errorf("error updating Redelegation into the table: %v: %w", err, rdb.ErrWrite)
		}
		if result.RowsAffected() != 1 {
			return fmt.Errorf("error updating Redelegation into the table: row updated: %v: %w", result.RowsAffected(), rdb.ErrWrite)
		}

	} else {
		// If there is not an existed record, then update the previous record's height range and insert a new record

		err := view.setRedelegationRecordHeightUpperBound(row)
		if err != nil {
			return fmt.Errorf("error updating Redelegation.Height upper bound: %v", err)
		}
		err = view.Insert(row)
		if err != nil {
			return fmt.Errorf("error inserting a new record for this Redelegation: %v", err)
		}

	}

	return nil
}

func (view *RedelegationsView) checkIfRedelegationRecordExistByHeightLowerBound(row RedelegationRow) (bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"COUNT(*)",
		).
		From("view_vd_redelegations").
		Where(
			"delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND lower(height) = ?",
			row.DelegatorAddress,
			row.ValidatorSrcAddress,
			row.ValidatorDstAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("error building 'checking Redelegation at specific lower(height) sql: %v: %w", err, rdb.ErrPrepare)
	}

	var count int64
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&count,
	); err != nil {
		return false, fmt.Errorf("error scanning count: %v: %w", err, rdb.ErrQuery)
	}

	return count == 1, nil
}

func (view *RedelegationsView) Delete(row RedelegationRow) error {

	return view.setRedelegationRecordHeightUpperBound(row)
}

func (view *RedelegationsView) setRedelegationRecordHeightUpperBound(row RedelegationRow) error {

	// Set the upper bound for record height: `[<PREVIOUS-LOWER-BOUND>, row.Height)`.
	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`
		UPDATE view_vd_redelegations
		SET height = int8range(lower(height), ?, '[)')
		WHERE delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND height @> ?::int8
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building Redelegation upper(height) update sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		row.Height,
		row.DelegatorAddress,
		row.ValidatorSrcAddress,
		row.ValidatorDstAddress,
		row.Height,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error executing Redelegation upper(height)update sql: %v: %w", execErr, rdb.ErrWrite)
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
			"delegator_address = ? AND validator_src_address = ? AND validator_dst_address = ? AND height @> ?::int8",
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
			"validator_src_address = ? AND height @> ?::int8",
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
			&redelegation.ValidatorDstAddress,
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

func (view *RedelegationsView) ListBySrcValidatorWithPagination(
	validatorSrcAddress string,
	height int64,
	pagination *pagination_appinterface.Pagination,
) ([]RedelegationRow, *pagination_appinterface.PaginationResult, error) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"delegator_address",
			"validator_dst_address",
			"entries",
		).
		From(
			"view_vd_redelegations",
		).
		Where(
			"validator_src_address = ? AND height @> ?::int8",
			validatorSrcAddress,
			height,
		)

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building redelegation select by src validator SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing redelegation select by src validator SQL: %v: %w", err, rdb.ErrQuery)
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
			&redelegation.ValidatorDstAddress,
			&entriesJSON,
		); err != nil {
			return nil, nil, fmt.Errorf("error scanning redelegation row: %v: %w", err, rdb.ErrQuery)
		}

		if err = jsoniter.UnmarshalFromString(entriesJSON, &redelegation.Entries); err != nil {
			return nil, nil, fmt.Errorf("error unmarshalling RedelegationRow.Entries JSON: %v: %w", err, rdb.ErrQuery)
		}

		redelegations = append(redelegations, redelegation)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return redelegations, paginationResult, nil
}

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
		Height:              creationHeight,
		DelegatorAddress:    delegatorAddress,
		ValidatorSrcAddress: validatorSrcAddress,
		ValidatorDstAddress: validatorDstAddress,
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
