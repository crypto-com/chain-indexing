package ibc_channel_message_test

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
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message/view"
	usecase_event "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func NewIBCChannelMessageProjection(rdbConn rdb.Conn) *ibc_channel_message.IBCChannelMessage {

	return ibc_channel_message.NewIBCChannelMessage(
		nil,
		rdbConn,
	)
}

func NewMockRDbConn() *test.MockRDbConn {
	mock := test.NewMockRDBbConn()
	mock.On("ToHandle").Return(&rdb.Handle{
		Runner:      mock,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: sq.StatementBuilderType{},
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

func TestIBCChannelMessage_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Events   []entity_event.Event
		MockFunc func(events []entity_event.Event) []*testify_mock.Mock
	}{
		{
			Name: "HandleMsgIBCChannelOpenInit",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCChannelOpenInit{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_CHANNEL_OPEN_INIT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenInitParams{
						ChannelID: "ChannelID",
						RawMsgChannelOpenInit: ibc_model.RawMsgChannelOpenInit{
							Signer: "Signer",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCChannelOpenInit)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "ChannelID",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MessageType:     "MsgChannelOpenInit",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "ChannelID"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "ChannelID", "MsgChannelOpenInit"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenTry",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCChannelOpenTry{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_CHANNEL_OPEN_TRY,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenTryParams{
						ChannelID: "ChannelID",
						RawMsgChannelOpenTry: ibc_model.RawMsgChannelOpenTry{
							Signer: "Signer",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCChannelOpenTry)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "ChannelID",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MessageType:     "MsgChannelOpenTry",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "ChannelID"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "ChannelID", "MsgChannelOpenTry"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenAck",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCChannelOpenAck{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_CHANNEL_OPEN_ACK,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenAckParams{
						RawMsgChannelOpenAck: ibc_model.RawMsgChannelOpenAck{
							ChannelID: "ChannelID",
							Signer:    "Signer",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCChannelOpenAck)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "ChannelID",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MessageType:     "MsgChannelOpenAck",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "ChannelID"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "ChannelID", "MsgChannelOpenAck"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCChannelOpenConfirm",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCChannelOpenConfirm{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_CHANNEL_OPEN_CONFIRM,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgChannelOpenConfirmParams{
						RawMsgChannelOpenConfirm: ibc_model.RawMsgChannelOpenConfirm{
							ChannelID: "ChannelID",
							Signer:    "Signer",
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCChannelOpenConfirm)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "ChannelID",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MessageType:     "MsgChannelOpenConfirm",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "ChannelID"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "ChannelID", "MsgChannelOpenConfirm"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTransfer",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCTransferTransfer{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_TRANSFER_TRANSFER,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTransferParams{
						RawMsgTransfer: ibc_model.RawMsgTransfer{
							SourceChannel: "SourceChannel",
							Sender:        "Sender",
							Receiver:      "Receiver",
							Token: ibc_model.MsgTransferToken{
								Denom:  "Denom",
								Amount: uint64(1000),
							},
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCTransferTransfer)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "SourceChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeSender:     primptr.String("Sender"),
						MaybeReceiver:   primptr.String("Receiver"),
						MaybeDenom:      primptr.String("Denom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgTransfer",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "SourceChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "SourceChannel", "MsgTransfer"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket WITHOUT error",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCRecvPacket{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_RECV_PACKET,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgRecvPacketParams{
						RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							Success: true,
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewUint64(1000),
							},
						},
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("MaybeResult"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCRecvPacket)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "DestinationChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeError:      nil,
						MaybeSuccess:    primptr.Bool(true),
						MaybeSender:     primptr.String("Sender"),
						MaybeReceiver:   primptr.String("Receiver"),
						MaybeDenom:      primptr.String("Denom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgRecvPacket",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "DestinationChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "DestinationChannel", "MsgRecvPacket"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCRecvPacket WITH error",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCRecvPacket{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_RECV_PACKET,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgRecvPacketParams{
						RawMsgRecvPacket: ibc_model.RawMsgRecvPacket{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							Success: false,
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewUint64(1000),
							},
						},
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: nil,
							MaybeError:  primptr.String("MaybeError"),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCRecvPacket)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "DestinationChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeError:      primptr.String("MaybeError"),
						MaybeSuccess:    primptr.Bool(false),
						MaybeSender:     primptr.String("Sender"),
						MaybeReceiver:   primptr.String("Receiver"),
						MaybeDenom:      primptr.String("Denom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgRecvPacket",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "DestinationChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "DestinationChannel", "MsgRecvPacket"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement WITHOUT error",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCAcknowledgement{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_ACKNOWLEDGEMENT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							Success: true,
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewUint64(1000),
							},
							MaybeError: nil,
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCAcknowledgement)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "SourceChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeError:      nil,
						MaybeSuccess:    primptr.Bool(true),
						MaybeSender:     primptr.String("Sender"),
						MaybeReceiver:   primptr.String("Receiver"),
						MaybeDenom:      primptr.String("Denom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgAcknowledgement",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "SourceChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "SourceChannel", "MsgAcknowledgement"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCAcknowledgement WITH error",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCAcknowledgement{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_ACKNOWLEDGEMENT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							Success: false,
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "Sender",
								Receiver: "Receiver",
								Denom:    "Denom",
								Amount:   json.NewUint64(1000),
							},
							MaybeError: primptr.String("MaybeError"),
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCAcknowledgement)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "SourceChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeError:      primptr.String("MaybeError"),
						MaybeSuccess:    primptr.Bool(false),
						MaybeSender:     primptr.String("Sender"),
						MaybeReceiver:   primptr.String("Receiver"),
						MaybeDenom:      primptr.String("Denom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgAcknowledgement",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "SourceChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "SourceChannel", "MsgAcknowledgement"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTimeout",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCTimeout{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_TIMEOUT,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutParams{
						RawMsgTimeout: ibc_model.RawMsgTimeout{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
							RefundReceiver: "RefundReceiver",
							RefundDenom:    "RefundDenom",
							RefundAmount:   1000,
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCTimeout)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "SourceChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeReceiver:   primptr.String("RefundReceiver"),
						MaybeDenom:      primptr.String("RefundDenom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgTimeout",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "SourceChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "SourceChannel", "MsgTimeout"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
					return nil
				}

				return mocks
			},
		},
		{
			Name: "HandleMsgIBCTimeoutOnClose",
			Events: []entity_event.Event{
				&usecase_event.MsgIBCTimeoutOnClose{
					MsgBase: usecase_event.NewMsgBase(usecase_event.MsgBaseParams{
						MsgName: usecase_event.MSG_IBC_TIMEOUT_ON_CLOSE,
						Version: 1,
						MsgCommonParams: usecase_event.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   true,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTimeoutOnCloseParams{
						RawMsgTimeoutOnClose: ibc_model.RawMsgTimeoutOnClose{
							Signer: "Signer",
							Packet: ibc_model.Packet{
								SourceChannel:      "SourceChannel",
								DestinationChannel: "DestinationChannel",
							},
						},
						MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
							RefundReceiver: "RefundReceiver",
							RefundDenom:    "RefundDenom",
							RefundAmount:   1000,
						},
					},
				},
			},
			MockFunc: func(events []entity_event.Event) (mocks []*testify_mock.Mock) {
				typedEvent := events[0].(*usecase_event.MsgIBCTimeoutOnClose)

				mockIBCChannelMessageView := &view.MockIBCChannelMessageView{}
				mocks = append(mocks, &mockIBCChannelMessageView.Mock)
				mockIBCChannelMessageView.
					On("Insert", &view.IBCChannelMessageRow{
						ChannelID:       "SourceChannel",
						BlockHeight:     int64(1),
						BlockTime:       utctime.UTCTime{},
						TransactionHash: "TxHash",
						MaybeRelayer:    primptr.String("Signer"),
						MaybeReceiver:   primptr.String("RefundReceiver"),
						MaybeDenom:      primptr.String("RefundDenom"),
						MaybeAmount:     primptr.String("1000"),
						MessageType:     "MsgTimeoutOnClose",
						Message:         typedEvent,
					}).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessages = func(handle *rdb.Handle) view.IBCChannelMessages {
					return mockIBCChannelMessageView
				}

				mockIBCChannelMessageTotalView := &view.MockIBCChannelMessageTotalView{}
				mocks = append(mocks, &mockIBCChannelMessageTotalView.Mock)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:-", "SourceChannel"), int64(1)).
					Return(nil)
				mockIBCChannelMessageTotalView.
					On("Increment", fmt.Sprintf("%s:%s", "SourceChannel", "MsgTimeoutOnClose"), int64(1)).
					Return(nil)

				ibc_channel_message.NewIBCChannelMessagesTotal = func(handle *rdb.Handle) view.IBCChannelMessagesTotal {
					return mockIBCChannelMessageTotalView
				}

				ibc_channel_message.UpdateLastHandledEventHeight = func(_ *ibc_channel_message.IBCChannelMessage, _ *rdb.Handle, _ int64) error {
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

		mocks := tc.MockFunc(tc.Events)
		mocks = append(mocks, &mockRDbConn.Mock)
		mocks = append(mocks, &mockTx.Mock)

		projection := NewIBCChannelMessageProjection(mockRDbConn)
		err := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, err)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}

		fmt.Println(tc.Name, "Passed")
	}
}
