package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/coin"

	"github.com/crypto-com/chainindex/internal/primptr"
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
	Describe("MsgSubmitCommunityPoolSpendProposal", func() {
		It("should parse Msg commands when there is gov.MsgSubmitCommunityPoolSpendProposal in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitCommunityPoolSpendProposal(
					event.MsgCommonParams{
						BlockHeight: int64(483830),
						TxHash:      "4305B12B7135955080A96D9A1C33455FBC704A7B88E6C74E884AF5A01FB80B3B",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitCommunityPoolSpendProposalParams{
						MaybeProposalId: primptr.String("5"),
						Content: model.MsgSubmitCommunityPoolSpendProposalContent{
							Type:             "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal",
							Title:            "Community Pool Spend",
							Description:      "Pay me some Cro!",
							RecipientAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
							Amount:           coin.MustNewCoinFromString("1"),
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.MustNewCoinFromString("2"),
					},
				),
			}))
		})
	})
})
