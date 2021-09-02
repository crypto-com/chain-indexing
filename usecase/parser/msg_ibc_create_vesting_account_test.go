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
	Describe("MsgIBCExec", func() {
		It("should parse Msg commands when there is MsgExec in the transaction", func() {
			expected := `{
            "name": "MsgCreateVestingAccountCreated",
            "version": 1,
            "height": 193118,
            "uuid": "{UUID}",
            "msgName": "MsgCreateVestingAccount",
            "txHash": "6DDA0564ED4ADB89AC8518CDB3B990D7959EBD5775D44AB2A3381E38722A21C7",
            "msgIndex": 0,
            "params": {
                "@type": "/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
                "fromAddress": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
                "toAddress": "tcro155kkac0jc05dxceylphwyhx2dt868u9t6jk6g5",
                "amount": [
                    {
                        "denom": "basetcro",
                        "amount": "100000000"
                    }
                ],
                "endTime": 1,
                "delayed": false
            }
        }`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_VESTING_ACCOUNT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_VESTING_ACCOUNT_BLOCK_RESULTS_RESP,
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
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgCreateVestingAccount"))

			untypedEvent, _ := cmd.Exec()
			createMsgCreateVestingAccount := untypedEvent.(*event.MsgIBCCreateVestingAccount)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgCreateVestingAccount)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgCreateVestingAccount.UUID(),
					-1,
				),
			))
		})
	})
})
