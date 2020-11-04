package event

type Manager interface {
	// Get latest Event sequence
	GetLatestEventSeq() *int64

	GetEventBySeq(seq int64) (Event, error)

	// Insert an event and assign an Id to it
	InsertEvent(evt Event) error
}
