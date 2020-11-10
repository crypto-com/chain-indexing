package test

import "github.com/stretchr/testify/mock"

type MockEvent struct {
	mock.Mock
}

func NewMockEvent() *MockEvent {
	return &MockEvent{}
}

func (evt *MockEvent) Height() int64 {
	mockArgs := evt.Called()
	return mockArgs.Get(0).(int64)
}
func (evt *MockEvent) Name() string {
	mockArgs := evt.Called()
	return mockArgs.String(0)
}
func (evt *MockEvent) Version() int {
	mockArgs := evt.Called()
	return mockArgs.Int(0)
}
func (evt *MockEvent) Id() string {
	mockArgs := evt.Called()
	return mockArgs.String(0)
}
func (evt *MockEvent) String() string {
	mockArgs := evt.Called()
	return mockArgs.String(0)
}
