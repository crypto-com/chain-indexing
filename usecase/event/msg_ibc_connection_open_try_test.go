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

	Describe("En/DecodeMsgIBCConnectionOpenTry", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgConnectionOpenTry ibc_model.RawMsgConnectionOpenTryTendermintClient
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgConnectionOpenTry,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenTry",
  "client_id": "07-tendermint-0",
  "previous_connection_id": "",
  "client_state": {
    "@type": "/ibc.lightclients.tendermint.v1.ClientState",
    "chain_id": "devnet-2",
    "trust_level": {
      "numerator": "1",
      "denominator": "3"
    },
    "trusting_period": "1209600s",
    "unbonding_period": "1814400s",
    "max_clock_drift": "5s",
    "frozen_height": {
      "revision_number": "0",
      "revision_height": "0"
    },
    "latest_height": {
      "revision_number": "2",
      "revision_height": "4"
    },
    "proof_specs": [
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 33,
          "min_prefix_length": 4,
          "max_prefix_length": 12,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0
      },
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 32,
          "min_prefix_length": 1,
          "max_prefix_length": 1,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0
      }
    ],
    "upgrade_path": [
      "upgrade",
      "upgradedIBCState"
    ],
    "allow_update_after_expiry": false,
    "allow_update_after_misbehaviour": false
  },
  "counterparty": {
    "client_id": "07-tendermint-0",
    "connection_id": "connection-0",
    "prefix": {
      "key_prefix": "aWJj"
    }
  },
  "delay_period": "0",
  "counterparty_versions": [
    {
      "identifier": "1",
      "features": [
        "ORDER_ORDERED",
        "ORDER_UNORDERED"
      ]
    }
  ],
  "proof_height": {
    "revision_number": "1",
    "revision_height": "8"
  },
  "proof_init": "Cq4CCqsCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASUgoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgBIhgKDzA3LXRlbmRlcm1pbnQtMBoFCgNpYmMaCwgBGAEgASoDAAIMIisIARIEAgQOIBohIKC9H3IwOG4Tpg/Nymt8yMllJEjWt9aeYeuMwflKc4zFIisIARIEBAgOIBohIHV1PdysEO47G2t5+mlniRDH1Pad0QTeX4qIqIhnTO0wIikIARIlBg4OIIIVttNVy/4Nk+sDJg6RVumpww3erCRz3wr8gpVu3e76ICIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "proof_client": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMhIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAIQBEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAg4iKwgBEgQCBA4gGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQEBg4gGiEg3quYtD6jLHzeIroI7j0I7LcgdqygTR1W4dnk3us6cJQiKwgBEgQIFA4gGiEgpk5CYIA5vc4apf6bk8arEpuObvNs1tN4ml/j+Ji0ZykK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "proof_consensus": "CvYCCvMCCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMi00EoYBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElQKDAjGmMaEBhCYnc+kARIiCiB5FDU4YJfh5rt8nXe3L5vEpRz0QglvmY7w9AUf/jWVtBog494NKzI3oC6cIMNPnuBPafWGH7wuJyKgEcqQN/xnp+waCwgBGAEgASoDAAIOIisIARIEAgQOIBohIK7bCMEuGbJOWa02nJkzHLbqFqogpoHwkmOl4CWI2/KHIikIARIlBAYOIBAW83IGMWt6nb7HNyJbW0Y3gaTLyiY73bO6+fGVfsn/ICIrCAESBAYODiAaISDAshei6GmugrMmo3DjCOIosuZkuRkcEPudaBJlzPzcSCIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "consensus_height": {
    "revision_number": "2",
    "revision_height": "4"
  },
  "signer": "cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgConnectionOpenTryParams{
				MsgConnectionOpenTryBaseParams: anyRawMsgConnectionOpenTry.MsgConnectionOpenTryBaseParams,
				MaybeTendermintClientState:     &anyRawMsgConnectionOpenTry.TendermintClientState,
				ConnectionID:                   anyConnectionId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenTry(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_TRY_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenTry)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_TRY_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.Counterparty).To(Equal(anyParams.Counterparty))
			Expect(typedEvent.Params.MaybeTendermintClientState).To(Equal(anyParams.MaybeTendermintClientState))
			Expect(typedEvent.Params.DelayPeriod).To(Equal(anyParams.DelayPeriod))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgConnectionOpenTry ibc_model.RawMsgConnectionOpenTryTendermintClient
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgConnectionOpenTry,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenTry",
  "client_id": "07-tendermint-0",
  "previous_connection_id": "",
  "client_state": {
    "@type": "/ibc.lightclients.tendermint.v1.ClientState",
    "chain_id": "devnet-2",
    "trust_level": {
      "numerator": "1",
      "denominator": "3"
    },
    "trusting_period": "1209600s",
    "unbonding_period": "1814400s",
    "max_clock_drift": "5s",
    "frozen_height": {
      "revision_number": "0",
      "revision_height": "0"
    },
    "latest_height": {
      "revision_number": "2",
      "revision_height": "4"
    },
    "proof_specs": [
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 33,
          "min_prefix_length": 4,
          "max_prefix_length": 12,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0
      },
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 32,
          "min_prefix_length": 1,
          "max_prefix_length": 1,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0
      }
    ],
    "upgrade_path": [
      "upgrade",
      "upgradedIBCState"
    ],
    "allow_update_after_expiry": false,
    "allow_update_after_misbehaviour": false
  },
  "counterparty": {
    "client_id": "07-tendermint-0",
    "connection_id": "connection-0",
    "prefix": {
      "key_prefix": "aWJj"
    }
  },
  "delay_period": "0",
  "counterparty_versions": [
    {
      "identifier": "1",
      "features": [
        "ORDER_ORDERED",
        "ORDER_UNORDERED"
      ]
    }
  ],
  "proof_height": {
    "revision_number": "1",
    "revision_height": "8"
  },
  "proof_init": "Cq4CCqsCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASUgoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgBIhgKDzA3LXRlbmRlcm1pbnQtMBoFCgNpYmMaCwgBGAEgASoDAAIMIisIARIEAgQOIBohIKC9H3IwOG4Tpg/Nymt8yMllJEjWt9aeYeuMwflKc4zFIisIARIEBAgOIBohIHV1PdysEO47G2t5+mlniRDH1Pad0QTeX4qIqIhnTO0wIikIARIlBg4OIIIVttNVy/4Nk+sDJg6RVumpww3erCRz3wr8gpVu3e76ICIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "proof_client": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMhIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAIQBEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAg4iKwgBEgQCBA4gGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQEBg4gGiEg3quYtD6jLHzeIroI7j0I7LcgdqygTR1W4dnk3us6cJQiKwgBEgQIFA4gGiEgpk5CYIA5vc4apf6bk8arEpuObvNs1tN4ml/j+Ji0ZykK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "proof_consensus": "CvYCCvMCCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMi00EoYBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElQKDAjGmMaEBhCYnc+kARIiCiB5FDU4YJfh5rt8nXe3L5vEpRz0QglvmY7w9AUf/jWVtBog494NKzI3oC6cIMNPnuBPafWGH7wuJyKgEcqQN/xnp+waCwgBGAEgASoDAAIOIisIARIEAgQOIBohIK7bCMEuGbJOWa02nJkzHLbqFqogpoHwkmOl4CWI2/KHIikIARIlBAYOIBAW83IGMWt6nb7HNyJbW0Y3gaTLyiY73bO6+fGVfsn/ICIrCAESBAYODiAaISDAshei6GmugrMmo3DjCOIosuZkuRkcEPudaBJlzPzcSCIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
  "consensus_height": {
    "revision_number": "2",
    "revision_height": "4"
  },
  "signer": "cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgConnectionOpenTryParams{
				MsgConnectionOpenTryBaseParams: anyRawMsgConnectionOpenTry.MsgConnectionOpenTryBaseParams,
				MaybeTendermintClientState:     &anyRawMsgConnectionOpenTry.TendermintClientState,
				ConnectionID:                   anyConnectionId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenTry(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_TRY_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenTry)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_TRY_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.Counterparty).To(Equal(anyParams.Counterparty))
			Expect(typedEvent.Params.MaybeTendermintClientState).To(Equal(anyParams.MaybeTendermintClientState))
			Expect(typedEvent.Params.DelayPeriod).To(Equal(anyParams.DelayPeriod))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})
	})
})
