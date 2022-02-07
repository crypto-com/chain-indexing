package view

import (
	"errors"
	"fmt"

	pagination_appinterface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type Delegations interface {
	Insert(row DelegationRow) error
	Update(row DelegationRow) error
	Delete(row DelegationRow) error

	FindBy(validatorAddress string, delegatorAddress string, height int64) (DelegationRow, bool, error)

	ListByValidatorAddr(
		validatorAddress string,
		height int64,
		pagination *pagination_appinterface.Pagination,
	) ([]DelegationRow, *pagination_appinterface.PaginationResult, error)
	ListByDelegatorAddr(
		delegatorAddress string,
		height int64,
		pagination *pagination_appinterface.Pagination,
	) ([]DelegationRow, *pagination_appinterface.PaginationResult, error)
}

type DelegationsView struct {
	rdb *rdb.Handle
}

func NewDelegationsView(handle *rdb.Handle) Delegations {
	return &DelegationsView{
		handle,
	}
}

func (view *DelegationsView) Insert(row DelegationRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_delegations",
		).
		Columns(
			"height",
			"validator_address",
			"delegator_address",
			"shares",
		).
		Values(
			fmt.Sprintf("[%v,)", row.Height),
			row.ValidatorAddress,
			row.DelegatorAddress,
			row.Shares.String(),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building delegation insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting delegation into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting delegation into the table: rows inserted: %v: %w", result.RowsAffected(), rdb.ErrWrite)
	}

	return nil
}

func (view *DelegationsView) Update(row DelegationRow) error {

	// Check if there is an record's lower bound start with this height.
	found, err := view.checkIfDelegationRecordExistByHeightLowerBound(row)
	if err != nil {
		return fmt.Errorf("error in checking new delegation record existence at this height: %v", err)
	}

	if found {
		// If there is a record lower(height) == row.Height, then update the existed one

		sql, sqlArgs, err := view.rdb.StmtBuilder.
			Update(
				"view_vd_delegations",
			).
			SetMap(map[string]interface{}{
				"shares": row.Shares.String(),
			}).
			Where(
				"validator_address = ? AND delegator_address = ? AND height @> ?::int8",
				row.ValidatorAddress,
				row.DelegatorAddress,
				row.Height,
			).
			ToSql()

		if err != nil {
			return fmt.Errorf("error building delegation update sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		result, err := view.rdb.Exec(sql, sqlArgs...)
		if err != nil {
			return fmt.Errorf("error updating delegation into the table: %v: %w", err, rdb.ErrWrite)
		}
		if result.RowsAffected() != 1 {
			return fmt.Errorf("error updating delegation into the table: row updated: %v: %w", result.RowsAffected(), rdb.ErrWrite)
		}

	} else {
		// Else, update the previous record's height range, then insert a new record

		err := view.setDelegationRecordHeightUpperBound(row)
		if err != nil {
			return fmt.Errorf("error updating delegation.Height upper bound: %v", err)
		}
		err = view.Insert(row)
		if err != nil {
			return fmt.Errorf("error inserting a new record for this delegation: %v", err)
		}

	}

	return nil
}

func (view *DelegationsView) checkIfDelegationRecordExistByHeightLowerBound(row DelegationRow) (bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"COUNT(*)",
		).
		From("view_vd_delegations").
		Where(
			"validator_address = ? AND delegator_address = ? AND lower(height) = ?",
			row.ValidatorAddress,
			row.DelegatorAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("error building 'checking delegation at specific lower(height) sql: %v: %w", err, rdb.ErrPrepare)
	}

	var count int64
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&count,
	); err != nil {
		return false, fmt.Errorf("error scanning count: %v: %w", err, rdb.ErrQuery)
	}

	return count == 1, nil
}

func (view *DelegationsView) Delete(row DelegationRow) error {

	return view.setDelegationRecordHeightUpperBound(row)
}

func (view *DelegationsView) setDelegationRecordHeightUpperBound(row DelegationRow) error {

	// Set the upper bound for record height: `[<PREVIOUS-LOWER-BOUND>, row.Height)`.
	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`
		UPDATE view_vd_delegations
		SET height = int8range(lower(height), ?, '[)')
		WHERE validator_address = ? AND delegator_address = ? AND height @> ?::int8
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building delegation upper(height) update sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		row.Height,
		row.ValidatorAddress,
		row.DelegatorAddress,
		row.Height,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error executing delegation upper(height)update sql: %v: %w", execErr, rdb.ErrWrite)
	}

	return nil
}

func (view *DelegationsView) FindBy(
	validatorAddress string,
	delegatorAddress string,
	height int64,
) (DelegationRow, bool, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"shares",
		).
		From("view_vd_delegations").
		Where(
			"validator_address = ? AND delegator_address = ? AND height @> ?::int8",
			validatorAddress,
			delegatorAddress,
			height,
		).
		ToSql()
	if err != nil {
		return DelegationRow{}, false, fmt.Errorf("error building select delegation sql: %v: %w", err, rdb.ErrPrepare)
	}

	var delegation DelegationRow
	delegation.Height = height
	delegation.ValidatorAddress = validatorAddress
	delegation.DelegatorAddress = delegatorAddress

	var sharesInString string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&sharesInString,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			// When the row is not found, do not return error
			return DelegationRow{}, false, nil
		}
		return DelegationRow{}, false, fmt.Errorf("error scanning delegation: %v: %w", err, rdb.ErrQuery)
	}

	delegation.Shares, err = coin.NewDecFromStr(sharesInString)
	if err != nil {
		return DelegationRow{}, false, fmt.Errorf("error parsing shares to coin.Dec: %v: %w", err, rdb.ErrQuery)
	}

	return delegation, true, nil
}

func (view *DelegationsView) ListByValidatorAddr(
	validatorAddress string,
	height int64,
	pagination *pagination_appinterface.Pagination,
) ([]DelegationRow, *pagination_appinterface.PaginationResult, error) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"delegator_address",
			"shares",
		).
		From(
			"view_vd_delegations",
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
		return nil, nil, fmt.Errorf("error building delegations select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing delegations select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	delegations := make([]DelegationRow, 0)
	for rowsResult.Next() {
		var delegation DelegationRow
		delegation.Height = height
		delegation.ValidatorAddress = validatorAddress

		var sharesInString string
		if err = rowsResult.Scan(
			&delegation.DelegatorAddress,
			&sharesInString,
		); err != nil {
			return nil, nil, fmt.Errorf("error scanning delegation row: %v: %w", err, rdb.ErrQuery)
		}

		delegation.Shares, err = coin.NewDecFromStr(sharesInString)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing shares to coin.Dec: %v: %w", err, rdb.ErrQuery)
		}

		delegations = append(delegations, delegation)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return delegations, paginationResult, nil
}

func (view *DelegationsView) ListByDelegatorAddr(
	delegatorAddress string,
	height int64,
	pagination *pagination_appinterface.Pagination,
) ([]DelegationRow, *pagination_appinterface.PaginationResult, error) {

	stmtBuilder := view.rdb.StmtBuilder.
		Select(
			"validator_address",
			"shares",
		).
		From(
			"view_vd_delegations",
		).
		Where(
			"delegator_address = ? AND height @> ?::int8",
			delegatorAddress,
			height,
		)

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		view.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building delegations select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := view.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing delegations select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	delegations := make([]DelegationRow, 0)
	for rowsResult.Next() {
		var delegation DelegationRow
		delegation.Height = height
		delegation.DelegatorAddress = delegatorAddress

		var sharesInString string
		if err = rowsResult.Scan(
			&delegation.ValidatorAddress,
			&sharesInString,
		); err != nil {
			return nil, nil, fmt.Errorf("error scanning delegation row: %v: %w", err, rdb.ErrQuery)
		}

		delegation.Shares, err = coin.NewDecFromStr(sharesInString)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing shares to coin.Dec: %v: %w", err, rdb.ErrQuery)
		}

		delegations = append(delegations, delegation)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return delegations, paginationResult, nil
}

type DelegationRow struct {
	Height           int64    `json:"height"`
	ValidatorAddress string   `json:"validatorAddress"`
	DelegatorAddress string   `json:"delegatorAddress"`
	Shares           coin.Dec `json:"shares"`
}
