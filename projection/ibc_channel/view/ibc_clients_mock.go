package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockIBCClientsView struct {
	mock.Mock
}

func NewMockIBCClientsView(_ *rdb.Handle) IBCClients {
	return &MockIBCClientsView{}
}

func (ibcClientsView *MockIBCClientsView) Insert(ibcClient *IBCClientRow) error {
	mockArgs := ibcClientsView.Called(ibcClient)
	return mockArgs.Error(0)
}

func (ibcClientsView *MockIBCClientsView) FindCounterpartyChainIDBy(clientID string) (string, error) {
	mockArgs := ibcClientsView.Called(clientID)
	result, _ := mockArgs.Get(0).(string)
	return result, mockArgs.Error(1)
}
