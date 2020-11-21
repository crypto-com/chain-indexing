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

var _ = Describe("ParseTxAccountTransferCommands", func() {
	Describe("MsgSend", func() {

		It("should parse CreateAccountTransfer commands when there is transfer event in transaction", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_SEND_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseTxAccountTransferCommands(
				blockResults.Height,
				blockResults.TxsResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			expectecBlockHeight := int64(377673)
			Expect(cmds).To(Equal([]command.Command{
				// Transaction fee
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
						Sender:    "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
						Amount:    coin.MustNewCoinFromString("8000000"),
					},
				),
				// MsgSend
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
						Sender:    "tcro1feqh6ad9ytjkr79kjk5nhnl4un3wez0ynurrwv",
						Amount:    coin.MustNewCoinFromString("1000000000"),
					},
				),
			}))
		})

		It("should parse Msg commands when there are multiple bank.MsgSend in one transaction", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.ONE_TX_TWO_MSG_SEND_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseTxAccountTransferCommands(
				blockResults.Height,
				blockResults.TxsResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			expectecBlockHeight := int64(343358)
			Expect(cmds).To(Equal([]command.Command{
				// Transaction fee
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
						Sender:    "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
						Amount:    coin.MustNewCoinFromString("1000"),
					},
				),
				// MsgSend
				command_usecase.NewCreateAccountTransfer(
					expectecBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3",
						Sender:    "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
						Amount:    coin.MustNewCoinFromString("2000"),
					},
				),
			}))
		})
	})
})
