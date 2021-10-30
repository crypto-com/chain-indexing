package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCChannelCloseInit", func() {
		It("should parse Msg commands when there is MsgChannelCloseInit in the transaction", func() {
			expected := `
			{
				"name": "MsgChannelCloseInitCreated",
				"version": 1,
				"height": 23,
				"uuid": "{UUID}",
				"msgName": "MsgChannelCloseInit",
				"txHash": "3E491C5B2404FC3C3C15A59630729355ED398BFB617DA7D61C1B548E98955F8D",
				"msgIndex": 0,
				"params": {
					"portId": "transfer",
					"channelId": "channel-1",
					"signer": "cro1t7yk3d4meeaqf5zfegv8p94wlfhpcnsftz55f7",
					"counterpartyPortId": "transfer",
					"counterpartyChannelId": "channel-0",
					"connectionId": "connection-1"
				}
			}`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_CLOSE_INIT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_CLOSE_INIT_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgIBCChannelCloseInit"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelCloseInit)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
		})
	})
})
