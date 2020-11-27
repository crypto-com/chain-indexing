package parser_test

import (
	"github.com/crypto-com/chainindex/internal/utctime"
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
	Describe("MsgSubmitSoftwareUpgradeProposal", func() {
		It("should parse gov.MsgSubmitCommunityPoolSpendProposal command with effective height in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_HEIGHT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_HEIGHT_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
					event.MsgCommonParams{
						BlockHeight: int64(484638),
						TxHash:      "6068B21BA119610D545939518D6399DFFC4914A0CEE0822D752813C4D68D272C",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitSoftwareUpgradeProposalParams{
						MaybeProposalId: primptr.String("6"),
						Content: model.MsgSubmitSoftwareUpgradeProposalContent{
							Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
							Title:       "Upgrade Title",
							Description: "Upgrade Description",
							Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
								Name:   "Upgrade Name",
								Time:   utctime.FromUnixNano(int64(-6795364578871345152)),
								Height: 100000000,
								Info:   "Upgrade Info",
							},
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.Zero(),
					},
				),
			}))
		})

		It("should parse gov.MsgSubmitCommunityPoolSpendProposal command with effective time in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_TIME_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_TIME_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
					event.MsgCommonParams{
						BlockHeight: int64(490333),
						TxHash:      "7422843783CADA0321F517CAD2F38CB63E3C2CBD126150272BC5D678C8869E92",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitSoftwareUpgradeProposalParams{
						MaybeProposalId: primptr.String("8"),
						Content: model.MsgSubmitSoftwareUpgradeProposalContent{
							Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
							Title:       "Upgrade Title",
							Description: "Upgrade Description",
							Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
								Name:   "Upgrade Name",
								Time:   utctime.FromUnixNano(int64(4102585445000000000)),
								Height: 0,
								Info:   "Upgrade Info",
							},
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.Zero(),
					},
				),
			}))
		})
	})
})
