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
	Describe("MsgWithdrawDelegatorReward and MsgWithdrawValidatorCommission", func() {
		It("should parse Msg commands when there is distribution.MsgWithdrawDelegatorReward and MsgWithdrawValidatorCommission in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _ := mustParseBlockResp(
				usecase_parser_test.TX_MSGS_WITHDRAW_DELEGATOR_REWARD_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESP,
			)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSGS_WITHDRAW_DELEGATOR_REWARD_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESULTS_RESP,
			)

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
						Amount:           coin.MustNewCoinFromString("33934701990"),
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
						Amount:           coin.MustNewCoinFromString("4161370358"),
					},
				),
			}))
		})

		It("should parse Msg commands when there is no reward withdraw in the MsgWithdrawDelegatorReward", func() {
			txDecoder := parser.NewTxDecoder("basetrcro")
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_WITHDRAW_DELEGATOR_REWARD_NO_REWARD_BLOCK_RESULTS_RESP,
			))

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
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
						Amount:           coin.Zero(),
					},
				),
			}))
		})
	})
})
