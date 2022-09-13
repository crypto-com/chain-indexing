package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgWithdrawDelegatorReward and MsgWithdrawValidatorCommission", func() {
		It("should parse Msg commands when there is distribution.MsgWithdrawDelegatorReward and MsgWithdrawValidatorCommission in the transaction", func() {
			block, _ := mustParseBlockResp(
				usecase_parser_test.TX_MSGS_WITHDRAW_DELEGATOR_REWARD_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESP,
			)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSGS_WITHDRAW_DELEGATOR_REWARD_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSGS_WITHDRAW_DELEGATOR_REWARD_WITHDRAW_VALIDATOR_COMMISSION_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgWithdrawDelegatorReward(
					event.MsgCommonParams{
						BlockHeight: int64(435599),
						TxHash:      "202A0C0F03760D523C2F64F07C527E789F3D87819CB7097B35422A7F14FA055A",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgWithdrawDelegatorRewardParams{
						DelegatorAddress: "tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
						ValidatorAddress: "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek",
						RecipientAddress: "tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
						Amount:           coin.MustParseCoinsNormalized("33934701990basetcro"),
					},
				),
				command_usecase.NewCreateMsgWithdrawValidatorCommission(
					event.MsgCommonParams{
						BlockHeight: int64(435599),
						TxHash:      "202A0C0F03760D523C2F64F07C527E789F3D87819CB7097B35422A7F14FA055A",
						TxSuccess:   true,
						MsgIndex:    1,
					},
					model.MsgWithdrawValidatorCommissionParams{
						ValidatorAddress: "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek",
						RecipientAddress: "tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
						Amount:           coin.MustParseCoinsNormalized("4161370358basetcro"),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4", "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"}))
		})

		It("should parse failed MsgWithdrawValidatorCommission in the transaction", func() {
			block, _ := mustParseBlockResp(
				usecase_parser_test.TX_FAILED_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESP,
			)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_FAILED_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_MSG_WITHDRAW_VALIDATOR_COMMISSION_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(2))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgWithdrawDelegatorReward(
					event.MsgCommonParams{
						BlockHeight: int64(804969),
						TxHash:      "CC5EE77B6CBCEA4DF26F5AC8FA06BA893D018602F03A09E9E02B8417B12C46ED",
						TxSuccess:   false,
						MsgIndex:    0,
					},
					model.MsgWithdrawDelegatorRewardParams{
						DelegatorAddress: "tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel",
						ValidatorAddress: "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu",
						RecipientAddress: "tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel",
						Amount:           coin.NewEmptyCoins(),
					},
				),
				command_usecase.NewCreateMsgWithdrawValidatorCommission(
					event.MsgCommonParams{
						BlockHeight: int64(804969),
						TxHash:      "CC5EE77B6CBCEA4DF26F5AC8FA06BA893D018602F03A09E9E02B8417B12C46ED",
						TxSuccess:   false,
						MsgIndex:    1,
					},
					model.MsgWithdrawValidatorCommissionParams{
						ValidatorAddress: "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu",
						RecipientAddress: "",
						Amount:           coin.NewEmptyCoins(),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel", "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu"}))

		})

		It("should parse Msg commands when there is no reward withdraw in the MsgWithdrawDelegatorReward", func() {
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgWithdrawDelegatorReward(
					event.MsgCommonParams{
						BlockHeight: int64(435599),
						TxHash:      "3643A4CA41EC52BCB5B10DB32EC9867B2FA6B6A7C48B4DE9D45E6EDBC39B31B5",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgWithdrawDelegatorRewardParams{
						DelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						ValidatorAddress: "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek",
						RecipientAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Amount:           coin.NewEmptyCoins(),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"}))
		})
	})
})
