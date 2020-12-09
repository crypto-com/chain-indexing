package rdbprojectionbase

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

const DEFAULT_TABLE = "projections"

// Table should have the following schema
// | Field                     | Data Type | Constraint  |
// | ------------------------- | --------- | ----------- |
// | id                        | VARCHAR   | PRIMARY KEY |
// | last_handled_event_height | INT64     | NOT NULL    |

// Base is a bas for projection which keeps track of last handled event height using relational
// database. It implements Id() and GetLastHandledEventHeight() of projection interface.
type Base struct {
	rdbHandle *rdb.Handle
	store     *Store

	projectionId string
}

// Create a new Base using table name in the RDb to keep the projection handling records
func NewRDbBase(rdbHandle *rdb.Handle, projectionId string) *Base {
	return &Base{
		rdbHandle: rdbHandle,
		store:     NewStore(DEFAULT_TABLE),

		projectionId: projectionId,
	}
}

// Create a new Base with customize table name in the RDb to keep the projection handling records
func NewRDbBaseWithTable(rdbHandle *rdb.Handle, projectionId string, table string) *Base {
	return &Base{
		rdbHandle: rdbHandle,
		store:     NewStore(table),

		projectionId: projectionId,
	}
}

// Implements projection.Id()
func (base *Base) Id() string {
	return base.projectionId
}

func (base *Base) UpdateLastHandledEventHeight(rdbHandle *rdb.Handle, height int64) error {
	if err := base.store.UpdateLastHandledEventHeight(rdbHandle, base.projectionId, height); err != nil {
		return err
	}
	return nil
}

// Implements projection.GetLastHandledEventHeight()
func (base *Base) GetLastHandledEventHeight() (*int64, error) {
	return base.store.GetLastHandledEventHeight(base.rdbHandle, base.projectionId)
}
