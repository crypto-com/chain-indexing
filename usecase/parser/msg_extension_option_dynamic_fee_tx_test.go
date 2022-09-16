package parser_test

import (
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
	Describe("ExtensionOptionDynamicFeeTx", func() {
		It("should parse Msg commands when there is ethermint.ExtensionOptionDynamicFeeTx in the transaction", func() {
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TX_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TX_BLOCK_RESULTS_RESP,
			)

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TXS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgEthereumTx(
				event.MsgCommonParams{
					BlockHeight: int64(5168311),
					TxHash:      "3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.DynamicFeeTxParams{
					RawDynamicFeeTx: model.RawDynamicFeeTx{
						Type: "/ethermint.evm.v1.MsgEthereumTx",
						Size: 0,
						Data: model.DynamicFeeTx{
							Type:      "/ethermint.evm.v1.DynamicFeeTx",
							ChainId:   "338",
							Nonce:     "115",
							GasTipCap: "5000000000000",
							GasFeeCap: "5000000000000",
							Gas:       "48834",
							To:        "0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a",
							Value:     "0",
							Data:      "CV6nswAAAAAAAAAAAAAAAPVk7wA0u3182EQSJ14eUTZNTXc0//////////////////////////////////////////8=",
							V:         "",
							R:         "HHigR02IN48gvEM991C/MOCmbVfOw9Sj804IhtI5z2Q=",
							S:         "SspHmiOAoM4bQdvAig9I8PQzLv5GTzElhcROppO0WpU=",
						},
						From: "",
						Hash: "0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53",
					},
				},
			)}))
			var emptyAddress []string
			Expect(possibleSignerAddresses).To(Equal(emptyAddress))
		})
	})
})
