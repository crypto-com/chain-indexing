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
	Describe("MsgNFTIssueDenom", func() {
		It("should parse command with effective height in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_NFT_ISSUE_DENOM_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_NFT_ISSUE_DENOM_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_NFT_ISSUE_DENOM_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "cro"
			bondingDenom := "basecro"

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
				command_usecase.NewCreateMsgNFTIssueDenom(
					event.MsgCommonParams{
						BlockHeight: int64(1972),
						TxHash:      "FF16D34DA9991673CB3BD648A58D067D011E0C569E0AC71142735F7D86D71549",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgNFTIssueDenomParams{
						DenomId:   "denomid",
						DenomName: "DenomName",
						Schema:    "Schema",
						Sender:    "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"}))
		})
	})
})
