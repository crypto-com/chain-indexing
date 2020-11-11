package test

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
)

type FakeEventStore struct{}

func NewFakeEventStore() *FakeEventStore {
	return &FakeEventStore{}
}

func (manager *FakeEventStore) GetLatestHeight() (*int64, error) {
	return nil, nil
}

func (manager *FakeEventStore) GetAllByHeight(seq int64) ([]entity_event.Event, error) {
	return []entity_event.Event{NewFakeEvent()}, nil
}

func (manager *FakeEventStore) Insert(evt entity_event.Event) error {
	return nil
}

func (manager *FakeEventStore) InsertAll(evts []entity_event.Event) error {
	return nil
}
