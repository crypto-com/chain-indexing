package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockAccountRawEventsTotalView struct {
	mock.Mock
}

func NewMockAccountRawEventsTotalView(_ *rdb.Handle) AccountRawEventsTotal {
	return &MockAccountRawEventsTotalView{}
}

func (view *MockAccountRawEventsTotalView) Set(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockAccountRawEventsTotalView) Increment(identity string, total int64) error {
	mockArgs := view.Called(identity, total)
	return mockArgs.Error(0)
}

func (view *MockAccountRawEventsTotalView) IncrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockAccountRawEventsTotalView) DecrementAll(identities []string, total int64) error {
	mockArgs := view.Called(identities, total)
	return mockArgs.Error(0)
}

func (view *MockAccountRawEventsTotalView) FindBy(identity string) (int64, error) {
	mockArgs := view.Called(identity)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockAccountRawEventsTotalView) SumBy(identities []string) (int64, error) {
	mockArgs := view.Called(identities)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
