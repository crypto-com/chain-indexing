package parser_test

import (
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
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
	Describe("MsgSetSendEnabled", func() {

		It("should parse Msg commands when there is bank.MsgSetSendEnabled in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SET_SEND_ENABLED_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_SET_SEND_ENABlED_BLOCK_RESULTS_RESP, &tendermint.Base64BlockResultEventAttributeDecoder{})

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_SET_SEND_ENABLED_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSetSendEnabled(
				event.MsgCommonParams{
					BlockHeight: int64(377673),
					TxHash:      "2A2A64A310B3D0E84C9831F4353E188A6E63BF451975C859DF40C54047AC6324",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgSetSendEnabled{
					Authority: "tcrc10d07y265gmmuvt4z0w9aw880jnsr700jrfqhyu",
					SendEnabled: []model.SendEnabled{
						{
							Denom:   "stake",
							Enabled: true,
						},
						{
							Denom:   "basetcro",
							Enabled: false,
						},
					},
					UseDefaultFor: []string{},
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcrc10d07y265gmmuvt4z0w9aw880jnsr700jrfqhyu"}))
		})

	})
})
