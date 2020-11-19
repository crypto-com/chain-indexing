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
	Describe("MsgDelegate", func() {

		It("should parse Msg commands when there is staking.MsgDelegate in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_DELEGATE_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_DELEGATE_BLOCK_RESULTS_RESP))

			cmds := parser.ParseMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgDelegate(
				event.MsgCommonParams{
					BlockHeight: int64(466543),
					TxHash:      "005BC5071A655A6219F7ECFE677E050866A33A174BC63A372A3B6208F4DE1F6C",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgDelegateParams{
					DelegatorAddress: "tcro1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxh5hy8r",
					ValidatorAddress: "tcrocncl1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxzt5alq",
					Amount:           coin.MustNewCoinFromString("27464382775"),
				},
			)}))
		})
	})
})
