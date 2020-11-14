package event

// Event must be JSON-completed. i.e. All important fields must be encoded to JSON
type Event interface {
	Height() int64

	Name() string
	Version() int

	// Unique Id that is assigned on event creation
	Id() string

	// Encode the event into JSON string payload
	ToJSON() (string, error)

	// String function returns the stringified event
	String() string
}
