package view

import (
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockValidatorActivitiesView struct {
	mock.Mock
}

func NewMockValidatorActivitiesView(_ *rdb.Handle) ValidatorActivities {
	return &MockValidatorActivitiesView{}
}
func (validatorActivitiesView *MockValidatorActivitiesView) Insert(validatorActivity *ValidatorActivityRow) error {
	mockArgs := validatorActivitiesView.Called(validatorActivity)
	return mockArgs.Error(0)
}

func (validatorActivitiesView *MockValidatorActivitiesView) InsertAll(validatorActivities []ValidatorActivityRow) error {
	mockArgs := validatorActivitiesView.Called(validatorActivities)
	return mockArgs.Error(0)
}
func (validatorActivitiesView *MockValidatorActivitiesView) List(
	filter ValidatorActivitiesListFilter,
	order ValidatorActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]ValidatorActivityRow, *pagination_interface.PaginationResult, error) {
	mockArgs := validatorActivitiesView.Called(filter, order, pagination)
	result0, _ := mockArgs.Get(0).([]ValidatorActivityRow)
	result1, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}
