package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockAccountMessagesTotalView struct {
	mock.Mock
}

func NewMockAccountMessagesTotalView(_ *rdb.Handle) AccountMessagesTotal {
	return &MockAccountMessagesTotalView{}
}

func (view *MockAccountMessagesTotalView) Set(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockAccountMessagesTotalView) Increment(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockAccountMessagesTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockAccountMessagesTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockAccountMessagesTotalView) FindBy(identity string) (int64, error) {
	mockArgs := view.Called(identity)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockAccountMessagesTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := view.Called(identities)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}