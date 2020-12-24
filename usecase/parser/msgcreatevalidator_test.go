package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgCreateValidator", func() {

		It("should parse Msg commands when there is staking.MsgCreateValidator in the transaction", func() {
			txDecoder := parser.NewTxDecoder("basetcro")
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_BLOCK_RESULTS_RESP))

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(5))
			thiscmd := cmds[2]
			Expect(thiscmd.Name()).To(Equal("CreateMsgCreateValidator"))

			description := model.MsgValidatorDescription{
				Moniker:         "Calvin Test Node",
				Identity:        "",
				Website:         "",
				SecurityContact: "",
				Details:         "",
			}
			commissionRates := model.MsgValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			}

			Expect(thiscmd).To(Equal(command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: int64(503978),
					TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description:       description,
					Commission:        commissionRates,
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					Pubkey:            "tcrocnclconspub1zcjduepqa5rksn4ds9u6jmmg4n86d9wct7wmj23pyqe6p7e252lffzqsgcvqxm5lc2",
					Amount:            coin.MustNewCoinFromString("10"),
				},
			)))
		})
	})
})
