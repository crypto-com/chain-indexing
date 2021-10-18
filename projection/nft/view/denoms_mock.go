package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockDenomsView struct {
	testify_mock.Mock
}

func (denomView *MockDenomsView) Insert(row *DenomRow) error {
	mockArgs := denomView.Called(row)
	return mockArgs.Error(0)
}

func (denomView *MockDenomsView) FindById(denomId string) (*DenomRow, error) {
	mockArgs := denomView.Called(denomId)
	denomRow, _ := mockArgs.Get(0).(*DenomRow)
	return denomRow, mockArgs.Error(1)
}

func (denomView *MockDenomsView) FindByName(denomName string) (*DenomRow, error) {
	mockArgs := denomView.Called(denomName)
	denomRow, _ := mockArgs.Get(0).(*DenomRow)
	return denomRow, mockArgs.Error(1)
}

func (denomView *MockDenomsView) List(
	filter DenomListFilter,
	order DenomListOrder,
	pagination *pagination_interface.Pagination,
) ([]DenomRow, *pagination_interface.PaginationResult, error) {
	mockArgs := denomView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]DenomRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}
