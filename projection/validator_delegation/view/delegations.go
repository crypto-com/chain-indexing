package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_appinterface "github.com/crypto-com/chain-indexing/appinterface/pagination"
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

	sql, sqlErr := view.rdb.StmtBuilder.ReplacePlaceholders(`INSERT INTO view_vd_delegations(
		height,
		validator_address,
		delegator_address,
		shares
	) SELECT
		?,
		validator_address,
		delegator_address,
		shares?
	FROM view_vd_delegations WHERE height = ?
	`)
	if sqlErr != nil {
		return fmt.Errorf("error building delegation clone sql: %v: %w", sqlErr, rdb.ErrBuildSQLStmt)
	}
	sqlArgs := []interface{}{
		currentHeight,
		previousHeight,
	}

	if _, execErr := view.rdb.Exec(sql, sqlArgs...); execErr != nil {
		return fmt.Errorf("error cloning delegation into the table: %v: %w", execErr, rdb.ErrWrite)
	}

	return nil
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
			row.Height,
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
		return fmt.Errorf("error inserting delegation into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *DelegationsView) Update(row DelegationRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Update(
			"view_vd_delegations",
		).
		SetMap(map[string]interface{}{
			"shares": row.Shares.String(),
		}).
		Where(
			"validator_address = ? AND delegator_address = ? AND height = ?",
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
		return fmt.Errorf("error updating delegation into the table: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *DelegationsView) Delete(row DelegationRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Delete(
			"view_vd_delegations",
		).
		Where(
			"validator_address = ? AND delegator_address = ? AND height = ?",
			row.ValidatorAddress,
			row.DelegatorAddress,
			row.Height,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building delegation deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting delegation from the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error deleting delegation into the table: no row deleted: %w", rdb.ErrWrite)
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
			"validator_address = ? AND delegator_address = ? AND height = ?",
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
			"validator_address = ? AND height = ?",
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
			"delegator_address = ? AND height = ?",
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

// Notes:
// - UNIQUE(validatorAddress, delegatorAddress, height)
// - INDEX(validatorAddress, height)
// - INDEX(delegatorAddress, height)
type DelegationRow struct {
	Height           int64    `json:"height"`
	ValidatorAddress string   `json:"validatorAddress"`
	DelegatorAddress string   `json:"delegatorAddress"`
	Shares           coin.Dec `json:"shares"`
}
