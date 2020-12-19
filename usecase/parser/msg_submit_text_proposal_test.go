package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/internal/primptr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSubmitTextProposal", func() {
		It("should parse gov.MsgSubmitTextProposal command  in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitTextProposal(
					event.MsgCommonParams{
						BlockHeight: int64(490470),
						TxHash:      "EE7D914AB846DF0D2E96F2EB32A843D3F8CCB0234962F53D6A8EBE7F024F9231",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitTextProposalParams{
						MaybeProposalId: primptr.String("9"),
						Content: model.MsgSubmitTextProposalContent{
							Type:        "/cosmos.gov.v1beta1.TextProposal",
							Title:       "Text Title",
							Description: "Text Description",
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.MustNewCoinFromString("2"),
					},
				),
			}))
		})
	})
})
