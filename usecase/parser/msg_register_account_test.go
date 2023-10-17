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
	Describe("MsgRegisterAccount", func() {
		It("should parse Msg commands when there is MsgRegisterAccount in the transaction", func() {
			expected := `{
  "name": "/chainmain.icaauth.v1.MsgRegisterAccount.Created",
  "version": 1,
  "height": 67461,
  "uuid": "{UUID}",
  "msgName": "/chainmain.icaauth.v1.MsgRegisterAccount",
  "txHash": "5FC4ECBEB0FFD3809F62052CADA161E542BAA1198EE94CCD0DF197EE4A885F35",
  "msgIndex": 0,
  "params": {
	"owner": "tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
    "connectionId": "connection-18",
	"version": "",
    "portId": "icacontroller-tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra",
    "channelId": "channel-48",
    "counterpartyPortId": "icahost",
    "counterpartyChannelId": ""
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REGISTER_ACCOUNT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_REGISTER_ACCOUNT_BLOCK_RESULTS_RESP,
			), &tendermint.Base64BlockResultEventAttributeDecoder{})

			tx := MustParseTxsResp(usecase_parser_test.TX_MSG_REGISTER_ACCOUNT_TXS_RESP)
			txs := []model.CosmosTxWithHash{*tx}

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
			Expect(cmd.Name()).To(Equal("/chainmain.icaauth.v1.MsgRegisterAccount.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgRegisterAccount)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[0]).To(Equal("tcro1np7ztcfeycqwhj0nr8hxfu0lfjz27telqx53ra"))
		})
	})
})
