package test

import (
	"github.com/stretchr/testify/mock"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
)

type MockProjection struct {
	mock.Mock
}

func NewMockProjection() *MockProjection {
	return &MockProjection{}
}

func (projection *MockProjection) Id() string {
	mockArgs := projection.Called()

	return mockArgs.String(0)
}

func (projection *MockProjection) GetEventsToListen() []string {
	mockArgs := projection.Called()

	return mockArgs.Get(0).([]string)
}

func (projection *MockProjection) GetLastHandledEventHeight() (*int64, error) {
	mockArgs := projection.Called()

	return mockArgs.Get(0).(*int64), mockArgs.Error(1)
}

func (projection *MockProjection) OnInit() error {
	mockArgs := projection.Called()

	return mockArgs.Error(0)
}

func (projection *MockProjection) HandleEvents(height int64, events []entity_event.Event) error {
	mockArgs := projection.Called(height, events)

	return mockArgs.Error(0)
}
