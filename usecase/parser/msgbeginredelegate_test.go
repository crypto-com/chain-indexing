package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgBeginRedelegate", func() {

		It("should parse Msg commands when there is staking.MsgBeginRedelegate in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
					Amount:              coin.MustNewCoinFromString("10000000000"),
				},
			)}))
		})
	})
})
