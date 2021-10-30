package view

import (
	testify_mock "github.com/stretchr/testify/mock"
)

type MockValidatorsView struct {
	testify_mock.Mock
}

func NewMockValidatorsView() Validators {
	return &MockValidatorsView{}
}

func (validatorsView *MockValidatorsView) Upsert(row *ValidatorRow) error {
	mockArgs := validatorsView.Called(row)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) Insert(row *ValidatorRow) error {
	mockArgs := validatorsView.Called(row)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) Update(row *ValidatorRow) error {
	mockArgs := validatorsView.Called(row)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) ListAll() (
	[]ValidatorRow,
	error,
) {
	mockArgs := validatorsView.Called()
	result1, _ := mockArgs.Get(0).([]ValidatorRow)
	return result1, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) FindLastBy(
	identity ValidatorIdentity,
) (
	*ValidatorRow,
	error,
) {
	mockArgs := validatorsView.Called(identity)
	result1, _ := mockArgs.Get(0).(*ValidatorRow)
	return result1, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) Count() (
	int64,
	error,
) {
	mockArgs := validatorsView.Called()
	result1, _ := mockArgs.Get(0).(int64)
	return result1, mockArgs.Error(1)
}
