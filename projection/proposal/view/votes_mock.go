package view

import (
	testify_mock "github.com/stretchr/testify/mock"

	pagination2 "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockVotesView struct {
	testify_mock.Mock
}

func NewMockVotesView() Votes {
	return &MockVotesView{}
}

func (votesView *MockVotesView) Insert(row *VoteRow) error {
	mockArgs := votesView.Called(row)
	return mockArgs.Error(0)
}

func (votesView *MockVotesView) FindByProposalIdVoter(
	proposalId string,
	voterAddress string,
) (
	[]VoteWithMonikerRow,
	error,
) {
	mockArgs := votesView.Called(proposalId, voterAddress)
	result1, _ := mockArgs.Get(0).([]VoteWithMonikerRow)
	return result1, mockArgs.Error(1)
}

func (votesView *MockVotesView) ListByProposalId(
	proposalId string,
	order VoteListOrder,
	pagination *pagination2.Pagination,
) (
	[]VoteWithMonikerRow,
	*pagination2.PaginationResult,
	error,
) {
	mockArgs := votesView.Called(proposalId, order, pagination)
	result1, _ := mockArgs.Get(0).([]VoteWithMonikerRow)
	result2, _ := mockArgs.Get(1).(*pagination2.PaginationResult)
	return result1, result2, mockArgs.Error(2)
}

func (votesView *MockVotesView) DeleteByProposalIdVoter(proposalId string, voterAddress string) (int64, error) {
	mockArgs := votesView.Called(proposalId, voterAddress)
	rowsAffected, _ := mockArgs.Get(0).(int64)
	return rowsAffected, mockArgs.Error(1)
}
