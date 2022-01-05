package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/crypto-com/chain-indexing/external/json"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Examples interface {
	Insert(*ExampleRow) error
	List(ExamplesListOrder, *pagination.Pagination) ([]ExampleRow, *pagination.PaginationResult, error)
}

type ExamplesView struct {
	rdb *rdb.Handle
}

func NewExamplesView(handle *rdb.Handle) Examples {
	return &ExamplesView{
		handle,
	}
}

func (examplesView *ExamplesView) Insert(example *ExampleRow) error {
	sql, sqlArgs, err := examplesView.rdb.StmtBuilder.
		Insert(
			"view_examples",
		).
		Columns(
			"address",
			"balance",
		).
		Values(
			example.Address,
			json.MustMarshalToString(example.Balance),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building examples insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := examplesView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting example into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting example into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (examplesView *ExamplesView) List(
	order ExamplesListOrder,
	pagination *pagination.Pagination,
) ([]ExampleRow, *pagination.PaginationResult, error) {
	stmtBuilder := examplesView.rdb.StmtBuilder.Select(
		"address",
		"balance",
	).From(
		"view_examples",
	)

	if order.Address == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("address DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("address")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		examplesView.rdb,
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building examples select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := examplesView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing examples select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	exampleRows := make([]ExampleRow, 0)
	for rowsResult.Next() {
		var example ExampleRow
		var balance string
		if err = rowsResult.Scan(
			&example.Address,
			&balance,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning example row: %v: %w", err, rdb.ErrQuery)
		}

		json.MustUnmarshalFromString(balance, &example.Balance)
		exampleRows = append(exampleRows, example)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return exampleRows, paginationResult, nil
}

type ExampleRow struct {
	Address string     `json:"address"`
	Balance coin.Coins `json:"balance"`
}

type ExamplesListOrder struct {
	Address view.ORDER
}
