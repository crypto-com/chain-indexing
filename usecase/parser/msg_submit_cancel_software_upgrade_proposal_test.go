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
	Describe("MsgSubmitCancelSoftwareUpgradeProposal", func() {
		It("should parse gov.MsgSubmitCancelSoftwareUpgradeProposal command  in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_TXS_RESP)
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
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitCancelSoftwareUpgradeProposal(
					event.MsgCommonParams{
						BlockHeight: int64(490470),
						TxHash:      "EE7D914AB846DF0D2E96F2EB32A843D3F8CCB0234962F53D6A8EBE7F024F9231",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitCancelSoftwareUpgradeProposalParams{
						MaybeProposalId: primptr.String("9"),
						Content: model.MsgSubmitCancelSoftwareUpgradeProposalContent{
							Type:        "/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal",
							Title:       "Cancel Upgrade Title",
							Description: "Cancel Upgrade Description",
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
