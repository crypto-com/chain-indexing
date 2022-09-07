package parser_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgExec", func() {
		It("should parse Msg commands when there is MsgEthereumTx in the transaction", func() {
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
			Expect(cmd.Name()).To(Equal("/ethermint.evm.v1.MsgEthereumTx.Create"))

			Expect(cmd).To(Equal(command_usecase.NewCreateMsgEthereumTx(
				event.MsgCommonParams{
					BlockHeight: int64(83178),
					TxHash:      "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgEthereumTxParams{
					RawMsgEthereumTx: model.RawMsgEthereumTx{

						Type: "/ethermint.evm.v1.MsgEthereumTx",
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
	})
})
