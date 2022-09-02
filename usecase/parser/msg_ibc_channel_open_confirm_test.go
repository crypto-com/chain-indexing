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
	Describe("MsgIBCChannelOpenConfirm", func() {
		It("should parse Msg commands when there is MsgChannelOpenConfirm in the transaction", func() {
			expected := `{
  "name": "/ibc.core.channel.v1.MsgChannelOpenConfirm.Created",
  "version": 1,
  "height": 20,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgChannelOpenConfirm",
  "txHash": "B5A78071FCC88BC2C10B0BD273E494367F8AA02AAA81773CBCA1DE4AA5A300A2",
  "msgIndex": 1,
  "params": {
	"portId": "transfer",
	"channelId": "channel-0",
	"proofAck": "CqcCCqQCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASMggDEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0wKgdpY3MyMC0xGgsIARgBIAEqAwACJCIrCAESBAIEJCAaISDDJbFDoMhjrfxsIL/UJIqSvmLPSt2ZctIzz6MHJoNY2CIrCAESBAQGJCAaISCImllkhkBTWazxA+Jbdmu6F0R4KEaa1c2mj5lacw3vgiIrCAESBAgQJCAaISD43xC1gJnUXhb8IHs8QqSgGlDixdSTsmEVREenZwjYSyIrCAESBAokJCAaISAXREhZL0sssIbwcKvpkVHTJ7vYuBAeAI7Jlj9WJq+ptArTAQrQAQoDaWJjEiAKVR+z19yBM4u74+hFtqRTEf5OMzGBuV0Erx6cMv6oghoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAf5bNtKbIj3TPLHsdxLDED53UgSUyQcDs3ycK8WvhKYNIiUIARIhAX9+sHvlA9MAeSOWPvQStTa/FNxe/ZFVdDVeiC8m3pqAIicIARIBARogB2CER0RiIRwqvNxiFAdL12xJmmt/Y/18HwrQ/ULquEQ=",
	"proofHeight": {
	  "revisionNumber": "1",
	  "revisionHeight": "19"
	},
	"signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",

    "counterpartyChannelId": "channel-0",
    "counterpartyPortId": "transfer",
    "connectionId": "connection-0"
  }
}`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_CONFIRM_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_CHANNEL_OPEN_CONFIRM_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_CHANNEL_OPEN_CONFIRM_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgChannelOpenConfirm.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCChannelOpenConfirm)

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
