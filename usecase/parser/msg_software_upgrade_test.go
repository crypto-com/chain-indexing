package parser_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSoftwareUpgrade", func() {
		It("should parse Msg commands when there is gov.MsgSoftwareUpgrade in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESULTS_RESP,
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_SOFTWARE_UPGRADE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSoftwareUpgrade(
				event.MsgCommonParams{
					BlockHeight: int64(1634),
					TxHash:      "BFEDD454DED949E0CD349BBFD8F518AED187214A69630445CDDEBF924A48F83C",
					TxSuccess:   true,
					MsgIndex:    0,
				},

				v1_model.MsgSoftwareUpgradeParams{
					MaybeProposalId: primptr.String("3"),
					Proposer:        "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
					InitialDeposit: coin.Coins{
						{
							Denom:  "basecro",
							Amount: coin.NewInt(10000),
						},
					},
					Metadata:  "ipfs://CID",
					Authority: "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
					Plan: v1_model.MsgSoftwareUpgradePlan{
						Name:   "name",
						Time:   utctime.MustParse(time.RFC3339, "0001-01-01T00:00:00Z"),
						Height: int64(1000),
						Info:   "info",
					},
				},
			),
				command_usecase.NewStartProposalVotingPeriod(int64(1634), "3"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"}))
		})
	})
})
