package ddd_event_test

import (
	"github.com/crypto-com/chainindex/ddd/event"
	"github.com/stretchr/testify/mock"
)

type MockManager struct {
	mock.Mock
}

func NewMockManager() *MockManager {
	return &MockManager{}
}

func (manager *MockManager) GetLatestEventSeq() *int64 {
	mockArgs := manager.Called()

	return mockArgs.Get(0).(*int64)
}

func (manager *MockManager) GetEventBySeq(seq int64) (event.Event, error) {
	mockArgs := manager.Called(seq)

	return mockArgs.Get(0).(event.Event), mockArgs.Error(1)
}

func (manager *MockManager) InsertEvent(evt event.Event) error {
	mockArgs := manager.Called(evt)

	return mockArgs.Error(0)
}
