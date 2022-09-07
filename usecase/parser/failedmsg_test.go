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
	Describe("Failed Msg", func() {
		It("should parse Msg Failed commands when the transaction has failed", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_WITH_FEE_TXS_RESP)
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
					BlockHeight: int64(420301),
					TxHash:      "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
					TxSuccess:   false,
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
	})
})
