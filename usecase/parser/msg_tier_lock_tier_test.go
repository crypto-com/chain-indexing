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
	Describe("MsgTierLockTier", func() {
		It("should parse tieredrewards.MsgLockTier command in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_TIER_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_TIER_BLOCK_RESULTS_RESP,
				&tendermint.Base64BlockResultEventAttributeDecoder{},
			)
			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_TIER_LOCK_TIER_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm, block.Height, blockResults, txs, "cro", "basecro",
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgTierLockTier(
					event.MsgCommonParams{
						BlockHeight: int64(100000),
						TxHash:      "A1B2C3D4E5F6A7B8C9D0E1F2A3B4C5D6E7F8A9B0C1D2E3F4A5B6C7D8E9F0A1B2",
						TxSuccess:   true,
						MsgIndex:    0,
						MsgVersion:  tmcosmosutils.CosmosAPIVersionV1,
					},
					model.MsgLockTierParams{
						RawMsgLockTier: model.RawMsgLockTier{
							Type:                   "/chainmain.tieredrewards.v1.MsgLockTier",
							Owner:                  "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84",
							Id:                     1,
							Amount:                 "100000basecro",
							ValidatorAddress:       "crocncl1kkqxv3szgh099xezt7y38t5anqzue4s3fhp2tm",
							TriggerExitImmediately: true,
						},
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw84"}))
		})
	})
})
