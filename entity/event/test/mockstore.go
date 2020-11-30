package test

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/stretchr/testify/mock"
)

type MockEventStore struct {
	mock.Mock
}

func NewMockEventStore() *MockEventStore {
	return &MockEventStore{}
}

func (manager *MockEventStore) GetLatestHeight() (*int64, error) {
	mockArgs := manager.Called()

	return mockArgs.Get(0).(*int64), mockArgs.Error(1)
}

func (manager *MockEventStore) GetAllByHeight(height int64) ([]entity_event.Event, error) {
	mockArgs := manager.Called(height)

	return mockArgs.Get(0).([]entity_event.Event), mockArgs.Error(1)
}

func (manager *MockEventStore) Insert(evt entity_event.Event) error {
	mockArgs := manager.Called(evt)

	return mockArgs.Error(0)
}

func (manager *MockEventStore) InsertAll(evts []entity_event.Event) error {
	mockArgs := manager.Called(evts)

	return mockArgs.Error(0)
}

func (manager *MockEventStore) EnsurePartitionTableExists(height int64) error {
	mockArgs := manager.Called(height)
	return mockArgs.Error(0)
}
