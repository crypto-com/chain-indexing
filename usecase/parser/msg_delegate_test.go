package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgDelegate", func() {

		It("should parse Msg commands when there is staking.MsgDelegate in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_DELEGATE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_DELEGATE_BLOCK_RESULTS_RESP)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_DELEGATE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgDelegate(
				event.MsgCommonParams{
					BlockHeight: int64(466543),
					TxHash:      "005BC5071A655A6219F7ECFE677E050866A33A174BC63A372A3B6208F4DE1F6C",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgDelegateParams{
					DelegatorAddress:   "tcro1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxh5hy8r",
					ValidatorAddress:   "tcrocncl1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxzt5alq",
					Amount:             coin.MustParseCoinNormalized("27464382775basetcro"),
					AutoClaimedRewards: coin.MustParseCoinNormalized("4082858866basetcro"),
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxh5hy8r"}))
		})
	})
})
