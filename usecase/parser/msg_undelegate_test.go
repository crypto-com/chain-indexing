package parser_test

import (
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgDelegate", func() {
		It("should parse Msg commands when there is staking.MsgUndelegate in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_UNDELEGATE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_UNDELEGATE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_UNDELEGATE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgUndelegate(
				event.MsgCommonParams{
					BlockHeight: int64(374371),
					TxHash:      "0F525EFC1DD9C319E9036C35CF1656E09480B308301BB3A46F850AE482A3875C",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgUndelegateParams{
					DelegatorAddress:      "tcro1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5w2ll64",
					ValidatorAddress:      "tcrocncl1j7pej8kplem4wt50p4hfvndhuw5jprxxxtenvr",
					Amount:                coin.MustParseCoinNormalized("1000000000basetcro"),
					MaybeUnbondCompleteAt: primptr.UTCTime(utctime.FromUnixNano(int64(1605152654000000000))),
					AutoClaimedRewards:    coin.MustNewCoin("basetcro", coin.NewInt(76688)),
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5w2ll64"}))
		})

		It("should parse MsgUndelegate command in failed transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_FAILED_MSG_UNDELEGATE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_FAILED_MSG_UNDELEGATE_BLOCK_RESULTS_RESP)

			tx := mustParseTxsResp(usecase_parser_test.TX_FAILED_MSG_UNDELEGATE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgUndelegate(
				event.MsgCommonParams{
					BlockHeight: int64(184399),
					TxHash:      "5285A9B475157E01540536299A2B5F505AC900159C268B3D90652557F9ACDE1E",
					TxSuccess:   false,
					MsgIndex:    0,
				},
				model.MsgUndelegateParams{
					DelegatorAddress:      "tcro1llst0cguh5azl9t8wr6mz5yzjuwukz7f67z7f6",
					ValidatorAddress:      "tcrocncl15e69kdrtczajjdlzyt2qgs5q2anc5qpmk2c68z",
					Amount:                coin.MustParseCoinNormalized("50000000000000basetcro"),
					MaybeUnbondCompleteAt: nil,
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1llst0cguh5azl9t8wr6mz5yzjuwukz7f67z7f6"}))
		})
	})
})
