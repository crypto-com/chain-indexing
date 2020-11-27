package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgUnjail", func() {

		It("should parse Msg commands when there is slashing.MsgUnjail in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_UNJAIL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_UNJAIL_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgUnjail(
				event.MsgCommonParams{
					BlockHeight: int64(381001),
					TxHash:      "2D2075ECAF45DB84052B6A98E84C5E7FE158AB5157364FD5A934FAC96B77C6B1",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgUnjailParams{
					ValidatorAddr: "tcrocncl1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5m4uxzk",
				},
			)}))
		})
	})
})
