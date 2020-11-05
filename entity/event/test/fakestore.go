package entity_event_test

import (
	"github.com/crypto-com/chainindex/entity/event"
)

type FakeManager struct{}

func NewFakeManager() *FakeManager {
	return &FakeManager{}
}

func (manager *FakeManager) GetLatestEventSeq() *int64 {
	return nil
}

func (manager *FakeManager) GetEventBySeq(seq int64) (event.Event, error) {
	return NewFakeEvent(), nil
}

func (manager *FakeManager) InsertEvent(evt event.Event) error {
	return nil
}
