package cosmosapp

import (
	"math/big"

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

func (conn *MockClient) Account(accountAddress string, cosmosAPIVersion string) (*Account, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(*Account)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Balances(accountAddress string, cosmosAPIVersion string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) BalanceByDenom(accountAddress string, denom string, cosmosAPIVersion string) (coin.Coin, error) {
	mockArgs := conn.Called(accountAddress, denom, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coin)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) BondedBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) RedelegatingBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) UnbondingBalance(accountAddress string, cosmosAPIVersion string) (coin.Coins, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) SupplyByDenom(denom string, cosmosAPIVersion string) (coin.Coin, error) {
	mockArgs := conn.Called(denom, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.Coin)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) TotalRewards(accountAddress string, cosmosAPIVersion string) (coin.DecCoins, error) {
	mockArgs := conn.Called(accountAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.DecCoins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Commission(validatorAddress string, cosmosAPIVersion string) (coin.DecCoins, error) {
	mockArgs := conn.Called(validatorAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.DecCoins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) CommunityPool(cosmosAPIVersion string) (coin.DecCoins, error) {
	mockArgs := conn.Called(cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(coin.DecCoins)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Validator(validatorAddress string, cosmosAPIVersion string) (*Validator, error) {
	mockArgs := conn.Called(validatorAddress, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(*Validator)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Delegation(delegator string, validator string, cosmosAPIVersion string) (*DelegationResponse, error) {
	mockArgs := conn.Called(delegator, validator, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(*DelegationResponse)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) TotalBondedBalance(cosmosAPIVersion string) (coin.Coin, error) {
	args := conn.Called(cosmosAPIVersion)
	return args.Get(0).(coin.Coin), args.Error(1)
}

func (conn *MockClient) CommunityTax(cosmosAPIVersion string) (*big.Float, error) {
	args := conn.Called(cosmosAPIVersion)
	return args.Get(0).(*big.Float), args.Error(1)
}

func (conn *MockClient) AnnualProvisions(cosmosAPIVersion string) (coin.DecCoin, error) {
	args := conn.Called(cosmosAPIVersion)
	return args.Get(0).(coin.DecCoin), args.Error(1)
}

func (conn *MockClient) Proposals(cosmosAPIVersion string) ([]Proposal, error) {
	args := conn.Called(cosmosAPIVersion)
	return args.Get(0).([]Proposal), args.Error(1)
}

func (conn *MockClient) ProposalById(id string, cosmosAPIVersion string) (Proposal, error) {
	mockArgs := conn.Called(id, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(Proposal)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) ProposalTally(id string, cosmosAPIVersion string) (Tally, error) {
	mockArgs := conn.Called(id, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(Tally)
	return result, mockArgs.Error(1)
}

func (conn *MockClient) Tx(txHash string, cosmosAPIVersion string) (*model.Tx, error) {
	mockArgs := conn.Called(txHash, cosmosAPIVersion)
	result, _ := mockArgs.Get(0).(*model.Tx)
	return result, mockArgs.Error(1)
}
