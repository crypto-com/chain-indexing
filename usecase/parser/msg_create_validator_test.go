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
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_BLOCK_RESP))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_BLOCK_RESULTS_RESP))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CREATE_VALIDATOR_TXS_RESP)
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
			thiscmd := cmds[0]
			Expect(thiscmd.Name()).To(Equal("/cosmos.staking.v1beta1.MsgCreateValidator.Create"))

			description := model.ValidatorDescription{
				Moniker:         "leo-node",
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
					BlockHeight: int64(76550),
					TxHash:      "1FE830F23A3C542547700AAB3D0E5106A0131B393260910F63EE3B5542E281EF",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgCreateValidatorParams{
					Description:       description,
					Commission:        commissionRates,
					MinSelfDelegation: "1",
					DelegatorAddress:  "tcro109ww3ss92v4vsaq470vvgw528mtqp98mq0vvp9",
					ValidatorAddress:  "tcrocncl109ww3ss92v4vsaq470vvgw528mtqp98m4s04ex",
					TendermintPubkey:  "Kpox5fS2po0sJUHmzllExuJ4uZ5nm0bbCp6UQKESsnE=",
					Amount:            coin.MustParseCoinNormalized("10000000000000basetcro"),
				},
			)))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro109ww3ss92v4vsaq470vvgw528mtqp98mq0vvp9", "tcrocncl109ww3ss92v4vsaq470vvgw528mtqp98m4s04ex"}))
		})
	})
})
