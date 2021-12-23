package view

import (
	"github.com/stretchr/testify/mock"
)

type MockBridgePendingActivitiesView struct {
	mock.Mock
}

var _ BridgePendingActivities = &MockBridgePendingActivitiesView{}

func NewMockBridgePendingActivitiesView() BridgePendingActivities {
	return &MockBridgePendingActivitiesView{}
}

func (view *MockBridgePendingActivitiesView) ListAllUnprocessedOutgoing() ([]BridgePendingActivityReadRow, error) {
	mockArgs := view.Called()
	row, _ := mockArgs.Get(0).([]BridgePendingActivityReadRow)
	return row, mockArgs.Error(1)
}

func (view *MockBridgePendingActivitiesView) ListAllUnprocessedIncoming() ([]BridgePendingActivityReadRow, error) {
	mockArgs := view.Called()
	row, _ := mockArgs.Get(0).([]BridgePendingActivityReadRow)
	return row, mockArgs.Error(1)
}

func (view *MockBridgePendingActivitiesView) List(
	filter BridgePendingActivitiesFilter,
	order BridgePendingActivitiesOrder,
) ([]BridgePendingActivityReadRow, error) {
	mockArgs := view.Called(filter, order)
	rows, _ := mockArgs.Get(0).([]BridgePendingActivityReadRow)
	return rows, mockArgs.Error(1)
}

func (view *MockBridgePendingActivitiesView) UpdateToProcessed(id int64) error {
	mockArgs := view.Called(id)
	return mockArgs.Error(0)
}

func (view *MockBridgePendingActivitiesView) Insert(activity *BridgePendingActivityInsertRow) error {
	mockArgs := view.Called(activity)
	return mockArgs.Error(0)
}
