package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgGrantBasicAllowance", func() {
		It("should parse Msg commands when there is MsgGrantAllowance in the transaction", func() {
			expected := `{
            "name": "/cosmos.feegrant.v1beta1.MsgGrantAllowance.Created",
            "version": 1,
            "height": 124056,
            "uuid": "{UUID}",
            "msgName": "/cosmos.feegrant.v1beta1.MsgGrantAllowance",
            "txHash": "1798B9B2694B891BF275DC79DF0C79FDF426D41BA498685C82A284A88207E36C",
            "msgIndex": 0,
            "params": {
                "maybeBasicAllowance": {
                    "@type": "/cosmos.feegrant.v1beta1.MsgGrantAllowance",
                    "granter": "tcro16u0wuyhc73tdw0m7qt3cmfeg68ug8g4hc4verc",
                    "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
                    "allowance": {
                        "@type": "/cosmos.feegrant.v1beta1.BasicAllowance",
                        "spendLimit": [],
                        "expiration": ""
                    }
                },
                "maybePeriodicAllowance": null,
                "maybeAllowedMsgAllowance": null
            }
        }`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_BASIC_ALLOWANCE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_BASIC_ALLOWANCE_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_GRANT_BASIC_ALLOWANCE_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/cosmos.feegrant.v1beta1.MsgGrantAllowance.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgGrantAllowanceEvent := untypedEvent.(*event.MsgGrantAllowance)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgGrantAllowanceEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgGrantAllowanceEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro16u0wuyhc73tdw0m7qt3cmfeg68ug8g4hc4verc"}))
		})
	})
})
