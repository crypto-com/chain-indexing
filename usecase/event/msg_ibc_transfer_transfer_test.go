package event_test

import (
	"github.com/crypto-com/chain-indexing/external/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCTransferTransfer", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyDestinationPort := "transfer"
			anyDestinationChannel := "channel-0"
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgTransfer ibc_model.RawMsgTransfer
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.applications.transfer.v1.MsgTransfer",
  "source_port": "transfer",
  "source_channel": "channel-0",
  "token": {
    "denom": "basecro",
    "amount": "1234"
  },
  "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
  "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "timeout_height": {
    "revision_number": "2",
    "revision_height": "1023"
  },
  "timeout_timestamp": "0"
}
`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(anyRawValue, &anyRawMsgTransfer)

			anyParams := ibc_model.MsgTransferParams{
				RawMsgTransfer: anyRawMsgTransfer,

				PacketSequence:     anyPacketSequence,
				DestinationPort:    anyDestinationPort,
				DestinationChannel: anyDestinationChannel,
				ChannelOrdering:    anyChannelOrdering,
				ConnectionID:       anyConnectionId,
			}

			event := event_usecase.NewMsgIBCTransferTransfer(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCTransferTransfer)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_TRANSFER_TRANSFER_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.SourcePort).To(Equal(anyParams.SourcePort))
			Expect(typedEvent.Params.SourceChannel).To(Equal(anyParams.SourceChannel))
			Expect(typedEvent.Params.Token).To(Equal(anyParams.Token))
			Expect(typedEvent.Params.Sender).To(Equal(anyParams.Sender))
			Expect(typedEvent.Params.Receiver).To(Equal(anyParams.Receiver))
			Expect(typedEvent.Params.TimeoutHeight).To(Equal(anyParams.TimeoutHeight))
			Expect(typedEvent.Params.TimeoutTimestamp).To(Equal(anyParams.TimeoutTimestamp))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.DestinationPort).To(Equal(anyParams.DestinationPort))
			Expect(typedEvent.Params.DestinationChannel).To(Equal(anyParams.DestinationChannel))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyDestinationPort := "transfer"
			anyDestinationChannel := "channel-0"
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgTransfer ibc_model.RawMsgTransfer
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.applications.transfer.v1.MsgTransfer",
  "source_port": "transfer",
  "source_channel": "channel-0",
  "token": {
    "denom": "basecro",
    "amount": "1234"
  },
  "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
  "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "timeout_height": {
    "revision_number": "2",
    "revision_height": "1023"
  },
  "timeout_timestamp": "0"
}
`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(anyRawValue, &anyRawMsgTransfer)

			anyParams := ibc_model.MsgTransferParams{
				RawMsgTransfer: anyRawMsgTransfer,

				PacketSequence:     anyPacketSequence,
				DestinationPort:    anyDestinationPort,
				DestinationChannel: anyDestinationChannel,
				ChannelOrdering:    anyChannelOrdering,
				ConnectionID:       anyConnectionId,
			}

			event := event_usecase.NewMsgIBCTransferTransfer(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_TRANSFER_TRANSFER_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCTransferTransfer)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_TRANSFER_TRANSFER_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.SourcePort).To(Equal(anyParams.SourcePort))
			Expect(typedEvent.Params.SourceChannel).To(Equal(anyParams.SourceChannel))
			Expect(typedEvent.Params.Token).To(Equal(anyParams.Token))
			Expect(typedEvent.Params.Sender).To(Equal(anyParams.Sender))
			Expect(typedEvent.Params.Receiver).To(Equal(anyParams.Receiver))
			Expect(typedEvent.Params.TimeoutHeight).To(Equal(anyParams.TimeoutHeight))
			Expect(typedEvent.Params.TimeoutTimestamp).To(Equal(anyParams.TimeoutTimestamp))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.DestinationPort).To(Equal(anyParams.DestinationPort))
			Expect(typedEvent.Params.DestinationChannel).To(Equal(anyParams.DestinationChannel))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})
	})
})
