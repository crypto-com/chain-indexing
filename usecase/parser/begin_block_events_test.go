package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("ParseBeginBlockEventsCommands", func() {
	Describe("MsgSend", func() {

		It("should return CreateAccountTransfer commands when there is transfer event in begin_block_events", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.BEGIN_BLOCK_TRANFER_EVENTS_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBeginBlockEventsCommands(
				blockResults.Height,
				blockResults.BeginBlockEvents,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			expectecBlockHeight := int64(377673)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
						Sender:    "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
						Amount:    coin.MustNewCoinFromString("17477215277"),
					},
				),
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
						Sender:    "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
						Amount:    coin.MustNewCoinFromString("17477255277"),
					},
				),
			}))
		})
	})
})
