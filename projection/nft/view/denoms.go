package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

const DENOMS_TABLE_NAME = "view_nft_denoms"

type Denoms struct {
	rdb *rdb.Handle
}

func NewDenoms(handle *rdb.Handle) *Denoms {
	return &Denoms{
		handle,
	}
}

func (denomsView *Denoms) Insert(denomRow *DenomRow) error {
	var err error

	sql, sqlArgs, err := denomsView.rdb.StmtBuilder.Insert(
		DENOMS_TABLE_NAME,
	).Columns(
		"denom_id",
		"name",
		"schema",
		"creator",
		"created_at",
	).Values(
		denomRow.DenomId,
		denomRow.Name,
		denomRow.Schema,
		denomRow.Creator,
		denomsView.rdb.Tton(&denomRow.CreatedAt),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building NFT denom insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := denomsView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting NFT denom into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting NFT denom into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (denomsView *Denoms) FindById(denomId string) (*DenomRow, error) {
	selectStmtBuilder := denomsView.rdb.StmtBuilder.Select(
		"denom_id",
		"name",
		"schema",
		"creator",
		"created_at",
	).From(
		DENOMS_TABLE_NAME,
	).Where(
		"denom_id = ?", denomId,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building NFT denom selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row DenomRow
	createdAtTimeReader := denomsView.rdb.NtotReader()

	if err = denomsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.DenomId,
		&row.Name,
		&row.Schema,
		&row.Creator,
		createdAtTimeReader.ScannableArg(),
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning NFT denom row: %v: %w", err, rdb.ErrQuery)
	}

	createdAt, parseCreatedAtErr := createdAtTimeReader.Parse()
	if parseCreatedAtErr != nil {
		return nil, fmt.Errorf("error parsing NFT denom creation time: %v: %w", parseCreatedAtErr, rdb.ErrQuery)
	}
	row.CreatedAt = *createdAt

	return &row, nil
}

func (denomsView *Denoms) FindByName(denomName string) (*DenomRow, error) {
	selectStmtBuilder := denomsView.rdb.StmtBuilder.Select(
		"denom_id",
		"name",
		"schema",
		"creator",
		"created_at",
	).From(
		DENOMS_TABLE_NAME,
	).Where(
		"name = ?", denomName,
	)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building NFT denom selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var row DenomRow
	createdAtTimeReader := denomsView.rdb.NtotReader()

	if err = denomsView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&row.DenomId,
		&row.Name,
		&row.Schema,
		&row.Creator,
		createdAtTimeReader.ScannableArg(),
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning NFT denom row: %v: %w", err, rdb.ErrQuery)
	}

	createdAt, parseCreatedAtErr := createdAtTimeReader.Parse()
	if parseCreatedAtErr != nil {
		return nil, fmt.Errorf("error parsing NFT denom creation time: %v: %w", parseCreatedAtErr, rdb.ErrQuery)
	}
	row.CreatedAt = *createdAt

	return &row, nil
}

func (denomsView *Denoms) List(
	filter DenomListFilter,
	order DenomListOrder,
	pagination *pagination_interface.Pagination,
) ([]DenomRow, *pagination_interface.PaginationResult, error) {
	stmtBuilder := denomsView.rdb.StmtBuilder.Select(
		"denom_id",
		"name",
		"schema",
		"creator",
		"created_at",
	).From(
		DENOMS_TABLE_NAME,
	)

	if filter.MaybeCreator != nil {
		stmtBuilder = stmtBuilder.Where("creator = ?", *filter.MaybeCreator)
	}

	if order.CreatedAt == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("id DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("id")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		denomsView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			totalView := NewDenomsTotal(rdbHandle)
			identifier := "-"
			if filter.MaybeCreator != nil {
				identifier = *filter.MaybeCreator
			}
			total, err := totalView.FindBy(identifier)
			if err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building NFT denoms list SQL: %v: %w", err, rdb.ErrPrepare)
	}

	rowsResult, err := denomsView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing NFT denoms list query: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	rows := make([]DenomRow, 0)
	for rowsResult.Next() {
		var row DenomRow
		createdAtTimeReader := denomsView.rdb.NtotReader()

		if scanErr := rowsResult.Scan(
			&row.DenomId,
			&row.Name,
			&row.Schema,
			&row.Creator,
			createdAtTimeReader.ScannableArg(),
		); scanErr != nil {
			if errors.Is(scanErr, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning NFT denom row: %v: %w", scanErr, rdb.ErrQuery)
		}

		createdAt, parseCreatedAtErr := createdAtTimeReader.Parse()
		if parseCreatedAtErr != nil {
			return nil, nil, fmt.Errorf("error parsing NFT denom creation time: %v: %w", parseCreatedAtErr, rdb.ErrQuery)
		}
		row.CreatedAt = *createdAt

		rows = append(rows, row)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return rows, paginationResult, nil
}

type DenomListFilter struct {
	MaybeCreator *string
}

type DenomListOrder struct {
	CreatedAt view.ORDER
}

type DenomRow struct {
	DenomId   string          `json:"denomId"`
	Name      string          `json:"denomName"`
	Schema    string          `json:"denomSchema"`
	Creator   string          `json:"denomCreator"`
	CreatedAt utctime.UTCTime `json:"denomCreatedAt"`
}
