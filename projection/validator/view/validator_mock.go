package view

import (
	"math/big"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/stretchr/testify/mock"
)

type MockValidatorsView struct {
	mock.Mock
}

func NewMockValidatorsView(_ *rdb.Handle) Validators {
	return &MockValidatorsView{}
}

func (validatorsView *MockValidatorsView) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, error) {
	mockArgs := validatorsView.Called(operatorAddress, consensusNodeAddress)
	result0, _ := mockArgs.Get(0).(bool)
	result1, _ := mockArgs.Get(1).(int64)
	return result0, result1, mockArgs.Error(2)
}

func (validatorsView *MockValidatorsView) Upsert(validator *ValidatorRow) error {
	mockArgs := validatorsView.Called(validator)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) Insert(validator *ValidatorRow) error {
	mockArgs := validatorsView.Called(validator)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) totalPower() (*big.Float, error) {
	mockArgs := validatorsView.Called()
	row, _ := mockArgs.Get(0).(*big.Float)
	return row, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) Search(keyword string) ([]ValidatorRow, error) {
	mockArgs := validatorsView.Called(keyword)
	rows, _ := mockArgs.Get(0).([]ValidatorRow)
	return rows, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) FindBy(identity ValidatorIdentity) (*ValidatorRow, error) {
	mockArgs := validatorsView.Called(identity)
	result1, _ := mockArgs.Get(0).(*ValidatorRow)
	return result1, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) Update(validator *ValidatorRow) error {
	mockArgs := validatorsView.Called(validator)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) UpdateAllValidatorUpTime(validators []ValidatorRow) error {
	mockArgs := validatorsView.Called(validators)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) ListAll(filter ValidatorsListFilter, order ValidatorsListOrder) ([]ValidatorRow, error) {
	mockArgs := validatorsView.Called(filter, order)
	rows, _ := mockArgs.Get(0).([]ValidatorRow)
	return rows, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) List(filter ValidatorsListFilter, order ValidatorsListOrder, pagination *pagination_interface.Pagination) ([]ListValidatorRow, *pagination.PaginationResult, error) {
	mockArgs := validatorsView.Called(filter, order, pagination)
	result0, _ := mockArgs.Get(0).([]ListValidatorRow)
	result1, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return result0, result1, mockArgs.Error(2)
}

func (validatorsView *MockValidatorsView) Count(filter CountFilter) (int64, error) {
	mockArgs := validatorsView.Called(filter)
	total, _ := mockArgs.Get(0).(int64)
	return total, mockArgs.Error(1)
}
