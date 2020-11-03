package ddd

type Event interface {
	Name() string
	Version() int

	Payload() interface{}
}
