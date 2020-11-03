package projectiontest

import (
	"github.com/crypto-com/chainindex/ddd"
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
func (projection *FakeProjection) GetLastHandledEventId() string {
	return ""
}
func (projection *FakeProjection) OnInit() error {
	return nil
}
func (projection *FakeProjection) HandleEvent(event ddd.Event) error {
	return nil
}
