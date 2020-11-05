package ddd_projection_test

import (
	"github.com/crypto-com/chainindex/ddd/event"
)

type FakeProjection struct{}

func NewFakeProjection() *FakeProjection {
	return &FakeProjection{}
}

func (projection *FakeProjection) Id() string {
	return "FakeProjection"
}
func (projection *FakeProjection) GetEventsToListen() []string {
	return []string{}
}
func (projection *FakeProjection) GetLastHandledEventSeq() *int64 {
	return nil
}
func (projection *FakeProjection) OnInit() error {
	return nil
}
func (projection *FakeProjection) HandleEvent(event event.Event) error {
	return nil
}
