package parser_test

import (
	"strings"

	"github.com/crypto-com/chainindex/usecase/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSetWithdrawAddress", func() {
		It("should parse Msg commands when there is distribution.MsgSetWithdrawAddress in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESULTS_RESP,
			))

			cmds := parser.ParseMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSetWithdrawAddress(
				event.MsgCommonParams{
					BlockHeight: int64(460060),
					TxHash:      "9C2501310E18EE69A7FE5CA1A684A0701C43BEB1A8D91EDA80CC598C924F9CBE",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgSetWithdrawAddressParams{
					DelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					WithdrawAddress:  "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				},
			)}))
		})
	})
})
