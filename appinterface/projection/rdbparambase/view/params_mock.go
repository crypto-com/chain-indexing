package view

import (
	"math/big"
	"time"

	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/types"
)

type MockParamsView struct {
	testify_mock.Mock
}

func NewMockParamsView() Params {
	return &MockParamsView{}
}

func (view *MockParamsView) Set(accessor types.ParamAccessor, value string) error {
	mockArgs := view.Called(accessor, value)
	return mockArgs.Error(0)
}

func (view *MockParamsView) FindBy(accessor types.ParamAccessor) (string, error) {
	mockArgs := view.Called(accessor)
	return mockArgs.String(0), mockArgs.Error(1)
}

func (view *MockParamsView) FindInt64By(accessor types.ParamAccessor) (int64, error) {
	mockArgs := view.Called(accessor)
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}

func (view *MockParamsView) FindBigIntBy(accessor types.ParamAccessor) (*big.Int, error) {
	mockArgs := view.Called(accessor)
	result, _ := mockArgs.Get(0).(*big.Int)
	return result, mockArgs.Error(1)
}

func (view *MockParamsView) FindUint64By(accessor types.ParamAccessor) (uint64, error) {
	mockArgs := view.Called(accessor)
	result, _ := mockArgs.Get(0).(uint64)
	return result, mockArgs.Error(1)
}

func (view *MockParamsView) FindDurationBy(accessor types.ParamAccessor) (time.Duration, error) {
	mockArgs := view.Called(accessor)
	result, _ := mockArgs.Get(0).(time.Duration)
	return result, mockArgs.Error(1)
}

func (view *MockParamsView) FindBoolBy(accessor types.ParamAccessor) (bool, error) {
	mockArgs := view.Called(accessor)
	return mockArgs.Bool(0), mockArgs.Error(1)
}
