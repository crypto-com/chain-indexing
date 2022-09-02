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
	Describe("MsgRevokeAllowance", func() {
		It("should parse Msg commands when there is MsgRevokeAllowance in the transaction", func() {
			expected := `{
            "name": "/cosmos.feegrant.v1beta1.MsgRevokeAllowance.Created",
            "version": 1,
            "height": 128553,
            "uuid": "{UUID}",
            "msgName": "/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
            "txHash": "002579A793A5ABD82FAF819FC77CC4FF765C550332151B4BA9A811D662ABD027",
            "msgIndex": 0,
            "params": {
                "@type": "/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
                "granter": "tcro16u0wuyhc73tdw0m7qt3cmfeg68ug8g4hc4verc",
                "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2"
            }
        }`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REVOKE_ALLOWANCE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REVOKE_ALLOWANCE_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_REVOKE_ALLOWANCE_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/cosmos.feegrant.v1beta1.MsgRevokeAllowance.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgRevokeAllowanceEvent := untypedEvent.(*event.MsgRevokeAllowance)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgRevokeAllowanceEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgRevokeAllowanceEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro16u0wuyhc73tdw0m7qt3cmfeg68ug8g4hc4verc"}))
		})
	})
})
