package tendermint_test

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/crypto-com/chainindex/appinterface/tendermint"
	. "github.com/crypto-com/chainindex/infrastructure/tendermint"
	"github.com/crypto-com/chainindex/internal/utctime"

	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("HTTPClient", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	It("should implement Client", func() {
		var _ tendermint.Client = NewHTTPClient("http://localhost:26657")
	})

	Describe("RawBlockResults", func() {
		It("should return nil Events when there are no transactions nor events", func() {
			anyBlockHeight := int64(1)
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", fmt.Sprintf("height=%d", anyBlockHeight)),
					ghttp.RespondWith(http.StatusOK, BLOCK_RESULTS_EMPTY_EVENTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL())

			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(*blockResults).To(Equal(usecase_model.BlockResults{
				Height:           anyBlockHeight,
				TxsResults:       []usecase_model.BlockResultsTxsResult{},
				BeginBlockEvents: []usecase_model.BlockResultsEvent{},
				EndBlockEvents:   []usecase_model.BlockResultsEvent{},
				ValidatorUpdates: []usecase_model.BlockResultsValidatorUpdate{},
				ConsensusParamUpdates: usecase_model.BlockResultsConsensusParamUpdates{
					Block: usecase_model.BlockResultsConsensusParamUpdatesBlock{
						MaxBytes: "22020096",
						MaxGas:   "-1",
					},
					Evidence: usecase_model.BlockResultsConsensusParamUpdatesEvidence{
						MaxAgeNumBlocks: "100000",
						MaxAgeDuration:  "172800000000000",
					},
					Validator: usecase_model.BlockResultsConsensusParamsUpdatesValidator{
						PubKeyTypes: []string{"ed25519"},
					},
				},
			}))
		})

		It("should return parsed block results", func() {
			anyBlockHeight := int64(3813)
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", fmt.Sprintf("height=%d", anyBlockHeight)),
					ghttp.RespondWith(http.StatusOK, BLOCK_RESULTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL())

			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(jsoniter.MarshalToString(blockResults)).To(Equal("{\"Height\":367216,\"TxsResults\":[{\"code\":0,\"data\":\"\\n\\n\\n\\u0008delegate\",\"log\":[{\"msgIndex\":0,\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu\"},{\"key\":\"amount\",\"value\":\"19302674761\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"1913979901basetcro\"}]}]}],\"info\":\"\",\"gasWanted\":\"200000\",\"gasUsed\":\"143179\",\"events\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\"},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"},{\"key\":\"amount\",\"value\":\"20000basetcro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"}]}],\"codespace\":\"\"},{\"code\":0,\"data\":\"\\n\\n\\n\\u0008delegate\",\"log\":[{\"msgIndex\":0,\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl10gsqs8jzdlrem80shp0x6wx0jw7qu7m8cd29y5\"},{\"key\":\"amount\",\"value\":\"17338013566\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"2310941249basetcro\"}]}]}],\"info\":\"\",\"gasWanted\":\"200000\",\"gasUsed\":\"143737\",\"events\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\"},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"},{\"key\":\"amount\",\"value\":\"20000basetcro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"}]}],\"codespace\":\"\"}],\"BeginBlockEvents\":[{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha\"},{\"key\":\"sender\",\"value\":\"tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq\"},{\"key\":\"amount\",\"value\":\"17449528321basetcro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq\"}]},{\"type\":\"mint\",\"attributes\":[{\"key\":\"bonded_ratio\",\"value\":\"0.000809196054376644\"},{\"key\":\"inflation\",\"value\":\"0.013755821936855184\"},{\"key\":\"annual_provisions\",\"value\":\"110133046994204576.138016526579386288\"},{\"key\":\"amount\",\"value\":\"17449528321\"}]}],\"EndBlockEvents\":[{\"type\":\"commission\",\"attributes\":[{\"key\":\"amount\",\"value\":\"87247841.605000000000000000basetcro\"},{\"key\":\"validator\",\"value\":\"tcrocncl18p07yvmphymscz6tl4a7zmh93g0k6vy72ww4s4\"}]},{\"type\":\"rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"872478416.050000000000000000basetcro\"},{\"key\":\"validator\",\"value\":\"tcrocncl18p07yvmphymscz6tl4a7zmh93g0k6vy72ww4s4\"}]}],\"ValidatorUpdates\":[{\"PubKey\":{\"type\":\"tendermint.crypto.PublicKey_Ed25519\",\"pubKey\":\"SE5zeTjcYPXVrfcOva61QWokSZFfQu2h316fR6bB2dY=\",\"address\":\"CA721C3A05F500838DDD1B16F4E2D2D09E463218\"},\"power\":138525202},{\"PubKey\":{\"type\":\"tendermint.crypto.PublicKey_Ed25519\",\"pubKey\":\"Epmo3U6yXlxSDQzWZ8yBPOMHw2R85lc26RK98Rlo0oM=\",\"address\":\"E067FCE33F7FDBD0CE4872F8E240A7AD6E654726\"},\"power\":112904113}],\"ConsensusParamUpdates\":{\"block\":{\"maxBytes\":\"22020096\",\"maxGas\":\"-1\"},\"evidence\":{\"maxAgeNumBlocks\":\"100000\",\"maxAgeDuration\":\"172800000000000\"},\"validator\":{\"pubKeyTypes\":[\"ed25519\"]}}}"))
		})
	})

	Describe("Block", func() {
		It("should return block response", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block", "height=100"),
					ghttp.RespondWith(http.StatusOK, BLOCK_JSON),
				),
			)

			blockHeight := int64(100)
			client := NewHTTPClient(server.URL())
			block, _, err := client.Block(blockHeight)
			Expect(err).To(BeNil())
			blockTime, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.195143319Z")
			signature0Time, _ := utctime.Parse("2006-01-02T15:04:05.00000000Z", "2020-10-15T09:33:42.18646236Z")
			signature1Time, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.195143319Z")
			signature2Time, _ := utctime.Parse("2006-01-02T15:04:05.000000000Z", "2020-10-15T09:33:42.206633743Z")

			Expect(*block).To(Equal(usecase_model.Block{
				Height:          blockHeight,
				Hash:            "82C25937191D1CF73BE9222CB04CE35B7A1366CC5BB08D9BB9AB457712E4F2D1",
				Time:            blockTime,
				AppHash:         "6AE0920938F76727054BC2531247632C5C0521E2B91EA3A9864EA4FF55023D77",
				ProposerAddress: "384E5F30F02538C0A34CBFF32F8D5554671C9029",
				Txs:             []string{},
				Signatures: []usecase_model.BlockSignature{
					{
						BlockIdFlag:      2,
						ValidatorAddress: "384E5F30F02538C0A34CBFF32F8D5554671C9029",
						Timestamp:        signature0Time,
						Signature:        "Lnjz+jTGTk6DzqvbvdjFIG5zyzzpioiN37/B2pKi6ac/Kf2a+kmfeQA3CXnPO5GY/AoImfbVcCQs4asTSDCqCA==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
						Timestamp:        signature1Time,
						Signature:        "OpYxiF7QTAaG4NG2/iHWvts8ISQED6F+pNU5Cv2ts8c8mFW9J+g0oig3IXhvVG011uQjcr0CVw5P0eOXx1vYDg==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "6D9E4B4995037D608E365CE90436C24580ABCC33",
						Timestamp:        signature2Time,
						Signature:        "jQd4JNrvX6DKmqDZ9VqoKtxRIxQHrvPWd4XW+ayrtVakiIMCWoVf1GMvxLbXYg68CyjmbuAX2VhCD0gSnj3pAw==",
					},
				},
			}))
		})
	})
})

const (
	BLOCK_RESULTS_EMPTY_EVENTS_JSON = `
{
	"jsonrpc": "2.0",
	"id": -1,
	"result": {
		"height": "1",
		"txs_results": null,
		"begin_block_events": null,
		"end_block_events": null,
		"validator_updates": null,
		"consensus_param_updates": {
		  "block": {
			"max_bytes": "22020096",
			"max_gas": "-1"
		  },
		  "evidence": {
			"max_age_num_blocks": "100000",
			"max_age_duration": "172800000000000"
		  },
		  "validator": {
			"pub_key_types": [
			  "ed25519"
			]
		  }
		}
	}
}`

	// This BlockResults is a modified version and is NOT real chain data
	BLOCK_RESULTS_JSON = `
	{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "367216",
    "txs_results": [
      {
        "code": 0,
        "data": "CgoKCGRlbGVnYXRl",
        "log": "[{\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu\"},{\"key\":\"amount\",\"value\":\"19302674761\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"1913979901basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "143179",
        "events": [
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFwbTI3ZGpjczVkanhqc3h3M3Vucmt2M20zanR4ZGV4azczaHFlbA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjAwMDBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzFwbTI3ZGpjczVkanhqc3h3M3Vucmt2M20zanR4ZGV4azczaHFlbA==",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      },
      {
        "code": 0,
        "data": "CgoKCGRlbGVnYXRl",
        "log": "[{\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl10gsqs8jzdlrem80shp0x6wx0jw7qu7m8cd29y5\"},{\"key\":\"amount\",\"value\":\"17338013566\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"2310941249basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "143737",
        "events": [
          {
            "type": "transfer",
            "attributes": [
              {
                "key": "cmVjaXBpZW50",
                "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "dGNybzEwZ3NxczhqemRscmVtODBzaHAweDZ3eDBqdzdxdTdtOGRqZnV1aA==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjAwMDBiYXNldGNybw==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNybzEwZ3NxczhqemRscmVtODBzaHAweDZ3eDBqdzdxdTdtOGRqZnV1aA==",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      }
    ],
    "begin_block_events": [
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc0NDk1MjgzMjFiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDA4MDkxOTYwNTQzNzY2NDQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM3NTU4MjE5MzY4NTUxODQ=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTEwMTMzMDQ2OTk0MjA0NTc2LjEzODAxNjUyNjU3OTM4NjI4OA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc0NDk1MjgzMjE=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODcyNDc4NDEuNjA1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHAwN3l2bXBoeW1zY3o2dGw0YTd6bWg5M2cwazZ2eTcyd3c0czQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODcyNDc4NDE2LjA1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHAwN3l2bXBoeW1zY3o2dGw0YTd6bWg5M2cwazZ2eTcyd3c0czQ=",
            "index": true
          }
        ]
      }
    ],
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "SE5zeTjcYPXVrfcOva61QWokSZFfQu2h316fR6bB2dY="
            }
          }
        },
        "power": "138525202"
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "Epmo3U6yXlxSDQzWZ8yBPOMHw2R85lc26RK98Rlo0oM="
            }
          }
        },
        "power": "112904113"
      }
    ],
    "consensus_param_updates": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`
	BLOCK_JSON = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "82C25937191D1CF73BE9222CB04CE35B7A1366CC5BB08D9BB9AB457712E4F2D1",
      "parts": {
        "total": 1,
        "hash": "0A19FD939EBCE493C3126D99A6133F968AF454CAAEEAA1AD20EC4A3CFD60F2D5"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "100",
        "time": "2020-10-15T09:33:42.195143319Z",
        "last_block_id": {
          "hash": "1532E4FFBDE4FE8CCDF5654A097D534B8C6E2EBC4473F36CFE314C0421970C2E",
          "parts": {
            "total": 1,
            "hash": "EEDFCCF098B1695CE939CF4E395AA8FC0EEC9F4673E1418B29E8928489BEF06A"
          }
        },
        "last_commit_hash": "E7437018C93A3BEA020C76DF06A209BFD9E2EEB6F01A5D4AF1EC64BA6488E6F5",
        "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "validators_hash": "C97D9876B52DAD0DA395EA0203EF1D85682198E82AB90D2CE60DB1FB25117C12",
        "next_validators_hash": "C97D9876B52DAD0DA395EA0203EF1D85682198E82AB90D2CE60DB1FB25117C12",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "6AE0920938F76727054BC2531247632C5C0521E2B91EA3A9864EA4FF55023D77",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "384E5F30F02538C0A34CBFF32F8D5554671C9029"
      },
      "data": {
        "txs": []
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "99",
        "round": 0,
        "block_id": {
          "hash": "1532E4FFBDE4FE8CCDF5654A097D534B8C6E2EBC4473F36CFE314C0421970C2E",
          "parts": {
            "total": 1,
            "hash": "EEDFCCF098B1695CE939CF4E395AA8FC0EEC9F4673E1418B29E8928489BEF06A"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "384E5F30F02538C0A34CBFF32F8D5554671C9029",
            "timestamp": "2020-10-15T09:33:42.18646236Z",
            "signature": "Lnjz+jTGTk6DzqvbvdjFIG5zyzzpioiN37/B2pKi6ac/Kf2a+kmfeQA3CXnPO5GY/AoImfbVcCQs4asTSDCqCA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
            "timestamp": "2020-10-15T09:33:42.195143319Z",
            "signature": "OpYxiF7QTAaG4NG2/iHWvts8ISQED6F+pNU5Cv2ts8c8mFW9J+g0oig3IXhvVG011uQjcr0CVw5P0eOXx1vYDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6D9E4B4995037D608E365CE90436C24580ABCC33",
            "timestamp": "2020-10-15T09:33:42.206633743Z",
            "signature": "jQd4JNrvX6DKmqDZ9VqoKtxRIxQHrvPWd4XW+ayrtVakiIMCWoVf1GMvxLbXYg68CyjmbuAX2VhCD0gSnj3pAw=="
          }
        ]
      }
    }
  }
}`
)
