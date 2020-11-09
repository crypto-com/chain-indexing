package rdbeventstore

import (
	"errors"
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/internal/primptr"
)

const DEFAULT_TABLE = "events"

// Table should have the following schema
// | Field                     | Data Type | Constraint  |
// | ------------------------- | --------- | ----------- |
// | id 					   | VARCHAR   | PRIMARY KEY |
// | height                    | INT64     | NOT NULL    |
// | name                      | VARCHAR   | NOT NULL    |
// | version                   | INT64     | NOT NULL    |
// | payload				   | JSONB	   | NOT NULL    |
type RDbEventStore struct {
	rdbHandle *rdb.Handle
}

func NewRDbEventStore(handle *rdb.Handle) *RDbEventStore {
	return &RDbEventStore{
		rdbHandle: handle,
	}
}

// GetLatestEventHeight returns latest event height, nil if no event is stored
func (store *RDbEventStore) GetLatestHeight() (*int64, error) {
	sql, args, err := store.rdbHandle.StmtBuilder.Select(
		"MAX(height)",
	).From(
		DEFAULT_TABLE,
	).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get latest event height selection SQL: %v", err)
	}

	var latestEventHeight int64
	if err := store.rdbHandle.QueryRow(sql, args...).Scan(&latestEventHeight); err != nil {
		if err == rdb.ErrNoRows {
			return nil, nil
		} else {
			return nil, fmt.Errorf("error executing get latest event height selection SQL: %v", err)
		}
	}

	return primptr.Int64(latestEventHeight), nil
}

func (store *RDbEventStore) GetAllByHeight(height int64) ([]entity_event.Event, error) {
	sql, args, err := store.rdbHandle.StmtBuilder.Select(
		"id", "height", "name", "version", "payload",
	).From(
		DEFAULT_TABLE,
	).Where("height = ?", height).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building get all events by height selection SQL: %v", err)
	}

	var events = make([]entity_event.Event, 0)
	rows, err := store.rdbHandle.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing get all events by height selection SQL: %v", err)
	}
	for rows.Next() {
		var (
			id      string
			height  int64
			name    string
			version int64
			payload string
		)

		if err := rows.Scan(&id, &height, &name, &version, &payload); err != nil {
			if err == rdb.ErrNoRows {
				return nil, nil
			} else {
				return nil, fmt.Errorf("error error executing get each event by height selection SQL: %v", err)
			}
		}
		
		// TODO: Missing event json deserializer
		//evt := struct{
		//	id string
		//	height  int64
		//	name    string
		//	version int64
		//	payload string
		//}{
		//	id,
		//	height,
		//	name,
		//	version,
		//	payload,
		//}
		//append(events, evt)
	}

	return events, nil
}

func (store *RDbEventStore) Insert(evt entity_event.Event) error {
	// Insert the incoming event
	sql, args, err := store.rdbHandle.StmtBuilder.Insert(
		DEFAULT_TABLE,
	).Columns(
		"id", "height", "name", "version", "payload",
	).Values(
		evt.Id(),
		evt.Height(),
		evt.Name(),
		evt.Version(),
		evt.String(),
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building event insertion SQL: %v", err)
	}

	execResult, err := store.rdbHandle.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error exectuing event insertion SQL: %v", err)
	}
	if execResult.RowsAffected() == 0 {
		return errors.New("error executing event insertion SQL: no rows inserted")
	}

	return nil
}

// InsertAll insert all events into store. It will rollback when the insert fails at any point.
func (store *RDbEventStore) InsertAll(evts []entity_event.Event) error {
	for _, evt := range evts {
		if err := store.Insert(evt); err != nil {
			return errors.New("error executing events batch insertion SQL")
		}
	}

	return nil
}
