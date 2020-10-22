package event

type Event struct {
	Name    string
	Version uint64
	Payload interface{}
}

// NewEvent creates a new event
func NewEvent(Name string, Version uint64, Payload interface{}) Event {
	return Event{
		Name,
		Version,
		Payload,
	}
}

// Store function will run store current event
func (evt Event) Store() error {
	// TODO: store the event and get ready for rollback this event store procedure
	// if a sequence of event storing failed in the middle
	return nil
}
