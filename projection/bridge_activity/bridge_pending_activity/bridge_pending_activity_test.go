package bridge_pending_activity_test

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/rdb/test"
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/json"
	test_logger "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_pending_activity"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/types"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func TestBridgePendingActivity_HandleEvents(t *testing.T) {
	testCases := []struct {
		Name     string
		Config   bridge_pending_activity.Config
		Events   []entity_event.Event
		MockFunc func() (mocks []*testify_mock.Mock, assertFunc func())
	}{
		{
			Name: "It should not handle IBC transfer from non-listened source channel",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
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
							SourcePort:    "transfer",
							SourceChannel: "channel-1",
							Token: ibc_model.MsgTransferToken{
								Denom:  "basecro",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Sender:   "from",
							Receiver: "to",
							TimeoutHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							TimeoutTimestamp: "",
						},
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
						PacketData: ibc_model.FungibleTokenPacketData{
							Sender:   "from",
							Receiver: "to",
							Denom:    "basecro",
							Amount:   json.NewNumericStringFromUint64(100),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgTransfer event on all listened source channels",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty0",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
					{
						ChainName:      "counterparty1",
						ChannelId:      "channel-1",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
							SourcePort:    "transfer",
							SourceChannel: "channel-0",
							Token: ibc_model.MsgTransferToken{
								Denom:  "basecro",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Sender:   "from",
							Receiver: "to",
							TimeoutHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							TimeoutTimestamp: "",
						},
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
						PacketData: ibc_model.FungibleTokenPacketData{
							Sender:   "from",
							Receiver: "to",
							Denom:    "basecro",
							Amount:   json.NewNumericStringFromUint64(100),
						},
					},
				},
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
							SourcePort:    "transfer",
							SourceChannel: "channel-1",
							Token: ibc_model.MsgTransferToken{
								Denom:  "basecro",
								Amount: json.NewNumericStringFromUint64(200),
							},
							Sender:   "from",
							Receiver: "to",
							TimeoutHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							TimeoutTimestamp: "",
						},
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
						PacketData: ibc_model.FungibleTokenPacketData{
							Sender:   "from",
							Receiver: "to",
							Denom:    "basecro",
							Amount:   json.NewNumericStringFromUint64(100),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty0",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(100),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)
				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-1;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty1",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-1"),
					Amount:                        coin.NewInt(200),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 2)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgTransfer event as outgoing pending activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
							SourcePort:    "transfer",
							SourceChannel: "channel-0",
							Token: ibc_model.MsgTransferToken{
								Denom:  "basecro",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Sender:   "from",
							Receiver: "to",
							TimeoutHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							TimeoutTimestamp: "",
						},
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
						PacketData: ibc_model.FungibleTokenPacketData{
							Sender:   "from",
							Receiver: "to",
							Denom:    "basecro",
							Amount:   json.NewNumericStringFromUint64(100),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(100),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle failed MsgTransfer event as outgoing failed activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgIBCTransferTransfer{
					MsgBase: event_usecase.NewMsgBase(event_usecase.MsgBaseParams{
						MsgName: event_usecase.MSG_IBC_TRANSFER_TRANSFER,
						Version: 1,
						MsgCommonParams: event_usecase.MsgCommonParams{
							BlockHeight: 1,
							TxHash:      "TxHash",
							TxSuccess:   false,
							MsgIndex:    0,
						},
					}),
					Params: ibc_model.MsgTransferParams{
						RawMsgTransfer: ibc_model.RawMsgTransfer{
							SourcePort:    "transfer",
							SourceChannel: "channel-0",
							Token: ibc_model.MsgTransferToken{
								Denom:  "basecro",
								Amount: json.NewNumericStringFromUint64(100),
							},
							Sender:   "from",
							Receiver: "to",
							TimeoutHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							TimeoutTimestamp: "",
						},
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
						PacketData: ibc_model.FungibleTokenPacketData{
							Sender:   "from",
							Receiver: "to",
							Denom:    "basecro",
							Amount:   json.NewNumericStringFromUint64(100),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;transactionId:TxHash;status:failedOnChain",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(100),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_FAILED_ON_CHAIN,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should not handle MsgRecvPacket event from non-listened destination channel",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-0",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-2",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							ProofCommitment: "proof",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success: true,
							MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
								Hash:  "hash",
								Denom: "basecro",
							},
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("result"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle successful MsgRecvPacket event as incoming confirmed activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-1",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-0",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							ProofCommitment: "proof",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success: true,
							MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
								Hash:  "hash",
								Denom: "basecro",
							},
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("result"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:counterparty;channel:channel-1;sequence:0",
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   "counterparty",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "this",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgRecvPacket event in all listened channels",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty0
						// this <-- channel-2 --- counterparty0
						ChainName:      "counterparty0",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
					{
						// this --- channel-1 --> counterparty1
						// this <-- channel-3 --- counterparty1
						ChainName:      "counterparty1",
						ChannelId:      "channel-1",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-2",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-0",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							ProofCommitment: "proof",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success: true,
							MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
								Hash:  "hash",
								Denom: "basecro",
							},
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("result"),
							MaybeError:  nil,
						},
					},
				},
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
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-3",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-1",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							ProofCommitment: "proof",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(2000),
							},
							Success: true,
							MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
								Hash:  "hash",
								Denom: "basecro",
							},
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: []byte("result"),
							MaybeError:  nil,
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:counterparty0;channel:channel-2;sequence:0",
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   "counterparty0",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "this",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)
				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:counterparty1;channel:channel-3;sequence:0",
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   "counterparty1",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "this",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-1"),
					Amount:                        coin.NewInt(2000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 2)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgRecvPacket event with error as incoming rejected activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
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
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-1",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-0",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							ProofCommitment: "proof",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success: false,
							MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
								Hash:  "hash",
								Denom: "basecro",
							},
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
						PacketAck: ibc_model.MsgRecvPacketPacketAck{
							MaybeResult: nil,
							MaybeError:  primptr.String("timeout"),
						},
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:counterparty;channel:channel-1;sequence:0",
					Direction:                     types.DIRECTION_INCOMING,
					FromChainId:                   "counterparty",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "this",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_REJECTED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should not handle MsgAcknowledgement event from non-listened source channel",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgIBCAcknowledgement{
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
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-2",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-2",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							Acknowledgement: "acknowledgement",
							ProofAcked:      "proofacked",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success:         false,
							Acknowledgement: "acknowledgement",
							MaybeError:      primptr.String("error"),
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle unsuccessful MsgAcknowledgement event as rejected response activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgIBCAcknowledgement{
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
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-0",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-1",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							Acknowledgement: "acknowledgement",
							ProofAcked:      "proofacked",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success:         false,
							Acknowledgement: "acknowledgement",
							MaybeError:      primptr.String("error"),
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_REJECTED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgAcknowledgement event as confirmed response activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgIBCAcknowledgement{
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
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-0",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-1",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							Acknowledgement: "acknowledgement",
							ProofAcked:      "proofacked",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success:         true,
							Acknowledgement: "acknowledgement",
							MaybeError:      nil,
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle MsgAcknowledgement event from all listened channels",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-2 --- counterparty
						ChainName:      "counterparty0",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
					{
						// this --- channel-1 --> counterparty1
						// this <-- channel-3 --- counterparty1
						ChainName:      "counterparty1",
						ChannelId:      "channel-1",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.MsgIBCAcknowledgement{
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
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-0",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-2",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							Acknowledgement: "acknowledgement",
							ProofAcked:      "proofacked",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(1000),
							},
							Success:         true,
							Acknowledgement: "acknowledgement",
							MaybeError:      nil,
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
					},
				},
				&event_usecase.MsgIBCAcknowledgement{
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
					Params: ibc_model.MsgAcknowledgementParams{
						RawMsgAcknowledgement: ibc_model.RawMsgAcknowledgement{
							Packet: ibc_model.Packet{
								Sequence:           "0",
								SourcePort:         "transfer",
								SourceChannel:      "channel-1",
								DestinationPort:    "transfer",
								DestinationChannel: "channel-3",
								Data:               "data=",
								TimeoutHeight: ibc_model.Height{
									RevisionNumber: 0,
									RevisionHeight: 0,
								},
								TimeoutTimestamp: "",
							},
							Acknowledgement: "acknowledgement",
							ProofAcked:      "proofacked",
							ProofHeight: ibc_model.Height{
								RevisionNumber: 0,
								RevisionHeight: 0,
							},
							Signer: "signer",
						},
						Application: "transfer",
						MessageType: "/ibc.applications.transfer.v1.MsgTransfer",
						MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
							FungibleTokenPacketData: ibc_model.FungibleTokenPacketData{
								Sender:   "from",
								Receiver: "to",
								Denom:    "basecro",
								Amount:   json.NewNumericStringFromUint64(2000),
							},
							Success:         true,
							Acknowledgement: "acknowledgement",
							MaybeError:      nil,
						},
						PacketSequence:  0,
						ChannelOrdering: "UNORDERED",
						ConnectionID:    "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty0",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(1000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)
				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("TxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-1;sequence:0",
					Direction:                     types.DIRECTION_RESPONSE,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty1",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-1"),
					Amount:                        coin.NewInt(2000),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_COUNTERPARTY_CONFIRMED,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 2)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should not handle CronosSendToIBCCreated event from non-listening source channel",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-2",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(100),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle CronosSendToIBCCreated event as pending outgoing activity",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty
						// this <-- channel-1 --- counterparty
						ChainName:      "counterparty",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-0",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(100),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-1",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("EthereumTxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(100),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 1)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should handle CronosSendToIBCCreated event from all listened channels",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty0
						// this <-- channel-2 --- counterparty0
						ChainName:      "counterparty0",
						ChannelId:      "channel-0",
						StartingHeight: 0,
					},
					{
						// this --- channel-1 --> counterparty0
						// this <-- channel-3 --- counterparty0
						ChainName:      "counterparty1",
						ChannelId:      "channel-1",
						StartingHeight: 0,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-0",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(100),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-1",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(200),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("EthereumTxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-0;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty0",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-0"),
					Amount:                        coin.NewInt(100),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)
				mockBridgePendingActivitiesView.On("Insert", &view.BridgePendingActivityInsertRow{
					BlockHeight:                   1,
					BlockTime:                     primptr.UTCTime(utctime.FromUnixNano(1)),
					MaybeTransactionId:            primptr.String("EthereumTxHash"),
					BridgeType:                    types.BRIDGE_TYPE_IBC,
					LinkId:                        "source:this;channel:channel-1;sequence:0",
					Direction:                     types.DIRECTION_OUTGOING,
					FromChainId:                   "this",
					MaybeFromAddress:              primptr.String("from"),
					MaybeFromSmartContractAddress: nil,
					ToChainId:                     "counterparty1",
					ToAddress:                     "to",
					MaybeToSmartContractAddress:   nil,
					MaybeChannelId:                primptr.String("channel-1"),
					Amount:                        coin.NewInt(200),
					MaybeDenom:                    primptr.String("basecro"),
					MaybeBridgeFeeAmount:          nil,
					MaybeBridgeFeeDenom:           nil,
					Status:                        types.STATUS_PENDING,
					IsProcessed:                   false,
				}).Return(nil)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 2)
				}

				return mocks, assertFunc
			},
		},
		{
			Name: "It should not handle CronosSendToIBCCreated event that has not reach listened channel starting height",
			Config: bridge_pending_activity.Config{
				ThisChainName: "this",
				CounterPartyChains: []bridge_pending_activity.CounterPartyChainConfig{
					{
						// this --- channel-0 --> counterparty0
						// this <-- channel-2 --- counterparty0
						ChainName:      "counterparty0",
						ChannelId:      "channel-0",
						StartingHeight: 100,
					},
					{
						// this --- channel-1 --> counterparty1
						// this <-- channel-3 --- counterparty1
						ChainName:      "counterparty1",
						ChannelId:      "channel-1",
						StartingHeight: 100,
					},
				},
			},
			Events: []entity_event.Event{
				&event_usecase.BlockCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.BLOCK_CREATED,
						Version:     1,
						BlockHeight: 1,
					}),
					Block: &usecase_model.Block{
						Height:          1,
						Hash:            "Hash",
						Time:            utctime.FromUnixNano(1),
						AppHash:         "AppHash",
						ProposerAddress: "ProposerAddress",
						Txs:             nil,
						Signatures:      nil,
						Evidences:       nil,
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-0",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(100),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-2",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
				&event_usecase.CronosSendToIBCCreated{
					Base: entity_event.NewBase(entity_event.BaseParams{
						Name:        event_usecase.CRONOS_SEND_TO_IBC_CREATED,
						Version:     0,
						BlockHeight: 1,
					}),
					Params: usecase_model.CronosSendToIBCParams{
						TxHash:         "TxHash",
						EthereumTxHash: "EthereumTxHash",
						SourcePort:     "transfer",
						SourceChannel:  "channel-1",
						Token: usecase_model.CronosSendToIBCToken{
							Denom:  "basecro",
							Amount: json.NewNumericStringFromUint64(200),
						},
						Sender:   "from",
						Receiver: "to",
						TimeoutHeight: usecase_model.CronosSendToIBCHeight{
							RevisionNumber: 0,
							RevisionHeight: 0,
						},
						TimeoutTimestamp:   "",
						PacketDataHex:      "packetdatahex",
						PacketSequence:     0,
						DestinationPort:    "transfer",
						DestinationChannel: "channel-3",
						ChannelOrdering:    "UNORDERED",
						ConnectionID:       "connection-0",
					},
				},
			},
			MockFunc: func() (mocks []*testify_mock.Mock, assertFunc func()) {
				mockBridgePendingActivitiesView := view.NewMockBridgePendingActivitiesView().(*view.MockBridgePendingActivitiesView)
				mocks = append(mocks, &mockBridgePendingActivitiesView.Mock)

				bridge_pending_activity.NewBridgePendingActivitiesView = func(_ *rdb.Handle) view.BridgePendingActivities {
					return mockBridgePendingActivitiesView
				}

				bridge_pending_activity.UpdateLastHandledEventHeight = func(_ *bridge_pending_activity.BridgePendingActivity, _ *rdb.Handle, _ int64) error {
					return nil
				}

				assertFunc = func() {
					mockBridgePendingActivitiesView.AssertNumberOfCalls(t, "Insert", 0)
				}

				return mocks, assertFunc
			},
		},
	}

	for _, tc := range testCases {
		fmt.Print(tc.Name)
		mockRDbConn := NewMockRDbConn()
		mockTx := NewMockRDbTx()
		mockRDbConn.On("Begin").Return(mockTx, nil)
		mocks, assertFunc := tc.MockFunc()

		projection := NewBridgePendingActivityProjection(mockRDbConn, tc.Config)
		onInitErr := projection.OnInit()
		assert.NoError(t, onInitErr)

		handleErr := projection.HandleEvents(1, tc.Events)
		assert.NoError(t, handleErr)

		for _, m := range mocks {
			m.AssertExpectations(t)
		}
		assertFunc()

		fmt.Println(" Passed")
	}
}

func NewBridgePendingActivityProjection(
	rdbConn rdb.Conn,
	config bridge_pending_activity.Config,
) *bridge_pending_activity.BridgePendingActivity {
	fakeLogger := test_logger.NewFakeLogger()
	return bridge_pending_activity.New(
		config,
		fakeLogger,
		rdbConn,
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
	mockTx.On("ToHandle").Return(&rdb.Handle{
		Runner:      mockTx,
		TypeConv:    &pg.PgxTypeConv{},
		StmtBuilder: pg.PostgresStmtBuilder,
	}).Maybe()
	mockTx.On("Rollback").Return(nil).Maybe()
	mockTx.On("Commit").Return(nil).Maybe()

	return mockTx
}
