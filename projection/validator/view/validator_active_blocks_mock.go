package view

import (
	"github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockValidatorActiveBlocksView struct {
	mock.Mock
}

func NewMockValidatorActiveBlocksView(_ *rdb.Handle) ValidatorActiveBlocks {
	return &MockValidatorActiveBlocksView{}
}

func (view *MockValidatorActiveBlocksView) UpdateValidatorsActiveBlocks(
	operatorAddressToSignedBlockFlagMap OperatorAddressToSignedBlockFlagMap,
	height int64,
	maxRecentUpTimeInBlocks int64,
) error {
	mockArgs := view.Called(operatorAddressToSignedBlockFlagMap, height, maxRecentUpTimeInBlocks)
	return mockArgs.Error(0)
}

func (view *MockValidatorActiveBlocksView) ClearValidatorActiveBLocks(operatorAddress string) error {
	mockArgs := view.Called(operatorAddress)
	return mockArgs.Error(0)

}
