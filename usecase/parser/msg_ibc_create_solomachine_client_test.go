package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCCreateSolomachineClient", func() {
		It("should parse Msg commands when there is MsgCreateClient in the transaction", func() {
			expected := `{
            "name": "/ibc.core.client.v1.MsgCreateClient.Created",
            "version": 1,
            "height": 14791,
            "uuid": "{UUID}",
            "msgName": "/ibc.core.client.v1.MsgCreateClient",
            "txHash": "0DB632201805BC399035F0B0CD0E1DABC061E5209D8A709EC0A2B29B1A306BA5",
            "msgIndex": 0,
            "params": {
                "maybeTendermintLightClient": null,
                "maybeSoloMachineLightClient": {
                    "clientState": {
                        "@type": "/ibc.lightclients.solomachine.v2.ClientState",
                        "sequence": 1,
                        "isFrozen": false,
                        "allowUpdateAfterProposal": false
                    },
                    "consensusState": {
                        "publicKey": {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "Ausp3XCdeFAWXHNeTBb3JZY56fMncZ/tpRnGfSFVdESm"
                        },
                        "diversifier": "solo-machine-diversifier",
                        "timestamp": 1629427754
                    }
                },
                "maybeLocalhostLightClient": null,
                "signer": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
                "clientId": "06-solomachine-0",
                "clientType": "06-solomachine"
            }
        }`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_SOLOMACHINE_CLIENT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CREATE_SOLOMACHINE_CLIENT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CREATE_SOLOMACHINE_CLIENT_TXS_RESP)
			txs := []model.Tx{*tx}

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

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
			Expect(cmd.Name()).To(Equal("/ibc.core.client.v1.MsgCreateClient.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgCreateClientEvent := untypedEvent.(*event.MsgIBCCreateClient)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgCreateClientEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgCreateClientEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy"}))
		})
	})
})
