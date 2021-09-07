package parser_test

import (
	"strings"

	"github.com/crypto-com/chain-indexing/projection/validator/constants"

	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("Parse Genesis", func() {
	It("should throw when parsing unknown fields with strict mode on", func() {
		genesisReader := strings.NewReader(usecase_parser_test.GENESIS_RESP)
		strict := true
		_, err := tendermint.ParseGenesisResp(genesisReader, strict)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("error decoding Tendermint genesis response"))
	})

	It("should throw when parsing unknown fields with strict mode on", func() {
		genesisReader := strings.NewReader(usecase_parser_test.GENESIS_TESTNET_CROESEID_3_RESP)
		strict := true

		_, err := tendermint.ParseGenesisResp(genesisReader, strict)
		Expect(err).To(BeNil())
	})

	It("should return genesis command corresponding to staking validator in genesis response", func() {
		strict := false
		rawGenesis := mustParseGenesisResp(usecase_parser_test.GENESIS_EXPORTED_RESP, strict)

		accountAddressPrefix := "tcro"
		cmds, err := ParseGenesisCommands(rawGenesis, accountAddressPrefix)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(4))
		Expect(cmds[0]).To(Equal(command_usecase.NewCreateGenesis(*rawGenesis)))
		Expect(cmds[1]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.UNBONDED,
					Description: model.ValidatorDescription{
						Moniker:         "sg42-node",
						Identity:        "",
						Website:         "",
						SecurityContact: "@gutz42:matrix.org",
						Details:         "",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1q435860mlxc8954ye4v6vghwge8rw5eq5newp7",
					ValidatorAddress:  "tcrocncl1q435860mlxc8954ye4v6vghwge8rw5eqpv6hea",
					TendermintPubkey:  "npeBO7O/zYRoGCwjTKf04ZBMkvwDWOF5FbiU6t3u2Kc=",
					Amount:            coin.MustParseCoinNormalized("499500000basetcro"),
					Jailed:            true,
				},
			),
		))
		Expect(cmds[2]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.UNBONDED,
					Description: model.ValidatorDescription{
						Moniker:         "bt",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.050000000000000000",
						MaxRate:       "1.000000000000000000",
						MaxChangeRate: "1.000000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro12ynscey3jv65trm0j02d42feg5epm6sv9z7lrh",
					ValidatorAddress:  "tcrocncl12ynscey3jv65trm0j02d42feg5epm6svsaaxm5",
					TendermintPubkey:  "ewmWXircBKmed5/MP2UDXn8c7uJFABHs/L/vJ5s1vOQ=",
					Amount:            coin.MustParseCoinNormalized("500000basetcro"),
					Jailed:            false,
				},
			),
		))
		Expect(cmds[3]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.BONDED,
					Description: model.ValidatorDescription{
						Moniker:         "GreatLakes-node - TeamThrive",
						Identity:        "",
						Website:         "",
						SecurityContact: "joppy@oh.rr.com",
						Details:         "Part of the TeamThrive Operation",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1qm8c62ewj99ufj34jgjk3uv3tu3a6jxvycyktp",
					ValidatorAddress:  "tcrocncl1qm8c62ewj99ufj34jgjk3uv3tu3a6jxv3880nz",
					TendermintPubkey:  "Rav9WljfqBIvZoh1lP/3EBhI9KalH6GxFGTsThWcPSs=",
					Amount:            coin.MustParseCoinNormalized("249484668742000basetcro"),
					Jailed:            false,
				},
			),
		))
	})

	It("should return genesis command corresponding to genesis response", func() {
		strict := false
		rawGenesis := mustParseGenesisResp(usecase_parser_test.GENESIS_RESP, strict)

		accountAddressPrefix := "tcro"
		cmds, err := ParseGenesisCommands(rawGenesis, accountAddressPrefix)
		Expect(err).To(BeNil())
		Expect(cmds).To(HaveLen(4))
		Expect(cmds[0]).To(Equal(command_usecase.NewCreateGenesis(*rawGenesis)))
		Expect(cmds[1]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.BONDED,
					Description: model.ValidatorDescription{
						Moniker:         "node0",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1n4t5q77kn9vf73s7ljs96m85jgg49yqpasmwm3",
					ValidatorAddress:  "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
					TendermintPubkey:  "Og8ZfQTHFgTBGD5qoyo5NpyJCJRddC+WuSPtyZtlE7E=",
					Amount:            coin.MustParseCoinNormalized("10000000000000basetcro"),
					Jailed:            false,
				},
			),
		))
		Expect(cmds[2]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.BONDED,
					Description: model.ValidatorDescription{
						Moniker:         "node2",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro15xr8daqzpu0wf8t6hx95zlxmqwzmf4eaph3yzv",
					ValidatorAddress:  "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60",
					TendermintPubkey:  "BuuPYme7R4eH/nWs2p+sS1UpCQwy+QJgBZuhGICH8Es=",
					Amount:            coin.MustParseCoinNormalized("10000000000000basetcro"),
					Jailed:            false,
				},
			),
		))
		Expect(cmds[3]).To(Equal(
			command_usecase.NewCreateGenesisValidator(
				genesis.CreateGenesisValidatorParams{
					Status: constants.BONDED,
					Description: model.ValidatorDescription{
						Moniker:         "node1",
						Identity:        "",
						Website:         "",
						SecurityContact: "",
						Details:         "",
					},
					Commission: model.ValidatorCommission{
						Rate:          "0.100000000000000000",
						MaxRate:       "0.200000000000000000",
						MaxChangeRate: "0.010000000000000000",
					},
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro197ujxhaeyyv309f39c0s2gn0af0pps5pden6h7",
					ValidatorAddress:  "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a",
					TendermintPubkey:  "wWw0e9tZcVmev/NyJlZv5Apd7U5IONoyx3U/9rD5fHI=",
					Amount:            coin.MustParseCoinNormalized("10000000000000basetcro"),
					Jailed:            false,
				},
			),
		))
	})
})
