package parser_test

import (
	"encoding/json"

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
	Describe("MsgSubmitParamChangeProposal", func() {
		It("should parse Msg commands when there is gov.MsgSubmitParamChangeProposal in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitParamChangeProposal(
					event.MsgCommonParams{
						BlockHeight: int64(470973),
						TxHash:      "88B44ABB7B4E81E58AE6F0739B989A42542940D7583D83265ACDBE712A267F4F",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitParamChangeProposalParams{
						MaybeProposalId: primptr.String("2"),
						Content: model.MsgSubmitParamChangeProposalContent{
							Type:        "/cosmos.params.v1beta1.ParameterChangeProposal",
							Title:       "Param-Change Voting Period",
							Description: "This is to change the voting time on Testnet to be 8 hours.",
							Changes: []model.MsgSubmitParamChangeProposalChange{
								{
									Subspace: "gov",
									Key:      "votingparams",
									Value:    json.RawMessage("\"{ \\\"voting_period\\\": \\\"28800000000000\\\" }\""),
								},
							},
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.MustNewCoinFromString("10"),
					},
				),
			}))
		})

		It("should return a command with nil proposal id when the gov.MsgSubmitParamChangeProposal transaction failed", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_FAILED_MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitParamChangeProposal(
					event.MsgCommonParams{
						BlockHeight: int64(475232),
						TxHash:      "1682BBF647E8E326F4581986D7334883EC9783F41D892E6E425F6E962E50296C",
						TxSuccess:   false,
						MsgIndex:    0,
					},
					model.MsgSubmitParamChangeProposalParams{
						MaybeProposalId: primptr.StringNil(),
						Content: model.MsgSubmitParamChangeProposalContent{
							Type:        "/cosmos.params.v1beta1.ParameterChangeProposal",
							Title:       "Param-Change Voting Period",
							Description: "This is to change the voting time on Testnet to be 8 hours.",
							Changes: []model.MsgSubmitParamChangeProposalChange{
								{
									Subspace: "gov",
									Key:      "votingparams",
									Value:    json.RawMessage("\"{ \\\"voting_period\\\": \\\"28800000000000\\\" }\""),
								},
							},
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.MustNewCoinFromString("10"),
					},
				),
			}))
		})
	})
})
