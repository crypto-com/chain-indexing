package view

import (
	"math/big"

	testify_mock "github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type MockValidatorsView struct {
	testify_mock.Mock
}

func NewMockValidatorsView(_ *rdb.Handle) Validators {
	return &MockValidatorsView{}
}

func (validatorsView *MockValidatorsView) LastJoinedBlockHeight(
	operatorAddress string,
	consensusNodeAddress string,
) (bool, int64, error) {
	mockArgs := validatorsView.Called(operatorAddress, consensusNodeAddress)
	rows, _ := mockArgs.Get(0).(int64)
	return true, rows, mockArgs.Error(0)
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
	rows, _ := mockArgs.Get(0).(big.Float)
	return &rows, mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) Search(keyword string) ([]ValidatorRow, error) {
	mockArgs := validatorsView.Called(keyword)
	rows, _ := mockArgs.Get(0).([]ValidatorRow)
	return rows, mockArgs.Error(1)
}

func (validatorsView *MockValidatorsView) FindBy(identity ValidatorIdentity) (*ValidatorRow, error) {
	mockArgs := validatorsView.Called(identity)
	row, _ := mockArgs.Get(0).(*ValidatorRow)
	return row, mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) Update(validator *ValidatorRow) error {
	mockArgs := validatorsView.Called(validator)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) UpdateAllValidatorUpTime(validators []ValidatorRow) error {
	mockArgs := validatorsView.Called(validators)
	return mockArgs.Error(0)
}

func (validatorsView *MockValidatorsView) ListAll(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
) ([]ValidatorRow, error) {
	mockArgs := validatorsView.Called(filter, order)
	rows, _ := mockArgs.Get(0).([]ValidatorRow)
	return rows, mockArgs.Error(2)
}

func (validatorsView *MockValidatorsView) List(
	filter ValidatorsListFilter,
	order ValidatorsListOrder,
	pagination *pagination_interface.Pagination,
) ([]ListValidatorRow, *pagination_interface.PaginationResult, error) {
	mockArgs := validatorsView.Called(filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]ListValidatorRow)
	paginationResult, _ := mockArgs.Get(1).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(2)
}

func (validatorsView *MockValidatorsView) Count(filter CountFilter) (int64, error) {
	mockArgs := validatorsView.Called()
	result, _ := mockArgs.Get(0).(int64)
	return result, mockArgs.Error(1)
}
