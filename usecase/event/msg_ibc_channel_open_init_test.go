package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/json"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCChannelOpenInit", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyChannelId := "channel-0"
			var anyRawMsgChannelOpenInit ibc_model.RawMsgChannelOpenInit
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenInit",
  "port_id": "transfer",
  "channel": {
    "state": "STATE_INIT",
    "ordering": "ORDER_UNORDERED",
    "counterparty": {
  	  "port_id": "transfer",
  	  "channel_id": ""
    },
    "connection_hops": [
  	  "connection-0"
    ],
    "version": "ics20-1"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawMsgChannelOpenInit)

			anyParams := ibc_model.MsgChannelOpenInitParams{
				RawMsgChannelOpenInit: anyRawMsgChannelOpenInit,

				ChannelID: anyChannelId,
			}

			event := event_usecase.NewMsgIBCChannelOpenInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Type).To(Equal(anyParams.Type))
			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.Channel).To(Equal(anyParams.Channel))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyChannelId := "channel-0"
			var anyRawMsgChannelOpenInit ibc_model.RawMsgChannelOpenInit
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenInit",
  "port_id": "transfer",
  "channel": {
    "state": "STATE_INIT",
    "ordering": "ORDER_UNORDERED",
    "counterparty": {
  	  "port_id": "transfer",
  	  "channel_id": ""
    },
    "connection_hops": [
  	  "connection-0"
    ],
    "version": "ics20-1"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawMsgChannelOpenInit)

			anyParams := ibc_model.MsgChannelOpenInitParams{
				RawMsgChannelOpenInit: anyRawMsgChannelOpenInit,

				ChannelID: anyChannelId,
			}

			event := event_usecase.NewMsgIBCChannelOpenInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_INIT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Type).To(Equal(anyParams.Type))
			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.Channel).To(Equal(anyParams.Channel))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
		})
	})
})
