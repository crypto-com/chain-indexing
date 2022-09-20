package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgExec", func() {
		It("should parse Msg commands when there is EthermintLegacyTx in the transaction", func() {
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ETHEREUM_TX_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_ETHEREUM_TX_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_ETHEREUM_TX_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "tcro"
			stakingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				block.Height,
				blockResults,
				txs,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("/ethermint.evm.v1.EthermintLegacyTx.Create"))

			Expect(cmd).To(Equal(command_usecase.NewCreateLegacyTx(
				event.MsgCommonParams{
					BlockHeight: int64(83178),
					TxHash:      "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.EthermintLegacyTxParams{
					RawLegacyTx: model.RawLegacyTx{
						Type: "/ethermint.evm.v1.EthermintLegacyTx",
						Size: 208,
						Data: model.LegacyTx{
							Type:     "/ethermint.evm.v1.LegacyTx",
							Nonce:    "130",
							GasPrice: "5000000000000",
							Gas:      "77595",
							To:       "0xAa53Dd6D234A0c431b39B9E90454666432869dc9",
							Value:    "0",
							Data:     "k4YIwgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABphbdB3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGFt0HcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXPXkUYA==",
							V:        "Asc=",
							R:        "GWDX+kHcVNVKp5K2lG+/zHAOJI8yR6lYZ2GW4kYgEhE=",
							S:        "PqawF/sgoCKiOnqN9al9x9AAWOS2uKaW5Dq+cg74Lgg=",
						},
						From: "",
						Hash: "0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff",
					},
				},
			)))
			var emptyAddress []string
			Expect(possibleSignerAddresses).To(Equal(emptyAddress))
		})

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
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("/ethermint.evm.v1.EthermintLegacyTx.Create"))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgExtensionOptionDynamicFeeTxTx(
				event.MsgCommonParams{
					BlockHeight: int64(5168311),
					TxHash:      "3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.EthermintDynamicFeeTxParams{
					RawDynamicFeeTx: model.RawDynamicFeeTx{
						Type: "/ethermint.evm.v1.EthermintLegacyTx",
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
