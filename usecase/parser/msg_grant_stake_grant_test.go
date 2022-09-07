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
	Describe("MsgGrantStakeGrant", func() {
		It("should parse Msg commands when there is MsgGrant (StakeAuthorization) in the transaction", func() {
			expected := `{
				"name": "/cosmos.authz.v1beta1.MsgGrant.Created",
				"version": 1,
				"height": 170108,
				"uuid": "{UUID}",
				"msgName": "/cosmos.authz.v1beta1.MsgGrant",
				"txHash": "D8AE71B4C05B7A220114F17347D6A66ADBFE75C51279E4541E911284A2BE7E04",
				"msgIndex": 0,
				"params": {
					"maybeSendGrant": null,
					"maybeStakeGrant": {
						"@type": "/cosmos.authz.v1beta1.MsgGrant",
						"granter": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
						"grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
						"grant": {
							"authorization": {
								"maxTokens": { "denom": "basetcro", "amount": "400000000" },
								"allowList": {
									"address": ["tcrocncl163tv59yzgeqcap8lrsa2r4zk580h8ddr5a0sdd"]
								},
								"authorizationType": "AUTHORIZATION_TYPE_DELEGATE"
							},
							"expiration": "2022-08-29T16:41:37Z"
						}
					},
					"maybeGenericGrant": null
				}
			}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_STAKE_GRANT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_GRANT_STAKE_GRANT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_GRANT_STAKE_GRANT_TXS_RESP)
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
