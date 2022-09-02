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
	Describe("MsgGrantSendGrant", func() {
		It("should parse Msg commands when there is MsgGrant in the transaction", func() {
			expected := `{
            "name": "/cosmos.authz.v1beta1.MsgGrant.Created",
            "version": 1,
            "height": 128465,
            "uuid": "{UUID}",
            "msgName": "/cosmos.authz.v1beta1.MsgGrant",
            "txHash": "928E45A1D77FD01EA4EA8A3A20A19D0A69F5AA5A259D8AB5D956FF0BF6811034",
            "msgIndex": 0,
            "params": {
                "maybeSendGrant": {
                    "@type": "/cosmos.authz.v1beta1.MsgGrant",
                    "granter": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
                    "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
                    "grant": {
                        "authorization": {
                            "spendLimit": [
                                {
                                    "denom": "basetcro",
                                    "amount": "400000000"
                                }
                            ]
                        },
                        "expiration": "2022-08-27T03:14:35Z"
                    }
                },
                "maybeStakeGrant": null,
                "maybeGenericGrant": null
            }
        }`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_SEND_GRANT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_SEND_GRANT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_GRANT_SEND_GRANT_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/cosmos.authz.v1beta1.MsgGrant.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgGrantEvent := untypedEvent.(*event.MsgGrant)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgGrantEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgGrantEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9"}))
		})
	})
})
