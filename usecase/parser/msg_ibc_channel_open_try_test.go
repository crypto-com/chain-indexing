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
	Describe("MsgIBCChannelOpenInit", func() {
		It("should parse Msg commands when there is MsgChannelOpenInit in the transaction", func() {
			expected := `{
  "name": "/ibc.core.channel.v1.MsgChannelOpenTry.Created",
  "version": 1,
  "height": 16,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgChannelOpenTry",
  "txHash": "CB12276476D61977AB1625CB5BBA9C1890EF3D9592DB7A9A5C5EC32D941A7B81",
  "msgIndex": 1,
  "params": {
	"portId": "transfer",
	"previousChannelId": "",
	"channel": {
	  "state": "STATE_TRYOPEN",
	  "ordering": "ORDER_UNORDERED",
	  "counterparty": {
		"portId": "transfer",
		"channelId": "channel-0"
	  },
	  "connectionHops": [
		"connection-0"
	  ],
	  "version": "ics20-1"
	},
	"counterpartyVersion": "ics20-1",
	"proofInit": "CpwCCpkCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASJwgBEAEaCgoIdHJhbnNmZXIiDGNvbm5lY3Rpb24tMCoHaWNzMjAtMRoLCAEYASABKgMAAhwiKwgBEgQCBBwgGiEgxASpVteatK9vbv42ysU8IwbBDKqExJmoe4EnC2fQrXUiKwgBEgQEBhwgGiEgiJpZZIZAU1ms8QPiW3ZruhdEeChGmtXNpo+ZWnMN74IiKwgBEgQGDBwgGiEgSQK1Wy6XrzzpRUSz2wsxVsDh5vfLrpytpIczSYPOgAIiKwgBEgQKIBwgGiEgF0RIWS9LLLCG8HCr6ZFR0ye72LgQHgCOyZY/ViavqbQK0wEK0AEKA2liYxIgmynSRECR0yOhKgOMfUyJJPhlPUA8hz+XLmkskcr/hDQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyIlCAESIQEnssIok+Atq0LVaSnmcE7wnNM2RsSHOJ6LH++HW5h6dSIlCAESIQHWh6c3n369c+m6hB5/PHNc5x5CXz7DhZ0+Y+GEB4CpBCInCAESAQEaIGnwKIkp/LcnctdNSlJgxCiIV9a0KxIlpZYtxybthsge",
	"proofHeight": {
	  "revisionNumber": "1",
	  "revisionHeight": "15"
	},
	"signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",

    "channelId": "channel-0",
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_TRY_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_TRY_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CHANNEL_OPEN_TRY_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgChannelOpenTry.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelOpenTry)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"))
		})
	})
})
