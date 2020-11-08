package rdbbase

import "github.com/crypto-com/chainindex/appinterface/rdb"

// Projection store is a projection handling record store interface. An interface allows easy
// testing
type RDbStore interface {
	// UpdateLastHandledEventHeight update last handled event height of projection id to provided
	// height
	UpdateLastHandledEventHeight(rdbHandle *rdb.Handle, projectionId string, height int64) error
	// GetLastHandledEventHeight returns the last handled event height, nil if no event has been
	// handled
	GetLastHandledEventHeight(rdbHandle *rdb.Handle, projectionId string) (*int64, error)
}
