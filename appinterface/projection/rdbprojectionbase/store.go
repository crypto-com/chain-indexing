package rdbprojectionbase

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/primptr"
)

// Store is an implementation of the RDbStore
type Store struct {
	table string
}

func NewStore(table string) *Store {
	return &Store{
		table,
	}
}

// UpdateLastHandledEventHeight update last handled event height of projection id to provided
// height
func (impl *Store) UpdateLastHandledEventHeight(rdbHandle *rdb.Handle, projectionId string, height int64) error {
	lastHandledEventHeight, err := impl.GetLastHandledEventHeight(rdbHandle, projectionId)
	if err != nil {
		return fmt.Errorf("error checking projection record existence: %v", err)
	}

	if lastHandledEventHeight == nil {
		// Insert projection record with the updated height
		sql, args, sqlErr := rdbHandle.StmtBuilder.Insert(
			impl.table,
		).Columns(
			"id", "last_handled_event_height",
		).Values(projectionId, height).ToSql()
		if sqlErr != nil {
			return fmt.Errorf("error building last handled event height insertion SQL: %v", sqlErr)
		}

		execResult, sqlErr := rdbHandle.Exec(sql, args...)
		if sqlErr != nil {
			return fmt.Errorf("error exectuing last handled event height insertion SQL: %v", sqlErr)
		}
		if execResult.RowsAffected() == 0 {
			return errors.New("error executing last handled event height insertion SQL: no rows inserted")
		}

		return nil
	}

	// Update existing projection record with the updated height
	sql, args, err := rdbHandle.StmtBuilder.Update(
		impl.table,
	).SetMap(map[string]interface{}{
		"id":                        projectionId,
		"last_handled_event_height": height,
	}).Where(
		"id = ?", projectionId,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building last handled event height update SQL: %v", err)
	}

	execResult, err := rdbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing last handled event height update SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing last handled event height update SQL: no rows updated")
	}

	return nil
}

// GetLastHandledEventHeight returns the last handled event height, nil if no event has been
// handled
func (impl *Store) GetLastHandledEventHeight(rdbHandle *rdb.Handle, projectionId string) (*int64, error) {
	sql, args, err := rdbHandle.StmtBuilder.Select(
		"last_handled_event_height",
	).From(
		impl.table,
	).Where("id = ?", projectionId).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building last handled event height selection SQL: %v", err)
	}

	var lastHandledEventHeight int64
	if err := rdbHandle.QueryRow(sql, args...).Scan(&lastHandledEventHeight); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, nil
		} else {
			return nil, fmt.Errorf("error building last handled event height selection SQL: %v", err)
		}
	}

	return primptr.Int64(lastHandledEventHeight), nil
}
