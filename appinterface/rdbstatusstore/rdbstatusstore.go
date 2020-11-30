package rdbstatusstore

import "github.com/crypto-com/chain-indexing/appinterface/rdb"

type RDbStatusStore interface {
	// UpdateLastHandledEventHeight update last indexed block height
	UpdateLastIndexedBlockHeight(rdbHandle *rdb.Handle, height int64) error
	// GetLastHandledEventHeight returns the last indexed block height, nil if no block has been indexed
	GetLastIndexedBlockHeight() (*int64, error)
}
