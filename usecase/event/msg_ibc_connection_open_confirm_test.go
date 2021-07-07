package event_test

import (
	"time"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/internal/must"
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

	Describe("En/DecodeMsgIBCConnectionOpenConfirm", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "tendermint-0"
			anyCounterpartyClientId := "tendermint-1"
			anyCounterpartyCconnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgConnectionOpenConfirm ibc_model.RawMsgConnectionOpenConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgConnectionOpenConfirm,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
  "connection_id": "connection-0",
  "proof_ack": "CrwCCrkCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgDIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhYiKwgBEgQCBBYgGiEgoL0fcjA4bhOmD83Ka3zIyWUkSNa31p5h64zB+UpzjMUiKwgBEgQECBYgGiEge8ROC6gJ/hvxo0nddRPSKTbuIWiteu7Fbhs59h7ImF8iKQgBEiUGDhYgV421a0/wlqTvyG2E9dyQAqOGy0gV3DwNOBpJ0YluyPYgIikIARIlCBgWILhLJzoETx8Ak2Xp4wuB75gso6yrXyn1M/DsizNa1JTbIArTAQrQAQoDaWJjEiAO3+Pnp8M0UIwSpChoUC35Bc39UWoKO58iK4AdECFfaxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAdCcOLwBd3qyMTRIUXZvtvDINpEb8SM8Pht3OxmrabIIIiUIARIhAU4hRda6BIRjBSBxdB+U4LUrwyCjdlk1HhOhdGwMgM4pIicIARIBARogVcvkOvJpiz0tN0H450uAl9LNEWX8wQmU9PavrXqJ2hU=",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "12"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgConnectionOpenConfirmParams{
				RawMsgConnectionOpenConfirm: anyRawMsgConnectionOpenConfirm,

				ClientID:                 anyClientId,
				CounterpartyClientID:     anyCounterpartyClientId,
				CounterpartyConnectionID: anyCounterpartyCconnectionId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.ProofACK).To(Equal(anyParams.ProofACK))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.CounterpartyClientID).To(Equal(anyParams.CounterpartyClientID))
			Expect(typedEvent.Params.CounterpartyConnectionID).To(Equal(anyParams.CounterpartyConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "tendermint-0"
			anyCounterpartyClientId := "tendermint-1"
			anyCounterpartyCconnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgConnectionOpenConfirm ibc_model.RawMsgConnectionOpenConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgConnectionOpenConfirm,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
  "connection_id": "connection-0",
  "proof_ack": "CrwCCrkCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgDIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhYiKwgBEgQCBBYgGiEgoL0fcjA4bhOmD83Ka3zIyWUkSNa31p5h64zB+UpzjMUiKwgBEgQECBYgGiEge8ROC6gJ/hvxo0nddRPSKTbuIWiteu7Fbhs59h7ImF8iKQgBEiUGDhYgV421a0/wlqTvyG2E9dyQAqOGy0gV3DwNOBpJ0YluyPYgIikIARIlCBgWILhLJzoETx8Ak2Xp4wuB75gso6yrXyn1M/DsizNa1JTbIArTAQrQAQoDaWJjEiAO3+Pnp8M0UIwSpChoUC35Bc39UWoKO58iK4AdECFfaxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAdCcOLwBd3qyMTRIUXZvtvDINpEb8SM8Pht3OxmrabIIIiUIARIhAU4hRda6BIRjBSBxdB+U4LUrwyCjdlk1HhOhdGwMgM4pIicIARIBARogVcvkOvJpiz0tN0H450uAl9LNEWX8wQmU9PavrXqJ2hU=",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "12"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgConnectionOpenConfirmParams{
				RawMsgConnectionOpenConfirm: anyRawMsgConnectionOpenConfirm,

				ClientID:                 anyClientId,
				CounterpartyClientID:     anyCounterpartyClientId,
				CounterpartyConnectionID: anyCounterpartyCconnectionId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_CONFIRM_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.ProofACK).To(Equal(anyParams.ProofACK))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.CounterpartyClientID).To(Equal(anyParams.CounterpartyClientID))
			Expect(typedEvent.Params.CounterpartyConnectionID).To(Equal(anyParams.CounterpartyConnectionID))
		})
	})
})
