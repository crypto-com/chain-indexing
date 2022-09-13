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
  "name": "/ibc.core.channel.v1.MsgChannelOpenInit.Created",
  "version": 1,
  "height": 14,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgChannelOpenInit",
  "txHash": "310C7A5DE69CABB2175A2CBD417A2BFBA105C030941663F3F8809BBA2A6D810D",
  "msgIndex": 0,
  "params": {
	"portId": "transfer",
	"channel": {
	  "state": "STATE_INIT",
	  "ordering": "ORDER_UNORDERED",
	  "counterparty": {
		"portId": "transfer",
		"channelId": ""
	  },
	  "connectionHops": [
		"connection-0"
	  ],
	  "version": "ics20-1"
	},
	"signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",

    "channelId": "channel-0",
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_INIT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_INIT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CHANNEL_OPEN_INIT_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgChannelOpenInit.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelOpenInit)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[0]).To(Equal("cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"))
		})
	})
})
