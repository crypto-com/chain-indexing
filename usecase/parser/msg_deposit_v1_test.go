package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("v1.MsgDeposit", func() {
		It("should parse gov.v1.MsgDeposit command with effective height in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_DEPOSIT_V1_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_DEPOSIT_V1_BLOCK_RESULTS_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_DEPOSIT_V1_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

			accountAddressPrefix := "crc"
			bondingDenom := "basecro"

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
				command_usecase.NewCreateMsgDepositV1(
					event.MsgCommonParams{
						BlockHeight: int64(5066),
						TxHash:      "8F549DCBED33FF8909C83C6FE957D262F833962ACFE70C53E969DE8F1D657B7C",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model_gov_v1.MsgDepositParams{
						ProposalId: "3",
						Depositor:  "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
						Amount:     coin.MustParseCoinsNormalized("100000basecro"),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"}))

		})
	})
})
