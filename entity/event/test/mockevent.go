package test

import "github.com/stretchr/testify/mock"

type MockEvent struct {
	mock.Mock
}

func NewMockEvent() *MockEvent {
	return &MockEvent{}
}

func (event *MockEvent) Height() int64 {
	mockArgs := event.Called()
	return mockArgs.Get(0).(int64)
}
func (event *MockEvent) Name() string {
	mockArgs := event.Called()
	return mockArgs.String(0)
}
func (event *MockEvent) Version() int {
	mockArgs := event.Called()
	return mockArgs.Int(0)
}
func (event *MockEvent) UUID() string {
	mockArgs := event.Called()
	return mockArgs.String(0)
}
func (event *MockEvent) ToJSON() (string, error) {
	mockArgs := event.Called()
	return mockArgs.String(0), mockArgs.Error(1)
}
func (event *MockEvent) String() string {
	mockArgs := event.Called()
	return mockArgs.String(0)
}
