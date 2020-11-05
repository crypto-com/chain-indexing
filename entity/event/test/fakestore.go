package entity_event_test

import (
	"github.com/crypto-com/chainindex/entity/event"
)

type FakeStore struct{}

func NewFakeStore() *FakeStore {
	return &FakeStore{}
}

func (manager *FakeStore) GetLatestHeight() *int64 {
	return nil
}

func (manager *FakeStore) GetByHeight(seq int64) (event.Event, error) {
	return NewFakeEvent(), nil
}

func (manager *FakeStore) Insert(evt event.Event) error {
	return nil
}

func (manager *FakeStore) InsertAll(evts []event.Event) error {
	return nil
}
