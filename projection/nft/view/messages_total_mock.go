package view

import (
	testify_mock "github.com/stretchr/testify/mock"
)

type MockMessagesTotalView struct {
	testify_mock.Mock
}

func (totalView *MockMessagesTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := totalView.Called(identities, total)
	return mockArgs.Error(0)
}

func (totalView *MockMessagesTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := totalView.Called(identities)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}
