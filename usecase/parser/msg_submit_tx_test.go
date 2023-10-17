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
	Describe("MsgSubmitTx", func() {
		It("should parse Msg commands when there is MsgSubmitTx in the transaction", func() {
			expected := `{
  "name": "/chainmain.icaauth.v1.MsgSubmitTx.Created",
  "version": 1,
  "height": 67975,
  "uuid": "{UUID}",
  "msgName": "/chainmain.icaauth.v1.MsgSubmitTx",
  "txHash": "A4AB8B9882379DC79EA19FD1CFCF53C1A8A252ECA557F39E4B316164DD27CF33",
  "msgIndex": 0,
  "params": {
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"connectionId": "connection-18",
	"msgs": [
	  {
		"@type": "/cosmos.bank.v1beta1.MsgSend"
	  }
	],
	"timeoutDuration": "300s",
	"sequence": "1",
	"sourcePort": "icacontroller-tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
	"sourceChannel": "channel-48",
	"destinationPort": "icahost",
	"destinationChannel": "channel-1",
	"data": "{\"data\":\"CsgBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEqcBCit0Y3JvMW05MDJ3bmozOHE3NmR4ZjQyNXRjdjVnemU5cmdlZWh0NjdlbGRmEit0Y3JvMTQ0anRtN3J4OHk3d2p3aDh0dWZla2U4bjI2ZHY0cGR4cDB6bW0wGksKRGliYy82QjVBNjY0QkYwQUY0RjcxQjJGMEJBQTMzMTQxRTJGMTMyMTI0MkZCRDVEMTk3NjJGNTQxRUM5NzFBQ0IwODY1EgMxMDA=\",\"memo\":\"\",\"type\":\"TYPE_EXECUTE_TX\"}",
	"timeoutHeight": {
	  "revisionNumber": "0",
      "revisionHeight": "0"
	},
	"timeoutTimestamp": "1668771093164765856"
  }
}
`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_SUBMIT_TX_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_SUBMIT_TX_BLOCK_RESULTS_RESP,
			), &tendermint.Base64BlockResultEventAttributeDecoder{})

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_SUBMIT_TX_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

			accountAddressPrefix := "tcro"
			stakingDenom := "basetcro"

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
			Expect(cmd.Name()).To(Equal("/chainmain.icaauth.v1.MsgSubmitTx.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgSubmitTx)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[0]).To(Equal("tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra"))
		})
	})
})
