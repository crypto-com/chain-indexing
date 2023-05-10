package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockValidatorBlockCommitmentsView struct {
	mock.Mock
}

func NewMockValidatorBlockCommitmentsView(_ *rdb.Handle) ValidatorBlockCommitments {
	return &MockValidatorBlockCommitmentsView{}
}
func (commitmentsView *MockValidatorBlockCommitmentsView) InsertAll(commitments []ValidatorBlockCommitmentRow) error {
	mockArgs := commitmentsView.Called(commitments)
	return mockArgs.Error(0)
}
func (commitmentsView *MockValidatorBlockCommitmentsView) Insert(commitment ValidatorBlockCommitmentRow) error {
	mockArgs := commitmentsView.Called(commitment)
	return mockArgs.Error(0)
}
func (commitmentsView *MockValidatorBlockCommitmentsView) List(
	filter ValidatorBlockCommitmentsListFilter,
	pagination *pagination_interface.Pagination,
) ([]ListValidatorBlockCommitmentRow, *pagination.PaginationResult, error) {
	mockArgs := commitmentsView.Called(filter, pagination)
	result0, _ := mockArgs.Get(0).([]ListValidatorBlockCommitmentRow)
	result1, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}
