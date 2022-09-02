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
	Describe("MsgIBCConnectionOpenAck", func() {
		It("should parse Msg commands when there is MsgIBCConnectionOpenAck in the transaction", func() {
			expected := `{
  "name": "/ibc.core.connection.v1.MsgConnectionOpenAck.Created",
  "version": 1,
  "height": 11,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.connection.v1.MsgConnectionOpenAck",
  "txHash": "831EDBD3F114BDA5F49276A9AD591F4A3B8B82073101B142CDFA02A2AA21889B",
  "msgIndex": 1,
  "params": {
	"connectionId": "connection-0",
	"counterpartyConnectionId": "connection-0",
	"version": {
	  "identifier": "1",
	  "features": [
		"ORDER_ORDERED",
		"ORDER_UNORDERED"
	  ]
	},
	"proofHeight": {
	  "revisionNumber": "2",
	  "revisionHeight": "11"
	},
	"proofTry": "CroCCrcCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgCIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhIiKQgBEiUCBBQg82rv+J5yJTsqmzR59NDsb2QSM+WLnXIG6KcNLkM+OkggIikIARIlBAgUID4MsCa4ABOa69QPz6AbX5/lzcJdKUsavb5rnsXVeBlKICIrCAESBAYOFCAaISDTUETUzwU2dxU/8zMto7/LcYNlrxNPaK70aHoUWiD38SIpCAESJQgUFCAd54gQJQY2sQyko0wVJ2M3tABQQOuK2JsO4oj298eF0iAK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proofClient": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMRIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAEQCEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAhQiKwgBEgQCBBQgGiEgJfqlX0WUAPCdhTWwIU4QDSykCXpE1UJWobmRPN3UTBsiKwgBEgQEBhQgGiEgkmNVfGb0FEP31ym+IQvt8UFwE/+FwjzUm5ko3UmI+2YiKwgBEgQIFBQgGiEg4xjkOtDX3dH/OhsfvOIzyBDw1UI6TKSzjRcDP5QMI2UK0wEK0AEKA2liYxIg71e3Rs2XG2Ipt/mZEWvdHjEb4Kae7U4q/d+ZLiqq3qoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQGXUquofWyM7JLfJMs08veWDcXyNrjfuFddUpyfjogFiiIlCAESIQEDZ2s3aYqk1e0B4capTiXk7mWYJhxm4+LdAiwISABZFCInCAESAQEaIOh6nzuULJTKLW7LXM9MNkDflz0qO91qfguYPH7BNGRQ",
	"proofConsensus": "CvUCCvICCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMS04EoUBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElMKCwi/gYaHBhDAyrNBEiIKIGrO3+I45yc7lT3lJhHz6S3kQL4vPCw6A6lQb7jL31L0GiDDvmZM+0VcRAU839+33ipkg7bBliRUVb9FPTh4xYCTSxoLCAEYASABKgMAAhQiKQgBEiUCBBQgBcFP0jQiTKHo9UDsNzOd7DE0+3kPyfthUzsi6kPNnAsgIisIARIEBAgUIBohIJNbKob+2Y6eG2iS472XaQGWG2v9nJvqc3l+E+dJa1DYIisIARIEBg4UIBohINNQRNTPBTZ3FT/zMy2jv8txg2WvE09orvRoehRaIPfxIikIARIlCBQUIB3niBAlBjaxDKSjTBUnYze0AFBA64rYmw7iiPb3x4XSIArTAQrQAQoDaWJjEiDvV7dGzZcbYim3+ZkRa90eMRvgpp7tTir935kuKqreqhoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAZdSq6h9bIzskt8kyzTy95YNxfI2uN+4V11SnJ+OiAWKIiUIARIhAQNnazdpiqTV7QHhxqlOJeTuZZgmHGbj4t0CLAhIAFkUIicIARIBARog6HqfO5QslMotbstcz0w2QN+XPSo73Wp+C5g8fsE0ZFA=",
	"consensusHeight": {
	  "revisionNumber": "1",
	  "revisionHeight": "8"
	},
	"signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
	"maybeTendermintClientState": {
	  "@type": "/ibc.lightclients.tendermint.v1.ClientState",
	  "chainId": "devnet-1",
	  "trustLevel": {
		"numerator": "1",
		"denominator": "3"
	  },
	  "trustingPeriod": "336h0m0s",
	  "unbondingPeriod": "504h0m0s",
	  "maxClockDrift": "5s",
	  "frozenHeight": {
		"revisionNumber": "0",
		"revisionHeight": "0"
	  },
	  "latestHeight": {
		"revisionNumber": "1",
		"revisionHeight": "8"
	  },
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
			"childOrder": [
			  0,
			  1
			],
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
			"childOrder": [
			  0,
			  1
			],
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
	  "upgradePath": [
		"upgrade",
		"upgradedIBCState"
	  ],
	  "allowUpdateAfterExpiry": false,
	  "allowUpdateAfterMisbehaviour": false
	},

    "clientId": "07-tendermint-0",
    "counterpartyClientId": "07-tendermint-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_ACK_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_ACK_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CONNECTION_OPEN_ACK_TXS_RESP)
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
			cmd := cmds[1]
			Expect(cmd.Name()).To(Equal("/ibc.core.connection.v1.MsgConnectionOpenAck.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCConnectionOpenAck)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"))
		})
	})
})
