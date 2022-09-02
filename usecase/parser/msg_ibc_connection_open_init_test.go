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
	Describe("MsgIBCConnectionOpenInit", func() {
		It("should parse Msg commands when there is MsgIBCConnectionOpenInit in the transaction", func() {
			expected := `{
  "name": "/ibc.core.connection.v1.MsgConnectionOpenInit.Created",
  "version": 1,
  "height": 6,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.connection.v1.MsgConnectionOpenInit",
  "txHash": "F2B7D61BA783E6CDD9FE5825EBF7770688F6F45C482CB78ACB51E84B06FC643E",
  "msgIndex": 0,
  "params": {
    "clientId": "07-tendermint-0",
    "counterparty": {
  	  "clientId": "07-tendermint-0",
  	  "connectionId": "",
  	  "prefix": { "keyPrefix": "aWJj" }
    },
    "version": {
  	  "identifier": "1",
  	  "features": ["ORDER_ORDERED", "ORDER_UNORDERED"]
    },
    "delayPeriod": "0",
    "signer": "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f",
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CONNECTION_OPEN_INIT_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.connection.v1.MsgConnectionOpenInit.Create"))

			untypedEvent, _ := cmd.Exec()
			createMsgCreateClientEvent := untypedEvent.(*event.MsgIBCConnectionOpenInit)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(createMsgCreateClientEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					createMsgCreateClientEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses).To(Equal([]string{"cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f"}))
		})
	})
})
