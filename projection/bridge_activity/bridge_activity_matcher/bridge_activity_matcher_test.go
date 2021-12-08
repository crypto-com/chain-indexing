package bridge_activity_matcher_test

import (
	"fmt"
	"testing"

	applogger "github.com/crypto-com/chain-indexing/external/logger"

	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_activity_matcher"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
)

func TestBridgeActivityMatcher_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Config   bridge_activity_matcher.Config
		MockFunc func() (mocks []*testify_mock.Mock, assertFunc func())
	}{
		{
			Name:   "It should not handle IBC transfer from non-listened source channel",
			Config: bridge_activity_matcher.Config{},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockThisRdbHandle := NewMockRdbHandle()
				mockThisRDbConn := NewMockRDbConn(mockThisRdbHandle)
				mockThisTx := NewMockRDbTx(mockThisRdbHandle)
				mockThisRDbConn.On("Begin").Return(mockThisTx, nil)

				mockCryptoOrgChainRdbHandle := NewMockRdbHandle()
				mockCryptoOrgChainRDbConn := NewMockRDbConn(mockCryptoOrgChainRdbHandle)
				mockCryptoOrgChainTx := NewMockRDbTx(mockCryptoOrgChainRdbHandle)
				mockCryptoOrgChainRDbConn.On("Begin").Return(mockCryptoOrgChainTx, nil)

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCryptoOrgChainBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCryptoOrgChainBridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRdbHandle {
						return mockBridgeActivitiesView
					} else if conn == mockCryptoOrgChainRdbHandle {
						panic("unexpected Crypto.org Chain RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRdbHandle {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCryptoOrgChainRdbHandle {
						return mockCryptoOrgChainBridgePendingActivitiesView
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewPgxConnPool = func(
					_ *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					return mockCryptoOrgChainRDbConn, nil
				}

				assertFunc = func() {
					mockThisBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
	}

	for _, tc := range testCases {
		fmt.Print(tc.Name)
		mocks, assertFunc := tc.MockFunc()

		projection := NewBridgePendingActivityProjection(mockRDbConn, tc.Config)
		projection.OnInit()
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}
		assertFunc()

		fmt.Println(" Passed")
	}
}

func NewBridgeActivityMatcherProjection(
	rdbConn rdb.Conn,
	config bridge_activity_matcher.Config,
) *bridge_activity_matcher.BridgeActivityMatcher {
	return bridge_activity_matcher.NewWithConfig(
		&config,
		nil,
		rdbConn,
		nil,
	)
}

func NewMockRDbConn(handle *rdb.Handle) *test.MockRDbConn {
	mock := test.NewMockRDbConn()
	mock.On("ToHandle").Return(handle)

	return mock
}

func NewMockRDbTx(handle *rdb.Handle) *test.MockRDbTx {
	mockTx := &test.MockRDbTx{}
	mockTx.On("ToHandle").Return(handle).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}

func NewMockRdbHandle() *rdb.Handle {
	return &rdb.Handle{
		Runner:   &test.MockRDbTx{},
		TypeConv: &pg.PgxTypeConv{},
		StmtBuilder: &rdb.StatementBuilder{
			StatementBuilderType: sq.StatementBuilderType{},
			PlaceholderFormat:    nil,
		},
	}
}
