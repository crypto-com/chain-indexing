package view

import (
	"github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockVotesTotalView struct {
	mock.Mock
}

func NewMockVotesTotalView(_ *rdb.Handle) VotesTotal {
	return &MockVotesTotalView{}
}

func (view *MockVotesTotalView) Set(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockVotesTotalView) Increment(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockVotesTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockVotesTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockVotesTotalView) FindBy(identity string) (int64, error) {
	mockArgs := view.Called(identity)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockVotesTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := view.Called(identities)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
