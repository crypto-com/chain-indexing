package parser_test

import (
	"regexp"
	"strings"

	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

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
  "name": "MsgConnectionOpenInitCreated",
  "version": 1,
  "height": 6,
  "uuid": "{UUID}",
  "msgName": "MsgIBCConnectionOpenInit",
  "txHash": "F2B7D61BA783E6CDD9FE5825EBF7770688F6F45C482CB78ACB51E84B06FC643E",
  "msgIndex": 0,
  "clientId": "07-tendermint-0",
  "counterparty": {
    "clientId": "07-tendermint-0",
    "connectionId": "",
    "prefix": { "keyPrefix": "YVdKag==" }
  },
  "connectionVersion": {
    "identifier": "1",
    "features": ["ORDER_ORDERED", "ORDER_UNORDERED"]
  },
  "delayPeriod": "0",
  "signer": "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f",
  "connectionId": "connection-0"
}`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_INIT_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"
			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgConnectionOpenInit"))

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
		})
	})
})
