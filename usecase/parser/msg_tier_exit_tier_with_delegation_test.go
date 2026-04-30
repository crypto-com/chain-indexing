package parser_test

import (
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgTierExitTierWithDelegation", func() {
		It("should parse tieredrewards.MsgExitTierWithDelegation command in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_TIER_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_TIER_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
			)
			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_TIER_EXIT_TIER_WITH_DELEGATION_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm, block.Height, blockResults, txs, "cro", "basecro",
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgTierExitTierWithDelegation(
					event.MsgCommonParams{
						BlockHeight: int64(100000),
						TxHash:      "E1F2A3B4C5D6E7F8A9B0C1D2E3F4A5B6C7D8E9F0A1B2C3D4E5F6A7B8C9D0E1F2",
						TxSuccess:   true,
						MsgIndex:    0,
						MsgVersion:  tmcosmosutils.CosmosAPIVersionV1,
					},
					model.MsgExitTierWithDelegationParams{
						RawMsgExitTierWithDelegation: model.RawMsgExitTierWithDelegation{
							Type:       "/chainmain.tieredrewards.v1.MsgExitTierWithDelegation",
							Owner:      "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
							PositionId: "789",
							Amount:     "30000basecro",
						},
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84"}))
		})
	})
})
