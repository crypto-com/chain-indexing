package parser_test

import (
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSubmitProopsal", func() {
		It("should parse Msg commands when there is gov.MsgSoftwareUpgrade in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SOFTWARE_UPGRADE_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSubmitProposal(
				event.MsgCommonParams{
					BlockHeight: int64(1634),
					TxHash:      "BFEDD454DED949E0CD349BBFD8F518AED187214A69630445CDDEBF924A48F83C",
					TxSuccess:   true,
					MsgIndex:    0,
				},

				model_gov_v1.MsgSubmitProposalParams{
					MaybeProposalId: primptr.String("3"),
					Proposer:        "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
					InitialDeposit: coin.Coins{
						{
							Denom:  "basecro",
							Amount: coin.NewInt(10000),
						},
					},
					Metadata: "ipfs://CID",
					Messages: []interface{}{
						map[string]interface{}{
							"@type":     "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
							"authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
							"plan": map[string]interface{}{
								"name":                  "name",
								"time":                  "0001-01-01T00:00:00Z",
								"height":                "1000",
								"info":                  "info",
								"upgraded_client_state": nil,
							},
						},
					},
				},
			),
				command_usecase.NewStartProposalVotingPeriod(int64(1634), "3"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"}))
		})

		It("should parse Msg commands when there is gov.MsgCancelUpgrade in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_CANCEL_UPGRADE_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_CANCEL_UPGRADE_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
			)

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_CANCEL_UPGRADE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSubmitProposal(
				event.MsgCommonParams{
					BlockHeight: int64(2171),
					TxHash:      "7360D83D7C9FB1B04D73546757200A375C46ECC1B591F77E1E2BA2666BACD710",
					TxSuccess:   true,
					MsgIndex:    0,
				},

				model_gov_v1.MsgSubmitProposalParams{
					MaybeProposalId: primptr.String("4"),
					Proposer:        "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
					InitialDeposit: coin.Coins{
						{
							Denom:  "basecro",
							Amount: coin.NewInt(1000),
						},
					},
					Metadata: "ipfs://CID",
					Messages: []interface{}{
						map[string]interface{}{
							"@type":     "/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
							"authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
						},
					},
				},
			),
				command_usecase.NewStartProposalVotingPeriod(int64(2171), "4"),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp"}))
		})

		It("should parse Msg commands when there is gov.SoftwareUpgradeProposal in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_EXEC_LEGACY_CONTENT_V1_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSubmitProposal(
				event.MsgCommonParams{
					BlockHeight: int64(6580),
					TxHash:      "DBAFA8C7C7F3A39C8162E2463E560822554A89A79DAB550882270125902AF39C",
					TxSuccess:   true,
					MsgIndex:    0,
				},

				model_gov_v1.MsgSubmitProposalParams{
					MaybeProposalId: primptr.String("6"),
					Proposer:        "crc12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
					InitialDeposit: coin.Coins{
						{
							Denom:  "basecro",
							Amount: coin.NewInt(1000000),
						},
					},
					Metadata: "ipfs://CID",
					Messages: []interface{}{
						map[string]interface{}{
							"@type": "/cosmos.gov.v1.MsgExecLegacyContent",
							"content": map[string]interface{}{
								"@type":       "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
								"title":       "title",
								"description": "description",
								"plan": map[string]interface{}{
									"name":                  "test",
									"time":                  "0001-01-01T00:00:00Z",
									"height":                "10000",
									"info":                  "info",
									"upgraded_client_state": nil,
								},
							},
							"authority": "crc10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd",
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
