package parser_test

import (
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/usecase/coin"
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
		It("should parse gov.MsgSubmitTextProposal and NewStartProposalVotingPeriod command in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_AND_START_VOTING_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitTextProposal(
					event.MsgCommonParams{
						BlockHeight: int64(2036),
						TxHash:      "AAE71A09FFCDC3DD9D26CF2580CF26C38DDE2C6D7CBBF4264D295086B7E24148",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitTextProposalParams{
						MaybeProposalId: primptr.String("2"),
						Content: model.MsgSubmitTextProposalContent{
							Type:        "/cosmos.gov.v1beta1.TextProposal",
							Title:       "Test Proposal",
							Description: "My awesome proposal",
						},
						ProposerAddress: "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
						InitialDeposit:  coin.MustParseCoinsNormalized("20000basecro"),
					},
				),
				command_usecase.NewStartProposalVotingPeriod(int64(2036), "2"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"}))
		})
	})
})
