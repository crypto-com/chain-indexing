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

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_EXEC_BLOCK_RESULTS_RESP,
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
			Expect(cmd.Name()).To(Equal("CreateMsgExec"))

			untypedEvent, _ := cmd.Exec()
			createMsgExecEvent := untypedEvent.(*event.MsgIBCExec)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgExecEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgExecEvent.UUID(),
					-1,
				),
			))
		})
	})
})
