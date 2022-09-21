package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockAccountRawTransactionsView struct {
	testify_mock.Mock
}

func NewMockAccountRawTransactionsView(_ *rdb.Handle) AccountRawTransactions {
	return &MockAccountRawTransactionsView{}
}

func (accountRawTransactionsView *MockAccountRawTransactionsView) InsertAll(
	rows []AccountRawTransactionRow,
) error {
	mockArgs := accountRawTransactionsView.Called(rows)
	return mockArgs.Error(0)
}

func (accountRawTransactionsView *MockAccountRawTransactionsView) List(
	filter AccountRawTransactionsListFilter,
	order AccountRawTransactionsListOrder,
	pagination *pagination_interface.Pagination,
) ([]AccountRawTransactionRow, *pagination_interface.PaginationResult, error) {
	mockArgs := accountRawTransactionsView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]AccountRawTransactionRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}
