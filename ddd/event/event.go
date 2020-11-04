package event

type Event interface {
	// Returns event sequence, it is possible that event has no seq on creation and it could be
	// assigned by event managers or until DB persistance
	MaybeSeq() *int64
	SetSeq(seq int64)
	// Unique Id that is assigned on event creation
	Id() string

	Name() string
	Version() int

	Payload() interface{}

	String() string
}
