package view

import (
	testify_mock "github.com/stretchr/testify/mock"
)

type MockTokensTotalView struct {
	testify_mock.Mock
}

func (totalView *MockTokensTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := totalView.Called(identities, total)
	return mockArgs.Error(0)
}

func (totalView *MockTokensTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := totalView.Called(identities, total)
	return mockArgs.Error(0)
}

func (totalView *MockTokensTotalView) FindBy(identity string) (int64, error) {
	mockArgs := totalView.Called(identity)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}
