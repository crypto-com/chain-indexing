package polling

import (
	"fmt"
	"errors"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type Status struct {
	rdb *rdb.Handle
}

func NewStatus(handle *rdb.Handle) *Status {
	return &Status{
		handle,
	}
}

func (view *Status) Insert(statusid string, statusvalue string) error {
	var err error

	var sql string
	sql, _, err = view.rdb.StmtBuilder.Insert(
		"view_status",
	).Columns(
		"key",
		"value",
	).Values("?", "?").Suffix("ON CONFLICT(key) DO UPDATE SET VALUE=EXCLUDED.value").ToSql()

	if err != nil {
		return fmt.Errorf("error building blocks insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}
	result, err := view.rdb.Exec(sql, statusid, statusvalue)
	if err != nil {
		return fmt.Errorf("error inserting view_status into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting view_status into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}
func (view *Status) FindBy(statusid string) (string, error) {
	var err error

	selectStmtBuilder := view.rdb.StmtBuilder.Select(
		"value",
	).From("view_status").Where("key = ?", statusid)

	sql, sqlArgs, err := selectStmtBuilder.ToSql()

	if err != nil {
		return "", fmt.Errorf("error building status selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var found string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&found,
	); err != nil {
		if errors.Is(err,rdb.ErrNoRows) {
			return "", rdb.ErrNoRows
		}
		return "", fmt.Errorf("error scanning block row: %v: %w", err, rdb.ErrQuery)
	}

	return found, nil
}
