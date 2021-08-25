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
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgCreateValidator", func() {

		It("should parse Msg commands when there is staking.MsgCreateValidator in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_V0_43_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_V0_43_BLOCK_RESULTS_RESP))
			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			thiscmd := cmds[0]
			Expect(thiscmd.Name()).To(Equal("CreateMsgCreateValidator"))

			description := model.ValidatorDescription{
				Moniker:         "c4_node_testing",
				Identity:        "",
				Website:         "",
				SecurityContact: "",
				Details:         "",
			}
			commissionRates := model.ValidatorCommission{
				Rate:          "0.100000000000000000",
				MaxRate:       "0.200000000000000000",
				MaxChangeRate: "0.010000000000000000",
			}

			Expect(thiscmd).To(Equal(command_usecase.NewCreateMsgCreateValidator(
				event.MsgCommonParams{
					BlockHeight: int64(3106),
					TxHash:      "2F5390FB7F7570D463993A2A97CABCBC764286FA2CE052721D8771AA8B8D3853",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description:       description,
					Commission:        commissionRates,
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro1k4nz2yalsuzuzxexcxuughp33cfqls7e6utarx",
					ValidatorAddress:  "tcrocncl1k4nz2yalsuzuzxexcxuughp33cfqls7e0rgym9",
					TendermintPubkey:  "hTpJIeVSSo1PNU87B8y1vMvpyv3ShMmu2I0vYON1w2o=",
					Amount:            coin.MustParseCoinNormalized("100000000000basetcro"),
				},
			)))
		})
	})
})
