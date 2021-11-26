package ibc_channel_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	ibc_channel_view "github.com/crypto-com/chain-indexing/projection/ibc_channel/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func NewIBCChannelProjection(rdbConn rdb.Conn) *ibc_channel.IBCChannel {

	return ibc_channel.NewIBCChannel(
		nil,
		rdbConn,
		&ibc_channel.Config{
			EnableTxMsgTrace: false,
		},
		nil,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:   mock,
		TypeConv: &pg.PgxTypeConv{},
		StmtBuilder: &rdb.StatementBuilder{
			StatementBuilderType: sq.StatementBuilderType{},
			PlaceholderFormat:    nil,
		},
	})

	return mock
}

func NewMockRDbTx() *test.MockRDbTx {
	mockTx := &test.MockRDbTx{}
	mockTx.On("ToHandle").Return(nil).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}

func TestIBCChannel_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func() []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgIBCCreateClient",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCCreateClient{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CREATE_CLIENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgCreateClientParams{
						MaybeTendermintLightClient: &ibc_model.TendermintLightClient{
							TendermintClientState: ibc_model.TendermintLightClientState{
								ChainID: "ChainID",
							},
							TendermintLightClientConsensusState: ibc_model.TendermintLightClientConsensusState{},
						},
						ClientID: "ClientID",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcClientsView := ibc_channel_view.NewMockIBCClientsView(nil).(*ibc_channel_view.MockIBCClientsView)
				mocks = append(mocks, &mockIbcClientsView.Mock)

				ibc_channel.NewIBCClients = func(_ *rdb.Handle) ibc_channel_view.IBCClients {
					return mockIbcClientsView
				}
				mockIbcClientsView.On(
					"Insert",
					&ibc_channel_view.IBCClientRow{
						ClientID:            "ClientID",
						CounterpartyChainID: "ChainID",
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenInit",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCConnectionOpenInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenInitParams{
						ConnectionID: "ConnectionID",
						RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
							ClientID: "ClientID",
							Counterparty: ibc_model.ConnectionCounterparty{
								ConnectionID: "CounterpartyConnectionID",
								ClientID:     "CounterpartyClientID",
							},
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcClientsView := ibc_channel_view.NewMockIBCClientsView(nil).(*ibc_channel_view.MockIBCClientsView)
				mocks = append(mocks, &mockIbcClientsView.Mock)

				ibc_channel.NewIBCClients = func(_ *rdb.Handle) ibc_channel_view.IBCClients {
					return mockIbcClientsView
				}

				mockIbcClientsView.On(
					"FindCounterpartyChainIDBy",
					"ClientID",
				).Return("CounterpartyChainID", nil)

				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On(
					"Insert",
					&ibc_channel_view.IBCConnectionRow{
						ConnectionID:             "ConnectionID",
						ClientID:                 "ClientID",
						CounterpartyConnectionID: "CounterpartyConnectionID",
						CounterpartyClientID:     "CounterpartyClientID",
						CounterpartyChainID:      "CounterpartyChainID",
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenTry",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCConnectionOpenTry{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_TRY,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenTryParams{
						ConnectionID: "ConnectionID",
						MsgConnectionOpenTryBaseParams: ibc_model.MsgConnectionOpenTryBaseParams{
							ClientID: "ClientID",
							Counterparty: ibc_model.ConnectionCounterparty{
								ConnectionID: "CounterpartyConnectionID",
								ClientID:     "CounterpartyClientID",
							},
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcClientsView := ibc_channel_view.NewMockIBCClientsView(nil).(*ibc_channel_view.MockIBCClientsView)
				mocks = append(mocks, &mockIbcClientsView.Mock)

				ibc_channel.NewIBCClients = func(_ *rdb.Handle) ibc_channel_view.IBCClients {
					return mockIbcClientsView
				}

				mockIbcClientsView.On(
					"FindCounterpartyChainIDBy",
					"ClientID",
				).Return("CounterpartyChainID", nil)

				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On(
					"Insert",
					&ibc_channel_view.IBCConnectionRow{
						ConnectionID:             "ConnectionID",
						ClientID:                 "ClientID",
						CounterpartyConnectionID: "CounterpartyConnectionID",
						CounterpartyClientID:     "CounterpartyClientID",
						CounterpartyChainID:      "CounterpartyChainID",
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenConfirm",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCConnectionOpenConfirm{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenConfirmParams{
						ClientID:                 "ClientID",
						CounterpartyClientID:     "CounterpartyClientID",
						CounterpartyConnectionID: "CounterpartyConnectionID",
						RawMsgConnectionOpenConfirm: ibc_model.RawMsgConnectionOpenConfirm{
							ConnectionID: "ConnectionID",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcClientsView := ibc_channel_view.NewMockIBCClientsView(nil).(*ibc_channel_view.MockIBCClientsView)
				mocks = append(mocks, &mockIbcClientsView.Mock)

				ibc_channel.NewIBCClients = func(_ *rdb.Handle) ibc_channel_view.IBCClients {
					return mockIbcClientsView
				}

				mockIbcClientsView.On(
					"FindCounterpartyChainIDBy",
					"ClientID",
				).Return("CounterpartyChainID", nil)

				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On(
					"Update",
					&ibc_channel_view.IBCConnectionRow{
						ConnectionID:             "ConnectionID",
						ClientID:                 "ClientID",
						CounterpartyConnectionID: "CounterpartyConnectionID",
						CounterpartyClientID:     "CounterpartyClientID",
						CounterpartyChainID:      "CounterpartyChainID",
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCConnectionOpenAck",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCConnectionOpenAck{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CONNECTION_OPEN_ACK,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgConnectionOpenAckParams{
						ClientID:             "ClientID",
						CounterpartyClientID: "CounterpartyClientID",
						MsgConnectionOpenAckBaseParams: ibc_model.MsgConnectionOpenAckBaseParams{
							ConnectionID:             "ConnectionID",
							CounterpartyConnectionID: "CounterpartyConnectionID",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcClientsView := ibc_channel_view.NewMockIBCClientsView(nil).(*ibc_channel_view.MockIBCClientsView)
				mocks = append(mocks, &mockIbcClientsView.Mock)

				ibc_channel.NewIBCClients = func(_ *rdb.Handle) ibc_channel_view.IBCClients {
					return mockIbcClientsView
				}

				mockIbcClientsView.On(
					"FindCounterpartyChainIDBy",
					"ClientID",
				).Return("CounterpartyChainID", nil)

				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On(
					"Update",
					&ibc_channel_view.IBCConnectionRow{
						ConnectionID:             "ConnectionID",
						ClientID:                 "ClientID",
						CounterpartyConnectionID: "CounterpartyConnectionID",
						CounterpartyClientID:     "CounterpartyClientID",
						CounterpartyChainID:      "CounterpartyChainID",
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenInit",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelOpenInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenInitParams{
						ChannelID:    "ChannelID",
						ConnectionID: "ConnectionID",
						RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
							PortID: "PortID",
							Channel: ibc_model.Channel{
								Ordering: "Ordering",
								Counterparty: ibc_model.ChannelCounterparty{
									PortID:    "CounterpartyPortID",
									ChannelID: "CounterpartyChannelID",
								},
							},
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On("FindCounterpartyChainIDBy", "ConnectionID").Return("CounterpartyChainID", nil)

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"Insert",
					&ibc_channel_view.IBCChannelRow{
						ChannelID:                 "ChannelID",
						PortID:                    "PortID",
						ConnectionID:              "ConnectionID",
						CounterpartyChannelID:     "CounterpartyChannelID",
						CounterpartyPortID:        "CounterpartyPortID",
						CounterpartyChainID:       "CounterpartyChainID",
						Status:                    "NOT_ESTABLISHED",
						PacketOrdering:            "Ordering",
						LastInPacketSequence:      0,
						LastOutPacketSequence:     0,
						TotalRelayInCount:         0,
						TotalRelayOutCount:        0,
						TotalRelayOutSuccessCount: 0,
						TotalRelayOutSuccessRate:  0,
						CreatedAtBlockTime:        utctime.UTCTime{},
						CreatedAtBlockHeight:      0,
						Verified:                  false,
						Description:               "",
						LastActivityBlockTime:     utctime.UTCTime{},
						LastActivityBlockHeight:   0,
						BondedTokens: ibc_channel_view.BondedTokens{
							OnThisChain:         []ibc_channel_view.BondedToken{},
							OnCounterpartyChain: []ibc_channel_view.BondedToken{},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenTry",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelOpenTry{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_TRY,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenTryParams{
						ChannelID:    "ChannelID",
						ConnectionID: "ConnectionID",
						RawMsgChannelOpenTry: ibc_model.RawMsgChannelOpenTry{
							PortID: "PortID",
							Channel: ibc_model.Channel{
								Ordering: "Ordering",
								Counterparty: ibc_model.ChannelCounterparty{
									PortID:    "CounterpartyPortID",
									ChannelID: "CounterpartyChannelID",
								},
							},
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On("FindCounterpartyChainIDBy", "ConnectionID").Return("CounterpartyChainID", nil)

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"Insert",
					&ibc_channel_view.IBCChannelRow{
						ChannelID:                 "ChannelID",
						PortID:                    "PortID",
						ConnectionID:              "ConnectionID",
						CounterpartyChannelID:     "CounterpartyChannelID",
						CounterpartyPortID:        "CounterpartyPortID",
						CounterpartyChainID:       "CounterpartyChainID",
						Status:                    "NOT_ESTABLISHED",
						PacketOrdering:            "Ordering",
						LastInPacketSequence:      0,
						LastOutPacketSequence:     0,
						TotalRelayInCount:         0,
						TotalRelayOutCount:        0,
						TotalRelayOutSuccessCount: 0,
						TotalRelayOutSuccessRate:  0,
						CreatedAtBlockTime:        utctime.UTCTime{},
						CreatedAtBlockHeight:      0,
						Verified:                  false,
						Description:               "",
						LastActivityBlockTime:     utctime.UTCTime{},
						LastActivityBlockHeight:   0,
						BondedTokens: ibc_channel_view.BondedTokens{
							OnThisChain:         []ibc_channel_view.BondedToken{},
							OnCounterpartyChain: []ibc_channel_view.BondedToken{},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenAck",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelOpenAck{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_ACK,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenAckParams{
						CounterpartyPortID: "CounterpartyPortID",
						ConnectionID:       "ConnectionID",
						RawMsgChannelOpenAck: ibc_model.RawMsgChannelOpenAck{
							PortID:                "PortID",
							ChannelID:             "ChannelID",
							CounterpartyChannelID: "CounterpartyChannelID",
							CounterpartyVersion:   "CounterpartyVersion",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On("FindCounterpartyChainIDBy", "ConnectionID").Return("CounterpartyChainID", nil)

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateFactualColumns",
					&ibc_channel_view.IBCChannelRow{
						ChannelID:                 "ChannelID",
						PortID:                    "PortID",
						ConnectionID:              "ConnectionID",
						CounterpartyChannelID:     "CounterpartyChannelID",
						CounterpartyPortID:        "CounterpartyPortID",
						CounterpartyChainID:       "CounterpartyChainID",
						Status:                    "",
						PacketOrdering:            "",
						LastInPacketSequence:      0,
						LastOutPacketSequence:     0,
						TotalRelayInCount:         0,
						TotalRelayOutCount:        0,
						TotalRelayOutSuccessCount: 0,
						TotalRelayOutSuccessRate:  0,
						CreatedAtBlockTime:        utctime.UTCTime{},
						CreatedAtBlockHeight:      1,
						Verified:                  false,
						Description:               "",
						LastActivityBlockTime:     utctime.UTCTime{},
						LastActivityBlockHeight:   0,
						BondedTokens: ibc_channel_view.BondedTokens{
							OnThisChain:         []ibc_channel_view.BondedToken(nil),
							OnCounterpartyChain: []ibc_channel_view.BondedToken(nil),
						},
					},
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateStatus",
					"ChannelID",
					"OPENED",
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenConfirm",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelOpenConfirm{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenConfirmParams{
						CounterpartyPortID:    "CounterpartyPortID",
						ConnectionID:          "ConnectionID",
						CounterpartyChannelID: "CounterpartyChannelID",
						RawMsgChannelOpenConfirm: ibc_model.RawMsgChannelOpenConfirm{
							PortID:    "PortID",
							ChannelID: "ChannelID",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcConnectionsView := ibc_channel_view.NewMockIBCConnectionsView(nil).(*ibc_channel_view.MockIBCConnectionsView)
				mocks = append(mocks, &mockIbcConnectionsView.Mock)

				ibc_channel.NewIBCConnections = func(_ *rdb.Handle) ibc_channel_view.IBCConnections {
					return mockIbcConnectionsView
				}

				mockIbcConnectionsView.On("FindCounterpartyChainIDBy", "ConnectionID").Return("CounterpartyChainID", nil)

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateFactualColumns",
					&ibc_channel_view.IBCChannelRow{
						ChannelID:                 "ChannelID",
						PortID:                    "PortID",
						ConnectionID:              "ConnectionID",
						CounterpartyChannelID:     "CounterpartyChannelID",
						CounterpartyPortID:        "CounterpartyPortID",
						CounterpartyChainID:       "CounterpartyChainID",
						Status:                    "",
						PacketOrdering:            "",
						LastInPacketSequence:      0,
						LastOutPacketSequence:     0,
						TotalRelayInCount:         0,
						TotalRelayOutCount:        0,
						TotalRelayOutSuccessCount: 0,
						TotalRelayOutSuccessRate:  0,
						CreatedAtBlockTime:        utctime.UTCTime{},
						CreatedAtBlockHeight:      1,
						Verified:                  false,
						Description:               "",
						LastActivityBlockTime:     utctime.UTCTime{},
						LastActivityBlockHeight:   0,
						BondedTokens: ibc_channel_view.BondedTokens{
							OnThisChain:         []ibc_channel_view.BondedToken(nil),
							OnCounterpartyChain: []ibc_channel_view.BondedToken(nil),
						},
					},
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateStatus",
					"ChannelID",
					"OPENED",
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTransferTransfer",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCTransferTransfer{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TRANSFER_TRANSFER,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTransferParams{
						RawMsgTransfer: ibc_model.RawMsgTransfer{
							SourceChannel: "SourceChannel",
							SourcePort:    "SourcePort",
						},
						PacketData: ibc_model.FungibleTokenPacketData{
							Amount: json.NewNumericStringFromUint64(100),
							Denom:  "DENOM",
						},
						DestinationChannel: "DestinationChannel",
						DestinationPort:    "DestinationPort",
						PacketSequence:     1,
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"Increment",
					"SourceChannel",
					"total_relay_out_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateTotalRelayOutSuccessRate",
					"SourceChannel",
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateSequence",
					"SourceChannel",
					"last_out_packet_sequence",
					uint64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"SourceChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"FindBondedTokensBy",
					"SourceChannel",
				).Return(
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
					nil,
				)

				mockIbcChannelsView.On(
					"UpdateBondedTokens",
					"SourceChannel",
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
							{
								Denom:  "DestinationPort/DestinationChannel/DENOM",
								Amount: coin.NewInt(100),
							},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket WITHOUT error",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCRecvPacket{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_RECV_PACKET,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgRecvPacketParams{
						RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
							Packet: ibc_model.Packet{
								Sequence:           "PacketSequence",
								DestinationPort:    "DestinationPort",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Denom:  "DENOM",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Success: true,
						},
						PacketSequence: 1,
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("test"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"Increment",
					"DestinationChannel",
					"total_relay_in_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateSequence",
					"DestinationChannel",
					"last_in_packet_sequence",
					uint64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"DestinationChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"FindBondedTokensBy",
					"DestinationChannel",
				).Return(
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
					nil,
				)

				mockIbcDenomHashMappingView := ibc_channel_view.NewMockIBCDenomHashMapping(nil).(*ibc_channel_view.MockIBCDenomHashMappingView)
				mocks = append(mocks, &mockIbcDenomHashMappingView.Mock)

				ibc_channel.NewIBCDenomHashMapping = func(_ *rdb.Handle) ibc_channel_view.IBCDenomHashMapping {
					return mockIbcDenomHashMappingView
				}

				mockIbcDenomHashMappingView.On("IfDenomExist", "DestinationPort/DestinationChannel/DENOM").Return(false, nil)

				mockIbcDenomHashMappingView.On(
					"Insert",
					&ibc_channel_view.IBCDenomHashMappingRow{
						Denom: "DestinationPort/DestinationChannel/DENOM",
						Hash:  "293BDA936C232AE821E94F4CA8B95D60C09D13D2D84B0C027B2C250158A95381",
					},
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateBondedTokens",
					"DestinationChannel",
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
							{
								Denom:  "DestinationPort/DestinationChannel/DENOM",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket WITH error",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCRecvPacket{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_RECV_PACKET,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgRecvPacketParams{
						RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
							Packet: ibc_model.Packet{
								Sequence:           "PacketSequence",
								DestinationPort:    "DestinationPort",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Denom:  "DENOM",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Success: false,
						},
						PacketSequence: 1,
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: nil,
							MaybeError:  primptr.String("MaybeError"),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"Increment",
					"DestinationChannel",
					"total_relay_in_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateSequence",
					"DestinationChannel",
					"last_in_packet_sequence",
					uint64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"DestinationChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement WITHOUT error",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCAcknowledgement{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_ACKNOWLEDGEMENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:      "PacketSequence",
								SourceChannel: "SourceChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Denom:  "DENOM",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Success:    true,
							MaybeError: nil,
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"SourceChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"Increment",
					"SourceChannel",
					"total_relay_out_success_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateTotalRelayOutSuccessRate",
					"SourceChannel",
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement WITH error",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCAcknowledgement{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_ACKNOWLEDGEMENT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "PacketSequence",
								SourceChannel:      "SourceChannel",
								DestinationPort:    "DestinationPort",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Denom:  "DENOM",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Success:    false,
							MaybeError: primptr.String("MaybeError"),
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"SourceChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"Increment",
					"SourceChannel",
					"total_relay_out_success_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateTotalRelayOutSuccessRate",
					"SourceChannel",
				).Return(nil)

				mockIbcChannelsView.On(
					"FindBondedTokensBy",
					"SourceChannel",
				).Return(
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
							{
								Denom:  "DestinationPort/DestinationChannel/DENOM",
								Amount: coin.NewInt(100),
							},
						},
					},
					nil,
				)

				mockIbcChannelsView.On(
					"UpdateBondedTokens",
					"SourceChannel",
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgTimeout",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCTimeout{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TIMEOUT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutParams{
						RawMsgTimeout: ibc_model.RawMsgTimeout{
							Packet: ibc_model.Packet{
								Sequence:           "PacketSequence",
								SourceChannel:      "SourceChannel",
								DestinationPort:    "DestinationPort",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
							RefundDenom:  "DENOM",
							RefundAmount: json.NewNumericStringFromUint64(100),
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"SourceChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"Increment",
					"SourceChannel",
					"total_relay_out_success_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateTotalRelayOutSuccessRate",
					"SourceChannel",
				).Return(nil)

				mockIbcChannelsView.On(
					"FindBondedTokensBy",
					"SourceChannel",
				).Return(
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
							{
								Denom:  "DestinationPort/DestinationChannel/DENOM",
								Amount: coin.NewInt(100),
							},
						},
					},
					nil,
				)

				mockIbcChannelsView.On(
					"UpdateBondedTokens",
					"SourceChannel",
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgTimeoutOnClose",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCTimeoutOnClose{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TIMEOUT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutOnCloseParams{
						RawMsgTimeoutOnClose: ibc_model.RawMsgTimeoutOnClose{
							Packet: ibc_model.Packet{
								Sequence:           "PacketSequence",
								SourceChannel:      "SourceChannel",
								DestinationPort:    "DestinationPort",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
							RefundDenom:  "DENOM",
							RefundAmount: json.NewNumericStringFromUint64(100),
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {
				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"SourceChannel",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"Increment",
					"SourceChannel",
					"total_relay_out_success_count",
					int64(1),
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateTotalRelayOutSuccessRate",
					"SourceChannel",
				).Return(nil)

				mockIbcChannelsView.On(
					"FindBondedTokensBy",
					"SourceChannel",
				).Return(
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
							{
								Denom:  "DestinationPort/DestinationChannel/DENOM",
								Amount: coin.NewInt(100),
							},
						},
					},
					nil,
				)

				mockIbcChannelsView.On(
					"UpdateBondedTokens",
					"SourceChannel",
					&ibc_channel_view.BondedTokens{
						OnThisChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnThisChainDenom",
								Amount: coin.NewInt(100),
							},
						},
						OnCounterpartyChain: []ibc_channel_view.BondedToken{
							{
								Denom:  "OnCounterpartyChainDenom",
								Amount: coin.NewInt(1000),
							},
						},
					},
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelCloseInit",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelCloseInit{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelCloseInitParams{
						CounterpartyPortID:    "CounterpartyPortID",
						ConnectionID:          "ConnectionID",
						CounterpartyChannelID: "CounterpartyChannelID",
						RawMsgChannelCloseInit: ibc_model.RawMsgChannelCloseInit{
							PortID:    "PortID",
							ChannelID: "ChannelID",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateStatus",
					"ChannelID",
					"CLOSED",
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"ChannelID",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelCloseConfirm",
			Events: []entity_event.Event{
				&event_usecase.MsgIBCChannelCloseConfirm{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelCloseConfirmParams{
						CounterpartyPortID:    "CounterpartyPortID",
						ConnectionID:          "ConnectionID",
						CounterpartyChannelID: "CounterpartyChannelID",
						RawMsgChannelCloseConfirm: ibc_model.RawMsgChannelCloseConfirm{
							PortID:    "PortID",
							ChannelID: "ChannelID",
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock) {

				mockIbcChannelsView := ibc_channel_view.NewMockIBCChannelsView(nil).(*ibc_channel_view.MockIBCChannelsView)
				mocks = append(mocks, &mockIbcChannelsView.Mock)

				ibc_channel.NewIBCChannels = func(_ *rdb.Handle) ibc_channel_view.IBCChannels {
					return mockIbcChannelsView
				}

				mockIbcChannelsView.On(
					"UpdateStatus",
					"ChannelID",
					"CLOSED",
				).Return(nil)

				mockIbcChannelsView.On(
					"UpdateLastActivityTimeAndHeight",
					"ChannelID",
					utctime.UTCTime{},
					int64(1),
				).Return(nil)

				ibc_channel.UpdateLastHandledEventHeight = func(_ *ibc_channel.IBCChannel, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks := tc.MockFunc()

		projection := NewIBCChannelProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
