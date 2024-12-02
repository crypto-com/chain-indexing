package bridge_activity_matcher_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	testify_mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	test_logger "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_activity_matcher"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
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
			Name: "It should insert pending bridge activity when there is unprocessed outgoing bridge activity in the current chain",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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
			Name: "It should insert pending bridge activity when there is unprocessed outgoing bridge activity from the counterparty chain",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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
			Name: "It should insert pending bridge activitiy when there is unprocessed outgoing bridge activities ONLY in one of the counterparty chains",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name: "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty",
						},
					},
					{
						Name: "CounterpartyChain2",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty2",
						},
					},
				},
			},
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

				mockCounterparty2RdbHandle := NewMockRdbHandle()
				mockCounterparty2RDbConn := NewMockRDbConn(mockCounterparty2RdbHandle)
				mockCounterparty2Tx := NewMockRDbTx(mockCounterparty2RdbHandle)
				mockCounterparty2RDbConn.On("Begin").Return(mockCounterparty2Tx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					config *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					if config.Host == "postgres://counterparty" {
						return mockCounterpartyRDbConn, nil
					} else if config.Host == "postgres://counterparty2" {
						return mockCounterparty2RDbConn, nil
					}
					panic(fmt.Sprintf("unrecognized counterparty RDb host: %s", config.Host))
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterparty2BridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterparty2BridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						panic("unexpected Counterparty Chain 2 RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						return mockCounterparty2BridgePendingActivitiesView
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
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{},
					nil,
				)
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
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
			Name: "It should insert pending bridge activities when there are unprocessed outgoing bridge activities from multiple counterparty chains",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name: "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty",
						},
					},
					{
						Name: "CounterpartyChain2",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty2",
						},
					},
				},
			},
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

				mockCounterparty2RdbHandle := NewMockRdbHandle()
				mockCounterparty2RDbConn := NewMockRDbConn(mockCounterparty2RdbHandle)
				mockCounterparty2Tx := NewMockRDbTx(mockCounterparty2RdbHandle)
				mockCounterparty2RDbConn.On("Begin").Return(mockCounterparty2Tx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					config *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					if config.Host == "postgres://counterparty" {
						return mockCounterpartyRDbConn, nil
					} else if config.Host == "postgres://counterparty2" {
						return mockCounterparty2RDbConn, nil
					}
					panic(fmt.Sprintf("unrecognized counterparty RDb host: %s", config.Host))
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterparty2BridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterparty2BridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						panic("unexpected Counterparty Chain 2 RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						return mockCounterparty2BridgePendingActivitiesView
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
				unprocessedOutgoingRow2 := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   1,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1000)),
						MaybeTransactionId:            primptr.String("txid"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        "link-id",
						Direction:                     types.DIRECTION_OUTGOING,
						FromChainId:                   "from-chain-2",
						MaybeFromAddress:              primptr.String("from-address"),
						MaybeFromSmartContractAddress: nil,
						ToChainId:                     "to-chain",
						ToAddress:                     "to-address",
						MaybeToSmartContractAddress:   nil,
						MaybeChannelId:                primptr.String("channel-1"),
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
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{
						unprocessedOutgoingRow2,
					},
					nil,
				)
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
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
				mockBridgeActivitiesView.On("Insert", &view.BridgeActivityInsertRow{
					BridgeType:                           unprocessedOutgoingRow2.BridgeType,
					SourceBlockHeight:                    unprocessedOutgoingRow2.BlockHeight,
					SourceBlockTime:                      unprocessedOutgoingRow2.BlockTime,
					SourceTransactionId:                  *unprocessedOutgoingRow2.MaybeTransactionId,
					SourceChain:                          unprocessedOutgoingRow2.FromChainId,
					SourceAddress:                        *unprocessedOutgoingRow2.MaybeFromAddress,
					MaybeSourceSmartContractAddress:      unprocessedOutgoingRow2.MaybeFromSmartContractAddress,
					MaybeDestinationBlockHeight:          nil,
					MaybeDestinationBlockTime:            nil,
					MaybeDestinationTransactionId:        nil,
					DestinationChain:                     unprocessedOutgoingRow2.ToChainId,
					DestinationAddress:                   unprocessedOutgoingRow2.ToAddress,
					MaybeDestinationSmartContractAddress: unprocessedOutgoingRow2.MaybeToSmartContractAddress,
					MaybeChannelId:                       unprocessedOutgoingRow2.MaybeChannelId,
					LinkId:                               unprocessedOutgoingRow2.LinkId,
					Amount:                               unprocessedOutgoingRow2.Amount,
					MaybeDenom:                           unprocessedOutgoingRow2.MaybeDenom,
					MaybeBridgeFeeAmount:                 unprocessedOutgoingRow2.MaybeBridgeFeeAmount,
					MaybeBridgeFeeDenom:                  unprocessedOutgoingRow2.MaybeBridgeFeeDenom,
					Status:                               types.STATUS_PENDING,
				}).Return(nil)
				mockCounterpartyBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					unprocessedOutgoingRow.Id,
				).Return(nil)
				mockCounterparty2BridgePendingActivitiesView.On(
					"UpdateToProcessed",
					unprocessedOutgoingRow.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Insert", 2)
					mockCounterpartyBridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
					mockCounterparty2BridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name: "Given outgoing bridge activity, It should update the activity status when there is unprocessed incoming bridge activity at the counterparty chain",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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
		{
			Name: "Given outgoing bridge activities, It should update the activity status when there are corresponding unprocessed incoming bridge activities at multiple counterparty chain",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name: "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty",
						},
					},
					{
						Name: "CounterpartyChain2",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{
							Host: "postgres://counterparty2",
						},
					},
				},
			},
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

				mockCounterparty2RdbHandle := NewMockRdbHandle()
				mockCounterparty2RDbConn := NewMockRDbConn(mockCounterparty2RdbHandle)
				mockCounterparty2Tx := NewMockRDbTx(mockCounterparty2RdbHandle)
				mockCounterparty2RDbConn.On("Begin").Return(mockCounterparty2Tx, nil)

				bridge_activity_matcher.NewPgxConnPool = func(
					config *pg.PgxConnPoolConfig, _ applogger.Logger,
				) (rdb.Conn, error) {
					if config.Host == "postgres://counterparty" {
						return mockCounterpartyRDbConn, nil
					} else if config.Host == "postgres://counterparty2" {
						return mockCounterparty2RDbConn, nil
					}
					panic(fmt.Sprintf("unrecognized counterparty RDb host: %s", config.Host))
				}

				mockThisBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterpartyBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockCounterparty2BridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mockBridgeActivitiesView := view.NewMockBridgeActivitiesView().(*view.MockBridgeActivitiesView)
				mocks = append(mocks, &mockThisBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterpartyBridgePendingActivitiesView.Mock)
				mocks = append(mocks, &mockCounterparty2BridgePendingActivitiesView.Mock)

				bridge_activity_matcher.NewBridgeActivitiesView = func(conn *rdb.Handle) view.BridgeActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockBridgeActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						panic("unexpected Counterparty Chain RDb handle")
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						panic("unexpected Counterparty Chain 2 RDb handle")
					}
					panic("unrecognized RDb handle")
				}

				bridge_activity_matcher.NewBridgePendingActivitiesView = func(conn *rdb.Handle) view.BridgePendingActivities {
					if conn == mockThisRDbConn.ToHandle() {
						return mockThisBridgePendingActivitiesView
					} else if conn == mockCounterpartyRDbConn.ToHandle() {
						return mockCounterpartyBridgePendingActivitiesView
					} else if conn == mockCounterparty2RDbConn.ToHandle() {
						return mockCounterparty2BridgePendingActivitiesView
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
				thisOutgoingRow2 := view.BridgeActivityReadRow{
					BridgeActivityInsertRow: view.BridgeActivityInsertRow{
						BridgeType:                           types.BRIDGE_TYPE_IBC,
						SourceBlockHeight:                    3,
						SourceBlockTime:                      primptr.UTCTime(utctime.FromUnixNano(3000)),
						SourceTransactionId:                  "txid3",
						SourceChain:                          "from-chain",
						SourceAddress:                        "from-address",
						MaybeSourceSmartContractAddress:      nil,
						MaybeDestinationBlockHeight:          nil,
						MaybeDestinationBlockTime:            nil,
						MaybeDestinationTransactionId:        nil,
						DestinationChain:                     "to-chain-2",
						DestinationAddress:                   "to-address",
						MaybeDestinationSmartContractAddress: nil,
						MaybeChannelId:                       nil,
						LinkId:                               "link-id-2",
						Amount:                               coin.NewInt(100),
						MaybeDenom:                           primptr.String("basecro"),
						MaybeBridgeFeeAmount:                 nil,
						MaybeBridgeFeeDenom:                  nil,
						Status:                               types.STATUS_PENDING,
					},
					Id:        1,
					UUID:      "uuid-2",
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(5000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(5000)),
				}
				counterpartyIncomingRow2 := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   4,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(4000)),
						MaybeTransactionId:            primptr.String("txid4"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        thisOutgoingRow2.LinkId,
						Direction:                     types.DIRECTION_INCOMING,
						FromChainId:                   thisOutgoingRow2.SourceChain,
						MaybeFromAddress:              primptr.String(thisOutgoingRow2.SourceAddress),
						MaybeFromSmartContractAddress: thisOutgoingRow2.MaybeSourceSmartContractAddress,
						ToChainId:                     thisOutgoingRow2.DestinationChain,
						ToAddress:                     thisOutgoingRow2.DestinationAddress,
						MaybeToSmartContractAddress:   thisOutgoingRow2.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                primptr.String("channel-1"),
						Amount:                        thisOutgoingRow2.Amount,
						MaybeDenom:                    thisOutgoingRow2.MaybeDenom,
						MaybeBridgeFeeAmount:          thisOutgoingRow2.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:           thisOutgoingRow2.MaybeBridgeFeeDenom,
						Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
						IsProcessed:                   false,
					},
					Id:        1,
					CreatedAt: primptr.UTCTime(utctime.FromUnixNano(5000)),
					UpdatedAt: primptr.UTCTime(utctime.FromUnixNano(5000)),
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
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterparty2BridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{
						counterpartyIncomingRow2,
					},
					nil,
				)
				mockBridgeActivitiesView.On("FindByLinkId", thisOutgoingRow.LinkId).Return(
					thisOutgoingRow, nil,
				)
				mockBridgeActivitiesView.On("FindByLinkId", thisOutgoingRow2.LinkId).Return(
					thisOutgoingRow2, nil,
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
				mockBridgeActivitiesView.On("Update", &view.BridgeActivityReadRow{
					BridgeActivityInsertRow: view.BridgeActivityInsertRow{
						BridgeType:                           thisOutgoingRow2.BridgeType,
						SourceBlockHeight:                    thisOutgoingRow2.SourceBlockHeight,
						SourceBlockTime:                      thisOutgoingRow2.SourceBlockTime,
						SourceTransactionId:                  thisOutgoingRow2.SourceTransactionId,
						SourceChain:                          thisOutgoingRow2.SourceChain,
						SourceAddress:                        thisOutgoingRow2.SourceAddress,
						MaybeSourceSmartContractAddress:      thisOutgoingRow2.MaybeSourceSmartContractAddress,
						MaybeDestinationBlockHeight:          primptr.Int64(counterpartyIncomingRow2.BlockHeight),
						MaybeDestinationBlockTime:            counterpartyIncomingRow2.BlockTime,
						MaybeDestinationTransactionId:        counterpartyIncomingRow2.MaybeTransactionId,
						DestinationChain:                     thisOutgoingRow2.DestinationChain,
						DestinationAddress:                   thisOutgoingRow2.DestinationAddress,
						MaybeDestinationSmartContractAddress: thisOutgoingRow2.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                       thisOutgoingRow2.MaybeChannelId,
						LinkId:                               thisOutgoingRow2.LinkId,
						Amount:                               thisOutgoingRow2.Amount,
						MaybeDenom:                           thisOutgoingRow2.MaybeDenom,
						MaybeBridgeFeeAmount:                 thisOutgoingRow2.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:                  thisOutgoingRow2.MaybeBridgeFeeDenom,
						Status:                               counterpartyIncomingRow2.Status,
					},

					Id:        thisOutgoingRow2.Id,
					UUID:      thisOutgoingRow2.UUID,
					CreatedAt: thisOutgoingRow2.CreatedAt,
					UpdatedAt: thisOutgoingRow2.UpdatedAt,
				}).Return(nil)
				mockCounterpartyBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					counterpartyIncomingRow.Id,
				).Return(nil)
				mockCounterparty2BridgePendingActivitiesView.On(
					"UpdateToProcessed",
					counterpartyIncomingRow2.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Update", 2)
					mockCounterpartyBridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
					mockCounterparty2BridgePendingActivitiesView.AssertNumberOfCalls(t, "UpdateToProcessed", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name: "Given outgoing bridge activity at counterparty chain, It should update the activity status when there is unprocessed incoming bridge activity at the local chain",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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

				counterpartyOutgoingRow := view.BridgeActivityReadRow{
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
				thisIncomingRow := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   2,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(2000)),
						MaybeTransactionId:            primptr.String("txid2"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        counterpartyOutgoingRow.LinkId,
						Direction:                     types.DIRECTION_INCOMING,
						FromChainId:                   counterpartyOutgoingRow.SourceChain,
						MaybeFromAddress:              primptr.String(counterpartyOutgoingRow.SourceAddress),
						MaybeFromSmartContractAddress: counterpartyOutgoingRow.MaybeSourceSmartContractAddress,
						ToChainId:                     counterpartyOutgoingRow.DestinationChain,
						ToAddress:                     counterpartyOutgoingRow.DestinationAddress,
						MaybeToSmartContractAddress:   counterpartyOutgoingRow.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                primptr.String("channel-1"),
						Amount:                        counterpartyOutgoingRow.Amount,
						MaybeDenom:                    counterpartyOutgoingRow.MaybeDenom,
						MaybeBridgeFeeAmount:          counterpartyOutgoingRow.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:           counterpartyOutgoingRow.MaybeBridgeFeeDenom,
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
					[]view.BridgePendingActivityReadRow{
						thisIncomingRow,
					},
					nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockBridgeActivitiesView.On("FindByLinkId", counterpartyOutgoingRow.LinkId).Return(
					counterpartyOutgoingRow, nil,
				)

				mockBridgeActivitiesView.On("Update", &view.BridgeActivityReadRow{
					BridgeActivityInsertRow: view.BridgeActivityInsertRow{
						BridgeType:                           counterpartyOutgoingRow.BridgeType,
						SourceBlockHeight:                    counterpartyOutgoingRow.SourceBlockHeight,
						SourceBlockTime:                      counterpartyOutgoingRow.SourceBlockTime,
						SourceTransactionId:                  counterpartyOutgoingRow.SourceTransactionId,
						SourceChain:                          counterpartyOutgoingRow.SourceChain,
						SourceAddress:                        counterpartyOutgoingRow.SourceAddress,
						MaybeSourceSmartContractAddress:      counterpartyOutgoingRow.MaybeSourceSmartContractAddress,
						MaybeDestinationBlockHeight:          primptr.Int64(thisIncomingRow.BlockHeight),
						MaybeDestinationBlockTime:            thisIncomingRow.BlockTime,
						MaybeDestinationTransactionId:        thisIncomingRow.MaybeTransactionId,
						DestinationChain:                     counterpartyOutgoingRow.DestinationChain,
						DestinationAddress:                   counterpartyOutgoingRow.DestinationAddress,
						MaybeDestinationSmartContractAddress: counterpartyOutgoingRow.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                       counterpartyOutgoingRow.MaybeChannelId,
						LinkId:                               counterpartyOutgoingRow.LinkId,
						Amount:                               counterpartyOutgoingRow.Amount,
						MaybeDenom:                           counterpartyOutgoingRow.MaybeDenom,
						MaybeBridgeFeeAmount:                 counterpartyOutgoingRow.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:                  counterpartyOutgoingRow.MaybeBridgeFeeDenom,
						Status:                               thisIncomingRow.Status,
					},

					Id:        counterpartyOutgoingRow.Id,
					UUID:      counterpartyOutgoingRow.UUID,
					CreatedAt: counterpartyOutgoingRow.CreatedAt,
					UpdatedAt: counterpartyOutgoingRow.UpdatedAt,
				}).Return(nil)
				mockThisBridgePendingActivitiesView.On(
					"UpdateToProcessed",
					thisIncomingRow.Id,
				).Return(nil)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Update", 1)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name: "It should do nothing when there is unprocessed incoming bridge activity at the counterparty chain but the corresponding outgoing activity at this chain is not inserted yet",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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
					nil, rdb.ErrNoRows,
				)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Update", 0)
					mockCounterpartyBridgePendingActivitiesView.AssertNumberOfCalls(
						t, "UpdateToProcessed", 0,
					)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
		{
			Name: "It should do nothing when there is unprocessed incoming bridge activity at this chain but the corresponding outgoing activity at counterparty chain at this chain is not inserted yet",
			Config: bridge_activity_matcher.Config{
				CounterpartyChains: []bridge_activity_matcher.CounterpartyChainConfig{
					{
						Name:     "CounterpartyChain",
						Database: bridge_activity_matcher.CounterpartyDatabaseConfig{},
					},
				},
			},
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

				counterpartyOutgoingRow := view.BridgeActivityReadRow{
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
				thisIncomingRow := view.BridgePendingActivityReadRow{
					BridgePendingActivityInsertRow: view.BridgePendingActivityInsertRow{
						BlockHeight:                   2,
						BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(2000)),
						MaybeTransactionId:            primptr.String("txid2"),
						BridgeType:                    types.BRIDGE_TYPE_IBC,
						LinkId:                        counterpartyOutgoingRow.LinkId,
						Direction:                     types.DIRECTION_INCOMING,
						FromChainId:                   counterpartyOutgoingRow.SourceChain,
						MaybeFromAddress:              primptr.String(counterpartyOutgoingRow.SourceAddress),
						MaybeFromSmartContractAddress: counterpartyOutgoingRow.MaybeSourceSmartContractAddress,
						ToChainId:                     counterpartyOutgoingRow.DestinationChain,
						ToAddress:                     counterpartyOutgoingRow.DestinationAddress,
						MaybeToSmartContractAddress:   counterpartyOutgoingRow.MaybeDestinationSmartContractAddress,
						MaybeChannelId:                primptr.String("channel-1"),
						Amount:                        counterpartyOutgoingRow.Amount,
						MaybeDenom:                    counterpartyOutgoingRow.MaybeDenom,
						MaybeBridgeFeeAmount:          counterpartyOutgoingRow.MaybeBridgeFeeAmount,
						MaybeBridgeFeeDenom:           counterpartyOutgoingRow.MaybeBridgeFeeDenom,
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
					[]view.BridgePendingActivityReadRow{
						thisIncomingRow,
					},
					nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedOutgoing").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockCounterpartyBridgePendingActivitiesView.On("ListAllUnprocessedIncoming").Return(
					[]view.BridgePendingActivityReadRow{}, nil,
				)
				mockBridgeActivitiesView.On("FindByLinkId", thisIncomingRow.LinkId).Return(
					nil, rdb.ErrNoRows,
				)

				assertFunc = func() {
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
					mockBridgeActivitiesView.AssertNumberOfCalls(t, "Update", 0)
					mockCounterpartyBridgePendingActivitiesView.AssertNumberOfCalls(
						t, "UpdateToProcessed", 0,
					)
				}

				return mockThisRDbConn, mocks, assertFunc
			},
		},
	}

	for _, tc := range testCases {
		fmt.Print(tc.Name)
		mockThisRDbConn, mocks, assertFunc := tc.MockFunc(&tc.Config)

		projection := NewBridgeActivityMatcherProjection(mockThisRDbConn, tc.Config)
		onInitErr := projection.OnInit()
		require.NoError(t, onInitErr)

		execErr := projection.Exec()
		require.NoError(t, execErr)

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
