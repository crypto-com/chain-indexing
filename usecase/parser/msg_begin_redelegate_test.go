package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
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
	Describe("MsgBeginRedelegate", func() {
		It("should parse Msg commands when there is staking.MsgBeginRedelegate in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESULTS_RESP)
			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgBeginRedelegate(
				event.MsgCommonParams{
					BlockHeight: int64(374394),
					TxHash:      "97171BB77771E1288E86756B8EFEDB958B8B778C91ED1AF047A98BE540D70A01",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgBeginRedelegateParams{
					DelegatorAddress:    "tcro1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5w2ll64",
					ValidatorSrcAddress: "tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					ValidatorDstAddress: "tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
					Amount:              coin.MustParseCoinNormalized("10000000000basetcro"),
					AutoClaimedRewards:  coin.NewCoin("basetcro", coin.NewInt(281334)),
				},
			)}))
		})
	})
})
