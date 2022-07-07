package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgRevoke", func() {
		It("should parse Msg commands when there is MsgRevoke in the transaction", func() {
			expected := `{
            "name": "MsgRevokeCreated",
            "version": 1,
            "height": 123731,
            "uuid": "{UUID}",
            "msgName": "MsgRevoke",
            "txHash": "9921D7DDC530DB81B0A5FD1163678757A7B7B7D8ED78C2B4BE433BFFD30C1228",
            "msgIndex": 0,
            "params": {
                "@type": "/cosmos.authz.v1beta1.MsgRevoke",
                "granter": "tcro1vurfhqf0j2jgfpjahlja6g6uq2ts2r60swm2d9",
                "grantee": "tcro15zh5tn7xjdecu4zjclsmlnlht5ead2mx84gau2",
                "msgTypeUrl": "/cosmos.bank.v1beta1.MsgSend"
            }
        }`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REVOKE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REVOKE_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"

			pm := usecase_parser_test.InitParserManager()

			cmds, _, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgRevoke"))

			untypedEvent, _ := cmd.Exec()
			createMsgRevokeEvent := untypedEvent.(*event.MsgRevoke)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgRevokeEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgRevokeEvent.UUID(),
					-1,
				),
			))
		})
	})
})
