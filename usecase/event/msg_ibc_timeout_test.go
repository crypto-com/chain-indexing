package event_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"time"

	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/internal/must"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCTimeout", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyApplication := "transfer"
			anyMessageType := "MsgTransfer"
			anyRefundReceiver := "cro1s7cu28403gzdvy5tttyskm3zxjejxcv63espre"
			anyRefundDenom := "basecro"
			anyRefundAmount := uint64(1)
			anyPacketTimeoutHeight := ibc_model.Height{
				RevisionNumber: 4,
				RevisionHeight: 6182017,
			}
			anyPacketTimeoutTimestamp := uint64(1620753450655319559)

			var anyRawValue map[string]interface{}
			var anyRawMsg ibc_model.RawMsgTimeout
			msgRecvPacketDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsg,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgTimeout",
  "packet": {
	"sequence": "5",
	"source_port": "transfer",
	"source_channel": "channel-9",
	"destination_port": "transfer",
	"destination_channel": "channel-109",
	"data": "eyJhbW91bnQiOiIxIiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjb3Ntb3Mxenp6NXhjd3l6bWs5dHBreTR0ZTkwM3ZzeGg1NXNmc2x2ZnplY2YiLCJzZW5kZXIiOiJjcm8xczdjdTI4NDAzZ3pkdnk1dHR0eXNrbTN6eGplanhjdjYzZXNwcmUifQ==",
	"timeout_height": {
	  "revision_number": "4",
	  "revision_height": "6182017"
	},
	"timeout_timestamp": "1620753450655319559"
  },
  "proof_unreceived": "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",
  "proof_height": {
	"revision_number": "4",
	"revision_height": "6185877"
  },
  "next_sequence_recv": "5",
  "signer": "cro1q040rm026jmpfmxdsj6q9phm9tdceepnsau6me"
}`, &anyRawValue)
			must.Do(msgRecvPacketDecoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgTimeoutParams{
				RawMsgTimeout: anyRawMsg,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
					RefundReceiver: anyRefundReceiver,
					RefundDenom:    anyRefundDenom,
					RefundAmount:   anyRefundAmount,
				},

				PacketTimeoutHeight:    anyPacketTimeoutHeight,
				PacketTimeoutTimestamp: anyPacketTimeoutTimestamp,
				PacketSequence:         anyPacketSequence,

				ChannelOrdering: anyChannelOrdering,
			}

			event := event_usecase.NewMsgIBCTimeout(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_TIMEOUT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCTimeout)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_TIMEOUT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.RawMsgTimeout).To(Equal(anyParams.RawMsgTimeout))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))
			Expect(typedEvent.Params.MaybeMsgTransfer).To(Equal(anyParams.MaybeMsgTransfer))

			Expect(typedEvent.Params.PacketTimeoutHeight).To(Equal(anyParams.PacketTimeoutHeight))
			Expect(typedEvent.Params.PacketTimeoutTimestamp).To(Equal(anyParams.PacketTimeoutTimestamp))
			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyApplication := "transfer"
			anyMessageType := "MsgTransfer"
			anyRefundReceiver := "cro1s7cu28403gzdvy5tttyskm3zxjejxcv63espre"
			anyRefundDenom := "basecro"
			anyRefundAmount := uint64(1)
			anyPacketTimeoutHeight := ibc_model.Height{
				RevisionNumber: 4,
				RevisionHeight: 6182017,
			}
			anyPacketTimeoutTimestamp := uint64(1620753450655319559)

			var anyRawValue map[string]interface{}
			var anyRawMsg ibc_model.RawMsgTimeout
			msgRecvPacketDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsg,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgTimeout",
  "packet": {
	"sequence": "5",
	"source_port": "transfer",
	"source_channel": "channel-9",
	"destination_port": "transfer",
	"destination_channel": "channel-109",
	"data": "eyJhbW91bnQiOiIxIiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjb3Ntb3Mxenp6NXhjd3l6bWs5dHBreTR0ZTkwM3ZzeGg1NXNmc2x2ZnplY2YiLCJzZW5kZXIiOiJjcm8xczdjdTI4NDAzZ3pkdnk1dHR0eXNrbTN6eGplanhjdjYzZXNwcmUifQ==",
	"timeout_height": {
	  "revision_number": "4",
	  "revision_height": "6182017"
	},
	"timeout_timestamp": "1620753450655319559"
  },
  "proof_unreceived": "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",
  "proof_height": {
	"revision_number": "4",
	"revision_height": "6185877"
  },
  "next_sequence_recv": "5",
  "signer": "cro1q040rm026jmpfmxdsj6q9phm9tdceepnsau6me"
}`, &anyRawValue)
			must.Do(msgRecvPacketDecoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgTimeoutParams{
				RawMsgTimeout: anyRawMsg,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeMsgTransfer: &ibc_model.MsgTimeoutMsgTransfer{
					RefundReceiver: anyRefundReceiver,
					RefundDenom:    anyRefundDenom,
					RefundAmount:   anyRefundAmount,
				},

				PacketTimeoutHeight:    anyPacketTimeoutHeight,
				PacketTimeoutTimestamp: anyPacketTimeoutTimestamp,
				PacketSequence:         anyPacketSequence,

				ChannelOrdering: anyChannelOrdering,
			}

			event := event_usecase.NewMsgIBCTimeout(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_TIMEOUT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCTimeout)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_TIMEOUT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.RawMsgTimeout).To(Equal(anyParams.RawMsgTimeout))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))
			Expect(typedEvent.Params.MaybeMsgTransfer).To(Equal(anyParams.MaybeMsgTransfer))

			Expect(typedEvent.Params.PacketTimeoutHeight).To(Equal(anyParams.PacketTimeoutHeight))
			Expect(typedEvent.Params.PacketTimeoutTimestamp).To(Equal(anyParams.PacketTimeoutTimestamp))
			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
		})
	})
})
