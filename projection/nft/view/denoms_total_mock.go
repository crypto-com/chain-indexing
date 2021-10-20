package view

import (
	testify_mock "github.com/stretchr/testify/mock"
)

type MockDenomsTotalView struct {
	testify_mock.Mock
}

func (totalView *MockDenomsTotalView) Increment(identity string, total int64) error {
	mockArgs := totalView.Called(identity, total)
	return mockArgs.Error(0)
}

func (totalView *MockDenomsTotalView) FindBy(identity string) (int64, error) {
	mockArgs := totalView.Called(identity)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}
