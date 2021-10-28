package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockTransactionsView struct {
	testify_mock.Mock
}

func NewMockTransactionsView(_ *rdb.Handle) BlockTransactions {
	return &MockTransactionsView{}
}

func (transactionsView *MockTransactionsView) InsertAll(transactions []TransactionRow) error {
	mockArgs := transactionsView.Called(transactions)
	return mockArgs.Error(0)
}

func (transactionsView *MockTransactionsView) Insert(row *TransactionRow) error {
	mockArgs := transactionsView.Called(row)
	return mockArgs.Error(0)
}

func (transactionsView *MockTransactionsView) FindByHash(txHash string) (*TransactionRow, error) {
	mockArgs := transactionsView.Called(txHash)
	row, _ := mockArgs.Get(0).(*TransactionRow)
	return row, mockArgs.Error(1)
}

func (transactionsView *MockTransactionsView) List(
	filter TransactionsListFilter,
	order TransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]TransactionRow, *pagination_interface.PaginationResult, error) {
	mockArgs := transactionsView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]TransactionRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (transactionsView *MockTransactionsView) Search(keyword string) ([]TransactionRow, error) {
	mockArgs := transactionsView.Called(keyword)
	rows, _ := mockArgs.Get(0).([]TransactionRow)
	return rows, mockArgs.Error(1)
}

func (transactionsView *MockTransactionsView) Count() (int64, error) {
	mockArgs := transactionsView.Called()
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
