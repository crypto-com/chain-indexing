package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockTransactionsView struct {
	testify_mock.Mock
}

func NewMockTransactionsView(_ *rdb.Handle) BlockRawTransactions {
	return &MockTransactionsView{}
}

func (transactionsView *MockTransactionsView) InsertAll(transactions []RawTransactionRow) error {
	mockArgs := transactionsView.Called(transactions)
	return mockArgs.Error(0)
}

func (transactionsView *MockTransactionsView) Insert(row *RawTransactionRow) error {
	mockArgs := transactionsView.Called(row)
	return mockArgs.Error(0)
}

func (transactionsView *MockTransactionsView) FindByHash(txHash string) (*RawTransactionRow, error) {
	mockArgs := transactionsView.Called(txHash)
	row, _ := mockArgs.Get(0).(*RawTransactionRow)
	return row, mockArgs.Error(1)
}

func (transactionsView *MockTransactionsView) List(
	filter RawTransactionsListFilter,
	order RawTransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]RawTransactionRow, *pagination_interface.PaginationResult, error) {
	mockArgs := transactionsView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]RawTransactionRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (transactionsView *MockTransactionsView) Search(keyword string) ([]RawTransactionRow, error) {
	mockArgs := transactionsView.Called(keyword)
	rows, _ := mockArgs.Get(0).([]RawTransactionRow)
	return rows, mockArgs.Error(1)
}

func (transactionsView *MockTransactionsView) Count() (int64, error) {
	mockArgs := transactionsView.Called()
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
