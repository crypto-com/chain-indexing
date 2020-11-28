package rdb

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

// RDbPaginationBuilder builds pagination query based on the pagination data. It supports both SQL and StmtBuilder
type RDbPaginationBuilder struct {
	*pagination_interface.Pagination

	customTotalQueryFn CustomTotalQueryFn
	rdbHandle          *Handle
}

func NewRDbPaginationBuilder(pagination *pagination_interface.Pagination, rdbHandle *Handle) *RDbPaginationBuilder {
	return &RDbPaginationBuilder{
		pagination,

		nil,
		rdbHandle,
	}
}

func (pagination *RDbPaginationBuilder) WithCustomTotalQueryFn(fn CustomTotalQueryFn) *RDbPaginationBuilder {
	pagination.customTotalQueryFn = fn

	return pagination
}

func (pagination *RDbPaginationBuilder) BuildStmt(stmtBuilder sq.SelectBuilder) *RDbPaginationStmtBuilder {
	return &RDbPaginationStmtBuilder{
		Pagination: pagination.Pagination,

		rdbHandle:          pagination.rdbHandle,
		stmtBuilder:        stmtBuilder,
		customTotalQueryFn: pagination.customTotalQueryFn,
	}
}

func (pagination *RDbPaginationBuilder) BuildSQL(sql string, args ...interface{}) *RDbPaginationSQLBuilder {
	return &RDbPaginationSQLBuilder{
		Pagination: pagination.Pagination,

		rdbHandle: pagination.rdbHandle,
		sql:       sql,
		args:      args,
	}
}

// RDbPaginationBuilder is the concrete implementation to builds the pagination for StmtBuilder
type RDbPaginationStmtBuilder struct {
	*pagination_interface.Pagination

	rdbHandle          *Handle
	stmtBuilder        sq.SelectBuilder
	customTotalQueryFn CustomTotalQueryFn
}

func (pagination *RDbPaginationStmtBuilder) ToStmtBuilder() sq.SelectBuilder {
	switch pagination.Type() {
	case pagination_interface.PAGINATION_OFFSET:

		params := pagination.OffsetParams()
		return pagination.stmtBuilder.Suffix("LIMIT ? OFFSET ?", params.Limit, params.Offset())
	}

	return pagination.stmtBuilder
}

func (pagination *RDbPaginationStmtBuilder) Result() (*pagination_interface.PaginationResult, error) {
	switch pagination.Type() {
	case pagination_interface.PAGINATION_OFFSET:
		return pagination.offsetResult()
	}

	return nil, nil
}

func (pagination *RDbPaginationStmtBuilder) offsetResult() (*pagination_interface.PaginationResult, error) {
	var err error

	sql, sqlArgs, err := pagination.stmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building total count select SQL: %v", err)
	}

	var total int64
	if pagination.customTotalQueryFn != nil {
		total, err = pagination.customTotalQueryFn(pagination.rdbHandle, pagination.stmtBuilder)
		if err != nil {
			return nil, fmt.Errorf("error executing custom total query function: %v", err)
		}
	} else {
		// No new parameter is introduced in the prepared statement. Should not
		// violate security
		// nolint:gosec
		sql = fmt.Sprintf("SELECT COUNT(*) FROM (%s) counting_table", sql)

		if err = pagination.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&total); err != nil {
			return nil, fmt.Errorf("error executing total count select SQL: %v", err)
		}
	}

	return pagination.OffsetResult(total), nil
}

// RDbPaginationBuilder is the concrete implementation to builds the pagination for SQL query
type RDbPaginationSQLBuilder struct {
	*pagination_interface.Pagination

	rdbHandle Runner
	sql       string
	args      []interface{}
}

func (pagination *RDbPaginationSQLBuilder) ToSQL() (string, []interface{}) {
	switch pagination.Type() {
	case pagination_interface.PAGINATION_OFFSET:
		lastPlacerHolder := len(pagination.args)
		params := pagination.OffsetParams()

		args := make([]interface{}, 0, len(pagination.args))
		copy(args, pagination.args)
		return fmt.Sprintf("%s LIMIT $%d OFFSET $%d",
			pagination.sql, lastPlacerHolder+1, lastPlacerHolder+2,
		), append(pagination.args, params.Limit, params.Offset())
	}

	return pagination.sql, pagination.args
}

func (pagination *RDbPaginationSQLBuilder) Result() (*pagination_interface.PaginationResult, error) {
	switch pagination.Type() {
	case pagination_interface.PAGINATION_OFFSET:
		return pagination.offsetResult()
	}

	return nil, nil
}

func (pagination *RDbPaginationSQLBuilder) offsetResult() (*pagination_interface.PaginationResult, error) {
	var err error

	// No new parameter is introduced in the prepare statement. Hence not
	// violating security
	// nolint:gosec
	sql := fmt.Sprintf("SELECT COUNT(*) FROM (%s) counting_table", pagination.sql)

	var total int64
	if err = pagination.rdbHandle.QueryRow(sql, pagination.args...).Scan(&total); err != nil {
		return nil, fmt.Errorf("error executing total count select SQL: %v", err)
	}

	return pagination.OffsetResult(total), nil
}

type CustomTotalQueryFn = func(conn *Handle, originalSelectBuilder sq.SelectBuilder) (int64, error)

//"SELECT reltuples::bigint FROM pg_class where relname='$1';
