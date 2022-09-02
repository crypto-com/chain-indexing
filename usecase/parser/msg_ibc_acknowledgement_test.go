package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCAcknowledgement", func() {
		It("should parse Msg commands when there is MsgIBCAcknowledgement in the transaction", func() {
			expected := `{
  "name": "/ibc.core.channel.v1.MsgAcknowledgement.Created",
  "version": 1,
  "height": 28,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgAcknowledgement",
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
    "messageType": "/ibc.applications.transfer.v1.MsgTransfer",
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

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgAcknowledgement.Create"))

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
			Expect(possibleSignerAddresses[1]).To(Equal("cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"))
		})

		It("should parse Msg commands when there is MsgIBCAcknowledgement in the transaction, and MsgIBCAcknowledgement.maybeMsgTransfer.Error is NOT empty string", func() {
			expected := `{
	"name": "/ibc.core.channel.v1.MsgAcknowledgement.Created",
	"version": 1,
	"height": 127749,
	"uuid": "{UUID}",
	"msgName": "/ibc.core.channel.v1.MsgAcknowledgement",
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
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
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

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_ERROR_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_ERROR_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_ERROR_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgAcknowledgement.Create"))

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
			Expect(possibleSignerAddresses[1]).To(Equal("tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6"))
		})

		It("should return two MsgIBCAcknowledgement with different success values when there are two MsgAcknowledgement of same packet sequence in the block, and in the second messages the fungible_token_packet event is missing", func() {
			expectedFirstMsgAck := `{
	"name": "/ibc.core.channel.v1.MsgAcknowledgement.Created",
	"version": 1,
	"height": 68821,
	"uuid": "{UUID}",
	"msgName": "/ibc.core.channel.v1.MsgAcknowledgement",
	"txHash": "0F1205B05DDDC0B848AD50FE35335C993FEDE7EAD161F2099818B0A03321B97F",
	"msgIndex": 0,
	"params": {
		"packet": {
			"sequence": "933",
			"sourcePort": "transfer",
			"sourceChannel": "channel-0",
			"destinationPort": "transfer",
			"destinationChannel": "channel-44",
			"data": "eyJhbW91bnQiOiI5OTc5Mzk2NzUwMCIsImRlbm9tIjoidHJhbnNmZXIvY2hhbm5lbC0wL2Jhc2Vjcm8iLCJyZWNlaXZlciI6ImNybzFkOWozOHdka2Rwd2Y2MjN1emw3czlucXcwOHVwMzBjd2t2MmNtbCIsInNlbmRlciI6ImNyYzF6ZGFlcHE5d3NraGY1bDV2MGoyd20waHI0a3VjanV3NXllbDI1dSJ9",
			"timeoutHeight": {
				"revisionNumber": "0",
				"revisionHeight": "0"
			},
			"timeoutTimestamp": "1636802327974618671"
		},
		"acknowledgement": "eyJyZXN1bHQiOiJBUT09In0=",
		"proofAcked": "CvIHCu8HCjVhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtNDQvc2VxdWVuY2VzLzkzMxIgCPdVftUYJv4Y2EUSvyTsdQAe268hI6R333KgqfNkCnwaDggBGAEgASoGAAKcxYQDIi4IARIHAgScxYQDIBohIDwQccgiMyrjLbcO8vPDN9d3axcaYEUa1R5jTbc9QrEUIiwIARIoBAicxYQDIJGWtfcKpISRSjNnTDmj7yfu2sYN7ylgVs8CsQmS+vZoICIsCAESKAYQnMWEAyDqIfPb8+Y9G25V4zTBgfYqEJpDH+14RXV3JjrTW+df2CAiLggBEgcIHJzFhAMgGiEgEH2fJo7kkzfUAdGGInk3u8gQSzsa+NQrQrmn5bWewbsiLAgBEigKLpzFhAMgxvf+3tGUbTd2hmvjvEipzpRg9PAN/Kx4HnX7qa/QGs8gIiwIARIoDFycxYQDIAkoKqt5WEDBrKES7vGpFkzWy4hi3L4VwZWlVFwrC1e8ICItCAESKQ7EAZzFhAMg9FqcFtOC9laytmdPUwKyyNoiwVuctCR6zxHQMR6jXl0gIi0IARIpEM4CnMWEAyAPR1g6VGLnyuFCnjd0Qj/JwbSQeKfVAovu+orlmkp3kyAiLQgBEikS1AOcxYQDILefyRWK+YSCuiVB3I2wwMDqVp16Sgyzz7ZmN75Lv3MMICItCAESKRTqB5zFhAMgDug2QuDrvM0oN2IDF1XJKr1XjxykD3SJR3/pVeDiIVwgIi0IARIpFooQnMWEAyCf5HTrSGpUSnJt9shB6Fo7yc/jM7XJUNRBNIRmypMelCAiLQgBEikY3BycxYQDIKFRwO1Q3eK2suvY0jiTVX6wKO7+ET6k3hC+e5+8pIBAICItCAESKRq+NJzFhAMgt8C/50Q0KYN5cW5MvvbnIS30/SLIyr+FO+/1ndgeiMcgIi0IARIpHOhfnMWEAyAHXT2SdgE8GcynTSUMRHBSTlvhiCGZjVuiTc55LvKI3SAiLggBEioepqUBnMWEAyDDwDasfVEFnpB16dEB6oTx3NuiDde72z8+wrGFgIEgXCAiLggBEioglsYCnMWEAyB9ye9kzIq5x74bFoFth70N2lCRzVyE7nzMN+h/ZRLDNyAiMAgBEgki4J4EnMWEAyAaISAKSYySjL8NwUPrqNv1B6piWwT2UGi0hO/RU9FC4hA/dyIwCAESCSSMlgicxYQDIBohIIHy7B0D/O1whzgbJ5YyCqC9BTiVEh4OQPdVpEwEGV65IjAIARIJKKLME5zFhAMgGiEgaxhG9Cvrw3YI2vEuL5aX7dF0gHsG+BrIbIzCXuddjg0K/AEK+QEKA2liYxIg+YaXZdDxHjzJJk0pV2Gpd0ZC96RM5nwKb+NMKozNbNMaCQgBGAEgASoBACIlCAESIQHGFYo+eJC7BDuwbQgF9JWTxmomsU5OToBz+rgm3rPoyiIlCAESIQFNau0HEtyE0UZQMBEACzeDPjN7NGmlnudP7Bj5/JnIVSIlCAESIQGwRJ6T3arUBmLdqgJBi0plcbY2gFLfK4mgRtbDp5+0/SInCAESAQEaIFzyRJ+I9dxVu3tUnJky7PZCHnVdfz2qDyAyPf+Oa1CBIicIARIBARog0gxBfMi3VGo2ZD71VCYvXMTGeOt/tSI6qNkOvMmZtdA=",
		"proofHeight": {
			"revisionNumber": "1",
			"revisionHeight": "3182927"
		},
		"signer": "crc1yzl6cnq3f66ew24d7u97vmp45nkckhwg4ak8hl",
		"application": "transfer",
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
		"maybeMsgTransfer": {
			"sender": "crc1zdaepq9wskhf5l5v0j2wm0hr4kucjuw5yel25u",
			"receiver": "cro1d9j38wdkdpwf623uzl7s9nqw08up30cwkv2cml",
			"denom": "transfer/channel-0/basecro",
			"amount": "99793967500",
			"success": true,
			"acknowledgement": "{ACKNOWLEDGEMENT}",
			"error": null
		},
		"packetSequence": "933",
		"channelOrdering": "ORDER_UNORDERED",
		"connectionId": "connection-0"
	}		
}
`

			expectedSecondMsgAck := `{
	"name": "MsgAlreadyRelayedAcknowledgement.Created",
	"version": 1,
	"height": 68821,
	"uuid": "{UUID}",
	"msgName": "MsgAlreadyRelayedAcknowledgement",
	"txHash": "961A752199E991F5653FB8773DF6A20F1381C3C7B3FAD2D817988BE6640FCF4E",
	"msgIndex": 1,
	"params": {
		"packet": {
			"sequence": "933",
			"sourcePort": "transfer",
			"sourceChannel": "channel-0",
			"destinationPort": "transfer",
			"destinationChannel": "channel-44",
			"data": "eyJhbW91bnQiOiI5OTc5Mzk2NzUwMCIsImRlbm9tIjoidHJhbnNmZXIvY2hhbm5lbC0wL2Jhc2Vjcm8iLCJyZWNlaXZlciI6ImNybzFkOWozOHdka2Rwd2Y2MjN1emw3czlucXcwOHVwMzBjd2t2MmNtbCIsInNlbmRlciI6ImNyYzF6ZGFlcHE5d3NraGY1bDV2MGoyd20waHI0a3VjanV3NXllbDI1dSJ9",
			"timeoutHeight": {
				"revisionNumber": "0",
				"revisionHeight": "0"
			},
			"timeoutTimestamp": "1636802327974618671"
		},
		"acknowledgement": "eyJyZXN1bHQiOiJBUT09In0=",
		"proofAcked": "CvIHCu8HCjVhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtNDQvc2VxdWVuY2VzLzkzMxIgCPdVftUYJv4Y2EUSvyTsdQAe268hI6R333KgqfNkCnwaDggBGAEgASoGAAKcxYQDIi4IARIHAgScxYQDIBohIDwQccgiMyrjLbcO8vPDN9d3axcaYEUa1R5jTbc9QrEUIiwIARIoBAicxYQDIJGWtfcKpISRSjNnTDmj7yfu2sYN7ylgVs8CsQmS+vZoICIsCAESKAYQnMWEAyDqIfPb8+Y9G25V4zTBgfYqEJpDH+14RXV3JjrTW+df2CAiLggBEgcIHJzFhAMgGiEgEH2fJo7kkzfUAdGGInk3u8gQSzsa+NQrQrmn5bWewbsiLAgBEigKLpzFhAMgxvf+3tGUbTd2hmvjvEipzpRg9PAN/Kx4HnX7qa/QGs8gIiwIARIoDFycxYQDIAkoKqt5WEDBrKES7vGpFkzWy4hi3L4VwZWlVFwrC1e8ICItCAESKQ7EAZzFhAMg9FqcFtOC9laytmdPUwKyyNoiwVuctCR6zxHQMR6jXl0gIi0IARIpEM4CnMWEAyAPR1g6VGLnyuFCnjd0Qj/JwbSQeKfVAovu+orlmkp3kyAiLQgBEikS1AOcxYQDILefyRWK+YSCuiVB3I2wwMDqVp16Sgyzz7ZmN75Lv3MMICItCAESKRTqB5zFhAMgDug2QuDrvM0oN2IDF1XJKr1XjxykD3SJR3/pVeDiIVwgIi0IARIpFooQnMWEAyCf5HTrSGpUSnJt9shB6Fo7yc/jM7XJUNRBNIRmypMelCAiLQgBEikY3BycxYQDIKFRwO1Q3eK2suvY0jiTVX6wKO7+ET6k3hC+e5+8pIBAICItCAESKRq+NJzFhAMgt8C/50Q0KYN5cW5MvvbnIS30/SLIyr+FO+/1ndgeiMcgIi0IARIpHOhfnMWEAyAHXT2SdgE8GcynTSUMRHBSTlvhiCGZjVuiTc55LvKI3SAiLggBEioepqUBnMWEAyDDwDasfVEFnpB16dEB6oTx3NuiDde72z8+wrGFgIEgXCAiLggBEioglsYCnMWEAyB9ye9kzIq5x74bFoFth70N2lCRzVyE7nzMN+h/ZRLDNyAiMAgBEgki4J4EnMWEAyAaISAKSYySjL8NwUPrqNv1B6piWwT2UGi0hO/RU9FC4hA/dyIwCAESCSSMlgicxYQDIBohIIHy7B0D/O1whzgbJ5YyCqC9BTiVEh4OQPdVpEwEGV65IjAIARIJKKLME5zFhAMgGiEgaxhG9Cvrw3YI2vEuL5aX7dF0gHsG+BrIbIzCXuddjg0K/AEK+QEKA2liYxIg+YaXZdDxHjzJJk0pV2Gpd0ZC96RM5nwKb+NMKozNbNMaCQgBGAEgASoBACIlCAESIQHGFYo+eJC7BDuwbQgF9JWTxmomsU5OToBz+rgm3rPoyiIlCAESIQFNau0HEtyE0UZQMBEACzeDPjN7NGmlnudP7Bj5/JnIVSIlCAESIQGwRJ6T3arUBmLdqgJBi0plcbY2gFLfK4mgRtbDp5+0/SInCAESAQEaIFzyRJ+I9dxVu3tUnJky7PZCHnVdfz2qDyAyPf+Oa1CBIicIARIBARog0gxBfMi3VGo2ZD71VCYvXMTGeOt/tSI6qNkOvMmZtdA=",
		"proofHeight": {
			"revisionNumber": "1",
			"revisionHeight": "3182927"
		},
		"signer": "crc1aaxs058pksrq8cx3k0nrxv60p2a9c7nq527949",
		"application": "transfer",
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
   		"maybeMsgTransfer": {
			"sender": "crc1zdaepq9wskhf5l5v0j2wm0hr4kucjuw5yel25u",
			"receiver": "cro1d9j38wdkdpwf623uzl7s9nqw08up30cwkv2cml",
			"denom": "transfer/channel-0/basecro",
			"amount": "99793967500",
			"acknowledgement": "",
			"error": null
		},
		"packetSequence": "933",
		"channelOrdering": "ORDER_UNORDERED",
		"connectionId": "connection-0"
	}		
}
`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_BLOCK_RESULTS_RESP,
			))

			tx1 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_1)
			tx2 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_2)
			tx3 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_3)
			tx4 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_4)
			tx5 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_5)
			tx6 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_6)
			tx7 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_7)
			tx8 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_8)
			tx9 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_9)
			tx10 := mustParseTxsResp(usecase_parser_test.TX_MSG_ACKNOWLEDGEMENT_DUPLICATE_PACKET_SEQUENCE_TXS_RESP_10)
			txs := []model.Tx{*tx1, *tx2, *tx3, *tx4, *tx5, *tx6, *tx7, *tx8, *tx9, *tx10}

			accountAddressPrefix := "crc"
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
			Expect(cmds).To(HaveLen(11))
			Expect(cmds[6].Name()).To(Equal("/ibc.core.client.v1.MsgUpdateClient.Create"))
			Expect(cmds[7].Name()).To(Equal("/ibc.core.channel.v1.MsgAcknowledgement.Create"))
			Expect(cmds[8].Name()).To(Equal("/ibc.core.client.v1.MsgUpdateClient.Create"))
			Expect(cmds[9].Name()).To(Equal("CreateMsgAlreadyRelayedIBCAcknowledgement"))

			firstMsgAckCmd := cmds[7]
			secondMsgAckCmd := cmds[9]

			regex, _ := regexp.Compile("\n?\r?\\s?")

			untypedFirstMsgAckEvent, _ := firstMsgAckCmd.Exec()
			typedFistMsgAckEvent := untypedFirstMsgAckEvent.(*event.MsgIBCAcknowledgement)
			expectedFirstMsgAckWithUUID := strings.Replace(
				regex.ReplaceAllString(expectedFirstMsgAck, ""),
				"{UUID}",
				typedFistMsgAckEvent.UUID(),
				-1,
			)
			expectedFirstMsgAckWithAcknowledgement := strings.Replace(
				expectedFirstMsgAckWithUUID,
				"{ACKNOWLEDGEMENT}",
				"result:\\\"\\\\001\\\" ",
				-1,
			)
			Expect(json.MustMarshalToString(typedFistMsgAckEvent)).To(Equal(expectedFirstMsgAckWithAcknowledgement))

			untypedSecondMsgAckEvent, _ := secondMsgAckCmd.Exec()
			typedSecondMsgAckCmd := untypedSecondMsgAckEvent.(*event.MsgAlreadyRelayedIBCAcknowledgement)
			expectedSecondMsgAckWithUUID := strings.Replace(
				regex.ReplaceAllString(expectedSecondMsgAck, ""),
				"{UUID}",
				typedSecondMsgAckCmd.UUID(),
				-1,
			)
			Expect(json.MustMarshalToString(typedSecondMsgAckCmd)).To(Equal(expectedSecondMsgAckWithUUID))
			Expect(possibleSignerAddresses[0]).To(Equal("crc1yzl6cnq3f66ew24d7u97vmp45nkckhwg4ak8hl"))
			Expect(possibleSignerAddresses[2]).To(Equal("crc1aaxs058pksrq8cx3k0nrxv60p2a9c7nq527949"))
		})
	})

})
