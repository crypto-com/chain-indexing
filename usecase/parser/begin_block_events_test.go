package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
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
			expectedBlockHeight := int64(377673)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateAccountTransfer(
					expectedBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
						Sender:    "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
						Amount:    coin.MustParseCoinsNormalized("17477215277basetcro"),
					},
				),
				command_usecase.NewCreateMint(
					expectedBlockHeight,
					model.MintParams{
						BondedRatio:      "0.000821761419299675",
						Inflation:        "0.013777334128586270",
						AnnualProvisions: "110307793770097823.255979052891494880",
						Amount:           "17477215277",
					},
				),
				command_usecase.NewCreateAccountTransfer(
					expectedBlockHeight,
					model.AccountTransferParams{
						Recipient: "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
						Sender:    "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
						Amount:    coin.MustParseCoinsNormalized("17477255277basetcro"),
					},
				),
				// should not double count proposer reward and the same amount block reward events
				command_usecase.NewCreateBlockProposerReward(
					expectedBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					coin.MustParseDecCoins("868550031.392766344419273056basetcro"),
				),
				command_usecase.NewCreateBlockCommission(
					expectedBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					coin.MustParseDecCoins("86855003.139276634441927306basetcro"),
				),
				command_usecase.NewCreateBlockCommission(
					expectedBlockHeight,
					"tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
					coin.MustParseDecCoins("459938524.284156813832125321basetcro"),
				),
				command_usecase.NewCreateBlockReward(
					expectedBlockHeight,
					"tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
					coin.MustParseDecCoins("919877048.568313627664250642basetcro"),
				),
				// proposer get both proposer reward and block reward
				command_usecase.NewCreateBlockCommission(
					expectedBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					coin.MustParseDecCoins("59324118.921629850151833479basetcro"),
				),
				command_usecase.NewCreateBlockReward(
					expectedBlockHeight,
					"tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					coin.MustParseDecCoins("593241189.216298501518334791basetcro"),
				),
			}))
		})

		It("should return ValidatorSlashed and ValidatorJailed command base on missing signature events", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.BEGIN_BLOCK_SLASH_MISSING_SIGNATURES_EVENT_BLOCK_RESULTS_RESP)
			cmds, err := parser.ParseBeginBlockEventsCommands(
				blockResults.Height,
				blockResults.BeginBlockEvents,
			)
			Expect(err).To(BeNil())
			expectedBlockHeight := int64(168481)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewSlashValidator(
					expectedBlockHeight,
					model.SlashValidatorParams{
						ConsensusNodeAddress: "crocnclcons18vyrj3se6cvdryk3vp09x88n46894ukjywnfmy",
						SlashedPower:         "17274617",
						Reason:               "missing_signature",
					},
				),
				command_usecase.NewJailValidator(
					expectedBlockHeight,
					"crocnclcons18vyrj3se6cvdryk3vp09x88n46894ukjywnfmy",
					"missing_signature",
				),
				command_usecase.NewSlashValidator(
					expectedBlockHeight,
					model.SlashValidatorParams{
						ConsensusNodeAddress: "crocnclcons10wy4k3dkjd3htgleaxlzmakh0k4h8ql23cpusx",
						SlashedPower:         "9902032",
						Reason:               "missing_signature",
					},
				),
				command_usecase.NewJailValidator(
					expectedBlockHeight,
					"crocnclcons10wy4k3dkjd3htgleaxlzmakh0k4h8ql23cpusx",
					"missing_signature",
				),
			}))
		})

		It("should return ValidatorSlashed and ValidatorJailed command base on double sign events", func() {
			blockResults := mustParseBlockResultsResp(usecase_parser_test.BEGIN_BLOCK_SLASH_DOUBLE_SIGN_EVENT_BLOCK_RESULTS_RESP)
			cmds, err := parser.ParseBeginBlockEventsCommands(
				blockResults.Height,
				blockResults.BeginBlockEvents,
			)
			Expect(err).To(BeNil())
			expectedBlockHeight := int64(165487)
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewSlashValidator(
					expectedBlockHeight,
					model.SlashValidatorParams{
						ConsensusNodeAddress: "crocnclcons1fnht46350sxm2e6w8ma3as2l6wdl78rfym8gku",
						SlashedPower:         "15600000",
						Reason:               "double_sign",
					},
				),
				command_usecase.NewJailValidator(
					expectedBlockHeight,
					"crocnclcons1fnht46350sxm2e6w8ma3as2l6wdl78rfym8gku",
					"double_sign",
				),
			}))
		})
	})
})
