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

	Describe("En/DecodeMsgIBCConnectionOpenAck", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "tendermint-0"
			anyCounterpartyClientId := "tendermint-1"
			var anyRawMsgConnectionOpenAck ibc_model.RawMsgConnectionOpenAckTendermintClient
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenAck",
  "connection_id": "connection-0",
  "counterparty_connection_id": "connection-0",
	"version": {
	  "identifier": "1",
	  "features": [
		"ORDER_ORDERED",
		"ORDER_UNORDERED"
	  ]
	},
	"client_state": {
	  "@type": "/ibc.lightclients.tendermint.v1.ClientState",
	  "chain_id": "devnet-1",
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
		"revision_number": "1",
		"revision_height": "8"
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
	"proof_height": {
	  "revision_number": "2",
	  "revision_height": "11"
	},
	"proof_try": "CroCCrcCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgCIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhIiKQgBEiUCBBQg82rv+J5yJTsqmzR59NDsb2QSM+WLnXIG6KcNLkM+OkggIikIARIlBAgUID4MsCa4ABOa69QPz6AbX5/lzcJdKUsavb5rnsXVeBlKICIrCAESBAYOFCAaISDTUETUzwU2dxU/8zMto7/LcYNlrxNPaK70aHoUWiD38SIpCAESJQgUFCAd54gQJQY2sQyko0wVJ2M3tABQQOuK2JsO4oj298eF0iAK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proof_client": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMRIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAEQCEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAhQiKwgBEgQCBBQgGiEgJfqlX0WUAPCdhTWwIU4QDSykCXpE1UJWobmRPN3UTBsiKwgBEgQEBhQgGiEgkmNVfGb0FEP31ym+IQvt8UFwE/+FwjzUm5ko3UmI+2YiKwgBEgQIFBQgGiEg4xjkOtDX3dH/OhsfvOIzyBDw1UI6TKSzjRcDP5QMI2UK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proof_consensus": "CvUCCvICCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMS04EoUBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElMKCwi/gYaHBhDAyrNBEiIKIGrO3+I45yc7lT3lJhHz6S3kQL4vPCw6A6lQb7jL31L0GiDDvmZM+0VcRAU839+33ipkg7bBliRUVb9FPTh4xYCTSxoLCAEYASABKgMAAhQiKQgBEiUCBBQgBcFP0jQiTKHo9UDsNzOd7DE0+3kPyfthUzsi6kPNnAsgIisIARIEBAgUIBohIJNbKob+2Y6eG2iS472XaQGWG2v9nJvqc3l+E+dJa1DYIisIARIEBg4UIBohINNQRNTPBTZ3FT/zMy2jv8txg2WvE09orvRoehRaIPfxIikIARIlCBQUIB3niBAlBjaxDKSjTBUnYze0AFBA64rYmw7iiPb3x4XSIArTAQrQAQoDaWJjEiDvV7dGzZcbYim3+ZkRa90eMRvgpp7tTir935kuKqreqhoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAZdSq6h9bIzskt8kyzTy95YNxfI2uN+4V11SnJ+OiAWKIiUIARIhAQNnazdpiqTV7QHhxqlOJeTuZZgmHGbj4t0CLAhIAFkUIicIARIBARog6HqfO5QslMotbstcz0w2QN+XPSo73Wp+C5g8fsE0ZFA=",
	"consensus_height": {
	  "revision_number": "1",
	  "revision_height": "8"
	},
	"signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawMsgConnectionOpenAck)

			anyParams := ibc_model.MsgConnectionOpenAckParams{
				MsgConnectionOpenAckBaseParams: anyRawMsgConnectionOpenAck.MsgConnectionOpenAckBaseParams,
				MaybeTendermintClientState:     &anyRawMsgConnectionOpenAck.TendermintClientState,

				ClientID:             anyClientId,
				CounterpartyClientID: anyCounterpartyClientId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenAck(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenAck)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.CounterpartyConnectionID).To(Equal(anyParams.CounterpartyConnectionID))
			Expect(typedEvent.Params.Version).To(Equal(anyParams.Version))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.ProofTry).To(Equal(anyParams.ProofTry))
			Expect(typedEvent.Params.ProofClient).To(Equal(anyParams.ProofClient))
			Expect(typedEvent.Params.ProofConsensus).To(Equal(anyParams.ProofConsensus))
			Expect(typedEvent.Params.ConsensusHeight).To(Equal(anyParams.ConsensusHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.MaybeTendermintClientState).To(Equal(anyParams.MaybeTendermintClientState))

			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.CounterpartyClientID).To(Equal(anyParams.CounterpartyClientID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "tendermint-0"
			anyCounterpartyClientId := "tendermint-1"
			var anyRawMsgConnectionOpenAck ibc_model.RawMsgConnectionOpenAckTendermintClient
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.connection.v1.MsgConnectionOpenAck",
  "connection_id": "connection-0",
  "counterparty_connection_id": "connection-0",
	"version": {
	  "identifier": "1",
	  "features": [
		"ORDER_ORDERED",
		"ORDER_UNORDERED"
	  ]
	},
	"client_state": {
	  "@type": "/ibc.lightclients.tendermint.v1.ClientState",
	  "chain_id": "devnet-1",
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
		"revision_number": "1",
		"revision_height": "8"
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
	"proof_height": {
	  "revision_number": "2",
	  "revision_height": "11"
	},
	"proof_try": "CroCCrcCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgCIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhIiKQgBEiUCBBQg82rv+J5yJTsqmzR59NDsb2QSM+WLnXIG6KcNLkM+OkggIikIARIlBAgUID4MsCa4ABOa69QPz6AbX5/lzcJdKUsavb5rnsXVeBlKICIrCAESBAYOFCAaISDTUETUzwU2dxU/8zMto7/LcYNlrxNPaK70aHoUWiD38SIpCAESJQgUFCAd54gQJQY2sQyko0wVJ2M3tABQQOuK2JsO4oj298eF0iAK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proof_client": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMRIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAEQCEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAhQiKwgBEgQCBBQgGiEgJfqlX0WUAPCdhTWwIU4QDSykCXpE1UJWobmRPN3UTBsiKwgBEgQEBhQgGiEgkmNVfGb0FEP31ym+IQvt8UFwE/+FwjzUm5ko3UmI+2YiKwgBEgQIFBQgGiEg4xjkOtDX3dH/OhsfvOIzyBDw1UI6TKSzjRcDP5QMI2UK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proof_consensus": "CvUCCvICCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMS04EoUBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElMKCwi/gYaHBhDAyrNBEiIKIGrO3+I45yc7lT3lJhHz6S3kQL4vPCw6A6lQb7jL31L0GiDDvmZM+0VcRAU839+33ipkg7bBliRUVb9FPTh4xYCTSxoLCAEYASABKgMAAhQiKQgBEiUCBBQgBcFP0jQiTKHo9UDsNzOd7DE0+3kPyfthUzsi6kPNnAsgIisIARIEBAgUIBohIJNbKob+2Y6eG2iS472XaQGWG2v9nJvqc3l+E+dJa1DYIisIARIEBg4UIBohINNQRNTPBTZ3FT/zMy2jv8txg2WvE09orvRoehRaIPfxIikIARIlCBQUIB3niBAlBjaxDKSjTBUnYze0AFBA64rYmw7iiPb3x4XSIArTAQrQAQoDaWJjEiDvV7dGzZcbYim3+ZkRa90eMRvgpp7tTir935kuKqreqhoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAZdSq6h9bIzskt8kyzTy95YNxfI2uN+4V11SnJ+OiAWKIiUIARIhAQNnazdpiqTV7QHhxqlOJeTuZZgmHGbj4t0CLAhIAFkUIicIARIBARog6HqfO5QslMotbstcz0w2QN+XPSo73Wp+C5g8fsE0ZFA=",
	"consensus_height": {
	  "revision_number": "1",
	  "revision_height": "8"
	},
	"signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawMsgConnectionOpenAck)

			anyParams := ibc_model.MsgConnectionOpenAckParams{
				MsgConnectionOpenAckBaseParams: anyRawMsgConnectionOpenAck.MsgConnectionOpenAckBaseParams,
				MaybeTendermintClientState:     &anyRawMsgConnectionOpenAck.TendermintClientState,

				ClientID:             anyClientId,
				CounterpartyClientID: anyCounterpartyClientId,
			}

			event := event_usecase.NewMsgIBCConnectionOpenAck(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenAck)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_ACK_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.CounterpartyConnectionID).To(Equal(anyParams.CounterpartyConnectionID))
			Expect(typedEvent.Params.Version).To(Equal(anyParams.Version))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.ProofTry).To(Equal(anyParams.ProofTry))
			Expect(typedEvent.Params.ProofClient).To(Equal(anyParams.ProofClient))
			Expect(typedEvent.Params.ProofConsensus).To(Equal(anyParams.ProofConsensus))
			Expect(typedEvent.Params.ConsensusHeight).To(Equal(anyParams.ConsensusHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.MaybeTendermintClientState).To(Equal(anyParams.MaybeTendermintClientState))

			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.CounterpartyClientID).To(Equal(anyParams.CounterpartyClientID))
		})
	})
})
