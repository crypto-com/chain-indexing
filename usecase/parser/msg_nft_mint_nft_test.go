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
	Describe("MsgNFTMintNFT", func() {
		It("should parse command with effective height in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_NFT_MINT_NFT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_NFT_MINT_NFT_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_NFT_MINT_NFT_TXS_RESP)
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
				command_usecase.NewCreateMsgNFTMintNFT(
					event.MsgCommonParams{
						BlockHeight: int64(11090),
						TxHash:      "5CC860DC00862A729C463BA414F13F2AB84908304DCA906C64365D26E40063C0",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgNFTMintNFTParams{
						DenomId:   "denomid",
						TokenId:   "tokenid123",
						TokenName: "tokenName",
						URI:       "tokenUri",
						Data:      "{\"title\":\"Vettel First Run\",\"type\":\"object\",\"properties\":{\"name\":\"Vettel First Run\",\"description\":\"Four-time Formula 1 world champion Sebastian Vettel made his name as one of the sport’s fastest and most fearless drivers. Vettel’s inaugural run in the AMR21 was also his first time out for Aston Martin Cognizant Formula One™ Team – his fifth Formula 1 outfit following stints at BMW Sauber, Scuderia Toro Rosso, Red Bull Racing and Scuderia Ferrari.\",\"image\":\"https://d2vi0z68k5oxnr.cloudfront.net/ac155986-d54c-4e23-957a-628325e051a4/original.mp4\",\"attributes\":[{\"trait_type\":\"Type\",\"value\":\"Human\"},{\"trait_type\":\"Hair Style\",\"value\":\"Bald\"},{\"trait_type\":\"Hat\",\"value\":\"Backwards Cap\"},{\"trait_type\":\"Hat Color\",\"value\":\"Gray\"},{\"trait_type\":\"Shirt\",\"value\":\"Skull Tee\"},{\"trait_type\":\"Overshirt\",\"value\":\"Athletic Jacket\"},{\"trait_type\":\"Overshirt Color\",\"value\":\"Red\"},{\"trait_type\":\"Pants\",\"value\":\"Cargo Pants\"},{\"trait_type\":\"Pants Color\",\"value\":\"Camo\"},{\"trait_type\":\"Shoes\",\"value\":\"Workboots\"}]}}",
						Sender:    "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
						Recipient: "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd",
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"}))
		})
	})
})
