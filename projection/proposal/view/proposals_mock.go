package view

import (
	"math/big"

	testify_mock "github.com/stretchr/testify/mock"

	pagination2 "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockProposalsView struct {
	testify_mock.Mock
}

func NewMockProposalsView() Proposals {
	return &MockProposalsView{}
}

func (proposalsView *MockProposalsView) Insert(row *ProposalRow) error {
	mockArgs := proposalsView.Called(row)
	return mockArgs.Error(0)
}

func (proposalsView *MockProposalsView) IncrementTotalVoteBy(
	proposalId uint64,
	voteToAdd *big.Int,
) error {
	mockArgs := proposalsView.Called(proposalId, voteToAdd)
	return mockArgs.Error(0)
}

func (proposalsView *MockProposalsView) Update(row *ProposalRow) error {
	mockArgs := proposalsView.Called(row)
	return mockArgs.Error(0)
}

func (proposalsView *MockProposalsView) FindById(
	proposalId string,
) (
	[]ProposalWithMonikerRow, error,
) {
	mockArgs := proposalsView.Called(proposalId)
	result1, _ := mockArgs.Get(0).([]ProposalWithMonikerRow)
	return result1, mockArgs.Error(1)
}

func (proposalsView *MockProposalsView) List(
	filter ProposalListFilter,
	order ProposalListOrder,
	pagination *pagination2.Pagination,
) (
	[]ProposalWithMonikerRow,
	*pagination2.PaginationResult,
	error,
) {
	mockArgs := proposalsView.Called(filter, order, pagination)
	result1, _ := mockArgs.Get(0).([]ProposalWithMonikerRow)
	result2, _ := mockArgs.Get(1).(*pagination2.PaginationResult)
	return result1, result2, mockArgs.Error(2)
}
