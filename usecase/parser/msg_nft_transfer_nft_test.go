package parser_test

import (
	"github.com/hashicorp/go-version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgNFTTransferNFT", func() {
		It("should parse command with effective height in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_NFT_TRANSFER_NFT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_NFT_TRANSFER_NFT_BLOCK_RESULTS_RESP,
			)
			accountAddressPrefix := "cro"
			bondingDenom := "basecro"

			anyVersion := version.Must(version.NewVersion("v0.43"))

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				bondingDenom,
				anyVersion,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgNFTTransferNFT(
					event.MsgCommonParams{
						BlockHeight: int64(13103),
						TxHash:      "8CF41CCC69DCE6B784B4A37B12017EE5A18A2018E17D2B6CEC3E06F4DFD7DFB1",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgNFTTransferNFTParams{
						DenomId:   "denomid",
						TokenId:   "tokenid2",
						Sender:    "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
						Recipient: "cro10ar7dpv4g95r3gj06pn52f64vwq7v9kdh0g58a",
					},
				),
			}))
		})
	})
})
