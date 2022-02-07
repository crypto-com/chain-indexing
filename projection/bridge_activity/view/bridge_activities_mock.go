package view

import (
	"github.com/stretchr/testify/mock"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

type MockBridgeActivitiesView struct {
	mock.Mock
}

var _ BridgeActivities = &MockBridgeActivitiesView{}

func NewMockBridgeActivitiesView() BridgeActivities {
	return &MockBridgeActivitiesView{}
}

func (view *MockBridgeActivitiesView) FindByLinkId(linkId string) (BridgeActivityReadRow, error) {
	mockArgs := view.Called(linkId)
	row, _ := mockArgs.Get(0).(BridgeActivityReadRow)
	return row, mockArgs.Error(1)
}

func (view *MockBridgeActivitiesView) FindBy(filter BridgeActivitiesFindByFilter) (BridgeActivityReadRow, error) {
	mockArgs := view.Called(filter)
	row, _ := mockArgs.Get(0).(BridgeActivityReadRow)
	return row, mockArgs.Error(1)
}

func (view *MockBridgeActivitiesView) ListByChainAddress(
	chain string,
	address string,
	order BridgeActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error) {
	mockArgs := view.Called(chain, address, order, pagination)
	rows, _ := mockArgs.Get(0).([]BridgeActivityReadRow)
	paginationResult, _ := mockArgs.Get(0).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(1)
}

func (view *MockBridgeActivitiesView) List(
	addressFilter BridgeActivitiesListAddressFilter,
	filter BridgeActivitiesListFilter,
	order BridgeActivitiesListOrder,
	pagination *pagination_interface.Pagination,
) ([]BridgeActivityReadRow, *pagination_interface.PaginationResult, error) {
	mockArgs := view.Called(addressFilter, filter, order, pagination)
	rows, _ := mockArgs.Get(0).([]BridgeActivityReadRow)
	paginationResult, _ := mockArgs.Get(0).(*pagination_interface.PaginationResult)
	return rows, paginationResult, mockArgs.Error(1)
}

func (view *MockBridgeActivitiesView) Insert(activity *BridgeActivityInsertRow) error {
	mockArgs := view.Called(activity)
	return mockArgs.Error(0)
}

func (view *MockBridgeActivitiesView) Update(activity *BridgeActivityReadRow) error {
	mockArgs := view.Called(activity)
	return mockArgs.Error(0)
}
