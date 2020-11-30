package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type ValidatorStats struct {
	rdbHandle *rdb.Handle
}

func NewValidatorStats(rdbHandle *rdb.Handle) *ValidatorStats {
	return &ValidatorStats{
		rdbHandle,
	}
}

func (view *ValidatorStats) Set(metrics string, value string) error {
	var err error

	var sql string
	var sqlArgs []interface{}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Select(
		"1",
	).From(
		"view_validator_stats",
	).Where(
		"metrics = ?", metrics,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error preparing metrics selection SQL: %v", err)
	}
	var placeholder int
	err = view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&placeholder)
	if err != nil {
		if err != rdb.ErrNoRows {
			return fmt.Errorf("error scanning metrics: %v", err)
		}
		sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Insert(
			"view_validator_stats",
		).Columns(
			"metrics",
			"value",
		).Values(metrics, value).ToSql()
		if err != nil {
			return fmt.Errorf("error building metrics insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
		}

		var execResult rdb.ExecResult
		if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
			return fmt.Errorf("error inserting metrics: %v", err)
		}
		if execResult.RowsAffected() != 1 {
			return errors.New("error inserting metrics: no rows inserted")
		}

		return nil
	}

	sql, sqlArgs, err = view.rdbHandle.StmtBuilder.Update(
		"view_validator_stats",
	).Set(
		"value", value,
	).Where(
		"metrics = ?", metrics,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building metrics update sql: %v", err)
	}

	var execResult rdb.ExecResult
	if execResult, err = view.rdbHandle.Exec(sql, sqlArgs...); err != nil {
		return fmt.Errorf("error updating metrics: %v", err)
	}
	if execResult.RowsAffected() != 1 {
		return errors.New("error updating metrics: no rows updated")
	}

	return nil
}

func (view *ValidatorStats) FindBy(metrics string) (string, error) {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.Select(
		"value",
	).From(
		"view_validator_stats",
	).Where(
		"metrics = ?", metrics,
	).ToSql()
	if err != nil {
		return "", fmt.Errorf("error preparing metrics selection SQL: %v", err)
	}

	var value string
	if err := view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&value); err != nil {
		if err == rdb.ErrNoRows {
			return "", nil
		}
		return "", fmt.Errorf("error getting metrics: %v", err)
	}

	return value, nil
}

type ValidatorStatsRow struct {
	Metrics string
	Value   string
}
