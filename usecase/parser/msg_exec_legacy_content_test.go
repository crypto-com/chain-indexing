package parser_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgExecLegacyContent", func() {
		It("should parse Msg commands when there is gov.SoftwareUpgradeProposal in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESULTS_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_EXEC_LEGACY_CONTENT_V1_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSubmitSoftwareUpgradeProposal(
				event.MsgCommonParams{
					BlockHeight: int64(6580),
					TxHash:      "DBAFA8C7C7F3A39C8162E2463E560822554A89A79DAB550882270125902AF39C",
					TxSuccess:   true,
					MsgIndex:    0,
				},

				model.MsgSubmitSoftwareUpgradeProposalParams{
					MaybeProposalId: primptr.String("6"),
					ProposerAddress: "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
					InitialDeposit: coin.Coins{
						{
							Denom:  "basecro",
							Amount: coin.NewInt(1000000),
						},
					},
					Content: model.MsgSubmitSoftwareUpgradeProposalContent{
						Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
						Title:       "title",
						Description: "description",
						Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
							Name:   "test",
							Time:   utctime.MustParse(time.RFC3339, "0001-01-01T00:00:00Z"),
							Height: int64(10000),
							Info:   "info",
						},
					},
				},
			),
				command_usecase.NewStartProposalVotingPeriod(int64(6580), "6"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"}))
		})
	})
})
