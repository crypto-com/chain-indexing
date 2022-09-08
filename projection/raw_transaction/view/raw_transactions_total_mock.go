package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockRawTransactionsTotalView struct {
	mock.Mock
}

func NewMockRawTransactionsTotalView(_ *rdb.Handle) RawTransactionsTotal {
	return &MockRawTransactionsTotalView{}
}

func (view *MockRawTransactionsTotalView) Set(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockRawTransactionsTotalView) Increment(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockRawTransactionsTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockRawTransactionsTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockRawTransactionsTotalView) FindBy(identity string) (int64, error) {
	mockArgs := view.Called(identity)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockRawTransactionsTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := view.Called(identities)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
