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
	Describe("MsgIBCConnectionOpenTry", func() {
		It("should parse Msg commands when there is MsgIBCConnectionOpenTry in the transaction", func() {
			expected := `{
  "name": "/ibc.core.connection.v1.MsgConnectionOpenTry.Created",
  "version": 1,
  "height": 7,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.connection.v1.MsgConnectionOpenTry",
  "txHash": "B803E2AD7B78C667EC242CD73315E6BCADE4B4B63DE5B71BB7055D5B454E9E12",
  "msgIndex": 1,
  "params": {
	"clientId": "07-tendermint-0",
	"previousConnectionId": "",
	"counterparty": {
	  "clientId": "07-tendermint-0",
	  "connectionId": "connection-0",
	  "prefix": {
		"keyPrefix": "aWJj"
	  }
	},
	"delayPeriod": "0",
	"counterpartyVersions": [
	  {
		"identifier": "1",
		"features": [
		  "ORDER_ORDERED",
		  "ORDER_UNORDERED"
		]
	  }
	],
	"proofHeight": {
	  "revisionNumber": "1",
	  "revisionHeight": "8"
	},
	"proofInit": "Cq4CCqsCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASUgoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgBIhgKDzA3LXRlbmRlcm1pbnQtMBoFCgNpYmMaCwgBGAEgASoDAAIMIisIARIEAgQOIBohIKC9H3IwOG4Tpg/Nymt8yMllJEjWt9aeYeuMwflKc4zFIisIARIEBAgOIBohIHV1PdysEO47G2t5+mlniRDH1Pad0QTeX4qIqIhnTO0wIikIARIlBg4OIIIVttNVy/4Nk+sDJg6RVumpww3erCRz3wr8gpVu3e76ICIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
	"proofClient": "CucCCuQCCiNjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jbGllbnRTdGF0ZRKoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMhIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAIQBEIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRoLCAEYASABKgMAAg4iKwgBEgQCBA4gGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQEBg4gGiEg3quYtD6jLHzeIroI7j0I7LcgdqygTR1W4dnk3us6cJQiKwgBEgQIFA4gGiEgpk5CYIA5vc4apf6bk8arEpuObvNs1tN4ml/j+Ji0ZykK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
	"proofConsensus": "CvYCCvMCCitjbGllbnRzLzA3LXRlbmRlcm1pbnQtMC9jb25zZW5zdXNTdGF0ZXMvMi00EoYBCi4vaWJjLmxpZ2h0Y2xpZW50cy50ZW5kZXJtaW50LnYxLkNvbnNlbnN1c1N0YXRlElQKDAjGmMaEBhCYnc+kARIiCiB5FDU4YJfh5rt8nXe3L5vEpRz0QglvmY7w9AUf/jWVtBog494NKzI3oC6cIMNPnuBPafWGH7wuJyKgEcqQN/xnp+waCwgBGAEgASoDAAIOIisIARIEAgQOIBohIK7bCMEuGbJOWa02nJkzHLbqFqogpoHwkmOl4CWI2/KHIikIARIlBAYOIBAW83IGMWt6nb7HNyJbW0Y3gaTLyiY73bO6+fGVfsn/ICIrCAESBAYODiAaISDAshei6GmugrMmo3DjCOIosuZkuRkcEPudaBJlzPzcSCIpCAESJQgUDiAzRa8KWN9tFt1iWkSU+zRa+LSal8CV5ewgt4W3zXjYRyAK0wEK0AEKA2liYxIg+76VX/C4YvtTKfW1ie/T9+NCuzZAfBhf01QJH8rrQeoaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEPUcrtUw6mvYoZwUhUfLiyrKXUuj0qkRI6iCeYKEo2qCIlCAESIQE0UuDWQ3XeVI3/NqZywxJAFuIHa9knsSJ0JwNZLRx9RiInCAESAQEaII2UrjCpSCW81DXyHpk7DKTomCBv04nLtEpKQcuu46jy",
	"consensusHeight": {
	  "revisionNumber": "2",
	  "revisionHeight": "4"
	},
	"signer": "cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss",
	"maybeTendermintClientState": {
	  "@type": "/ibc.lightclients.tendermint.v1.ClientState",
	  "chainId": "devnet-2",
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
		"revisionNumber": "2",
		"revisionHeight": "4"
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
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_TRY_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_TRY_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CONNECTION_OPEN_TRY_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.connection.v1.MsgConnectionOpenTry.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCConnectionOpenTry)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("cro1rhuxl5q4flq48fjxp0lhklhjmh5gjrpjvjqwss"))
		})
	})
})
