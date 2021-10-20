package view

import (
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/stretchr/testify/mock"
)

type MockAccountMessagesView struct {
	mock.Mock
}

func NewMockAccountMessagesView(_ *rdb.Handle) AccountMessages {
	return &MockAccountMessagesView{}
}

func (accountMessagesView *MockAccountMessagesView) Insert(messageRow *AccountMessageRow, accounts []string) error {
	if data, ok := messageRow.Data.(*event_usecase.MsgSend); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgMultiSend); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgSetWithdrawAddress); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgWithdrawDelegatorReward); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgWithdrawValidatorCommission); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgFundCommunityPool); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgSubmitParamChangeProposal); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgSubmitCommunityPoolSpendProposal); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgSubmitCancelSoftwareUpgradeProposal); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgDeposit); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgVote); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgCreateValidator); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgEditValidator); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgDelegate); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgUndelegate); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgBeginRedelegate); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgUnjail); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgNFTIssueDenom); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgNFTMintNFT); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgNFTTransferNFT); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgNFTEditNFT); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgNFTBurnNFT); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCCreateClient); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCUpdateClient); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCConnectionOpenInit); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCConnectionOpenAck); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCConnectionOpenTry); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCConnectionOpenConfirm); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCChannelOpenInit); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCChannelOpenAck); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCChannelOpenTry); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCChannelOpenConfirm); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCAcknowledgement); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCRecvPacket); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCTransferTransfer); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCTimeout); ok {
		data.EventUUID = "TESTUUID"
	} else if data, ok := messageRow.Data.(*event_usecase.MsgIBCTimeoutOnClose); ok {
		data.EventUUID = "TESTUUID"
	}
	mockArgs := accountMessagesView.Called(messageRow, accounts)
	return mockArgs.Error(0)
}

func (accountMessagesView *MockAccountMessagesView) List(
	filter AccountMessagesListFilter,
	order AccountMessagesListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountMessageRow, *pagination_interface.PaginationResult, error) {
	mockArgs := accountMessagesView.Called(filter, order, pagination)
	result0, _ := mockArgs.Get(0).([]AccountMessageRow)
	result1, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}
