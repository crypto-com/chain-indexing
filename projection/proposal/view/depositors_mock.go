package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination2 "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockDepositorsView struct {
	testify_mock.Mock
}

func NewMockDepositorsView() Depositors {
	return &MockDepositorsView{}
}

func (depositorsView *MockDepositorsView) Insert(row *DepositorRow) error {
	mockArgs := depositorsView.Called(row)
	return mockArgs.Error(0)
}

func (depositorsView *MockDepositorsView) ListByProposalId(
	proposalId string,
	order DepositorListOrder,
	pagination *pagination2.Pagination,
) (
	[]DepositorWithMonikerRow,
	*pagination2.PaginationResult,
	error,
) {
	mockArgs := depositorsView.Called(proposalId, order, pagination)
	result1, _ := mockArgs.Get(0).([]DepositorWithMonikerRow)
	result2, _ := mockArgs.Get(1).(*pagination2.PaginationResult)
	return result1, result2, mockArgs.Error(2)
}
