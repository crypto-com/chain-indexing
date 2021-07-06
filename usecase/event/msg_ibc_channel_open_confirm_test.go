package event_test

import (
	"time"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCTransferTransfer", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyChannelId := "channel-0"
			anyCounterpartyPortId := "transfer"
			anyConnectionId := "connection-0"
			var anyRawMsgChannelOpenConfirm ibc_model.RawMsgChannelOpenConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenConfirm,
			})
			decoder.Decode(`
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
`)

			anyParams := ibc_model.MsgChannelOpenConfirmParams{
				RawMsgChannelOpenConfirm: anyRawMsgChannelOpenConfirm,

				CounterpartyChannelID: anyCounterpartyChannelId,
				CounterpartyPortID:    anyCounterpartyPortId,
				ConnectionID:          anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Type).To(Equal(anyParams.Type))
			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyParams.CounterpartyPortID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyChannelId := "channel-0"
			anyCounterpartyPortId := "transfer"
			anyConnectionId := "connection-0"
			var anyRawMsgChannelOpenConfirm ibc_model.RawMsgChannelOpenConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenConfirm,
			})
			decoder.Decode(`
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
`)

			anyParams := ibc_model.MsgChannelOpenConfirmParams{
				RawMsgChannelOpenConfirm: anyRawMsgChannelOpenConfirm,

				CounterpartyChannelID: anyCounterpartyChannelId,
				CounterpartyPortID:    anyCounterpartyPortId,
				ConnectionID:          anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_CONFIRM_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Type).To(Equal(anyParams.Type))
			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyParams.CounterpartyPortID))
		})
	})
})
