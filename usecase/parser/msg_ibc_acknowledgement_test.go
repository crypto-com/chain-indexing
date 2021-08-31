package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCRecvPacket", func() {
		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction", func() {
			expected := `{
  "name": "MsgAcknowledgementCreated",
  "version": 1,
  "height": 28,
  "uuid": "{UUID}",
  "msgName": "MsgAcknowledgement",
  "txHash": "06B6CEE9FB786050A41BA57026B3FF629188C06FCDD97F59AF55D2CD40938CD3",
  "msgIndex": 1,
  "params": {
    "packet": {
      "sequence": "1",
      "sourcePort": "transfer",
      "sourceChannel": "channel-0",
      "destinationPort": "transfer",
      "destinationChannel": "channel-0",
      "data": "eyJhbW91bnQiOiIxMjM0IiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjcm8xZHVsd3FnY2RwZW1uOGMzNHNqZDkyZnhlcHo1cDBzcXBlZXZ3N2YiLCJzZW5kZXIiOiJjcm8xMHNuaGx2a3B1YzR4aHE4MnV5ZzVleDJlZXptbWY1ZWQ1dG1xc3YifQ==",
      "timeoutHeight": {
    	"revisionNumber": "2",
    	"revisionHeight": "1023"
      },
      "timeoutTimestamp": "0"
    },
	"acknowledgement": "eyJyZXN1bHQiOiJBUT09In0=",
    "proofAcked": "CscCCsQCCjJhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtMC9zZXF1ZW5jZXMvMRIgCPdVftUYJv4Y2EUSvyTsdQAe268hI6R333KgqfNkCnwaCwgBGAEgASoDAAI0IisIARIEAgQ0IBohICsdSIdU+v81lhDQ/dswq/j+y5pEx0RcBqDNNvGBURKNIisIARIEBAg0IBohIO5n1rXg/PKXipcxX5Uy+XyTZ3jGwpIB/RDNehpbIU0jIisIARIEBhA0IBohIIymA52HzmLEaKECJT3n3bOHLJhBuDBeQclE+lp93vOyIisIARIECBo0IBohIHQvJirT50kRamjMV0YCUFCu6329YWyKcegokMjXHxBQIisIARIECjI0IBohIHRaqiQwujUalYQJUdxvggy1gNyGKCYldtX1dq+fAjtwCtMBCtABCgNpYmMSIHdgqTmy3PkcFIK3GAEKyMa1kQJniPJbk/Hl3o6gyztLGgkIARgBIAEqAQAiJQgBEiEBRrIPWyoPk3eSgkugPkJ6eniFTnDkezXSqYrL5TiW+RMiJQgBEiEBb6XKyCWaqTgyGMhPJZ7TKQAjel7eQsEcAP3XGSsSBg4iJQgBEiEBpzj0MzNAMsi156BdbnVm9mYPLOVKKnxZmEfyj3pKw5kiJwgBEgEBGiA7xilGseIAH2wYkx1DUIexthQg/0UX9+RjCYJ6M0kL5Q==",
    "proofHeight": {
      "revisionNumber": "2",
      "revisionHeight": "27"
    },
    "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",

    "application": "transfer",
    "messageType": "MsgTransfer",
    "maybeMsgTransfer": {
      "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
      "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
      "denom": "basecro",
      "amount": "1234",
      "success": true,
      "acknowledgement": "{0xc0038ae7a0}",
			"error": ""
    },

    "packetSequence": "1",
    "channelOrdering": "ORDER_UNORDERED",
    "connectionId": "connection-0"
  }
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"
			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			cmd := cmds[1]
			Expect(cmd.Name()).To(Equal("CreateMsgIBCAcknowledgement"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCAcknowledgement)

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
