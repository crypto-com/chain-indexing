package view

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type OperationPerformanceLogs interface {
	Insert(row OperationPerformanceRow) error
}

type OperationPerformanceLogsView struct {
	rdb *rdb.Handle
}

func NewOperationPerformanceLogsView(handle *rdb.Handle) OperationPerformanceLogs {
	return &OperationPerformanceLogsView{
		handle,
	}
}

func (view *OperationPerformanceLogsView) Insert(row OperationPerformanceRow) error {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_vd_operation_performance_log",
		).
		Columns(
			"height",
			"action",
			"duration",
		).
		Values(
			row.Height,
			row.Action,
			row.Duration.Milliseconds(),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building operation performance insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting operation performance into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting operation performance into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type OperationPerformanceRow struct {
	Height   int64
	Action   string
	Duration time.Duration
}
