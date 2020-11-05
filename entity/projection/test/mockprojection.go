package ddd_projection_test

import (
	dddevent "github.com/crypto-com/chainindex/ddd/event"
	"github.com/stretchr/testify/mock"
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

func (projection *MockProjection) GetLastHandledEventSeq() *int64 {
	mockArgs := projection.Called()

	return mockArgs.Get(0).(*int64)
}

func (projection *MockProjection) OnInit() error {
	mockArgs := projection.Called()

	return mockArgs.Error(0)
}

func (projection *MockProjection) HandleEvent(event dddevent.Event) error {
	mockArgs := projection.Called(event)

	return mockArgs.Error(0)
}
