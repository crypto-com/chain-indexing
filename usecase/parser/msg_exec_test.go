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
	usecase_parser_test_msg_exec "github.com/crypto-com/chain-indexing/usecase/parser/test/tx_msg_exec"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgExec", func() {
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
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_SEND_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_SEND_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
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
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2"}))
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
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_DELEGATE_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
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
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2"}))

		})

		It("should parse Msg commands when there is MsgExec (inner message MsgSetWithdrawAddress) in the transaction", func() {
			expected := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 292,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "2DFF941C55DCA1FCDBDB7F38F84C9C7810C1E9F4709682D0570964D2BB79CD82",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress"
						}
					]
				}
			}
			`

			expectedInnerMsg := []string{`
			{
				"name": "MsgSetWithdrawAddressCreated",
				"version": 1,
				"height": 292,
				"uuid": "{UUID}",
				"msgName": "MsgSetWithdrawAddress",
				"txHash": "2DFF941C55DCA1FCDBDB7F38F84C9C7810C1E9F4709682D0570964D2BB79CD82",
				"msgIndex": 0,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"withdrawAddress": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25"
			}`}

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
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
			Expect(innerCmd.Name()).To(Equal("CreateMsgSetWithdrawAddress"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgSetWithdrawAddressEvent := untypedInnerEvent.(*event.MsgSetWithdrawAddress)

			Expect(json.MustMarshalToString(createMsgSetWithdrawAddressEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[0], ""),
					"{UUID}",
					createMsgSetWithdrawAddressEvent.UUID(),
					-1,
				),
			))

			Expect(possibleSignerAddresses).To(Equal([]string{"cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25"}))

		})

		It("should parse Msg commands when there is MsgExec (inner message MsgWithdrawDelegatorReward) in the transaction", func() {
			expected := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 313,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "711F6ED876F2EDDF927E8330621D734B1661D2A26132CFF141C9399266564176",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
						},
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
						}
					]
				}
			}
			`

			expectedInnerMsg := []string{`
			{
				"name": "MsgWithdrawDelegatorRewardCreated",
				"version": 1,
				"height": 313,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawDelegatorReward",
				"txHash": "711F6ED876F2EDDF927E8330621D734B1661D2A26132CFF141C9399266564176",
				"msgIndex": 0,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"validatorAddress": "crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "362332"
					}
				]
			}`,
				`{
				"name": "MsgWithdrawDelegatorRewardCreated",
				"version": 1,
				"height": 313,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawDelegatorReward",
				"txHash": "711F6ED876F2EDDF927E8330621D734B1661D2A26132CFF141C9399266564176",
				"msgIndex": 0,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"validatorAddress": "crocncl17cye5gmealmpgz3tclfjga7urdjxj6nah3u8w2",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": []
			}`}

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(3))

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
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawDelegatorReward"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgWithdrawDelegatorRewardEvent := untypedInnerEvent.(*event.MsgWithdrawDelegatorReward)
			Expect(json.MustMarshalToString(createMsgWithdrawDelegatorRewardEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[0], ""),
					"{UUID}",
					createMsgWithdrawDelegatorRewardEvent.UUID(),
					-1,
				),
			))

			innerCmd = cmds[2]
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawDelegatorReward"))

			untypedInnerEvent, _ = innerCmd.Exec()
			createMsgWithdrawDelegatorRewardEvent = untypedInnerEvent.(*event.MsgWithdrawDelegatorReward)

			Expect(json.MustMarshalToString(createMsgWithdrawDelegatorRewardEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[1], ""),
					"{UUID}",
					createMsgWithdrawDelegatorRewardEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n"}))

		})

		It("should parse Msg commands when there is MsgExec (inner message MsgWithdrawValidatorCommission) in the transaction", func() {
			expected := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 5991,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "575DE51BFF24CEE75654C1BE8EFE30B01E9269E58267A27B53F5338457B9B870",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
						},
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission"
						}
					]
				}
			}
			`

			expectedInnerMsg := []string{`
			{
				"name": "MsgWithdrawDelegatorRewardCreated",
				"version": 1,
				"height": 5991,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawDelegatorReward",
				"txHash": "575DE51BFF24CEE75654C1BE8EFE30B01E9269E58267A27B53F5338457B9B870",
				"msgIndex": 0,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"validatorAddress": "crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "1285329"
					}
				]
			}`,
				`{
				"name": "MsgWithdrawValidatorCommissionCreated",
				"version": 1,
				"height": 5991,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawValidatorCommission",
				"txHash": "575DE51BFF24CEE75654C1BE8EFE30B01E9269E58267A27B53F5338457B9B870",
				"msgIndex": 0,
				"validatorAddress": "crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "142814"
					}
				]
			}`}

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(3))

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
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawDelegatorReward"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgWithdrawDelegatorRewardEvent := untypedInnerEvent.(*event.MsgWithdrawDelegatorReward)
			Expect(json.MustMarshalToString(createMsgWithdrawDelegatorRewardEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[0], ""),
					"{UUID}",
					createMsgWithdrawDelegatorRewardEvent.UUID(),
					-1,
				),
			))

			innerCmd = cmds[2]
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawValidatorCommission"))

			untypedInnerEvent, _ = innerCmd.Exec()
			createMsgWithdrawValidatorCommissionEvent := untypedInnerEvent.(*event.MsgWithdrawValidatorCommission)

			Expect(json.MustMarshalToString(createMsgWithdrawValidatorCommissionEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[1], ""),
					"{UUID}",
					createMsgWithdrawValidatorCommissionEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25"}))

		})

		It("should parse Msg commands when there is MsgExec (inner message MsgFundCommunityPool) in the transaction", func() {
			expected := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 2569,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "6628CBD728424C24B58E1598800DFFB7FCD8CBECD212925171E933980B8E7143",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgFundCommunityPool"
						},
						{
							"@type": "/cosmos.distribution.v1beta1.MsgFundCommunityPool"
						}
					]
				}
			}
			`

			expectedInnerMsg := []string{`
			{
				"name": "MsgFundCommunityPoolCreated",
				"version": 1,
				"height": 2569,
				"uuid": "{UUID}",
				"msgName": "MsgFundCommunityPool",
				"txHash": "6628CBD728424C24B58E1598800DFFB7FCD8CBECD212925171E933980B8E7143",
				"msgIndex": 0,
				"depositor": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "10"
					}
				]
			}`,
				`{
				"name": "MsgFundCommunityPoolCreated",
				"version": 1,
				"height": 2569,
				"uuid": "{UUID}",
				"msgName": "MsgFundCommunityPool",
				"txHash": "6628CBD728424C24B58E1598800DFFB7FCD8CBECD212925171E933980B8E7143",
				"msgIndex": 0,
				"depositor": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "11"
					}
				]
			}`}

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_FUND_COMMUNITY_POOL_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MSG_FUND_COMMUNITY_POOL_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(3))

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
			Expect(innerCmd.Name()).To(Equal("CreateMsgFundCommunityPool"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgFundCommunityPoolEvent := untypedInnerEvent.(*event.MsgFundCommunityPool)

			Expect(json.MustMarshalToString(createMsgFundCommunityPoolEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[0], ""),
					"{UUID}",
					createMsgFundCommunityPoolEvent.UUID(),
					-1,
				),
			))

			innerCmd = cmds[2]
			Expect(innerCmd.Name()).To(Equal("CreateMsgFundCommunityPool"))

			untypedInnerEvent, _ = innerCmd.Exec()
			createMsgFundCommunityPoolEvent = untypedInnerEvent.(*event.MsgFundCommunityPool)

			Expect(json.MustMarshalToString(createMsgFundCommunityPoolEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg[1], ""),
					"{UUID}",
					createMsgFundCommunityPoolEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25"}))

		})

		It("should parse Msg commands when there are two MsgExec in the transaction", func() {

			expected := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 386,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "4A1DC2FE28E63C23BC5DA89FBAD1ECD62868B21E7B097300B5C064F1985CD9D4",
				"msgIndex": 0,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
						}
					]
				}
			}
			`

			expectedInnerMsg :=
				`{
				"name": "MsgWithdrawDelegatorRewardCreated",
				"version": 1,
				"height": 386,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawDelegatorReward",
				"txHash": "4A1DC2FE28E63C23BC5DA89FBAD1ECD62868B21E7B097300B5C064F1985CD9D4",
				"msgIndex": 0,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"validatorAddress": "crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": [
					{
						"denom": "basecro",
						"amount": "513305"
					}
				]
			}`

			expected2 := `
			{
				"name": "MsgExecCreated",
				"version": 1,
				"height": 386,
				"uuid": "{UUID}",
				"msgName": "MsgExec",
				"txHash": "4A1DC2FE28E63C23BC5DA89FBAD1ECD62868B21E7B097300B5C064F1985CD9D4",
				"msgIndex": 1,
				"params": {
					"@type": "/cosmos.authz.v1beta1.MsgExec",
					"grantee": "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25",
					"msgs": [
						{
							"@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
						}
					]
				}
			}
			`

			expectedInnerMsg2 :=
				`{
				"name": "MsgWithdrawDelegatorRewardCreated",
				"version": 1,
				"height": 386,
				"uuid": "{UUID}",
				"msgName": "MsgWithdrawDelegatorReward",
				"txHash": "4A1DC2FE28E63C23BC5DA89FBAD1ECD62868B21E7B097300B5C064F1985CD9D4",
				"msgIndex": 1,
				"delegatorAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"validatorAddress": "crocncl17cye5gmealmpgz3tclfjga7urdjxj6nah3u8w2",
				"recipientAddress": "cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n",
				"amount": []
			}`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MULT_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test_msg_exec.TX_MSG_EXEC_MULT_WITHDRAW_DELEGATOR_REWARD_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)

			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(4))

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
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawDelegatorReward"))

			untypedInnerEvent, _ := innerCmd.Exec()
			createMsgWithdrawDelegatorRewardEvent := untypedInnerEvent.(*event.MsgWithdrawDelegatorReward)

			Expect(json.MustMarshalToString(createMsgWithdrawDelegatorRewardEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg, ""),
					"{UUID}",
					createMsgWithdrawDelegatorRewardEvent.UUID(),
					-1,
				),
			))

			cmd = cmds[2]
			Expect(cmd.Name()).To(Equal("CreateMsgExec"))

			untypedEvent, _ = cmd.Exec()
			createMsgExecEvent = untypedEvent.(*event.MsgExec)

			regex, _ = regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgExecEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected2, ""),
					"{UUID}",
					createMsgExecEvent.UUID(),
					-1,
				),
			))

			innerCmd = cmds[3]
			Expect(innerCmd.Name()).To(Equal("CreateMsgWithdrawDelegatorReward"))

			untypedInnerEvent, _ = innerCmd.Exec()
			createMsgWithdrawDelegatorRewardEvent = untypedInnerEvent.(*event.MsgWithdrawDelegatorReward)

			Expect(json.MustMarshalToString(createMsgWithdrawDelegatorRewardEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expectedInnerMsg2, ""),
					"{UUID}",
					createMsgWithdrawDelegatorRewardEvent.UUID(),
					-1,
				),
			))

			Expect(possibleSignerAddresses).To(Equal([]string{"cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25", "cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25"}))

		})
	})
})
