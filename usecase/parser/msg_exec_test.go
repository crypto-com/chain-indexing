package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/external/logger/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	XDescribe("MsgExec", func() {
		It("should parse Msg commands when there is MsgExec (inner message MsgSend) in the transaction", func() {
			expected := `{
            "name": "MsgExecCreated",
            "version": 1,
            "height": 113382,
            "uuid": "{UUID}",
            "msgName": "MsgExec",
            "txHash": "0CE949FAB0CB8EFB6E80F8ED785A6313FE7C094C336D4A7E8630E7D81AECD946",
            "msgIndex": 0,
            "params": {
                "@type": "/cosmos.authz.v1beta1.MsgExec",
                "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
                "msgs": [
                    {
                        "@type": "/cosmos.bank.v1beta1.MsgSend"
                    }
                ]
            }
        }`

			expectedInnerMsg := `{
				"name": "MsgSendCreated",
				"version": 1,
				"height": 113382,
				"uuid": "{UUID}",
				"msgName": "MsgSend",
				"txHash": "0CE949FAB0CB8EFB6E80F8ED785A6313FE7C094C336D4A7E8630E7D81AECD946",
				"msgIndex": 0,
				"fromAddress": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
				"toAddress": "tcro1a93yfnsc3x7m0m445cjsvee2n7qz9c0purlzwq",
				"amount": [
					{
						"denom": "basetcro",
						"amount": "100000000"
					}
				]
			}`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_MSG_SEND_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_MSG_SEND_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()
			logger := test.NewFakeLogger()

			cmds, _, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				logger,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))

			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgExec"))

			untypedEvent, _ := cmd.Exec()
			createMsgExecEvent := untypedEvent.(*event.MsgExec)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgExecEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgExecEvent.UUID(),
					-1,
				),
			))

			innerCmd := cmds[1]
			Expect(innerCmd.Name()).To(Equal("CreateMsgSend"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgSendEvent := untypedInnerEvent.(*event.MsgSend)

			Expect(json.MustMarshalToString(createMsgSendEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg, ""),
					"{UUID}",
					createMsgSendEvent.UUID(),
					-1,
				),
			))

		})

		It("should parse Msg commands when there is MsgExec (inner message MsgDelegate) in the transaction", func() {
			expected := `{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 170493,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "AB1D25567EF5FC054375442B0B01728BA333972E685047A5C204DA4DC4A7324A",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
					"msgs": [{ "@type": "/cosmos.staking.v1beta1.MsgDelegate" }]
				}
			}`

			expectedInnerMsg := `{
				"name": "MsgDelegateCreated",
				"version": 1,
				"height": 170493,
				"uuid": "{UUID}",
				"msgName": "MsgDelegate",
				"txHash": "AB1D25567EF5FC054375442B0B01728BA333972E685047A5C204DA4DC4A7324A",
				"msgIndex": 0,
				"delegatorAddress": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
				"validatorAddress": "tcrocncl163tv59yzgeqcap8lrsa2r4zk580h8ddr5a0sdd",
				"amount": { "denom": "basetcro", "amount": "100000000" },
				"autoClaimedRewards": { "denom": "basecro", "amount": "0" }
			}`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()
			logger := test.NewFakeLogger()

			cmds, _, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				logger,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))

			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgExec"))

			untypedEvent, _ := cmd.Exec()
			createMsgExecEvent := untypedEvent.(*event.MsgExec)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgExecEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgExecEvent.UUID(),
					-1,
				),
			))

			innerCmd := cmds[1]
			Expect(innerCmd.Name()).To(Equal("CreateMsgDelegate"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgDelegateEvent := untypedInnerEvent.(*event.MsgDelegate)

			Expect(json.MustMarshalToString(createMsgDelegateEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg, ""),
					"{UUID}",
					createMsgDelegateEvent.UUID(),
					-1,
				),
			))

		})
	})
})
