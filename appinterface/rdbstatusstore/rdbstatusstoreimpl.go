package rdbstatusstore

import (
	"errors"
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/rdb"
)

const DEFAULT_TABLE = "service_status"

// Status table should have the following schema
// | Field                     | Data Type | Constraint  |
// | ------------------------- | --------- | ----------- |
// | last_indexed_block_height | INT64     | NOT NULL    |

// StatusStore implemented using relational database
type RDbStatusStoreImpl struct {
	rdbHandle *rdb.Handle

	table string
}

type LatestStatus struct {
	LastIndexedBlockHeight int64 `json:"last_indexed_block_height"`
}

func NewRDbStatusStoreImpl(handle *rdb.Handle) *RDbStatusStoreImpl {
	return &RDbStatusStoreImpl{
		rdbHandle: handle,

		table: DEFAULT_TABLE,
	}
}

// InitLatestStatus creates one row for initial latest status
func (impl *RDbStatusStoreImpl) InitLatestStatus() error {
	// Initial latest status defines here
	initialLastIndexedBlockHeight := 1

	// Insert initial latest status to the row
	sql, args, err := impl.rdbHandle.StmtBuilder.Insert(
		impl.table,
	).Columns(
		"last_indexed_block_height",
	).Values(initialLastIndexedBlockHeight).ToSql()

	if err != nil {
		return fmt.Errorf("error building getting row count selection SQL: %v", err)
	}

	execResult, err := impl.rdbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error exectuing initial latest status insertion SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing initial latest status insertion SQL: no rows inserted")
	}

	return nil
}

// CheckLatestStatusExist checks if the row exists, if not, call InitLatestStatus
func (impl *RDbStatusStoreImpl) CheckLatestStatusExist() error {
	sql, args, err := impl.rdbHandle.StmtBuilder.Select(
		"COUNT(*)",
	).From(impl.table).ToSql()
	if err != nil {
		return fmt.Errorf("error building getting row count selection SQL: %v", err)
	}

	var count int64
	if err := impl.rdbHandle.QueryRow(sql, args...).Scan(&count); err != nil {
		return fmt.Errorf("error querying row latest_status count: %v", err)
	}

	if count > 0 {
		return nil
	} else {
		return impl.InitLatestStatus()
	}
}

func (impl *RDbStatusStoreImpl) GetLastIndexedBlockHeight() (int64, error) {
	if err := impl.CheckLatestStatusExist(); err != nil {
		return 0, fmt.Errorf("error checking the latest_status row exist: %v", err)
	}

	sql, args, err := impl.rdbHandle.StmtBuilder.Select(
		"last_indexed_block_height",
	).From(impl.table).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building last indexed block height selection SQL: %v", err)
	}

	var lastIndexedBlockHeight int64
	if err := impl.rdbHandle.QueryRow(sql, args...).Scan(&lastIndexedBlockHeight); err != nil {
		return 0, fmt.Errorf("error querying last indexed block height: %v", err)
	}

	return lastIndexedBlockHeight, nil
}

func (impl *RDbStatusStoreImpl) UpdateLastIndexedBlockHeight(height int64) error {
	if err := impl.CheckLatestStatusExist(); err != nil {
		return fmt.Errorf("error checking the latest_status row exist: %v", err)
	}

	// TODO: pass the rdbHandle with params for transaction support
	sql, args, err := impl.rdbHandle.StmtBuilder.Update(
		impl.table,
	).Set(
		"last_indexed_block_height", height,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building last indexed block height update SQL: %v", err)
	}

	execResult, err := impl.rdbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing last indexed block height update SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing last indexed block height update SQL: no rows updated")
	}

	return nil
}
