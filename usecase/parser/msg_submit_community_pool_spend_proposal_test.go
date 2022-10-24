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
	Describe("MsgSubmitCommunityPoolSpendProposal", func() {
		It("should parse Msg commands when there is gov.MsgSubmitCommunityPoolSpendProposal in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

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
							Amount:           coin.MustParseCoinsNormalized("1basetcro"),
						},
						ProposerAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						InitialDeposit:  coin.MustParseCoinsNormalized("2basetcro"),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"}))
		})
	})
})
