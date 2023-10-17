package parser_test

import (
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
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
	Describe("v1.MsgVoteWeighted", func() {
		It("should parse gov.v1.MsgVoteWeighted command in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_VOTE_WEIGHTED_V1_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_VOTE_WEIGHTED_V1_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_VOTE_WEIGHTED_V1_TXS_RESP)
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
				command_usecase.NewCreateMsgVoteWeightedV1(
					event.MsgCommonParams{
						BlockHeight: int64(5284),
						TxHash:      "F258B015674DAFBE5D1788825690BC31D1F6254F599944AD4CDEE9A461179FF7",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model_gov_v1.MsgVoteWeightedParams{
						ProposalId: "1",
						Voter:      "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7",
						VoteOptions: []model_gov_v1.VoteOption{
							{
								Option: "VOTE_OPTION_YES",
								Weight: "0.200000000000000000",
							},
							{
								Option: "VOTE_OPTION_NO",
								Weight: "0.800000000000000000",
							},
						},
						Metadata: "",
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"}))
		})
	})
})
