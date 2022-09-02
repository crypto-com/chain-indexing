package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSend", func() {

		It("should parse Msg commands when there is bank.MsgSend in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_SEND_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSend(
				event.MsgCommonParams{
					BlockHeight: int64(377673),
					TxHash:      "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				event.MsgSendCreatedParams{
					FromAddress: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
					ToAddress:   "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
					Amount:      coin.MustParseCoinsNormalized("1000000000basetcro"),
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv"}))
		})

		It("should parse Msg commands when there are multiple bank.MsgSend in one transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSend(
				event.MsgCommonParams{
					BlockHeight: int64(343358),
					TxHash:      "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				event.MsgSendCreatedParams{
					FromAddress: "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
					ToAddress:   "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					Amount:      coin.MustParseCoinsNormalized("1000basetcro"),
				},
			), command_usecase.NewCreateMsgSend(
				event.MsgCommonParams{
					BlockHeight: int64(343358),
					TxHash:      "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416",
					TxSuccess:   true,
					MsgIndex:    1,
				},
				event.MsgSendCreatedParams{
					FromAddress: "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					ToAddress:   "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
					Amount:      coin.MustParseCoinsNormalized("2000basetcro"),
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3", "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"}))
		})
	})
})
