package parser_test

import (
	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chainindex/usecase/parser/test"
)

var _ = Describe("Parse Genesis", func() {
	It("should return genesis command corresponding to genesis response", func() {
		genesis := mustParseGenesisResp(usecase_parser_test.GENESIS_RESP)

		cmds, err := parser.ParseGenesisCommands(genesis)
		Expect(err).To(BeNil())
		Expect(cmds).To(Equal([]command.Command{
			command_usecase.NewCreateGenesis(*genesis),
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-0",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "jotaro",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.MsgValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					DelegatorAddress: "tcro16kqr009ptgken6qsxnzfnyjfsq6q97g3fxwppq",
					ValidatorAddress: "tcrocncl16kqr009ptgken6qsxnzfnyjfsq6q97g3uedcer",
					Pubkey:           "tcrocnclconspub1zcjduepq5xp88wqmrhkg3xuyl6vcey3d93kw6cdglkmq4ley3ysvjfx90jnqlvaxpc",
					Amount:           coin.MustNewCoinFromString("10000000000000"),
				},
			),
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-1",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "Stater",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.MsgValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					DelegatorAddress: "tcro1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dsdjun5p",
					ValidatorAddress: "tcrocncl1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dscdl2vz",
					Pubkey:           "tcrocnclconspub1zcjduepqp6el28dgz0fs6pm2265v4xmx0uys65zf5s2av6r5gh0hmcv5j64qg6zj4w",
					Amount:           coin.MustNewCoinFromString("10000000000000"),
				},
			),
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-2",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "Argenteus",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.MsgValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					DelegatorAddress: "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
					ValidatorAddress: "tcrocncl16yzcz3ty94awr7nr2txek9dp2klp2av9vh437s",
					Pubkey:           "tcrocnclconspub1zcjduepq9y742s8duh0eqenjdpc24k785c4en90gypn2vs2x3gxjd45xpp5svldq9u",
					Amount:           coin.MustNewCoinFromString("10000000000000"),
				},
			),
		}))
	})
})
