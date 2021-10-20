package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination2 "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockIBCChannelMessageView struct {
	testify_mock.Mock
}

func (ibcChannelMessagesView *MockIBCChannelMessageView) Insert(ibcChannelMessage *IBCChannelMessageRow) error {
	mockArgs := ibcChannelMessagesView.Called(ibcChannelMessage)
	return mockArgs.Error(0)
}

func (ibcChannelMessagesView *MockIBCChannelMessageView) ListByChannelID(
	channelID string,
	order IBCChannelMessagesListOrder,
	filter IBCChannelMessagesListFilter,
	pagination *pagination2.Pagination,
) (
	[]IBCChannelMessageRow,
	*pagination2.PaginationResult,
	error,
) {
	mockArgs := ibcChannelMessagesView.Called(channelID, order, filter, pagination)
	messages, _ := mockArgs.Get(0).([]IBCChannelMessageRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination2.PaginationResult)
	return messages, paginationResult, mockArgs.Error(3)
}
