package event

type Store interface {
	// GetLatestEventHeight returns latest event height, nil if no event is stored
	GetLatestHeight() (*int64, error)

	GetAllByHeight(height int64) ([]Event, error)

	Insert(evt Event) error

	// InsertAll insert all events into store. It will rollback when the insert fails at any point.
	InsertAll(evt []Event) error
}
