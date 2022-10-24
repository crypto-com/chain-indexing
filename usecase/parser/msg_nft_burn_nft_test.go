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
	Describe("MsgNFTBurnNFT", func() {
		It("should parse command with effective height in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_NFT_BURN_NFT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_NFT_BURN_NFT_BLOCK_RESULTS_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_NFT_BURN_NFT_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

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
				command_usecase.NewCreateMsgNFTBurnNFT(
					event.MsgCommonParams{
						BlockHeight: int64(17699),
						TxHash:      "63B42F5AC39D758E5590E7D54A0F811968D1C5C0420EA5162CE83CA6D6818AA5",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgNFTBurnNFTParams{
						DenomId: "denomid",
						TokenId: "tokenid4",
						Sender:  "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"}))
		})
	})
})
