package rdbstatusstore

type RDbStatusStore interface {
	// UpdateLastHandledEventHeight update last indexed block height
	UpdateLastIndexedBlockHeight(height int64) error
	// GetLastHandledEventHeight returns the last indexed block height, nil if no block has been indexed
	GetLastIndexedBlockHeight() (int64, error)
}
