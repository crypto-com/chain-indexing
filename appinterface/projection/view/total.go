package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/rdb"
)

// A "Total" compatible table should have the follow table schema
// | Column   | Type    | Constraint  |
// | -------- | ------- | ----------- |
// | identity | VARCHAR | PRIMARY KEY |
// | total    | BIGINT  | NOT NULL    |

// A generic record total tracker view
type Total struct {
	rdbHandle *rdb.Handle

	tableName string
}

func NewTotal(rdbHandle *rdb.Handle, tableName string) *Total {
	return &Total{
		rdbHandle,

		tableName,
	}
}

func (view *Total) Set(identity string, total int64) error {
	var err error

	var sql string
	var sqlArgs []interface{}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Select(
		"1",
	).From(
		view.tableName,
	).Where(
		"identity = ?", identity,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error preparing total selection SQL: %v", err)
	}
	var placeholder int
	err = view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&placeholder)
	if err != nil {
		if err != rdb.ErrNoRows {
			return fmt.Errorf("error scanning total: %v", err)
		}
		sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Insert(
			view.tableName,
		).Columns(
			"identity",
			"total",
		).Values(identity, total).ToSql()
		if err != nil {
			return fmt.Errorf("error building total insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		var execResult rdb.ExecResult
		if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
			return fmt.Errorf("error inserting total: %v", err)
		}
		if execResult.RowsAffected() != 1 {
			return errors.New("error inserting total: no rows inserted")
		}

		return nil
	}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Update(
		view.tableName,
	).Set(
		"total", total,
	).Where(
		"identity = ?", identity,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building total update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating total: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating total: no rows updated")
	}

	return nil
}

func (view *Total) Increment(identity string, incTotal int64) error {
	var err error

	var sql string
	var sqlArgs []interface{}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Select(
		"total",
	).From(
		view.tableName,
	).Where(
		"identity = ?", identity,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error preparing total selection SQL: %v", err)
	}

	var prevTotal int64
	err = view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&prevTotal)
	if err != nil {
		if err != rdb.ErrNoRows {
			return fmt.Errorf("error scanning total of identity: %v", err)
		}
		sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Insert(
			view.tableName,
		).Columns(
			"identity",
			"total",
		).Values(identity, incTotal).ToSql()
		if err != nil {
			return fmt.Errorf("error building total insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		var execResult rdb.ExecResult
		if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
			return fmt.Errorf("error inserting total: %v", err)
		}
		if execResult.RowsAffected() != 1 {
			return errors.New("error inserting total: no rows inserted")
		}

		return nil
	}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Update(
		view.tableName,
	).Set(
		"total", prevTotal+incTotal,
	).Where(
		"identity = ?", identity,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building total update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating total: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating total: no rows updated")
	}

	return nil
}

func (view *Total) FindBy(identity string) (int64, error) {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.Select(
		"total",
	).From(
		view.tableName,
	).Where(
		"identity = ?", identity,
	).ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error preparing total selection SQL: %v", err)
	}

	var total int64
	if err := view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&total); err != nil {
		if err == rdb.ErrNoRows {
			return int64(0), nil
		}
		return int64(0), fmt.Errorf("error getting total: %v", err)
	}

	return total, nil
}
