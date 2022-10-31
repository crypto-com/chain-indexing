package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockMessagesView struct {
	testify_mock.Mock
}

func (messagesView *MockMessagesView) Insert(row *MessageRow) error {
	mockArgs := messagesView.Called(row)
	return mockArgs.Error(0)
}

func (messagesView *MockMessagesView) List(
	filter MessagesListFilter,
	order MessagesListOrder,
	pagination *pagination_interface.Pagination,
) ([]MessageRow, *pagination_interface.PaginationResult, error) {
	mockArgs := messagesView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]MessageRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (messagesView *MockMessagesView) DeleteAllByDenomTokenIds(denomId string, tokenId string) (int64, error) {
	mockArgs := messagesView.Called(denomId, tokenId)
	rowsAffected, _ := mockArgs.Get(0).(int64)
	return rowsAffected, mockArgs.Error(1)
}

func (messagesView *MockMessagesView) BurnMessagesByToken(denomId string, maybeTokenId string) error {
	mockArgs := messagesView.Called(denomId, maybeTokenId)
	return mockArgs.Error(1)
}
