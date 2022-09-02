package parser_test

import (
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
	Describe("MsgVote", func() {
		It("should parse gov.MsgVote command in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_VOTE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_VOTE_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_VOTE_TXS_RESP)
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
				command_usecase.NewCreateMsgVote(
					event.MsgCommonParams{
						BlockHeight: int64(100),
						TxHash:      "6E6910024B74B16F3B9B14309D7F8CD89AF25E561F0FB3F56380F086218F1759",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgVoteParams{
						ProposalId: "1",
						Voter:      "cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7",
						Option:     "VOTE_OPTION_YES",
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1tg4xpryye2v4fp3smpfc3s2kqmvnrkwfyd63y7"}))
		})
	})
})
