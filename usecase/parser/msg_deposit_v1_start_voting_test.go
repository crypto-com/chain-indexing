package parser_test

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("v1.MsgDeposit", func() {
		It("should parse gov.v1.MsgDeposit command with effective height in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_DEPOSIT_V1_AND_START_VOTING_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_DEPOSIT_V1_AND_START_VOTING_BLOCK_RESULT_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_DEPOSIT_V1_AND_START_VOTING_TXS_RESP)
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
			Expect(cmds).To(HaveLen(2))

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgDepositV1(
					event.MsgCommonParams{
						BlockHeight: int64(5691),
						TxHash:      "D41222A53DDCF690879AF8A681B0EE55991EC8EA204533223A10706DD945D7AA",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					v1_model.MsgDepositParams{
						ProposalId: "5",
						Depositor:  "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
						Amount:     coin.MustParseCoinsNormalized("100000basecro"),
					},
				),
				command_usecase.NewStartProposalVotingPeriod(int64(5691), "5"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"}))
		})
	})
})
