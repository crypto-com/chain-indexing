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
	Describe("MsgIBCConnectionOpenConfirm", func() {
		It("should parse Msg commands when there is MsgIBCConnectionOpenAck in the transaction", func() {
			expected := `{
  "name": "/ibc.core.connection.v1.MsgConnectionOpenConfirm.Created",
  "version": 1,
  "height": 13,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.connection.v1.MsgConnectionOpenConfirm",
  "txHash": "745403DE00B4A8D112B3F4D2546A1D675FAED705E0123465A6734E8B1B7D869A",
  "msgIndex": 1,
  "params": {
	"connectionId": "connection-0",
	"proofAck": "CrwCCrkCChhjb25uZWN0aW9ucy9jb25uZWN0aW9uLTASYAoPMDctdGVuZGVybWludC0wEiMKATESDU9SREVSX09SREVSRUQSD09SREVSX1VOT1JERVJFRBgDIiYKDzA3LXRlbmRlcm1pbnQtMBIMY29ubmVjdGlvbi0wGgUKA2liYxoLCAEYASABKgMAAhYiKwgBEgQCBBYgGiEgoL0fcjA4bhOmD83Ka3zIyWUkSNa31p5h64zB+UpzjMUiKwgBEgQECBYgGiEge8ROC6gJ/hvxo0nddRPSKTbuIWiteu7Fbhs59h7ImF8iKQgBEiUGDhYgV421a0/wlqTvyG2E9dyQAqOGy0gV3DwNOBpJ0YluyPYgIikIARIlCBgWILhLJzoETx8Ak2Xp4wuB75gso6yrXyn1M/DsizNa1JTbIArTAQrQAQoDaWJjEiAO3+Pnp8M0UIwSpChoUC35Bc39UWoKO58iK4AdECFfaxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAdCcOLwBd3qyMTRIUXZvtvDINpEb8SM8Pht3OxmrabIIIiUIARIhAU4hRda6BIRjBSBxdB+U4LUrwyCjdlk1HhOhdGwMgM4pIicIARIBARogVcvkOvJpiz0tN0H450uAl9LNEWX8wQmU9PavrXqJ2hU=",
	"proofHeight": {
	  "revisionNumber": "1",
	  "revisionHeight": "12"
	},
	"signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",

    "clientId": "07-tendermint-0",
    "counterpartyClientId": "07-tendermint-0",
    "counterpartyConnectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_CONFIRM_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CONNECTION_OPEN_CONFIRM_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CONNECTION_OPEN_CONFIRM_TXS_RESP)
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
			Expect(cmds).To(HaveLen(2))
			cmd := cmds[1]
			Expect(cmd.Name()).To(Equal("/ibc.core.connection.v1.MsgConnectionOpenConfirm.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCConnectionOpenConfirm)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"))
		})
	})
})
