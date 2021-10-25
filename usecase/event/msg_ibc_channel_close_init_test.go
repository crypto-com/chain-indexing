package event_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/json"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCChannelCloseInit", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyPortID := "transfer"
			anyCounterpartyChannelID := "channel-0"
			anyConnectionID := "connection-1"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelCloseInit ibc_model.RawMsgChannelCloseInit
			json.MustUnmarshalFromString(`
			{
        "@type": "/ibc.core.channel.v1.MsgChannelCloseInit",
        "port_id": "transfer",
        "channel_id": "channel-1",
        "signer": "cro1t7yk3d4meeaqf5zfegv8p94wlfhpcnsftz55f7"
      }`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(
				anyRawValue,
				&anyRawMsgChannelCloseInit,
			)

			anyParams := ibc_model.MsgChannelCloseInitParams{
				RawMsgChannelCloseInit: anyRawMsgChannelCloseInit,

				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
				ConnectionID:          anyConnectionID,
			}

			event := event_usecase.NewMsgIBCChannelCloseInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelCloseInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyCounterpartyPortID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyCounterpartyChannelID))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyPortID := "transfer"
			anyCounterpartyChannelID := "channel-0"
			anyConnectionID := "connection-1"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelCloseInit ibc_model.RawMsgChannelCloseInit
			json.MustUnmarshalFromString(`
			{
        "@type": "/ibc.core.channel.v1.MsgChannelCloseInit",
        "port_id": "transfer",
        "channel_id": "channel-1",
        "signer": "cro1t7yk3d4meeaqf5zfegv8p94wlfhpcnsftz55f7"
      }`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(
				anyRawValue,
				&anyRawMsgChannelCloseInit,
			)

			anyParams := ibc_model.MsgChannelCloseInitParams{
				RawMsgChannelCloseInit: anyRawMsgChannelCloseInit,

				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
				ConnectionID:          anyConnectionID,
			}

			event := event_usecase.NewMsgIBCChannelCloseInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelCloseInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_CLOSE_INIT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyCounterpartyPortID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyCounterpartyChannelID))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyConnectionID))
		})
	})
})
