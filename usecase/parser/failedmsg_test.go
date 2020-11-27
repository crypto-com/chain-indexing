package parser_test

import (
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
	Describe("Failed Msg", func() {
		It("should parse Msg Failed commands when the transaction has failed", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_WITH_FEE_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
					Amount:      coin.MustNewCoinFromString("1000000000"),
				},
			)}))
		})
	})
})
