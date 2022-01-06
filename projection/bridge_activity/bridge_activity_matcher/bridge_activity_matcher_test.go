package bridge_activity_matcher_test

import (
	"fmt"
	test_logger "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	"github.com/crypto-com/chain-indexing/usecase/coin"
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
		MockFunc func(
			config *bridge_activity_matcher.Config,
		) (
			mockThisRDbConn *test.MockRDbConn,
			mocks []*testify_mock.Mock,
			assertFunc func(),
		)
	}{
		{
			Name:   "It should insert pending bridge activity when there is unprocessed outgoing bridge activity in the current chain",
			Config: bridge_activity_matcher.Config{},
			MockFunc: func(
				config *bridge_activity_matcher.Config,
			) (
				mockThisRDbConn *test.MockRDbConn,
				mocks []*testify_mock.Mock,
				assertFunc func(),
			) {
				mockThisRdbHandle := NewMockRdbHandle()
				mockThisRDbConn = NewMockRDbConn(mockThisRdbHandle)
				mockThisTx := NewMockRDbTx(mockThisRdbHandle)
				mockThisRDbConn.On("Begin").Return(mockThisTx, nil)

				mockCounterpartyRdbHandle := NewMockRdbHandle()
				mockCounterpartyRDbConn := NewMockRDbConn(mockCounterpartyRdbHandle)
				mockCounterpartyTx := NewMockRDbTx(mockCounterpartyRdbHandle)
				mockCounterpartyRDbConn.On("Begin").Return(mockCounterpartyTx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					_ *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					return mockCounterpartyRDbConn, nil
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					}
					panic("unrecognized RDb handle")
				}

				unprocessedOutgoingRow := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   1,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1000)),
						MaybeTransactionId:            primptr.String("txid"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        "link-id",
						Direction:                     types.DIRECTION_OUTGOING,
						FromChainId:                   "from-chain",
						MaybeFromAddress:              primptr.String("from-address"),
						MaybeFromSmartContractAddress: nil,
						ToChainId:                     "to-chain",
						ToAddress:                     "to-address",
						MaybeToSmartContractAddress:   nil,
						MaybeChannelId:                primptr.String("channel-0"),
						Amount:                        coin.NewInt(100),
						MaybeDenom:                    primptr.String("basecro"),
						MaybeBridgeFeeAmount:          nil,
						MaybeBridgeFeeDenom:           nil,
						Status:                        types.STATUS_PENDING,
						IsProcessed:                   false,
					},
					Id:        0,
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
				}
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{
						unprocessedOutgoingRow,
					},
					nil,
				)
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)

				mockBridgeActivitiesView.On("Insert", &view.BridgeActivityInsertRow{
					BridgeType:                           unprocessedOutgoingRow.BridgeType,
					SourceBlockHeight:                    unprocessedOutgoingRow.BlockHeight,
					SourceBlockTime:                      unprocessedOutgoingRow.BlockTime,
					SourceTransactionId:                  *unprocessedOutgoingRow.MaybeTransactionId,
					SourceChain:                          unprocessedOutgoingRow.FromChainId,
					SourceAddress:                        *unprocessedOutgoingRow.MaybeFromAddress,
					MaybeSourceSmartContractAddress:      unprocessedOutgoingRow.MaybeFromSmartContractAddress,
					MaybeDestinationBlockHeight:          nil,
					MaybeDestinationBlockTime:            nil,
					MaybeDestinationTransactionId:        nil,
					DestinationChain:                     unprocessedOutgoingRow.ToChainId,
					DestinationAddress:                   unprocessedOutgoingRow.ToAddress,
					MaybeDestinationSmartContractAddress: unprocessedOutgoingRow.MaybeToSmartContractAddress,
					MaybeChannelId:                       unprocessedOutgoingRow.MaybeChannelId,
					LinkId:                               unprocessedOutgoingRow.LinkId,
					Amount:                               unprocessedOutgoingRow.Amount,
					MaybeDenom:                           unprocessedOutgoingRow.MaybeDenom,
					MaybeBridgeFeeAmount:                 unprocessedOutgoingRow.MaybeBridgeFeeAmount,
					MaybeBridgeFeeDenom:                  unprocessedOutgoingRow.MaybeBridgeFeeDenom,
					Status:                               types.STATUS_PENDING,
				}).Return(nil)
				mockThisBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					unprocessedOutgoingRow.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
					mockThisBridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name:   "It should insert pending bridge activity when there is unprocessed outgoing bridge activity from the counterparty chain",
			Config: bridge_activity_matcher.Config{},
			MockFunc: func(
				config *bridge_activity_matcher.Config,
			) (
				mockThisRDbConn *test.MockRDbConn,
				mocks []*testify_mock.Mock,
				assertFunc func(),
			) {
				mockThisRdbHandle := NewMockRdbHandle()
				mockThisRDbConn = NewMockRDbConn(mockThisRdbHandle)
				mockThisTx := NewMockRDbTx(mockThisRdbHandle)
				mockThisRDbConn.On("Begin").Return(mockThisTx, nil)

				mockCounterpartyRdbHandle := NewMockRdbHandle()
				mockCounterpartyRDbConn := NewMockRDbConn(mockCounterpartyRdbHandle)
				mockCounterpartyTx := NewMockRDbTx(mockCounterpartyRdbHandle)
				mockCounterpartyRDbConn.On("Begin").Return(mockCounterpartyTx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					_ *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					return mockCounterpartyRDbConn, nil
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					}
					panic("unrecognized RDb handle")
				}

				unprocessedOutgoingRow := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   1,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1000)),
						MaybeTransactionId:            primptr.String("txid"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        "link-id",
						Direction:                     types.DIRECTION_OUTGOING,
						FromChainId:                   "from-chain",
						MaybeFromAddress:              primptr.String("from-address"),
						MaybeFromSmartContractAddress: nil,
						ToChainId:                     "to-chain",
						ToAddress:                     "to-address",
						MaybeToSmartContractAddress:   nil,
						MaybeChannelId:                primptr.String("channel-0"),
						Amount:                        coin.NewInt(100),
						MaybeDenom:                    primptr.String("basecro"),
						MaybeBridgeFeeAmount:          nil,
						MaybeBridgeFeeDenom:           nil,
						Status:                        types.STATUS_PENDING,
						IsProcessed:                   false,
					},
					Id:        0,
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
				}
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{
						unprocessedOutgoingRow,
					},
					nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)

				mockBridgeActivitiesView.On("Insert", &view.BridgeActivityInsertRow{
					BridgeType:                           unprocessedOutgoingRow.BridgeType,
					SourceBlockHeight:                    unprocessedOutgoingRow.BlockHeight,
					SourceBlockTime:                      unprocessedOutgoingRow.BlockTime,
					SourceTransactionId:                  *unprocessedOutgoingRow.MaybeTransactionId,
					SourceChain:                          unprocessedOutgoingRow.FromChainId,
					SourceAddress:                        *unprocessedOutgoingRow.MaybeFromAddress,
					MaybeSourceSmartContractAddress:      unprocessedOutgoingRow.MaybeFromSmartContractAddress,
					MaybeDestinationBlockHeight:          nil,
					MaybeDestinationBlockTime:            nil,
					MaybeDestinationTransactionId:        nil,
					DestinationChain:                     unprocessedOutgoingRow.ToChainId,
					DestinationAddress:                   unprocessedOutgoingRow.ToAddress,
					MaybeDestinationSmartContractAddress: unprocessedOutgoingRow.MaybeToSmartContractAddress,
					MaybeChannelId:                       unprocessedOutgoingRow.MaybeChannelId,
					LinkId:                               unprocessedOutgoingRow.LinkId,
					Amount:                               unprocessedOutgoingRow.Amount,
					MaybeDenom:                           unprocessedOutgoingRow.MaybeDenom,
					MaybeBridgeFeeAmount:                 unprocessedOutgoingRow.MaybeBridgeFeeAmount,
					MaybeBridgeFeeDenom:                  unprocessedOutgoingRow.MaybeBridgeFeeDenom,
					Status:                               types.STATUS_PENDING,
				}).Return(nil)
				mockCounterpartyBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					unprocessedOutgoingRow.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
					mockCounterpartyBridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name:   "Given outgoing bridge activity, It should update the activity status when there is unprocessed incoming bridge activity at the counterparty chain",
			Config: bridge_activity_matcher.Config{},
			MockFunc: func(
				config *bridge_activity_matcher.Config,
			) (
				mockThisRDbConn *test.MockRDbConn,
				mocks []*testify_mock.Mock,
				assertFunc func(),
			) {
				mockThisRdbHandle := NewMockRdbHandle()
				mockThisRDbConn = NewMockRDbConn(mockThisRdbHandle)
				mockThisTx := NewMockRDbTx(mockThisRdbHandle)
				mockThisRDbConn.On("Begin").Return(mockThisTx, nil)

				mockCounterpartyRdbHandle := NewMockRdbHandle()
				mockCounterpartyRDbConn := NewMockRDbConn(mockCounterpartyRdbHandle)
				mockCounterpartyTx := NewMockRDbTx(mockCounterpartyRdbHandle)
				mockCounterpartyRDbConn.On("Begin").Return(mockCounterpartyTx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					_ *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					return mockCounterpartyRDbConn, nil
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)

				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					}
					panic("unrecognized RDb handle")
				}

				thisOutgoingRow := view.BridgeActivityReadRow{
					BridgeActivityInsertRow: view.BridgeActivityInsertRow{
						BridgeType:                           types.BRIDGE_TYPE_IBC,
						SourceBlockHeight:                    1,
						SourceBlockTime:                      primptr.UTCTime(utctime.FromUnixNano(1000)),
						SourceTransactionId:                  "txid",
						SourceChain:                          "from-chain",
						SourceAddress:                        "from-address",
						MaybeSourceSmartContractAddress:      nil,
						MaybeDestinationBlockHeight:          nil,
						MaybeDestinationBlockTime:            nil,
						MaybeDestinationTransactionId:        nil,
						DestinationChain:                     "to-chain",
						DestinationAddress:                   "to-address",
						MaybeDestinationSmartContractAddress: nil,
						MaybeChannelId:                       nil,
						LinkId:                               "link-id",
						Amount:                               coin.NewInt(100),
						MaybeDenom:                           primptr.String("basecro"),
						MaybeBridgeFeeAmount:                 nil,
						MaybeBridgeFeeDenom:                  nil,
						Status:                               types.STATUS_PENDING,
					},
					Id:        0,
					UUID:      "uuid",
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
				}
				counterpartyIncomingRow := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   2,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(2000)),
						MaybeTransactionId:            primptr.String("txid2"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        thisOutgoingRow.LinkId,
						Direction:                     types.DIRECTION_INCOMING,
						FromChainId:                   thisOutgoingRow.SourceChain,
						MaybeFromAddress:              primptr.String(thisOutgoingRow.SourceAddress),
						MaybeFromSmartContractAddress: thisOutgoingRow.MaybeSourceSmartContractAddress,
						ToChainId:                     thisOutgoingRow.DestinationChain,
						ToAddress:                     thisOutgoingRow.DestinationAddress,
						MaybeToSmartContractAddress:   thisOutgoingRow.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                primptr.String("channel-1"),
						Amount:                        thisOutgoingRow.Amount,
						MaybeDenom:                    thisOutgoingRow.MaybeDenom,
						MaybeBridgeFeeAmount:          thisOutgoingRow.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:           thisOutgoingRow.MaybeBridgeFeeDenom,
						Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
						IsProcessed:                   false,
					},
					Id:        0,
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(2000)),
				}
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockThisBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{
						counterpartyIncomingRow,
					},
					nil,
				)
				mockBridgeActivitiesView.On("FindByLinkId", thisOutgoingRow.LinkId).Return(
					thisOutgoingRow, nil,
				)

				mockBridgeActivitiesView.On("Update", &view.BridgeActivityReadRow{
					BridgeActivityInsertRow: view.BridgeActivityInsertRow{
						BridgeType:                           thisOutgoingRow.BridgeType,
						SourceBlockHeight:                    thisOutgoingRow.SourceBlockHeight,
						SourceBlockTime:                      thisOutgoingRow.SourceBlockTime,
						SourceTransactionId:                  thisOutgoingRow.SourceTransactionId,
						SourceChain:                          thisOutgoingRow.SourceChain,
						SourceAddress:                        thisOutgoingRow.SourceAddress,
						MaybeSourceSmartContractAddress:      thisOutgoingRow.MaybeSourceSmartContractAddress,
						MaybeDestinationBlockHeight:          primptr.Int64(counterpartyIncomingRow.BlockHeight),
						MaybeDestinationBlockTime:            counterpartyIncomingRow.BlockTime,
						MaybeDestinationTransactionId:        counterpartyIncomingRow.MaybeTransactionId,
						DestinationChain:                     thisOutgoingRow.DestinationChain,
						DestinationAddress:                   thisOutgoingRow.DestinationAddress,
						MaybeDestinationSmartContractAddress: thisOutgoingRow.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                       thisOutgoingRow.MaybeChannelId,
						LinkId:                               thisOutgoingRow.LinkId,
						Amount:                               thisOutgoingRow.Amount,
						MaybeDenom:                           thisOutgoingRow.MaybeDenom,
						MaybeBridgeFeeAmount:                 thisOutgoingRow.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:                  thisOutgoingRow.MaybeBridgeFeeDenom,
						Status:                               counterpartyIncomingRow.Status,
					},

					Id:        thisOutgoingRow.Id,
					UUID:      thisOutgoingRow.UUID,
					CreatedAt: thisOutgoingRow.CreatedAt,
					UpdatedAt: thisOutgoingRow.UpdatedAt,
				}).Return(nil)
				mockCounterpartyBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					counterpartyIncomingRow.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Update", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		//{
		//	Name:   "It should do nothing when there is unprocessed incoming bridge activity at the counterparty chain but the corresponding outgoing activity is not inserted yet",
		//	Config: bridge_activity_matcher.Config{},
		//	MockFunc: func() (
		//		mockThisRDbConn *test.MockRDbConn,
		//		mocks []*testify_mock.Mock,
		//		assertFunc func(),
		//	) {
		//		return nil, []*testify_mock.Mock{}, func() {}
		//	},
		//},
	}

	for _, tc := range testCases {
		fmt.Print(tc.Name)
		mockThisRDbConn, mocks, assertFunc := tc.MockFunc(&tc.Config)

		projection := NewBridgeActivityMatcherProjection(mockThisRDbConn, tc.Config)
		onInitErr := projection.OnInit()
		assert.NoError(t, onInitErr)

		execErr := projection.Exec()
		assert.NoError(t, execErr)

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
	fakeLogger := test_logger.NewFakeLogger()
	return bridge_activity_matcher.New(
		config,
		fakeLogger,
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
