package rdbstatusstore

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DEFAULT_TABLE = "service_status"

// Status table should have the following schema
// | Field                     | Data Type | Constraint  |
// | ------------------------- | --------- | ----------- |
// | last_indexed_block_height | INT64     | NOT NULL    |

// status store implemented using relational database
type RDbStatusStore struct {
	selectRDbHandle *rdb.Handle

	table string
}

type LatestStatus struct {
	LastIndexedBlockHeight int64 `json:"last_indexed_block_height"`
}

func NewRDbStatusStore(rdbHandle *rdb.Handle) *RDbStatusStore {
	return &RDbStatusStore{
		selectRDbHandle: rdbHandle,

		table: DEFAULT_TABLE,
	}
}

// init initializes status store DB when it is first time running
func (impl *RDbStatusStore) init() error {
	var err error

	var exist bool
	if exist, err = impl.isStatusRowExist(); err != nil {
		return fmt.Errorf("error checking status row existence: %v", err)
	}

	if !exist {
		if err = impl.initStatusRow(); err != nil {
			return fmt.Errorf("error initializing status row: %v", err)
		}
	}

	return nil
}

// InitLatestStatus creates one row for initial latest status
func (impl *RDbStatusStore) initStatusRow() error {
	// Insert initial latest status to the row
	sql, args, err := impl.selectRDbHandle.StmtBuilder.Insert(
		impl.table,
	).Columns(
		"last_indexed_block_height",
	).Values(nil).ToSql()
	if err != nil {
		return fmt.Errorf("error building getting row count insertion SQL: %v", err)
	}

	execResult, err := impl.selectRDbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error inserting latest status SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing initial latest status insertion SQL: no rows inserted")
	}

	return nil
}

// isStatusRowExist returns true if the row exists
func (impl *RDbStatusStore) isStatusRowExist() (bool, error) {
	sql, args, err := impl.selectRDbHandle.StmtBuilder.Select(
		"COUNT(*)",
	).From(impl.table).ToSql()
	if err != nil {
		return false, fmt.Errorf("error building status row count selection SQL: %v", err)
	}

	var count int64
	if err := impl.selectRDbHandle.QueryRow(sql, args...).Scan(&count); err != nil {
		return false, fmt.Errorf("error querying status row count: %v", err)
	}

	return count > 0, nil
}

// Return last indexed block height, nil if no block has been indexed
func (impl *RDbStatusStore) GetLastIndexedBlockHeight() (*int64, error) {
	// lazy init
	if err := impl.init(); err != nil {
		return nil, fmt.Errorf("error initializing status store: %v", err)
	}

	sql, args, err := impl.selectRDbHandle.StmtBuilder.Select(
		"last_indexed_block_height",
	).From(impl.table).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building last indexed block height selection SQL: %v", err)
	}

	var lastIndexedBlockHeight *int64
	if err := impl.selectRDbHandle.QueryRow(sql, args...).Scan(&lastIndexedBlockHeight); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, nil
		} else {
			return nil, fmt.Errorf("error querying last indexed block height: %v", err)
		}
	}

	return lastIndexedBlockHeight, nil
}

func (impl *RDbStatusStore) UpdateLastIndexedBlockHeightWithRDbHandle(rdbHandle *rdb.Handle, height int64) error {
	// lazy init
	if err := impl.init(); err != nil {
		return fmt.Errorf("error initializing status store: %v", err)
	}

	sql, args, err := rdbHandle.StmtBuilder.Update(
		impl.table,
	).Set(
		"last_indexed_block_height", height,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building last indexed block height update SQL: %v", err)
	}

	execResult, err := rdbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing last indexed block height update SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing last indexed block height update SQL: no rows updated")
	}

	return nil
}
