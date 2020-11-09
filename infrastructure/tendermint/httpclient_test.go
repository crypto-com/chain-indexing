package tendermint_test

import (
	"github.com/crypto-com/chainindex/internal/primptr"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"

	"github.com/crypto-com/chainindex/appinterface/tendermint"
	. "github.com/crypto-com/chainindex/infrastructure/tendermint"
	"github.com/crypto-com/chainindex/internal/utctime"
	"github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("HTTPClient", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	It("should implement Client", func() {
		var _ tendermint.Client = NewHTTPClient("http://localhost:26657")
	})

	Describe("BlockResults", func() {
		It("should return nil Events when there are no transactions nor events", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", "height=1"),
					ghttp.RespondWith(http.StatusOK, BLOCK_RESULTS_EMPTY_EVENTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL())

			anyBlockHeight := int64(1)
			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(*blockResults).To(Equal(model.BlockResults{
				Height:           anyBlockHeight,
				TxsEvents:        nil,
				BeginBlockEvents: nil,
				ValidatorUpdates: nil,
			}))
		})

		It("should return parsed block results", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/block_results", "height=3813"),
					ghttp.RespondWith(http.StatusOK, BLOCK_RESULTS_JSON),
				),
			)

			client := NewHTTPClient(server.URL())

			anyBlockHeight := int64(3813)
			blockResults, err := client.BlockResults(anyBlockHeight)
			Expect(err).To(BeNil())
			Expect(*blockResults).To(Equal(model.BlockResults{
				Height: anyBlockHeight,
				TxsEvents: [][]model.BlockResultsEvent{
					{
						{
							Type: "valid_txs",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "ZmVl",
									Value: "MC4wMDAwMDQ2OQ==",
								},
								{
									Key:   "dHhpZA==",
									Value: "YmEyMjMwYTA0OTIyZDNmMDFkNDE5OTljNmFkYmUwNmZjOGE5ODQxM2IyZDU1YWM3ZjlhYzMwZmVjMzlmYzdiMg==",
								},
							},
						},
					},
					{
						{
							Type: "valid_txs",
							Attributes: []model.BlockResultsEventAttribute{
								{
									Key:   "ZmVl",
									Value: "MC4wMDAwMDQ2OQ==",
								},
								{
									Key:   "dHhpZA==",
									Value: "N2I1YWMzNmY2M2ZmYjZlZTEzMjg5ZDQ5ZDljZmRhY2UzZjhkMjM0NDY5ZWIwNTc1NWY1ZjVhYzgxYjNhNDVhNg==",
								},
							},
						},
					},
				},
				BeginBlockEvents: []model.BlockResultsEvent{
					{
						Type: "staking_change",
						Attributes: []model.BlockResultsEventAttribute{
							{
								Key:   "c3Rha2luZ19hZGRyZXNz",
								Value: "MHg2ZGJkNWI4ZmUwZGFkNDk0NDY1YWE3NTc0ZGVmYmE3MTFjMTg0MTAy",
							},
							{
								Key:   "c3Rha2luZ19vcHR5cGU=",
								Value: "cmV3YXJk",
							},
							{
								Key:   "c3Rha2luZ19kaWZm",
								Value: "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d",
							},
						},
					},
					{
						Type: "staking_change",
						Attributes: []model.BlockResultsEventAttribute{
							{
								Key:   "c3Rha2luZ19hZGRyZXNz",
								Value: "MHg2ZmMxZTMxMjRhN2VkMDdmMzcxMDM3OGI2OGY3MDQ2YzczMDAxNzlk",
							},
							{
								Key:   "c3Rha2luZ19vcHR5cGU=",
								Value: "cmV3YXJk",
							},
							{
								Key:   "c3Rha2luZ19kaWZm",
								Value: "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d",
							},
						},
					},
					{
						Type: "staking_change",
						Attributes: []model.BlockResultsEventAttribute{
							{
								Key:   "c3Rha2luZ19hZGRyZXNz",
								Value: "MHhiOGM2ODg2ZGEwOWUxMmRiOGFlYmZjODEwOGM2N2NlMmJhMDg2YWM2",
							},
							{
								Key:   "c3Rha2luZ19vcHR5cGU=",
								Value: "cmV3YXJk",
							},
							{
								Key:   "c3Rha2luZ19kaWZm",
								Value: "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d",
							},
						},
					},
					{
						Type: "reward",
						Attributes: []model.BlockResultsEventAttribute{
							{
								Key:   "bWludGVk",
								Value: "IjE0NTc2OTcwNjEwMDAwIg==",
							},
						},
					},
				},
				ValidatorUpdates: []model.BlockResultsValidator{
					{
						PubKey: model.BlockResultsValidatorPubKey{
							Type:    "tendermint.crypto.PublicKey_Ed25519",
							PubKey:  "rXhu7xhqYBtJftVLKxvKN0XnpyOzxFnUEfAhD1dEF/8=",
							Address: "34C725CABA703269B3F1D1A907A84DE5FEE96469",
						},
						Power: primptr.String("60000000"),
					},
					{
						PubKey: model.BlockResultsValidatorPubKey{
							Type:    "tendermint.crypto.PublicKey_Ed25519",
							PubKey:  "tDLheZJwsA8oYEwarR6/X+zAmNKMLHTVkh/fvcLqcwA=",
							Address: "D527DAECDE0501CF2E785A8DC0D9F4A64760F0BB",
						},
						Power: primptr.String("80000000"),
					},
				},
			}))
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

			Expect(*block).To(Equal(model.Block{
				Height:          blockHeight,
				Hash:            "82C25937191D1CF73BE9222CB04CE35B7A1366CC5BB08D9BB9AB457712E4F2D1",
				Time:            blockTime,
				AppHash:         "6AE0920938F76727054BC2531247632C5C0521E2B91EA3A9864EA4FF55023D77",
				ProposerAddress: "384E5F30F02538C0A34CBFF32F8D5554671C9029",
				Txs:             []string{},
				Signatures: []model.BlockSignature{
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
		"consensus_param_updates": null
	}
}`
	BLOCK_RESULTS_JSON = `
{
	"jsonrpc": "2.0",
	"id": -1,
	"result": {
		"height": "3813",
		"txs_results": [
			{
				"code": 0,
				"data": null,
				"log": "",
				"info": "",
				"gasWanted": "0",
				"gasUsed": "0",
				"events": [
					{
						"type": "valid_txs",
						"attributes": [
						{
							"key": "ZmVl",
							"value": "MC4wMDAwMDQ2OQ=="
						},
						{
							"key": "dHhpZA==",
							"value": "YmEyMjMwYTA0OTIyZDNmMDFkNDE5OTljNmFkYmUwNmZjOGE5ODQxM2IyZDU1YWM3ZjlhYzMwZmVjMzlmYzdiMg=="
						}
						]
					}
				],
				"codespace": ""
			},
			{
				"code": 0,
				"data": null,
				"log": "",
				"info": "",
				"gasWanted": "0",
				"gasUsed": "0",
				"events": [
					{
						"type": "valid_txs",
						"attributes": [
						{
							"key": "ZmVl",
							"value": "MC4wMDAwMDQ2OQ=="
						},
						{
							"key": "dHhpZA==",
							"value": "N2I1YWMzNmY2M2ZmYjZlZTEzMjg5ZDQ5ZDljZmRhY2UzZjhkMjM0NDY5ZWIwNTc1NWY1ZjVhYzgxYjNhNDVhNg=="
						}
						]
					}
				],
				"codespace": ""
			}
		],
		"begin_block_events": [
			{
				"type": "staking_change",
				"attributes": [
					{
						"key": "c3Rha2luZ19hZGRyZXNz",
						"value": "MHg2ZGJkNWI4ZmUwZGFkNDk0NDY1YWE3NTc0ZGVmYmE3MTFjMTg0MTAy"
					},
					{
						"key": "c3Rha2luZ19vcHR5cGU=",
						"value": "cmV3YXJk"
					},
					{
						"key": "c3Rha2luZ19kaWZm",
						"value": "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d"
					}
				]
			},
			{
				"type": "staking_change",
				"attributes": [
					{
						"key": "c3Rha2luZ19hZGRyZXNz",
						"value": "MHg2ZmMxZTMxMjRhN2VkMDdmMzcxMDM3OGI2OGY3MDQ2YzczMDAxNzlk"
					},
					{
						"key": "c3Rha2luZ19vcHR5cGU=",
						"value": "cmV3YXJk"
					},
					{
						"key": "c3Rha2luZ19kaWZm",
						"value": "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d"
					}
				]
			},
			{
				"type": "staking_change",
				"attributes": [
					{
						"key": "c3Rha2luZ19hZGRyZXNz",
						"value": "MHhiOGM2ODg2ZGEwOWUxMmRiOGFlYmZjODEwOGM2N2NlMmJhMDg2YWM2"
					},
					{
						"key": "c3Rha2luZ19vcHR5cGU=",
						"value": "cmV3YXJk"
					},
					{
						"key": "c3Rha2luZ19kaWZm",
						"value": "W3sia2V5IjoiQm9uZGVkIiwidmFsdWUiOiI0ODU4OTkwMjAzMzMzIn1d"
					}
				]
			},
			{
				"type": "reward",
				"attributes": [
					{
						"key": "bWludGVk",
						"value": "IjE0NTc2OTcwNjEwMDAwIg=="
					}
				]
			}
		],
		"end_block_events": [
			{
				"type": "block_filter",
				"attributes": [
					{
						"key": "ZXRoYmxvb20=",
						"value": "AAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
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
							"ed25519": "rXhu7xhqYBtJftVLKxvKN0XnpyOzxFnUEfAhD1dEF/8="
						}
					}
				},
				"power": "60000000"
			},
			{
				"pub_key": {
				"Sum": {
					"type": "tendermint.crypto.PublicKey_Ed25519",
						"value": {
							"ed25519": "tDLheZJwsA8oYEwarR6/X+zAmNKMLHTVkh/fvcLqcwA="
						}
					}
				},
				"power": "80000000"
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
