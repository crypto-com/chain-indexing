package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockAccountsView struct {
	mock.Mock
}

func NewMockAccountsView(_ *rdb.Handle) Accounts {
	return &MockAccountsView{}
}

func (accountsView *MockAccountsView) FindBy(identity *AccountIdentity) (*AccountRow, error) {
	mockArgs := accountsView.Called(identity)
	result, _ := mockArgs.Get(0).(*AccountRow)
	return result, mockArgs.Error(1)
}

func (accountsView *MockAccountsView) List(
	order AccountsListOrder,
	paginate *pagination.Pagination,
) ([]AccountRow, *pagination.PaginationResult, error) {
	mockArgs := accountsView.Called(order, paginate)
	result0, _ := mockArgs.Get(0).([]AccountRow)
	result1, _ := mockArgs.Get(1).(*pagination.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}

func (accountsView *MockAccountsView) IncrementUsableBalance(address string, denom string, increment int64) error {
	mockArgs := accountsView.Called(address, denom, increment)
	return mockArgs.Error(0)
}

func (accountsView *MockAccountsView) DecrementUsableBalance(address string, denom string, decrement int64) error {
	mockArgs := accountsView.Called(address, denom, decrement)
	return mockArgs.Error(0)
}

func (accountsView *MockAccountsView) InsertAccountEvent(event AccountEvent) error {
	mockArgs := accountsView.Called(event)
	return mockArgs.Error(0)
}
