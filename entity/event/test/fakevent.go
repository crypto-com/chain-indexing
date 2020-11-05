package entity_event_test

type FakeEvent struct{}

func NewFakeEvent() *FakeEvent {
	return &FakeEvent{}
}

func (evt *FakeEvent) Height() int64  { return 0 }
func (evt *FakeEvent) Name() string   { return "" }
func (evt *FakeEvent) Version() int   { return 0 }
func (evt *FakeEvent) Id() string     { return "" }
func (evt *FakeEvent) String() string { return "FakeEvent" }
