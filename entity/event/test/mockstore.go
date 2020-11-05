package entity_event_test

import (
	"github.com/stretchr/testify/mock"

	"github.com/crypto-com/chainindex/entity/event"
)

type MockStore struct {
	mock.Mock
}

func NewMockStore() *MockStore {
	return &MockStore{}
}

func (manager *MockStore) GetLatestHeight() *int64 {
	mockArgs := manager.Called()

	return mockArgs.Get(0).(*int64)
}

func (manager *MockStore) GetByHeight(height int64) ([]event.Event, error) {
	mockArgs := manager.Called(height)

	return mockArgs.Get(0).([]event.Event), mockArgs.Error(1)
}

func (manager *MockStore) Insert(evt event.Event) error {
	mockArgs := manager.Called(evt)

	return mockArgs.Error(0)
}

func (manager *MockStore) InsertAll(evts []event.Event) error {
	mockArgs := manager.Called(evts)

	return mockArgs.Error(0)
}
