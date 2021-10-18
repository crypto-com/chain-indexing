package view

import (
	testify_mock "github.com/stretchr/testify/mock"
)

type MockIBCChannelMessageTotalView struct {
	testify_mock.Mock
}

func (totalView *MockIBCChannelMessageTotalView) Increment(identity string, total int64) error {
	mockArgs := totalView.Called(identity, total)
	return mockArgs.Error(0)
}

func (totalView *MockIBCChannelMessageTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := totalView.Called(identities)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}
