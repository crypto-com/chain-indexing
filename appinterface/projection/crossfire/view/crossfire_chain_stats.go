package view

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type CrossfireChainStats struct {
	rdbHandle *rdb.Handle
}

const CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME = "view_crossfire_chain_stats"

func NewCrossfireChainStats(handle *rdb.Handle) *CrossfireChainStats {
	return &CrossfireChainStats{
		handle,
	}
}

func (view *CrossfireChainStats) Set(metric string, value int64) error {
	// UPSERT STATEMENT
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.
		Insert(CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME).
		Columns("metric", "value").
		Values(metric, value).
		Suffix("ON CONFLICT (metric) DO UPDATE SET value = EXCLUDED.value").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building value insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = view.rdbHandle.Exec(sql, sqlArgs...)

	if err != nil {
		return fmt.Errorf("error inserting value: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}


func (view *CrossfireChainStats) FindBy(metric string) (int64, error) {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.Select(
		"value",
	).From(
		CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME,
	).Where(
		"metric = ?", metric,
	).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error preparing key selection SQL: %v", err)
	}

	var value int64
	if err := view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&value); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("error getting key: %v", err)
	}

	return value, nil
}

func (view *CrossfireChainStats) IncrementOne(key string) error {
	value, err := view.FindBy(key)
	if err != nil {
		return fmt.Errorf("error getting value for %v: %v", key, err)
	}

	err = view.Set(key, value + 1)
	if err != nil {
		return fmt.Errorf("error setting increment for %v: %v", key, err)
	}

	return nil
}
