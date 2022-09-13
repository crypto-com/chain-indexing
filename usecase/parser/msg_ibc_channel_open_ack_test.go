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
	Describe("MsgIBCChannelOpenAck", func() {
		It("should parse Msg commands when there is MsgChannelOpenAck in the transaction", func() {
			expected := `{
  "name": "/ibc.core.channel.v1.MsgChannelOpenAck.Created",
  "version": 1,
  "height": 18,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgChannelOpenAck",
  "txHash": "E66CCD50C9946AC5DC2AEA6C332ADFCF6FEF6229DBC6AC793158DF7CC9C2CD16",
  "msgIndex": 1,
  "params": {
	"portId": "transfer",
	"channelId": "channel-0",
	"counterpartyChannelId": "channel-0",
	"counterpartyVersion": "ics20-1",
	"proofTry": "CqcCCqQCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASMggCEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0wKgdpY3MyMC0xGgsIARgBIAEqAwACICIrCAESBAIEICAaISCKn2oMLG+m7x551nh06PSUqF3nF6zuDAUZzHWN3Y7R6iIrCAESBAQGICAaISAl+qVfRZQA8J2FNbAhThANLKQJekTVQlahuZE83dRMGyIrCAESBAgQICAaISAnLE8YAY/vNnafNqIgxC5dnti4svEmGWT97i5E1czeVSIrCAESBAokICAaISD6l2Xp3knzZRXOdJhJ8ToqY0qLt6QgDUruuoeIyeXg6wrTAQrQAQoDaWJjEiDpdjgQlYa7NrP+ylMayPUCN/n48pTXiof8JywRXiIjxxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAaj29iVS1fIRuqzuZBN/XoI11+IAJTrJPqBHC1ZQFHsbIiUIARIhAdmJCKPzDC4kc5/SoZhd1InuuJbVmFMLbq5ugt5dhByhIicIARIBARogEqy+GTuqKMTr9SDspw2rkC/9N0sLQcUcY4c9336zn3M=",
	"proofHeight": {
	  "revisionNumber": "2",
	  "revisionHeight": "17"
	},
	"signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",

    "counterpartyPortId": "transfer",
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_ACK_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_ACK_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CHANNEL_OPEN_ACK_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgChannelOpenAck.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelOpenAck)

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
