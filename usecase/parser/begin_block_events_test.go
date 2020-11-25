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

		It("should return commands corresponding to events in begin_block_events", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.BEGIN_BLOCK_COMMON_EVENTS_BLOCK_RESULTS_RESP)

			cmds, err := parser.ParseBeginBlockEventsCommands(
				blockResults.Height,
				blockResults.BeginBlockEvents,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(9))
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
				command_usecase.NewCreateMint(
					expectecBlockHeight,
					model.MintParams{
						BondedRatio:      "0.000821761419299675",
						Inflation:        "0.013777334128586270",
						AnnualProvisions: "110307793770097823.255979052891494880",
						Amount:           "17477215277",
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
				// should not double count proposer reward and the same amount block reward events
				command_usecase.NewCreateBlockProposerReward(
					expectecBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					"868550031.392766344419273056basetcro",
				),
				command_usecase.NewCreateBlockCommission(
					expectecBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					"86855003.139276634441927306basetcro",
				),
				command_usecase.NewCreateBlockCommission(
					expectecBlockHeight,
					"tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
					"459938524.284156813832125321basetcro",
				),
				command_usecase.NewCreateBlockReward(
					expectecBlockHeight,
					"tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
					"919877048.568313627664250642basetcro",
				),
				// proposer get both proposer reward and block reward
				command_usecase.NewCreateBlockCommission(
					expectecBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					"59324118.921629850151833479basetcro",
				),
				command_usecase.NewCreateBlockReward(
					expectecBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					"593241189.216298501518334791basetcro",
				),
			}))
		})

		It("", func() {})
	})
})
