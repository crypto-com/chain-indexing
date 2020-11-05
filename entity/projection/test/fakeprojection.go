package ddd_projection_test

import (
	"github.com/crypto-com/chainindex/entity/event"
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
func (projection *FakeProjection) GetLastHandledEventHeight() *int64 {
	return nil
}
func (projection *FakeProjection) OnInit() error {
	return nil
}
func (projection *FakeProjection) HandleEvents(evts []event.Event) error {
	return nil
}
