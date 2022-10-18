package view

import (
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockAccountRawEventsView struct {
	mock.Mock
}

func NewMockAccountMessagesView(_ *rdb.Handle) AccountRawEvents {
	return &MockAccountRawEventsView{}
}

func (accountRawEventsView *MockAccountRawEventsView) Insert(messageRow *AccountRawEventRow) error {
	mockArgs := accountRawEventsView.Called(messageRow)
	return mockArgs.Error(0)
}

func (accountRawEventsView *MockAccountRawEventsView) InsertAll(messageRows []AccountRawEventRow) error {
	mockArgs := accountRawEventsView.Called(messageRows)
	return mockArgs.Error(0)
}

func (accountMessagesView *MockAccountRawEventsView) List(
	filter AccountRawEventsListFilter,
	order AccountRawEventsListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountRawEventRow, *pagination_interface.PaginationResult, error) {
	mockArgs := accountMessagesView.Called(filter, order, pagination)
	result0, _ := mockArgs.Get(0).([]AccountRawEventRow)
	result1, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}

func (accountMessagesView *MockAccountRawEventsView) AccountListByHeight(
	height int64,
) ([]string, error) {
	mockArgs := accountMessagesView.Called(height)
	result, _ := mockArgs.Get(0).([]string)
	return result, mockArgs.Error(1)
}
