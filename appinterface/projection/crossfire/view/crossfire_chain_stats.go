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

func (crossfireChainStatsView *CrossfireChainStats) Set(metric string, value string) error {
	// UPSERT STATEMENT
	sql, sqlArgs, err := crossfireChainStatsView.rdbHandle.StmtBuilder.
		Insert(CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME).
		Columns("metric", "value").
		Values(metric, value).
		Suffix("ON CONFLICT (metric) DO UPDATE SET value = EXCLUDED.value").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building value insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = crossfireChainStatsView.rdbHandle.Exec(sql, sqlArgs...)

	if err != nil {
		return fmt.Errorf("error inserting value: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

func (crossfireChainStatsView *CrossfireChainStats) Increment(metric string, value string) error {
	// Postgres UPSERT statement

	getMetric, err := crossfireChainStatsView.FindBy(metric)
	if err != nil {
		return fmt.Errorf("error getting value: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	if len(getMetric) == 0 {
		return fmt.Errorf("Got empty value!")
	}
	sql, sqlArgs, err := crossfireChainStatsView.rdbHandle.StmtBuilder.
		Insert(CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME+" AS totals").
		Columns("metric", "value").
		Values(metric, value).
		Suffix("ON CONFLICT (metric) DO UPDATE SET value = totals.value + EXCLUDED.value").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building value insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = crossfireChainStatsView.rdbHandle.Exec(sql, sqlArgs...)

	if err != nil {
		return fmt.Errorf("error inserting value: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

func (crossfireChainStatsView *CrossfireChainStats) FindBy(metric string) (string, error) {
	sql, sqlArgs, err := crossfireChainStatsView.rdbHandle.StmtBuilder.Select(
		"value",
	).From(
		CROSSFIRE_CHAIN_STATS_VIEW_TABLENAME,
	).Where(
		"metric = ?", metric,
	).ToSql()
	if err != nil {
		return "", fmt.Errorf("error preparing metric selection SQL: %v", err)
	}

	var value string
	if err := crossfireChainStatsView.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&value); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("error getting metric: %v", err)
	}

	return value, nil
}
