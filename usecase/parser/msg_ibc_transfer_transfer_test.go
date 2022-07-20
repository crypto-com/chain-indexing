package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCChannelOpenConfirm", func() {
		It("should parse Msg commands when there is MsgTransfer in the transaction", func() {
			expected := `{
  "name": "MsgTransferCreated",
  "version": 1,
  "height": 24,
  "uuid": "{UUID}",
  "msgName": "MsgTransfer",
  "txHash": "7A5D86E0B1A364106EE2F1B40431B15A8E1B6C4A2E09E831AB773A12F5F5A564",
  "msgIndex": 0,
  "params": {
    "sourcePort": "transfer",
    "sourceChannel": "channel-0",
    "token": {
      "denom": "basecro",
	  "amount": "1234"
    },
    "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
    "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
    "timeoutHeight": {
      "revisionNumber": "2",
      "revisionHeight": "1023"
    },
    "timeoutTimestamp": "0",

    "packetSequence": "1",
    "destinationPort": "transfer",
    "destinationChannel": "channel-0",
    "channelOrdering": "ORDER_UNORDERED",
    "connectionId": "connection-0",
    "packetData": {
      "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
      "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
      "denom": "basecro",
      "amount": "1234"
    }
  }
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_TRANSFER_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_TRANSFER_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddress, err := parser.ParseBlockTxsMsgToCommands(
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
			Expect(cmd.Name()).To(Equal("CreateMsgIBCTransferTransfer"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCTransferTransfer)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddress).To(Equal([]string{"cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"}))
		})

		It("should parse Msg commands when there is MsgTransfer with exceed uint64 amount in the transaction", func() {
			expected := `{
  "name": "MsgTransferFailed",
  "version": 1,
  "height": 140108,
  "uuid": "{UUID}",
  "msgName": "MsgTransfer",
  "txHash": "D924F6E1A16ACDFFBF0B5BFDECC8E010E8F8D746B379FFC63D477C472B4128B7",
  "msgIndex": 0,
  "params": {
    "sourcePort": "transfer",
    "sourceChannel": "channel-0",
    "token": {
      "denom": "basetcro",
	  "amount": "28836836236828398900"
    },
    "sender": "tcrc1v76r7u4uyr3ewdks8cqmuw7ca4lejvc89pxhev",
    "receiver": "tcro1558rfl8pnpk90rmryvy3ep03yslq4ar0lkvdnn",
    "timeoutHeight": {
      "revisionNumber": "4",
      "revisionHeight": "1042864"
    },
    "timeoutTimestamp": "1634894674620940059",

    "packetSequence": "0",
    "destinationPort": "",
    "destinationChannel": "",
    "channelOrdering": "",
    "connectionId": "",
    "packetData": {
      "sender": "",
      "receiver": "",
      "denom": "",
      "amount": null
    }
  }
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_TRANSFER_STRING_AMOUNT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_TRANSFER_STRING_AMOUNT_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddress, err := parser.ParseBlockTxsMsgToCommands(
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
			Expect(cmd.Name()).To(Equal("CreateMsgIBCTransferTransfer"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCTransferTransfer)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddress).To(Equal([]string{"tcrc1v76r7u4uyr3ewdks8cqmuw7ca4lejvc89pxhev"}))
		})
	})
})
