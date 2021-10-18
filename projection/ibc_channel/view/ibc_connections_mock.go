package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockIBCConnectionsView struct {
	mock.Mock
}

func NewMockIBCConnectionsView(_ *rdb.Handle) IBCConnections {
	return &MockIBCConnectionsView{}
}

func (ibcConnectionsView *MockIBCConnectionsView) Insert(ibcConnection *IBCConnectionRow) error {
	mockArgs := ibcConnectionsView.Called(ibcConnection)
	return mockArgs.Error(0)
}

func (ibcConnectionsView *MockIBCConnectionsView) FindCounterpartyChainIDBy(connectionID string) (string, error) {
	mockArgs := ibcConnectionsView.Called(connectionID)
	result, _ := mockArgs.Get(0).(string)
	return result, mockArgs.Error(1)
}