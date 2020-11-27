package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgDeposit", func() {
		It("should parse gov.MsgDeposit command with effective height in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_DEPOSIT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_DEPOSIT_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgDeposit(
					event.MsgCommonParams{
						BlockHeight: int64(493659),
						TxHash:      "52D0E10D951DA5124E899A988ED6D014AF750B951688C6A59853A97E4AFE69B4",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgDepositParams{
						ProposalId: "9",
						Depositor:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Amount:     coin.MustNewCoinFromString("2"),
					},
				),
			}))
		})
	})
})
