package event

type Event interface {
	Category() string
	CategorySequence() int64

	Name() string
	Version() int
	Payload() interface{}

	// Unique Id that is assigned on event creation
	Id() string

	// String function returns the stringified event
	String() string
}
