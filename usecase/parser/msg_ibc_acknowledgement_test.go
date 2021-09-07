package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser/test"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/internal/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
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
			"error": null
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
			cmds, err := ParseBlockResultsTxsMsgToCommands(
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

		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction, and MsgIBCRecvPacket.maybeMsgTransfer.Error is NOT empty string", func() {
			expected := `{
	"name": "MsgAcknowledgementCreated",
	"version": 1,
	"height": 127749,
	"uuid": "{UUID}",
	"msgName": "MsgAcknowledgement",
	"txHash": "AA4C3C2A27587AA769D9198BEBB42218419D61C46F1A2C9CE7BC10FE45F81A39",
	"msgIndex": 1,
	"params": {
		"packet": {
			"sequence": "10",
			"sourcePort": "transfer",
			"sourceChannel": "channel-3",
			"destinationPort": "transfer",
			"destinationChannel": "channel-0",
			"data": "eyJhbW91bnQiOiIxMDAwMDAwMDAiLCJkZW5vbSI6ImJhc2V0Y3JvIiwicmVjZWl2ZXIiOiIweDgwNjM5QTFGRTk0NEVDNTBDQjZBOUZCM0U5NEJFRjNCRTkwQTRCOEYiLCJzZW5kZXIiOiJ0Y3JvMXBkbjJuc245d2VzejZweDNsY2pzZ21wOHBlZmVkbnpwM2dtcDNxIn0=",
			"timeoutHeight": {
				"revisionNumber": "1",
				"revisionHeight": "16794"
			},
			"timeoutTimestamp": "1630030423261159012"
		},
		"acknowledgement": "eyJlcnJvciI6ImRlY29kaW5nIGJlY2gzMiBmYWlsZWQ6IHN0cmluZyBub3QgYWxsIGxvd2VyY2FzZSBvciBhbGwgdXBwZXJjYXNlIn0=",
		"proofAcked": "CrQDCrEDCjNhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtMC9zZXF1ZW5jZXMvMTASILHzro1BUrD+WXPMFD241fEm3zYnORy3SWwRUN82kMvGGg0IARgBIAEqBQACyPcBIi0IARIGAgTI9wEgGiEgyov7ua0WGcxfyDsT/xC0SyAcyvxUvbMnbo9937We3z8iLQgBEgYECMj3ASAaISDcXJBgYefzxoPVjmMHCXdiWBpZqydFKPI0YYOF9EnqCCItCAESBgYQyPcBIBohIEgflNLnHbMX64gvXsi0rQrEPA77OzfaRRN9D9+6YJkEIi0IARIGCCDI9wEgGiEgqmL/2zI6NqSbxK6kCLw9C7cJNvioQrPD8HAkp6g/OKoiLQgBEgYKQMj3ASAaISCYGFpX5+kJwG2Ygx6ShrN4GD1IaBFOrB+a8uvpZ9nHGSIuCAESBw6QAcj3ASAaISCsQaU5eHtrrRmBATcC7TInPFV6XSTYZNJnxsrjOBpzniIuCAESBxDUAcj3ASAaISDAeUvXFBaD/h6QP8IJ2SkmjMkeNVMYVZbxDOs9bVDumgr+AQr7AQoDaWJjEiB3S9Z2xt/ISObd2LE0TANDEteIykeGjjTO1RsTQmBBGBoJCAEYASABKgEAIicIARIBARog+VbefwQZr0EJzBl04fE3Iwq9K4y59Sd3XuzKGogXDyIiJQgBEiEBHhDMIutpz1N/dqAD7CKxN8Le8+xw92XPht9+FkWcIJgiJwgBEgEBGiD0VKCIon4Y8qH6ipROvGgCZpABP9UZsYnPNh1vPm1e4SIlCAESIQE+yovtgZYnFsGL3+iyu2e4ltJTc58nmZm3c/3URAx7IyInCAESAQEaIHVgcAqfN4KuqeY9pK7xJf8ARSjBPNtY2ZeCttQCEJOO",
		"proofHeight": {
			"revisionNumber": "1",
			"revisionHeight": "15845"
		},
		"signer": "tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6",
		"application": "transfer",
		"messageType": "MsgTransfer",
		"maybeMsgTransfer": {
			"sender": "tcro1pdn2nsn9wesz6px3lcjsgmp8pefednzp3gmp3q",
			"receiver": "0x80639A1FE944EC50CB6A9FB3E94BEF3BE90A4B8F",
			"denom": "basetcro",
			"amount": "100000000",
			"success": false,
			"acknowledgement": "{ACKNOWLEDGEMENT}",
			"error": "{ERROR}"
		},
		"packetSequence": "10",
		"channelOrdering": "ORDER_UNORDERED",
		"connectionId": "connection-3"
	}		
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_ERROR_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_ERROR_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"
			cmds, err := ParseBlockResultsTxsMsgToCommands(
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

			expectedWithUUID := strings.Replace(
				regex.ReplaceAllString(expected, ""),
				"{UUID}",
				typedEvent.UUID(),
				-1,
			)

			expectedWithError := strings.Replace(
				expectedWithUUID,
				"{ERROR}",
				"decoding bech32 failed: string not all lowercase or all uppercase",
				-1,
			)

			expectedWithAcknowledgement := strings.Replace(
				expectedWithError,
				"{ACKNOWLEDGEMENT}",
				"error:\\\"decoding bech32 failed: string not all lowercase or all uppercase\\\" ",
				-1,
			)

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(expectedWithAcknowledgement))
		})
	})
})
