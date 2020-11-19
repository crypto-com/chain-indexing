package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
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
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_BEGIN_REDELEGATE_BLOCK_RESULTS_RESP))

			cmds := parser.ParseMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
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
