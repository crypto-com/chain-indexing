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
	Describe("MsgSend", func() {

		It("should parse Msg commands when there is bank.MsgSend in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
					Amount:      coin.MustNewCoinFromString("1000000000"),
				},
			)}))
		})

		It("should parse Msg commands when there are multiple bank.MsgSend in one transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
					Amount:      coin.MustNewCoinFromString("1000"),
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
					Amount:      coin.MustNewCoinFromString("2000"),
				},
			)}))
		})
	})
})
