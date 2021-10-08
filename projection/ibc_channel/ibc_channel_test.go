package ibc_channel_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func NewIBCChannelProjection(rdbConn rdb.Conn) *ibc_channel.IBCChannel {

	return ibc_channel.NewIBCChannel(
		nil,
		rdbConn,
		ibc_channel.Config{
			EnableTxMsgTrace: false,
		},
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:      mock,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: sq.StatementBuilderType{},
	})

	return mock
}

func NewMockRDbTx() *test.MockRDbTx {
	mockTx := &test.MockRDbTx{}
	mockTx.On("ToHandle").Return(&rdb.Handle{
		Runner:      mockTx,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: pg.PostgresStmtBuilder,
	}).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}

func TestIBCChannel_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(mock *test.MockRDbConn) []*testify_mock.Mock
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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_clients (client_id,counterparty_chain_id) VALUES ($1,$2)",
					"ClientID",
					"ChainID",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var chainID string
				mockRowResult.On("Scan", &chainID).Return(nil)
				mockTx.On(
					"QueryRow",
					"SELECT counterparty_chain_id FROM view_ibc_clients WHERE client_id = $1",
					"ClientID",
				).Return(mockRowResult, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_connections (connection_id,client_id,counterparty_connection_id,counterparty_client_id,counterparty_chain_id) VALUES ($1,$2,$3,$4,$5)",
					"ConnectionID",
					"ClientID",
					"CounterpartyConnectionID",
					"CounterpartyClientID",
					"",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var chainID string
				mockRowResult.On("Scan", &chainID).Return(nil)
				mockTx.On(
					"QueryRow",
					"SELECT counterparty_chain_id FROM view_ibc_clients WHERE client_id = $1",
					"ClientID",
				).Return(mockRowResult, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_connections (connection_id,client_id,counterparty_connection_id,counterparty_client_id,counterparty_chain_id) VALUES ($1,$2,$3,$4,$5)",
					"ConnectionID",
					"ClientID",
					"CounterpartyConnectionID",
					"CounterpartyClientID",
					"",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
						ChannelID: "ChannelID",
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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_channels (channel_id,port_id,connection_id,counterparty_channel_id,counterparty_port_id,counterparty_chain_id,status,packet_ordering,last_in_packet_sequence,last_out_packet_sequence,total_transfer_in_count,total_transfer_out_count,total_transfer_out_success_count,total_transfer_out_success_rate,created_at_block_time,created_at_block_height,verified,description,last_activity_block_time,last_activity_block_height,bonded_tokens) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)",
					"ChannelID",
					"PortID",
					"",
					"CounterpartyChannelID",
					"CounterpartyPortID",
					"",
					false,
					"Ordering",
					int64(0),
					int64(0),
					int64(0),
					int64(0),
					int64(0),
					float64(0),
					int64(0),
					int64(0),
					false,
					"",
					int64(0),
					int64(0),
					"{\"onThisChain\":[],\"onCounterpartyChain\":[]}",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
						ChannelID: "ChannelID",
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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_channels (channel_id,port_id,connection_id,counterparty_channel_id,counterparty_port_id,counterparty_chain_id,status,packet_ordering,last_in_packet_sequence,last_out_packet_sequence,total_transfer_in_count,total_transfer_out_count,total_transfer_out_success_count,total_transfer_out_success_rate,created_at_block_time,created_at_block_height,verified,description,last_activity_block_time,last_activity_block_height,bonded_tokens) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)",
					"ChannelID",
					"PortID",
					"",
					"CounterpartyChannelID",
					"CounterpartyPortID",
					"",
					false,
					"Ordering",
					int64(0),
					int64(0),
					int64(0),
					int64(0),
					int64(0),
					float64(0),
					int64(0),
					int64(0),
					false,
					"",
					int64(0),
					int64(0),
					"{\"onThisChain\":[],\"onCounterpartyChain\":[]}",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var chainID string
				mockRowResult.On("Scan", &chainID).Return(nil)
				mockTx.On(
					"QueryRow",
					"SELECT counterparty_chain_id FROM view_ibc_connections WHERE connection_id = $1",
					"ConnectionID",
				).Return(mockRowResult, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET connection_id = $1, counterparty_chain_id = $2, counterparty_channel_id = $3, counterparty_port_id = $4, created_at_block_height = $5, created_at_block_time = $6 WHERE channel_id = $7",
					"ConnectionID",
					"",
					"CounterpartyChannelID",
					"CounterpartyPortID",
					int64(1),
					int64(0),
					"ChannelID",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET status = $1 WHERE channel_id = $2",
					true,
					"ChannelID",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var chainID string
				mockRowResult.On("Scan", &chainID).Return(nil)
				mockTx.On(
					"QueryRow",
					"SELECT counterparty_chain_id FROM view_ibc_connections WHERE connection_id = $1",
					"ConnectionID",
				).Return(mockRowResult, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET connection_id = $1, counterparty_chain_id = $2, counterparty_channel_id = $3, counterparty_port_id = $4, created_at_block_height = $5, created_at_block_time = $6 WHERE channel_id = $7",
					"ConnectionID",
					"",
					"CounterpartyChannelID",
					"CounterpartyPortID",
					int64(1),
					int64(0),
					"ChannelID",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET status = $1 WHERE channel_id = $2",
					true,
					"ChannelID",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

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
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET total_transfer_out_count = total_transfer_out_count+1 WHERE channel_id = $1",
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var totalTransferOutCount int64
				var totalTransferOutSuccessCount int64
				mockRowResult.On("Scan", &totalTransferOutCount, &totalTransferOutSuccessCount).Return(nil).Run(func(args testify_mock.Arguments) {
					totalTransferOutCount := args.Get(0).(*int64)
					totalTransferOutSuccessCount := args.Get(1).(*int64)
					*totalTransferOutCount = 1
					*totalTransferOutSuccessCount = 1
				})
				mockTx.On(
					"QueryRow",
					"SELECT total_transfer_out_count, total_transfer_out_success_count FROM view_ibc_channels WHERE channel_id = $1",
					"SourceChannel",
				).Return(mockRowResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET total_transfer_out_success_rate = $1 WHERE channel_id = $2",
					float64(1),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_out_packet_sequence = $1 WHERE channel_id = $2",
					uint64(1),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_activity_block_height = $1, last_activity_block_time = $2 WHERE channel_id = $3",
					int64(1),
					int64(0),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket",
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
								Amount: json.NewUint64(100),
							},
							Success: false,
						},
						PacketSequence: 1,
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("test"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET total_transfer_in_count = total_transfer_in_count+1 WHERE channel_id = $1",
					"DestinationChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_in_packet_sequence = $1 WHERE channel_id = $2",
					uint64(1),
					"DestinationChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_activity_block_height = $1, last_activity_block_time = $2 WHERE channel_id = $3",
					int64(1),
					int64(0),
					"DestinationChannel",
				).Return(mockExecResult, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var bondedTokensJSON string
				mockRowResult.On("Scan", &bondedTokensJSON).Return(nil).Run(func(args testify_mock.Arguments) {
					bondedTokensJSON := args.Get(0).(*string)
					*bondedTokensJSON = "{\"test\":\"\"}"
				})
				mockTx.On(
					"QueryRow",
					"SELECT bonded_tokens FROM view_ibc_channels WHERE channel_id = $1",
					"DestinationChannel",
				).Return(mockRowResult, nil)

				var existed bool
				mockRowResult.On("Scan", &existed).Return(nil).Run(func(args testify_mock.Arguments) {
					existed := args.Get(0).(*bool)
					*existed = false
				})
				mockTx.On(
					"QueryRow",
					"SELECT EXISTS ( SELECT * FROM view_ibc_denom_hash_mapping WHERE denom = $1 )",
					"DestinationPort/DestinationChannel/DENOM",
				).Return(mockRowResult, nil)

				mockTx.On(
					"Exec",
					"INSERT INTO view_ibc_denom_hash_mapping (denom,hash) VALUES ($1,$2)",
					"DestinationPort/DestinationChannel/DENOM",
					"293BDA936C232AE821E94F4CA8B95D60C09D13D2D84B0C027B2C250158A95381",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET bonded_tokens = $1 WHERE channel_id = $2",
					"{\"onThisChain\":[{\"denom\":\"DestinationPort/DestinationChannel/DENOM\",\"amount\":\"100\"}],\"onCounterpartyChain\":null}",
					"DestinationChannel",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement",
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
								Amount: json.NewUint64(100),
							},
							Success:    true,
							MaybeError: nil,
						},
						PacketSequence: 1,
					},
				},
			},
			MockFunc: func(mockConn *test.MockRDbConn) (mocks []*testify_mock.Mock) {
				mockTx := NewMockRDbTx()
				mocks = append(mocks, &mockTx.Mock)
				mockConn.On("Begin").Return(mockTx, nil)

				mockExecResult := &test.MockRDbExecResult{}
				mocks = append(mocks, &mockExecResult.Mock)
				mockExecResult.On("RowsAffected").Return(int64(1))
				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET total_transfer_out_success_count = total_transfer_out_success_count+1 WHERE channel_id = $1",
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_out_packet_sequence = $1 WHERE channel_id = $2",
					uint64(1),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET last_activity_block_height = $1, last_activity_block_time = $2 WHERE channel_id = $3",
					int64(1),
					int64(0),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockRowResult := &test.MockRDbRowResult{}
				mocks = append(mocks, &mockRowResult.Mock)
				var totalTransferOutCount int64
				var totalTransferOutSuccessCount int64
				mockRowResult.On("Scan", &totalTransferOutCount, &totalTransferOutSuccessCount).Return(nil).Run(func(args testify_mock.Arguments) {
					totalTransferOutCount := args.Get(0).(*int64)
					totalTransferOutSuccessCount := args.Get(1).(*int64)
					*totalTransferOutCount = 1
					*totalTransferOutSuccessCount = 1
				})
				mockTx.On(
					"QueryRow",
					"SELECT total_transfer_out_count, total_transfer_out_success_count FROM view_ibc_channels WHERE channel_id = $1",
					"SourceChannel",
				).Return(mockRowResult, nil)

				var bondedTokensJSON string
				mockRowResult.On("Scan", &bondedTokensJSON).Return(nil).Run(func(args testify_mock.Arguments) {
					bondedTokensJSON := args.Get(0).(*string)
					*bondedTokensJSON = "{\"test\":\"\"}"
				})
				mockTx.On(
					"QueryRow",
					"SELECT bonded_tokens FROM view_ibc_channels WHERE channel_id = $1",
					"SourceChannel",
				).Return(mockRowResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET total_transfer_out_success_rate = $1 WHERE channel_id = $2",
					float64(1),
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockTx.On(
					"Exec",
					"UPDATE view_ibc_channels SET bonded_tokens = $1 WHERE channel_id = $2",
					"{\"onThisChain\":null,\"onCounterpartyChain\":[{\"denom\":\"//DENOM\",\"amount\":\"100\"}]}",
					"SourceChannel",
				).Return(mockExecResult, nil)

				mockUpdateLastHandledEventHeight(mockTx)

				return mocks
			},
		},
	}

	for _, tc := range testCases {
		mockRDbConn := NewMockRDbConn()
		mocks := tc.MockFunc(mockRDbConn)
		mocks = append(mocks, &mockRDbConn.Mock)

		projection := NewIBCChannelProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}

func mockUpdateLastHandledEventHeight(mockTx *test.MockRDbTx) {
	mockRowResult := &test.MockRDbRowResult{}

	var lastHandledEventHeight int64
	mockRowResult.On("Scan", &lastHandledEventHeight).Return(nil)
	mockTx.On(
		"QueryRow",
		"SELECT last_handled_event_height FROM projections WHERE id = $1",
		"IBCChannel",
	).Return(mockRowResult, nil)

	mockExecResult := &test.MockRDbExecResult{}
	mockExecResult.On("RowsAffected").Return(int64(1))
	mockTx.On(
		"Exec",
		"UPDATE projections SET id = $1, last_handled_event_height = $2 WHERE id = $3",
		"IBCChannel",
		int64(1),
		"IBCChannel",
	).Return(mockExecResult, nil)
}
