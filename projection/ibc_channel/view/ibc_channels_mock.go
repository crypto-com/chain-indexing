package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel/types"
	"github.com/stretchr/testify/mock"
)

type MockIBCChannelsView struct {
	mock.Mock
}

func NewMockIBCChannelsView(_ *rdb.Handle) IBCChannels {
	return &MockIBCChannelsView{}
}

func (ibcChannelsView *MockIBCChannelsView) Insert(channel *IBCChannelRow) error {
	mockArgs := ibcChannelsView.Called(channel)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateFactualColumns(channel *IBCChannelRow) error {
	mockArgs := ibcChannelsView.Called(channel)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) Increment(channelID string, column string, increaseNumber int64) error {
	mockArgs := ibcChannelsView.Called(channelID, column, increaseNumber)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateSequence(channelID string, column string, sequence uint64) error {
	mockArgs := ibcChannelsView.Called(channelID, column, sequence)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateTotalTransferOutSuccessRate(channelID string) error {
	mockArgs := ibcChannelsView.Called(channelID)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateLastActivityTimeAndHeight(channelID string, time utctime.UTCTime, height int64) error {
	mockArgs := ibcChannelsView.Called(channelID, time, height)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateStatus(channelID string, status types.Status) error {
	mockArgs := ibcChannelsView.Called(channelID, status)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) UpdateBondedTokens(channelID string, bondedTokens *BondedTokens) error {
	mockArgs := ibcChannelsView.Called(channelID, bondedTokens)
	return mockArgs.Error(0)
}

func (ibcChannelsView *MockIBCChannelsView) FindBondedTokensBy(channelID string) (*BondedTokens, error) {
	mockArgs := ibcChannelsView.Called(channelID)
	result, _ := mockArgs.Get(0).(*BondedTokens)
	return result, mockArgs.Error(1)
}

func (ibcChannelsView *MockIBCChannelsView) FindBy(channelID string) (*IBCChannelRow, error) {
	mockArgs := ibcChannelsView.Called(channelID)
	result, _ := mockArgs.Get(0).(*IBCChannelRow)
	return result, mockArgs.Error(1)
}

func (ibcChannelsView *MockIBCChannelsView) List(
	order IBCChannelsListOrder,
	filter IBCChannelsListFilter,
	paginate *pagination.Pagination,
) (
	[]IBCChannelRow,
	*pagination.PaginationResult,
	error,
) {
	mockArgs := ibcChannelsView.Called(order, filter, paginate)
	result0, _ := mockArgs.Get(0).([]IBCChannelRow)
	result1, _ := mockArgs.Get(1).(*pagination.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}

func (ibcChannelsView *MockIBCChannelsView) ListChannelsGroupByChainId(
	order IBCChannelsListOrder,
	filter IBCChannelsListFilter,
	paginate *pagination.Pagination,
) (
	[]ChainChannels,
	*pagination.PaginationResult,
	error,
) {
	mockArgs := ibcChannelsView.Called(order, filter, paginate)
	result0, _ := mockArgs.Get(0).([]ChainChannels)
	result1, _ := mockArgs.Get(1).(*pagination.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}
