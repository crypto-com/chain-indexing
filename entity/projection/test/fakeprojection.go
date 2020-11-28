package test

import entity_event "github.com/crypto-com/chain-indexing/entity/event"

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
func (projection *FakeProjection) GetLastHandledEventHeight() (*int64, error) {
	return nil, nil
}
func (projection *FakeProjection) OnInit() error {
	return nil
}
func (projection *FakeProjection) HandleEvents(_ int64, _ []entity_event.Event) error {
	return nil
}
