package cosmosapp

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (conn *MockClient) Account(accountAddress string) (*Account, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(*Account)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Balances(accountAddress string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) BondedBalance(accountAddress string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) RedelegatingBalance(accountAddress string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) UnbondingBalance(accountAddress string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) TotalRewards(accountAddress string) (coin.DecCoins, error) {
	mockArgs := conn.Called(accountAddress)
	result, _ := mockArgs.Get(0).(coin.DecCoins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Commission(validatorAddress string) (coin.DecCoins, error) {
	mockArgs := conn.Called(validatorAddress)
	result, _ := mockArgs.Get(0).(coin.DecCoins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Validator(validatorAddress string) (*Validator, error) {
	mockArgs := conn.Called(validatorAddress)
	result, _ := mockArgs.Get(0).(*Validator)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Delegation(delegator string, validator string) (*DelegationResponse, error) {
	mockArgs := conn.Called(delegator, validator)
	result, _ := mockArgs.Get(0).(*DelegationResponse)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) TotalBondedBalance() (coin.Coin, error) {
	args := conn.Called()
	return args.Get(0).(coin.Coin), args.Error(1)
}

func (conn *MockClient) AnnualProvisions() (coin.DecCoin, error) {
	args := conn.Called()
	return args.Get(0).(coin.DecCoin), args.Error(1)
}

func (conn *MockClient) Proposals() ([]Proposal, error) {
	args := conn.Called()
	return args.Get(0).([]Proposal), args.Error(1)
}

func (conn *MockClient) ProposalById(id string) (Proposal, error) {
	mockArgs := conn.Called(id)
	result, _ := mockArgs.Get(0).(Proposal)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) ProposalTally(id string) (Tally, error) {
	mockArgs := conn.Called(id)
	result, _ := mockArgs.Get(0).(Tally)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Tx(txHash string) (*model.Tx, error) {
	mockArgs := conn.Called(txHash)
	result, _ := mockArgs.Get(0).(*model.Tx)
	return result, mockArgs.Error(1)
}
