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
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCChannelCloseConfirm", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyPortID := "transfer"
			anyCounterpartyChannelID := "channel-1"
			anyConnectionID := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelCloseConfirm ibc_model.RawMsgChannelCloseConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelCloseConfirm,
			})
			json.MustUnmarshalFromString(`
      {
        "@type": "/ibc.core.channel.v1.MsgChannelCloseConfirm",
        "port_id": "transfer",
        "channel_id": "channel-0",
        "proof_init": "Cv8CCvwCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTESMggEEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0xKgdpY3MyMC0xGgsIARgBIAEqAwACLiIpCAESJQIELiDoa/QgHw9tDnDiZlToyGCCtvdk/YcdNOtMNk4s7TESBiAiKwgBEgQEBi4gGiEgWkRCIWiKNqPdGeeaBlk1tUshQ314yl/2d0Dht7MZKTIiKwgBEgQGCi4gGiEgQZepEjXiUZCYCObv6YevxMpCE3NytECo0qCvAqMefj8iKwgBEgQIEi4gGiEgiNI7Cxb9YFfTvRC3QICXYJBUMG2Hd0PzF8UuHKL0G+AiKwgBEgQKLC4gGiEggTFR60ori9MUc86XOgPlQ8yINywZKHUcrwL1nHOCrFUiKwgBEgQMUi4gGiEg63yd2j19EPXHXJKMLbS35ceHlpLDlCK06cRIc9vxPoUK/gEK+wEKA2liYxIg9LWYWuKyPj5DLqVTHwqfaHiGnaYPpz14ploHwPgk8iQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyInCAESAQEaIPhZvf/MWZkK7MygO8J7NUbiwp6dqQN5Gf3gt6vxV489IicIARIBARogCD1iqt4WlG2/Sexg8bzDgcffB858+xTniEdtcFPnq/EiJQgBEiEBxwWsbBJAhv48/qrs9osH7NubOZjpZ2vlN0I+bSv+YMkiJwgBEgEBGiCwFtEy00C8HUDC67XhEg1JqHWPRCvNgukGVPcaXwl61Q==",
        "proof_height": { "revision_number": "2", "revision_height": "24" },
        "signer": "cro12cgecr4kmyylql6kerfpn7ff42weur7glq4uj3"
      }`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelCloseConfirmParams{
				RawMsgChannelCloseConfirm: anyRawMsgChannelCloseConfirm,

				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
				ConnectionID:          anyConnectionID,
			}

			event := event_usecase.NewMsgIBCChannelCloseConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelCloseConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.ProofInit).To(Equal(anyParams.ProofInit))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
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
			anyCounterpartyChannelID := "channel-1"
			anyConnectionID := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelCloseConfirm ibc_model.RawMsgChannelCloseConfirm
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelCloseConfirm,
			})
			json.MustUnmarshalFromString(`
			{
        "@type": "/ibc.core.channel.v1.MsgChannelCloseConfirm",
        "port_id": "transfer",
        "channel_id": "channel-0",
        "proof_init": "Cv8CCvwCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTESMggEEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0xKgdpY3MyMC0xGgsIARgBIAEqAwACLiIpCAESJQIELiDoa/QgHw9tDnDiZlToyGCCtvdk/YcdNOtMNk4s7TESBiAiKwgBEgQEBi4gGiEgWkRCIWiKNqPdGeeaBlk1tUshQ314yl/2d0Dht7MZKTIiKwgBEgQGCi4gGiEgQZepEjXiUZCYCObv6YevxMpCE3NytECo0qCvAqMefj8iKwgBEgQIEi4gGiEgiNI7Cxb9YFfTvRC3QICXYJBUMG2Hd0PzF8UuHKL0G+AiKwgBEgQKLC4gGiEggTFR60ori9MUc86XOgPlQ8yINywZKHUcrwL1nHOCrFUiKwgBEgQMUi4gGiEg63yd2j19EPXHXJKMLbS35ceHlpLDlCK06cRIc9vxPoUK/gEK+wEKA2liYxIg9LWYWuKyPj5DLqVTHwqfaHiGnaYPpz14ploHwPgk8iQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyInCAESAQEaIPhZvf/MWZkK7MygO8J7NUbiwp6dqQN5Gf3gt6vxV489IicIARIBARogCD1iqt4WlG2/Sexg8bzDgcffB858+xTniEdtcFPnq/EiJQgBEiEBxwWsbBJAhv48/qrs9osH7NubOZjpZ2vlN0I+bSv+YMkiJwgBEgEBGiCwFtEy00C8HUDC67XhEg1JqHWPRCvNgukGVPcaXwl61Q==",
        "proof_height": { "revision_number": "2", "revision_height": "24" },
        "signer": "cro12cgecr4kmyylql6kerfpn7ff42weur7glq4uj3"
      }`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelCloseConfirmParams{
				RawMsgChannelCloseConfirm: anyRawMsgChannelCloseConfirm,

				CounterpartyPortID:    anyCounterpartyPortID,
				CounterpartyChannelID: anyCounterpartyChannelID,
				ConnectionID:          anyConnectionID,
			}

			event := event_usecase.NewMsgIBCChannelCloseConfirm(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelCloseConfirm)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_CLOSE_CONFIRM_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.ProofInit).To(Equal(anyParams.ProofInit))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.CounterpartyPortID).To(Equal(anyCounterpartyPortID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyCounterpartyChannelID))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyConnectionID))
		})
	})
})
