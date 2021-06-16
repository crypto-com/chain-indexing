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

	Describe("En/DecodeMsgIBCUpdateClient", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "07-tendermint-0"
			anySigner := "cro1kkxdtp7yjtprvf2memjy8lmh4spfevea9yjxgn"
			anyRawHeader := `
{
  "@type": "/ibc.lightclients.tendermint.v1.Header",
  "signed_header": {
	"header": {
	  "version": {
		"block": "11",
		"app": "0"
	  },
	  "chain_id": "devnet-2",
	  "height": "10",
	  "time": "2021-06-04T06:32:50.317992Z",
	  "last_block_id": {
		"hash": "fjTfL3Lh/TBPfCgn+3hmM1xatobldIl4VIM4Rc6rFQM=",
		"part_set_header": {
		  "total": 1,
		  "hash": "tRlaaIW4NRm5o/HoovFoiGIvBnhBYYbmymXBanGzAd0="
		}
	  },
	  "last_commit_hash": "jZvjjDeaAPNDQyUiYxqgr2Lg1PD+7it5SdYryI8GvLc=",
	  "data_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	  "validators_hash": "JfgGoWDACxlYhafc1+k9EfoyeVKb5OKuPCClHSCPAEY=",
	  "next_validators_hash": "JfgGoWDACxlYhafc1+k9EfoyeVKb5OKuPCClHSCPAEY=",
	  "consensus_hash": "BICRvH3cKD93v7+R1zxE2ljD34qcvIZ0Bdi389qtoi8=",
	  "app_hash": "bJyiARV5/N+cnPPqdpdwInBfmm5BGBdjdC/yiChKoVo=",
	  "last_results_hash": "qY+I1bPrhzX4RpXbyx8pROu/2QKykkA6Aova8FBVDcw=",
	  "evidence_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	  "proposer_address": "b7Nb1HguYo+UMD7sUjToLIEEcKw="
	},
	"commit": {
	  "height": "10",
	  "round": 0,
	  "block_id": {
		"hash": "uZSTJ9c1V0rm3ivA2PSQYPvIIzzwzRvnUSYElD1q86c=",
		"part_set_header": {
		  "total": 1,
		  "hash": "9hgeGjUmYkYySwDxWMPb1kCheR4s68p6kvPo/5Zm6DM="
		}
	  },
	  "signatures": [
		{
		  "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
		  "validator_address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		  "timestamp": "2021-06-04T06:32:55.678031Z",
		  "signature": "vHyTqPllhI/7Ta1SRvMi4iVDRbJMISzSZQbyY7naO29duhUyeNSF/kOsBYven5f/1r3Z/PJlD6DERH1iPKDeCw=="
		}
	  ]
	}
  },
  "validator_set": {
	"validators": [
	  {
		"address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		"pub_key": {
		  "ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
		},
		"voting_power": "10000000000",
		"proposer_priority": "0"
	  }
	],
	"proposer": {
	  "address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
	  "pub_key": {
		"ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
	  },
	  "voting_power": "10000000000",
	  "proposer_priority": "0"
	},
	"total_voting_power": "10000000000"
  },
  "trusted_height": {
	"revision_number": "2",
	"revision_height": "5"
  },
  "trusted_validators": {
	"validators": [
	  {
		"address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		"pub_key": {
		  "ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
		},
		"voting_power": "10000000000",
		"proposer_priority": "0"
	  }
	],
	"proposer": {
	  "address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
	  "pub_key": {
		"ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
	  },
	  "voting_power": "10000000000",
	  "proposer_priority": "0"
	},
	"total_voting_power": "10000000000"
  }
}
`
			var anyHeader ibc_model.TendermintLightClientUpdateHeader
			json.MustUnmarshalFromString(anyRawHeader, &anyHeader)

			anyParams := ibc_model.MsgUpdateClientParams{
				MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{Header: anyHeader},
				ClientID:                         anyClientId,
				Signer:                           anySigner,
			}

			event := event_usecase.NewMsgIBCUpdateClient(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_UPDATE_CLIENT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCUpdateClient)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_UPDATE_CLIENT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.MaybeTendermintLightClientUpdate).NotTo(BeNil())
			Expect(typedEvent.MaybeTendermintLightClientUpdate.Header.Type).To(Equal(
				"/ibc.lightclients.tendermint.v1.Header",
			))
			Expect(typedEvent.MaybeTendermintLightClientUpdate.Header.TrustedHeight.RevisionHeight).To(Equal("5"))
			Expect(typedEvent.Signer).To(Equal(anyParams.Signer))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyClientId := "07-tendermint-0"
			anySigner := "cro1kkxdtp7yjtprvf2memjy8lmh4spfevea9yjxgn"
			anyRawHeader := `
{
  "@type": "/ibc.lightclients.tendermint.v1.Header",
  "signed_header": {
	"header": {
	  "version": {
		"block": "11",
		"app": "0"
	  },
	  "chain_id": "devnet-2",
	  "height": "10",
	  "time": "2021-06-04T06:32:50.317992Z",
	  "last_block_id": {
		"hash": "fjTfL3Lh/TBPfCgn+3hmM1xatobldIl4VIM4Rc6rFQM=",
		"part_set_header": {
		  "total": 1,
		  "hash": "tRlaaIW4NRm5o/HoovFoiGIvBnhBYYbmymXBanGzAd0="
		}
	  },
	  "last_commit_hash": "jZvjjDeaAPNDQyUiYxqgr2Lg1PD+7it5SdYryI8GvLc=",
	  "data_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	  "validators_hash": "JfgGoWDACxlYhafc1+k9EfoyeVKb5OKuPCClHSCPAEY=",
	  "next_validators_hash": "JfgGoWDACxlYhafc1+k9EfoyeVKb5OKuPCClHSCPAEY=",
	  "consensus_hash": "BICRvH3cKD93v7+R1zxE2ljD34qcvIZ0Bdi389qtoi8=",
	  "app_hash": "bJyiARV5/N+cnPPqdpdwInBfmm5BGBdjdC/yiChKoVo=",
	  "last_results_hash": "qY+I1bPrhzX4RpXbyx8pROu/2QKykkA6Aova8FBVDcw=",
	  "evidence_hash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	  "proposer_address": "b7Nb1HguYo+UMD7sUjToLIEEcKw="
	},
	"commit": {
	  "height": "10",
	  "round": 0,
	  "block_id": {
		"hash": "uZSTJ9c1V0rm3ivA2PSQYPvIIzzwzRvnUSYElD1q86c=",
		"part_set_header": {
		  "total": 1,
		  "hash": "9hgeGjUmYkYySwDxWMPb1kCheR4s68p6kvPo/5Zm6DM="
		}
	  },
	  "signatures": [
		{
		  "block_id_flag": "BLOCK_ID_FLAG_COMMIT",
		  "validator_address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		  "timestamp": "2021-06-04T06:32:55.678031Z",
		  "signature": "vHyTqPllhI/7Ta1SRvMi4iVDRbJMISzSZQbyY7naO29duhUyeNSF/kOsBYven5f/1r3Z/PJlD6DERH1iPKDeCw=="
		}
	  ]
	}
  },
  "validator_set": {
	"validators": [
	  {
		"address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		"pub_key": {
		  "ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
		},
		"voting_power": "10000000000",
		"proposer_priority": "0"
	  }
	],
	"proposer": {
	  "address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
	  "pub_key": {
		"ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
	  },
	  "voting_power": "10000000000",
	  "proposer_priority": "0"
	},
	"total_voting_power": "10000000000"
  },
  "trusted_height": {
	"revision_number": "2",
	"revision_height": "5"
  },
  "trusted_validators": {
	"validators": [
	  {
		"address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
		"pub_key": {
		  "ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
		},
		"voting_power": "10000000000",
		"proposer_priority": "0"
	  }
	],
	"proposer": {
	  "address": "b7Nb1HguYo+UMD7sUjToLIEEcKw=",
	  "pub_key": {
		"ed25519": "awslKyoIYmj2+4JGukqr32eCqj9JMtuSI8xBY102brY="
	  },
	  "voting_power": "10000000000",
	  "proposer_priority": "0"
	},
	"total_voting_power": "10000000000"
  }
}
`
			var anyHeader ibc_model.TendermintLightClientUpdateHeader
			json.MustUnmarshalFromString(anyRawHeader, &anyHeader)

			anyParams := ibc_model.MsgUpdateClientParams{
				MaybeTendermintLightClientUpdate: &ibc_model.TendermintLightClientUpdate{Header: anyHeader},
				ClientID:                         anyClientId,
				Signer:                           anySigner,
			}

			event := event_usecase.NewMsgIBCUpdateClient(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_UPDATE_CLIENT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCUpdateClient)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_UPDATE_CLIENT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.MaybeTendermintLightClientUpdate).NotTo(BeNil())
			Expect(typedEvent.MaybeTendermintLightClientUpdate.Header.Type).To(Equal(
				"/ibc.lightclients.tendermint.v1.Header",
			))
			Expect(typedEvent.MaybeTendermintLightClientUpdate.Header.TrustedHeight.RevisionHeight).To(Equal("5"))
			Expect(typedEvent.Signer).To(Equal(anyParams.Signer))
		})
	})
})
