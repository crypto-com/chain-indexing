package event_test

import (
	"time"

	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/internal/must"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCChannelOpenTry", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyChannelId := "channel-0"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelOpenTry ibc_model.RawMsgChannelOpenTry
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenTry,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenTry",
  "port_id": "transfer",
  "previous_channel_id": "",
  "channel": {
    "state": "STATE_TRYOPEN",
    "ordering": "ORDER_UNORDERED",
    "counterparty": {
      "port_id": "transfer",
      "channel_id": "channel-0"
    },
    "connection_hops": [
      "connection-0"
    ],
    "version": "ics20-1"
  },
  "counterparty_version": "ics20-1",
  "proof_init": "CpwCCpkCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASJwgBEAEaCgoIdHJhbnNmZXIiDGNvbm5lY3Rpb24tMCoHaWNzMjAtMRoLCAEYASABKgMAAhwiKwgBEgQCBBwgGiEgxASpVteatK9vbv42ysU8IwbBDKqExJmoe4EnC2fQrXUiKwgBEgQEBhwgGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQGDBwgGiEgSQK1Wy6XrzzpRUSz2wsxVsDh5vfLrpytpIczSYPOgAIiKwgBEgQKIBwgGiEgF0RIWS9LLLCG8HCr6ZFR0ye72LgQHgCOyZY/ViavqbQK0wEK0AEKA2liYxIgmynSRECR0yOhKgOMfUyJJPhlPUA8hz+XLmkskcr/hDQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEnssIok+Atq0LVaSnmcE7wnNM2RsSHOJ6LH++HW5h6dSIlCAESIQHWh6c3n369c+m6hB5/PHNc5x5CXz7DhZ0+Y+GEB4CpBCInCAESAQEaIGnwKIkp/LcnctdNSlJgxCiIV9a0KxIlpZYtxybthsge",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "15"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelOpenTryParams{
				RawMsgChannelOpenTry: anyRawMsgChannelOpenTry,

				ChannelID:    anyChannelId,
				ConnectionID: anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenTry(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenTry)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.PreviousChannelID).To(Equal(anyParams.PreviousChannelID))
			Expect(typedEvent.Params.Channel).To(Equal(anyParams.Channel))
			Expect(typedEvent.Params.CounterpartyVersion).To(Equal(anyParams.CounterpartyVersion))
			Expect(typedEvent.Params.ProofInit).To(Equal(anyParams.ProofInit))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyChannelId := "channel-0"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelOpenTry ibc_model.RawMsgChannelOpenTry
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenTry,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenTry",
  "port_id": "transfer",
  "previous_channel_id": "",
  "channel": {
    "state": "STATE_TRYOPEN",
    "ordering": "ORDER_UNORDERED",
    "counterparty": {
      "port_id": "transfer",
      "channel_id": "channel-0"
    },
    "connection_hops": [
      "connection-0"
    ],
    "version": "ics20-1"
  },
  "counterparty_version": "ics20-1",
  "proof_init": "CpwCCpkCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASJwgBEAEaCgoIdHJhbnNmZXIiDGNvbm5lY3Rpb24tMCoHaWNzMjAtMRoLCAEYASABKgMAAhwiKwgBEgQCBBwgGiEgxASpVteatK9vbv42ysU8IwbBDKqExJmoe4EnC2fQrXUiKwgBEgQEBhwgGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQGDBwgGiEgSQK1Wy6XrzzpRUSz2wsxVsDh5vfLrpytpIczSYPOgAIiKwgBEgQKIBwgGiEgF0RIWS9LLLCG8HCr6ZFR0ye72LgQHgCOyZY/ViavqbQK0wEK0AEKA2liYxIgmynSRECR0yOhKgOMfUyJJPhlPUA8hz+XLmkskcr/hDQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEnssIok+Atq0LVaSnmcE7wnNM2RsSHOJ6LH++HW5h6dSIlCAESIQHWh6c3n369c+m6hB5/PHNc5x5CXz7DhZ0+Y+GEB4CpBCInCAESAQEaIGnwKIkp/LcnctdNSlJgxCiIV9a0KxIlpZYtxybthsge",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "15"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelOpenTryParams{
				RawMsgChannelOpenTry: anyRawMsgChannelOpenTry,

				ChannelID:    anyChannelId,
				ConnectionID: anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenTry(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenTry)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_TRY_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.PreviousChannelID).To(Equal(anyParams.PreviousChannelID))
			Expect(typedEvent.Params.Channel).To(Equal(anyParams.Channel))
			Expect(typedEvent.Params.CounterpartyVersion).To(Equal(anyParams.CounterpartyVersion))
			Expect(typedEvent.Params.ProofInit).To(Equal(anyParams.ProofInit))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyChannelId))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyConnectionId))
		})
	})
})
