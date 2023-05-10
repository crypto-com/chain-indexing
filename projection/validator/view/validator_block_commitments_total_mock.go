package view

import (
	"github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockValidatorBlockCommitmentsTotalView struct {
	mock.Mock
}

func NewMockValidatorBlockCommitmentsTotal(_ *rdb.Handle) ValidatorBlockCommitmentsTotal {
	return &MockValidatorBlockCommitmentsTotalView{}
}

func (view *MockValidatorBlockCommitmentsTotalView) Set(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockValidatorBlockCommitmentsTotalView) Increment(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockValidatorBlockCommitmentsTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockValidatorBlockCommitmentsTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockValidatorBlockCommitmentsTotalView) FindBy(identity string) (int64, error) {
	mockArgs := view.Called(identity)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockValidatorBlockCommitmentsTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := view.Called(identities)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
