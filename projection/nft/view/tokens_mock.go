package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockTokensView struct {
	testify_mock.Mock
}

func (tokensView *MockTokensView) Insert(row *TokenRow) error {
	mockArgs := tokensView.Called(row)
	return mockArgs.Error(0)
}

func (tokensView *MockTokensView) Delete(denomId string, tokenId string) (int64, error) {
	mockArgs := tokensView.Called(denomId, tokenId)
	rowsAffected, _ := mockArgs.Get(0).(int64)
	return rowsAffected, mockArgs.Error(1)
}

func (tokensView *MockTokensView) FindById(denomId string, tokenId string) (*TokenRowWithDenomname, error) {
	mockArgs := tokensView.Called(denomId, tokenId)
	row, _ := mockArgs.Get(0).(*TokenRowWithDenomname)
	return row, mockArgs.Error(1)
}

func (tokensView *MockTokensView) Update(row TokenRow) error {
	mockArgs := tokensView.Called(row)
	return mockArgs.Error(0)
}

func (tokensView *MockTokensView) List(
	filter TokenListFilter,
	order TokenListOrder,
	pagination *pagination_interface.Pagination,
) ([]TokenRowWithDenomname, *pagination_interface.PaginationResult, error) {
	mockArgs := tokensView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]TokenRowWithDenomname)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (tokensView *MockTokensView) ListDrops(
	pagination *pagination_interface.Pagination,
) ([]string, *pagination_interface.PaginationResult, error) {
	mockArgs := tokensView.Called(pagination)
	rows, _ := mockArgs.Get(0).([]string)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (tokensView *MockTokensView) SoftDelete(denomId string, tokenId string) error {
	mockArgs := tokensView.Called(denomId, tokenId)
	return mockArgs.Error(0)
}
