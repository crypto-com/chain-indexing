package ddd

type Event interface {
	Name() string
	Version() int

	// Payload returns the event attached payload that needs to be stored
	Payload() interface{}
}
