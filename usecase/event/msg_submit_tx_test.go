package event_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/json"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
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

	Describe("En/DecodeMsgSubmitTx", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 0
			anySequence := "1"
			anySourcePort := "anySourcePort"
			anySourceChannel := "anySourceChannel"
			anyDestinationPort := "anyDestinationPort"
			anyDestinationChannel := "anyDestinationChannel"
			anyData := "anyData"
			anyTimeoutHeight := ibc_model.Height{
				RevisionNumber: 0,
				RevisionHeight: 0,
			}
			anyTimeoutTimestamp := "anyTimeoutTimestamp"

			var anyRawValue map[string]interface{}
			var anyRawMsgSubmitTx icaauth_model.MsgSubmitTx
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgSubmitTx,
			})
			json.MustUnmarshalFromString(`
{
	"@type": "/chainmain.icaauth.v1.MsgSubmitTx",
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"connectionId": "connection-18",
	"msgs": [
		{
			"@type": "/cosmos.bank.v1beta1.MsgSend",
			"from_address": "tcro1m902wnj38q76dxf425tcv5gze9rgeeht67eldf",
			"to_address": "tcro144jtm7rx8y7wjwh8tufeke8n26dv4pdxp0zmm0",
			"amount": [
				{
					"denom": "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
					"amount": "100"
				}
			]
		}
	],
	"timeoutDuration": "300s"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := icaauth_model.MsgSubmitTxParams{
				MsgSubmitTx: anyRawMsgSubmitTx,

				Packet: ibc_model.Packet{
					Sequence:           anySequence,
					SourcePort:         anySourcePort,
					SourceChannel:      anySourceChannel,
					DestinationPort:    anyDestinationPort,
					DestinationChannel: anyDestinationChannel,
					Data:               anyData,
					TimeoutHeight:      anyTimeoutHeight,
					TimeoutTimestamp:   anyTimeoutTimestamp,
				},
			}

			event := event_usecase.NewMsgSubmitTx(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_TX_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitTx)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_TX_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Packet).To(Equal(anyParams.Packet))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 0
			anySequence := "1"
			anySourcePort := "anySourcePort"
			anySourceChannel := "anySourceChannel"
			anyDestinationPort := "anyDestinationPort"
			anyDestinationChannel := "anyDestinationChannel"
			anyData := "anyData"
			anyTimeoutHeight := ibc_model.Height{
				RevisionNumber: 0,
				RevisionHeight: 0,
			}
			anyTimeoutTimestamp := "anyTimeoutTimestamp"

			var anyRawValue map[string]interface{}
			var anyRawMsgSubmitTx icaauth_model.MsgSubmitTx
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgSubmitTx,
			})
			json.MustUnmarshalFromString(`
{
	"@type": "/chainmain.icaauth.v1.MsgSubmitTx",
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"connectionId": "connection-18",
	"msgs": [
		{
			"@type": "/cosmos.bank.v1beta1.MsgSend",
			"from_address": "tcro1m902wnj38q76dxf425tcv5gze9rgeeht67eldf",
			"to_address": "tcro144jtm7rx8y7wjwh8tufeke8n26dv4pdxp0zmm0",
			"amount": [
				{
					"denom": "ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865",
					"amount": "100"
				}
			]
		}
	],
	"timeoutDuration": "300s"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := icaauth_model.MsgSubmitTxParams{
				MsgSubmitTx: anyRawMsgSubmitTx,

				Packet: ibc_model.Packet{
					Sequence:           anySequence,
					SourcePort:         anySourcePort,
					SourceChannel:      anySourceChannel,
					DestinationPort:    anyDestinationPort,
					DestinationChannel: anyDestinationChannel,
					Data:               anyData,
					TimeoutHeight:      anyTimeoutHeight,
					TimeoutTimestamp:   anyTimeoutTimestamp,
				},
			}

			event := event_usecase.NewMsgSubmitTx(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_TX_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitTx)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_TX_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.Packet).To(Equal(anyParams.Packet))
		})
	})
})
