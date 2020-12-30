package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("Parse Genesis", func() {
	It("should return genesis command corresponding to genesis response", func() {
		genesis := mustParseGenesisResp(usecase_parser_test.GENESIS_RESP)

		cmds, err := parser.ParseGenesisCommands(genesis)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(4))
		Expect(cmds[0]).To(Equal(command_usecase.NewCreateGenesis(*genesis)))
		Expect(cmds[1]).To(Equal(
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-0",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "node0",
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
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1n4t5q77kn9vf73s7ljs96m85jgg49yqpasmwm3",
					ValidatorAddress:  "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
					TendermintPubkey:  "Og8ZfQTHFgTBGD5qoyo5NpyJCJRddC+WuSPtyZtlE7E=",
					Amount:            coin.MustNewCoinFromString("10000000000000"),
				},
			),
		))
		Expect(cmds[2]).To(Equal(
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-1",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "node2",
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
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro15xr8daqzpu0wf8t6hx95zlxmqwzmf4eaph3yzv",
					ValidatorAddress:  "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60",
					TendermintPubkey:  "BuuPYme7R4eH/nWs2p+sS1UpCQwy+QJgBZuhGICH8Es=",
					Amount:            coin.MustNewCoinFromString("10000000000000"),
				},
			),
		))
		Expect(cmds[3]).To(Equal(
			command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: 0,
					TxHash:      "genesis-gentxs-2",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description: model.MsgValidatorDescription{
						Moniker:         "node1",
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
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro197ujxhaeyyv309f39c0s2gn0af0pps5pden6h7",
					ValidatorAddress:  "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a",
					TendermintPubkey:  "wWw0e9tZcVmev/NyJlZv5Apd7U5IONoyx3U/9rD5fHI=",
					Amount:            coin.MustNewCoinFromString("10000000000000"),
				},
			),
		))
	})
})
