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
	Describe("MsgIBCChannelCloseConfirm", func() {
		It("should parse Msg commands when there is MsgChannelCloseConfirm in the transaction", func() {
			expected := `
			{
				"name": "/ibc.core.channel.v1.MsgChannelCloseConfirm.Created",
				"version": 1,
				"height": 27,
				"uuid": "{UUID}",
				"msgName": "/ibc.core.channel.v1.MsgChannelCloseConfirm",
				"txHash": "A99265922CA5F9F2E2F647B822F37AA8845E3EFBC96F4378D9CA89DDF2BB9ECD",
				"msgIndex": 1,
				"params": {
					"portId": "transfer",
					"channelId": "channel-0",
					"proofInit": "Cv8CCvwCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTESMggEEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0xKgdpY3MyMC0xGgsIARgBIAEqAwACLiIpCAESJQIELiDoa/QgHw9tDnDiZlToyGCCtvdk/YcdNOtMNk4s7TESBiAiKwgBEgQEBi4gGiEgWkRCIWiKNqPdGeeaBlk1tUshQ314yl/2d0Dht7MZKTIiKwgBEgQGCi4gGiEgQZepEjXiUZCYCObv6YevxMpCE3NytECo0qCvAqMefj8iKwgBEgQIEi4gGiEgiNI7Cxb9YFfTvRC3QICXYJBUMG2Hd0PzF8UuHKL0G+AiKwgBEgQKLC4gGiEggTFR60ori9MUc86XOgPlQ8yINywZKHUcrwL1nHOCrFUiKwgBEgQMUi4gGiEg63yd2j19EPXHXJKMLbS35ceHlpLDlCK06cRIc9vxPoUK/gEK+wEKA2liYxIg9LWYWuKyPj5DLqVTHwqfaHiGnaYPpz14ploHwPgk8iQaCQgBGAEgASoBACIlCAESIQFGsg9bKg+Td5KCS6A+Qnp6eIVOcOR7NdKpisvlOJb5EyInCAESAQEaIPhZvf/MWZkK7MygO8J7NUbiwp6dqQN5Gf3gt6vxV489IicIARIBARogCD1iqt4WlG2/Sexg8bzDgcffB858+xTniEdtcFPnq/EiJQgBEiEBxwWsbBJAhv48/qrs9osH7NubOZjpZ2vlN0I+bSv+YMkiJwgBEgEBGiCwFtEy00C8HUDC67XhEg1JqHWPRCvNgukGVPcaXwl61Q==",
					"proofHeight": { "revisionNumber": "2", "revisionHeight": "24" },
					"signer": "cro12cgecr4kmyylql6kerfpn7ff42weur7glq4uj3",
					"counterpartyPortId": "transfer",
					"counterpartyChannelId": "channel-1",
					"connectionId": "connection-0"
				}
			}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_CLOSE_CONFIRM_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_CLOSE_CONFIRM_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CHANNEL_CLOSE_CONFIRM_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgChannelCloseConfirm.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelCloseConfirm)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("cro12cgecr4kmyylql6kerfpn7ff42weur7glq4uj3"))
		})
	})
})
