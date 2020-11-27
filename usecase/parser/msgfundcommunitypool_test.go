package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgFundCommunityPool", func() {
		It("should parse Msg commands when there is distribution.MsgFundCommunityPool in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgFundCommunityPool(
				event.MsgCommonParams{
					BlockHeight: int64(460662),
					TxHash:      "933052FD68B10549312F3CBA9AF4F4CC77536BE5ECD335638DD36ECCE681201E",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgFundCommunityPoolParams{
					Depositor: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					Amount:    coin.MustNewCoinFromString("1"),
				},
			)}))
		})
	})
})
