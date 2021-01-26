package view

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
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
	// Postgres UPSERT statement
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.
		Insert(view.tableName).
		Columns("identity", "total").
		Values(identity, total).
		Suffix("ON CONFLICT (identity) DO UPDATE SET total = EXCLUDED.total").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building total insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdbHandle.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting total: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating total: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *Total) Increment(identity string, total int64) error {
	// Postgres UPSERT statement
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.
		Insert(view.tableName+" AS totals").
		Columns("identity", "total").
		Values(identity, total).
		Suffix("ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building total insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdbHandle.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting total: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting total: no row inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *Total) IncrementAll(identities []string, total int64) error {
	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	totalIdentityCount := len(identities)
	for i, identity := range identities {
		if pendingRowCount == 0 {
			// Postgres UPSERT statement
			stmtBuilder = view.rdbHandle.StmtBuilder.
				Insert(view.tableName+" AS totals").
				Columns("identity", "total").
				Suffix("ON CONFLICT (identity) DO UPDATE SET total = totals.total + EXCLUDED.total")
		}

		stmtBuilder = stmtBuilder.Values(identity, total)
		pendingRowCount += 1

		if pendingRowCount == 500 || i+1 == totalIdentityCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building total insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := view.rdbHandle.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting total: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(totalIdentityCount) {
				return fmt.Errorf("error inserting total: no row inserted: %w", rdb.ErrWrite)
			}
			pendingRowCount = 0
		}
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
		if errors.Is(err, rdb.ErrNoRows) {
			return int64(0), nil
		}
		return int64(0), fmt.Errorf("error getting total: %v", err)
	}

	return total, nil
}

func (view *Total) SumBy(identities []string) (int64, error) {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.Select(
		"SUM(total)",
	).From(
		view.tableName,
	).Where(
		sq.Eq{"identity": identities},
	).ToSql()
	if err != nil {
		return int64(0), fmt.Errorf("error preparing total selection SQL: %v", err)
	}

	var total *int64
	if err := view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&total); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return int64(0), nil
		}
		return int64(0), fmt.Errorf("error getting total: %v", err)
	}
	if total == nil {
		return int64(0), nil
	}

	return *total, nil
}
