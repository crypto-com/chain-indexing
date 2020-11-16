package test

import "github.com/google/uuid"

type FakeEvent struct{}

func NewFakeEvent() *FakeEvent {
	return &FakeEvent{}
}

func (event *FakeEvent) Height() int64 { return 1 }
func (event *FakeEvent) Name() string  { return "FakeEvent" }
func (event *FakeEvent) Version() int  { return 0 }
func (event *FakeEvent) UUID() string  { return uuid.New().String() }
func (event *FakeEvent) ToJSON() (string, error) {
	return "\"" + event.UUID() + "\"", nil
}
func (event *FakeEvent) String() string { return "FakeEvent" }
