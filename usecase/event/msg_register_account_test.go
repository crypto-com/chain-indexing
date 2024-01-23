package event_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/json"
	icaauth_model "github.com/crypto-com/chain-indexing/usecase/model/icaauth"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/must"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgRegisterAccount", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 0
			anyChannelId := "channel-0"
			anyPortID := "anyPortID"
			anyCounterpartyPortID := "anyCounterpartyPortID"
			anyCounterpartyChannelID := "anyCounterpartyChannelID"

			var anyRawValue map[string]interface{}
			var anyRawMsgRegisterAccount icaauth_model.MsgRegisterAccount
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgRegisterAccount,
			})
			json.MustUnmarshalFromString(`
{
	"@type": "/icaauth.v1.MsgRegisterAccount",
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"connection_id": "connection-18",
	"version": ""
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := icaauth_model.MsgRegisterAccountParams{
				MsgRegisterAccount: anyRawMsgRegisterAccount,

				PortID:                anyPortID,
				ChannelID:             anyChannelId,
				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
			}

			event := event_usecase.NewMsgRegisterAccount(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_REGISTER_ACCOUNT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgRegisterAccount)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_REGISTER_ACCOUNT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyParams.CounterpartyPortID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 0
			anyChannelId := "channel-0"
			anyPortID := "anyPortID"
			anyCounterpartyPortID := "anyCounterpartyPortID"
			anyCounterpartyChannelID := "anyCounterpartyChannelID"

			var anyRawValue map[string]interface{}
			var anyRawMsgRegisterAccount icaauth_model.MsgRegisterAccount
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgRegisterAccount,
			})
			json.MustUnmarshalFromString(`
{
	"@type": "/icaauth.v1.MsgRegisterAccount",
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"connection_id": "connection-18",
	"version": ""
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := icaauth_model.MsgRegisterAccountParams{
				MsgRegisterAccount: anyRawMsgRegisterAccount,

				PortID:                anyPortID,
				ChannelID:             anyChannelId,
				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
			}

			event := event_usecase.NewMsgRegisterAccount(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_REGISTER_ACCOUNT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgRegisterAccount)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_REGISTER_ACCOUNT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyParams.CounterpartyPortID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
		})
	})
})
