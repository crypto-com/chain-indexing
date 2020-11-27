package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/event"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgMultiSend", func() {

		It("should parse Msg commands when there is bank.MsgMultiSend in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_MULTI_SEND_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_MULTI_SEND_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgMultiSend(
				event.MsgCommonParams{
					BlockHeight: int64(421481),
					TxHash:      "89D220EC5C551DA81EE6993D05B53F25912B072E7FFD863D93AA742923239A5F",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgMultiSendParams{
					Inputs: []model.MsgMultiSendInput{
						{
							Address: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
							Amount:  coin.MustNewCoinFromString("51"),
						},
					},
					Outputs: []model.MsgMultiSendOutput{
						{
							Address: "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
							Amount:  coin.MustNewCoinFromString("1"),
						},
						{
							Address: "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
							Amount:  coin.MustNewCoinFromString("20"),
						},
						{
							Address: "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
							Amount:  coin.MustNewCoinFromString("30"),
						},
					},
				},
			)}))
		})
	})
})
