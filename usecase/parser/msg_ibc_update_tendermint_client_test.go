package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCUpdateClient", func() {
		It("should parse Msg commands when there is MsgUpdateClient in the transaction", func() {
			expected := `{
  "name": "/ibc.core.client.v1.MsgUpdateClient.Created",
  "version": 1,
  "height": 7,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.client.v1.MsgUpdateClient",
  "txHash": "B803E2AD7B78C667EC242CD73315E6BCADE4B4B63DE5B71BB7055D5B454E9E12",
  "msgIndex": 0,
  "params": {
    "maybeTendermintLightClientUpdate": {
      "header": {
        "@type": "/ibc.lightclients.tendermint.v1.Header",
        "signedHeader": {
          "header": {
            "version": {
              "block": "11",
              "app": "0"
            },
            "chainId": "devnet-1",
            "height": "8",
            "time": "2021-05-04T18:02:58.538591Z",
            "lastBlockId": {
              "hash": "i+5baLXYrjDEde2ms5Xm/0lE0x1LnwhOSfAuaR0ICcM=",
              "partSetHeader": {
                "total": 1,
                "hash": "8oVYCsGpkahV5132UKj2W9UvVpls4+oUUaaoXVnocUs="
              }
            },
            "lastCommitHash": "++Md6XtfUIxYuNTd/C6cnyrVpUZWgcC1tb1h4iEdU9U=",
            "dataHash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
            "validatorsHash": "P0rvhgPNZYLAcrfqy78b3ZWRmWr2WWF8xyaDuWm+jfU=",
            "nextValidatorsHash": "P0rvhgPNZYLAcrfqy78b3ZWRmWr2WWF8xyaDuWm+jfU=",
            "consensusHash": "BICRvH3cKD93v7+R1zxE2ljD34qcvIZ0Bdi389qtoi8=",
            "appHash": "FYru2c/X7TL9cgW2CdT+RE9UOkek3ZzWON+VY/om/MM=",
            "lastResultsHash": "/wEBHCVZ+gZNJSrP3hvste0PUuxpJsXbjpyl2Ae1tYg=",
            "evidenceHash": "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
            "proposerAddress": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w="
          },
          "commit": {
            "height": "8",
            "round": 0,
            "blockId": {
              "hash": "tFQsxnaIAA5gMt7UQ8Jfhbogv9X7tCzTN64PdNzNqQI=",
              "partSetHeader": {
                "total": 1,
                "hash": "YDJUUy2Vs8EerqNeNo+GrSUitiNTe0sn1aVsnWCR2CY="
              }
            },
            "signatures": [
              {
                "blockIdFlag": "BLOCK_ID_FLAG_COMMIT",
                "validatorAddress": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w=",
                "timestamp": "2021-05-04T18:03:03.699921Z",
                "signature": "MO8IjHx/1UfeIxuZGhZwavxl7y/HCES+N3N0Sz7E8LBy7PLejg8Lxw5gvjvtcZumJYtMkmxuDfiUWHvgPlz5BQ=="
              }
            ]
          }
        },
        "validatorSet": {
          "validators": [
            {
              "address": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w=",
              "pubKey": {
                "ed25519": "XXkBQNXIhVrQqdi6Qzb4opk/0kgETITAm42jfdullFE="
              },
              "votingPower": "10000000000",
              "proposerPriority": "0"
            }
          ],
          "proposer": {
            "address": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w=",
            "pubKey": {
              "ed25519": "XXkBQNXIhVrQqdi6Qzb4opk/0kgETITAm42jfdullFE="
            },
            "votingPower": "10000000000",
            "proposerPriority": "0"
          },
          "totalVotingPower": "10000000000"
        },
        "trustedHeight": {
          "revisionNumber": "1",
          "revisionHeight": "5"
        },
        "trustedValidators": {
          "validators": [
            {
              "address": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w=",
              "pubKey": {
                "ed25519": "XXkBQNXIhVrQqdi6Qzb4opk/0kgETITAm42jfdullFE="
              },
              "votingPower": "10000000000",
              "proposerPriority": "0"
            }
          ],
          "proposer": {
            "address": "zRrIv8LAH8PC2vJmKNHMgg/Qz2w=",
            "pubKey": {
              "ed25519": "XXkBQNXIhVrQqdi6Qzb4opk/0kgETITAm42jfdullFE="
            },
            "votingPower": "10000000000",
            "proposerPriority": "0"
          },
          "totalVotingPower": "10000000000"
        }
      }
    },
    "maybeSoloMachineLightClientUpdate":null,
    "clientId": "07-tendermint-0",
    "clientType": "07-tendermint",
    "consensusHeight": {
      "revisionNumber": "1",
      "revisionHeight": "8"
    },
    "signer": "cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_UPDATE_TENDERMINT_CLIENT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_UPDATE_TENDERMINT_CLIENT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_UPDATE_TENDERMINT_CLIENT_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("/ibc.core.client.v1.MsgUpdateClient.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgCreateClientEvent := untypedEvent.(*event.MsgIBCUpdateClient)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgCreateClientEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgCreateClientEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[0]).To(Equal("cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss"))
		})
	})
})
