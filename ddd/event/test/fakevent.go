package ddd_event_test

type FakeEvent struct{}

func NewFakeEvent() *FakeEvent {
	return &FakeEvent{}
}

func (evt *FakeEvent) MaybeSeq() *int64     { return nil }
func (evt *FakeEvent) SetSeq(_seq int64)    {}
func (evt *FakeEvent) Id() string           { return "" }
func (evt *FakeEvent) Name() string         { return "" }
func (evt *FakeEvent) Version() int         { return 0 }
func (evt *FakeEvent) Payload() interface{} { return nil }
func (evt *FakeEvent) String() string       { return "FakeEvent" }
