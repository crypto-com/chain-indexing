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
	Describe("MsgIBCCreateTendermintClient", func() {
		It("should parse Msg commands when there is MsgCreateClient in the transaction", func() {
			expected := `{
  "name": "/ibc.core.client.v1.MsgCreateClient.Created",
  "version": 1,
  "height": 5,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.client.v1.MsgCreateClient",
  "txHash": "7E34A75D8063BADF7B93538C23C88DEEF1FF14E7BE7F13AD6AD34E228C64538D",
  "msgIndex": 0,
  "params": {
	  "maybeTendermintLightClient": {
		"clientState": {
		  "@type": "/ibc.lightclients.tendermint.v1.ClientState",
		  "chainId": "devnet-2",
		  "trustLevel": { "numerator": "1", "denominator": "3" },
		  "trustingPeriod": "336h0m0s",
		  "unbondingPeriod": "504h0m0s",
		  "maxClockDrift": "5s",
		  "frozenHeight": { "revisionNumber": "0", "revisionHeight": "0" },
		  "latestHeight": { "revisionNumber": "2", "revisionHeight": "2" },
		  "proofSpecs": [
			{
			  "leafSpec": {
				"hash": "SHA256",
				"prehashKey": "NO_HASH",
				"prehashValue": "SHA256",
				"length": "VAR_PROTO",
				"prefix": "AA=="
			  },
			  "innerSpec": {
				"childOrder": [0, 1],
				"childSize": 33,
				"minPrefixLength": 4,
				"maxPrefixLength": 12,
				"emptyChild": null,
				"hash": "SHA256"
			  },
			  "maxDepth": 0,
			  "minDepth": 0
			},
			{
			  "leafSpec": {
				"hash": "SHA256",
				"prehashKey": "NO_HASH",
				"prehashValue": "SHA256",
				"length": "VAR_PROTO",
				"prefix": "AA=="
			  },
			  "innerSpec": {
				"childOrder": [0, 1],
				"childSize": 32,
				"minPrefixLength": 1,
				"maxPrefixLength": 1,
				"emptyChild": null,
				"hash": "SHA256"
			  },
			  "maxDepth": 0,
			  "minDepth": 0
			}
		  ],
		  "upgradePath": ["upgrade", "upgradedIBCState"],
		  "allowUpdateAfterExpiry": false,
		  "allowUpdateAfterMisbehaviour": false
		},
		"consensusState": {
		  "@type": "/ibc.lightclients.tendermint.v1.ConsensusState",
		  "timestamp": "2021-05-04T18:02:36.089446Z",
		  "root": { "hash": "bVjiQ29+V522NVFdx1BiVJnIBJV8Y1pYe9psvxZFAWg=" },
		  "nextValidatorsHash": "E3DE0D2B3237A02E9C20C34F9EE04F69F5861FBC2E2722A011CA9037FC67A7EC"
		}
	  },
      "maybeSoloMachineLightClient": null,
      "maybeLocalhostLightClient": null,
	  "signer": "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f",
	  "clientId": "07-tendermint-0",
	  "clientType": "07-tendermint"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_TENDERMINT_CLIENT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_TENDERMINT_CLIENT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CREATE_TENDERMINT_CLIENT_TXS_RESP)
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
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("/ibc.core.client.v1.MsgCreateClient.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgCreateClientEvent := untypedEvent.(*event.MsgIBCCreateClient)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgCreateClientEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgCreateClientEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f"}))
		})
	})
})
