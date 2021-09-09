package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

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
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SUBMIT_TEXT_PROPOSAL_BLOCK_RESULTS_RESP,
			)
			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgSubmitTextProposal(
					event.MsgCommonParams{
						BlockHeight: int64(874207),
						TxHash:      "579B97CD5B947C2FA0EC87EDD4DAA8BECF422B96A82E2C9DBFE15F9F6DB4109B",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgSubmitTextProposalParams{
						MaybeProposalId: primptr.String("10"),
						Content: model.MsgSubmitTextProposalContent{
							Type:        "/cosmos.gov.v1beta1.TextProposal",
							Title:       "A proposal test from crypto.bzh",
							Description: "This a description for the proposal",
						},
						ProposerAddress: "tcro14fnzv5g92s6f8dg534lccp4x5tylkvth7zcq0u",
						InitialDeposit:  coin.MustParseCoinsNormalized("1000000basetcro"),
					},
				),
			}))
		})
	})
})
