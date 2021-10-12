package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockIBCDenomHashMappingView struct {
	mock.Mock
}

func NewMockIBCDenomHashMapping(_ *rdb.Handle) IBCDenomHashMapping {
	return &MockIBCDenomHashMappingView{}
}

func (ibcDenomHashMappingView *MockIBCDenomHashMappingView) IfDenomExist(denom string) (bool, error) {
	mockArgs := ibcDenomHashMappingView.Called(denom)
	result, _ := mockArgs.Get(0).(bool)
	return result, mockArgs.Error(1)
}

func (ibcDenomHashMappingView *MockIBCDenomHashMappingView) Insert(ibcDenomHash *IBCDenomHashMappingRow) error {
	mockArgs := ibcDenomHashMappingView.Called(ibcDenomHash)
	return mockArgs.Error(0)
}

func (ibcDenomHashMappingView *MockIBCDenomHashMappingView) ListAll() (*[]IBCDenomHashMappingRow, error) {
	args := ibcDenomHashMappingView.Called()
	return args.Get(0).(*[]IBCDenomHashMappingRow), args.Error(1)
}