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
	Describe("MsgIBCRecvPacket", func() {
		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction", func() {
			expected := `{
  "name": "MsgRecvPacketCreated",
  "version": 1,
  "height": 26,
  "uuid": "{UUID}",
  "msgName": "MsgRecvPacket",
  "txHash": "C94E6D87ACC4DD809CC05B9F9773B32B0ECEF9E11B8DFF85DD8ADF4566AF9ED1",
  "msgIndex": 1,
  "params": {
    "packet": {
      "sequence": "1",
      "sourcePort": "transfer",
      "sourceChannel": "channel-0",
      "destinationPort": "transfer",
      "destinationChannel": "channel-0",
      "data": "eyJhbW91bnQiOiIxMjM0IiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjcm8xZHVsd3FnY2RwZW1uOGMzNHNqZDkyZnhlcHo1cDBzcXBlZXZ3N2YiLCJzZW5kZXIiOiJjcm8xMHNuaGx2a3B1YzR4aHE4MnV5ZzVleDJlZXptbWY1ZWQ1dG1xc3YifQ==",
      "timeoutHeight": {
    	"revisionNumber": "2",
    	"revisionHeight": "1023"
      },
      "timeoutTimestamp": "0"
    },
    "proofCommitment": "CpsCCpgCCjljb21taXRtZW50cy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTAvc2VxdWVuY2VzLzESIOL9pxFSPp6ZejEK+FtCDChW0bJhdl57P/d6zx+hVhgpGgsIARgBIAEqAwACMCIpCAESJQIEMCDj88ZEbojOy3mqsQvLWRQf8tOSvTpdOPsZA8AECkOeziAiKQgBEiUECDAglzLiPEahZdTWY2pN870XzhUMewaTClTsjv3dhk1e6+ggIisIARIECBYwIBohIMytK1VpzSkLoEFzfe4606li/jw6NhJDwdjFKOR/YVqfIikIARIlCiYwIMdfhADZpHkpq3IfEwlxQT8z0N9wSA4ASWUuYkFML8CzIArTAQrQAQoDaWJjEiAbbEgt1HZqlOsM1hT5uoYHf+4r+1wCNMmALauIR+pv+xoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAad0DU0QldmQc2csoMIRyWEWnRpubv1IwniOeASCc9f+IiUIARIhARFr3/h1b3vKOoGq7PWsyM0e7UKp3Wd9nNus5KlIGgiZIicIARIBARogXnlUV86r9evx40joRXDS41kvFZuFFW8KFgSUfQW/uRQ=",
    "proofHeight": {
      "revisionNumber": "1",
      "revisionHeight": "25"
    },
    "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",

    "application": "transfer",
    "messageType": "MsgTransfer",
    "maybeMsgTransfer": {
      "sender": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv",
      "receiver": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
      "denom": "basecro",
      "amount": "1234",
      "success": true,
      "maybeDenominationTrace": {
        "hash": "6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
        "denom": "ibc/6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86"
	  }
    },

    "packetSequence": "1",
    "channelOrdering": "ORDER_UNORDERED",
    "connectionId": "connection-0",
    "packetAck": {
      "result": "AQ==",
			"error": ""
    }
  }
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
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
			Expect(cmds).To(HaveLen(2))
			cmd := cmds[1]
			Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
		})

		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction and MsgIBCRecvPacket.PacketAck.Error is NOT empty string", func() {
			expected := `{
	"name": "MsgRecvPacketCreated",
	"version": 1,
	"height": 130418,
	"uuid": "{UUID}",
	"msgName": "MsgRecvPacket",
	"txHash": "852566591D525E2F3DF54602091D33CEF2D87389E6D9AC260EC3267CF2C4CFCE",
	"msgIndex": 1,
	"params": {
		"packet": {
			"sequence": "15",
			"sourcePort": "transfer",
			"sourceChannel": "channel-0",
			"destinationPort": "transfer",
			"destinationChannel": "channel-3",
			"data": "eyJhbW91bnQiOiIyMDAwMDAwMDI3MDAwNjY5MTIiLCJkZW5vbSI6InRyYW5zZmVyL2NoYW5uZWwtMC9iYXNldGNybyIsInJlY2VpdmVyIjoidGNybzFwZG4ybnNuOXdlc3o2cHgzbGNqc2dtcDhwZWZlZG56cDNnbXAzcSIsInNlbmRlciI6ImV0aDFzcDNlNThsZmduazlwam0ybjdlN2pqbDA4MDVzNWp1MDlhMm11eiJ9",
			"timeoutHeight": { "revisionNumber": "4", "revisionHeight": "131385" },
			"timeoutTimestamp": "1630044922314197883"
		},
		"proofCommitment": "Ct0DCtoDCjpjb21taXRtZW50cy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTAvc2VxdWVuY2VzLzE1EiAERWcc5FyvY7a8djaEv7HQsNcUewiPX2hb/wh5wiTGSBoNCAEYASABKgUAApiSAiIrCAESJwIEmJICIF9GRTXveymFeABx34JOs6v8JP9iLfngiJMP/GdXA+/hICItCAESBgQGmJICIBohIPbTrk6EJ8X5DxhymRoZgwtWIdCWUEDvERLzfc0ZAqsGIisIARInBgqYkgIgfM2RnR7xRSaAk7DlQY4mTL/lzjDKhmmGCM9HpGioKRsgIisIARInCBKYkgIgAM9p06enrxsvF7Z8di0BcflzxWJMPQmy7ms+JffBFDQgIisIARInCiKYkgIggZ7uGJ58VfVwL5IFlJSpabKVA7pDNRK3o7fuIgaLjRcgIisIARInDD6YkgIgB+mqgdJ455512JE83JX4pCOP13pFkrMquvqV1sWdDbggIi0IARIGDlyYkgIgGiEgZPGGmrep9q4a8PX6Ek1Wv83UXGewolZrkszRXnx81EIiLAgBEigQogKYkgIgsEzSj7HxhWH602n7KMIOoKc0SA3umMBF41xrrY2MzXggCv4BCvsBCgNpYmMSIOSmkJr9qPAVGDXCsZqAw0ueCKLqgkKKj0pe5meOKfMkGgkIARgBIAEqAQAiJwgBEgEBGiD5Vt5/BBmvQQnMGXTh8TcjCr0rjLn1J3de7MoaiBcPIiIlCAESIQEeEMwi62nPU392oAPsIrE3wt7z7HD3Zc+G334WRZwgmCInCAESAQEaIDeSU0S9LQzFtlD5HtmNben3CUPk499nQlTqstFtRh+WIiUIARIhAYPzQXwKVrN7RMrXpkE7zScm7vo16FsjuSm8hEQknrNJIicIARIBARogdWBwCp83gq6p5j2krvEl/wBFKME821jZl4K21AIQk44=",
		"proofHeight": { "revisionNumber": "1", "revisionHeight": "17549" },
		"signer": "tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6",
		"application": "transfer",
		"messageType": "MsgTransfer",
		"maybeMsgTransfer": {
			"sender": "eth1sp3e58lfgnk9pjm2n7e7jjl0805s5ju09a2muz",
			"receiver": "tcro1pdn2nsn9wesz6px3lcjsgmp8pefednzp3gmp3q",
			"denom": "transfer/channel-0/basetcro",
			"amount": "200000002700066912",
			"success": true,
			"maybeDenominationTrace": null
		},
		"packetSequence": "15",
		"channelOrdering": "ORDER_UNORDERED",
		"connectionId": "connection-3",
		"packetAck": {
			"result": null,
			"error": "{PACKET_ACK_ERROR}"
		}
	}
}
`

			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_PACKET_ACK_ERROR_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_PACKET_ACK_ERROR_BLOCK_RESULTS_RESP,
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
			Expect(cmds).To(HaveLen(2))
			cmd := cmds[1]
			Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			expectedWithUUID := strings.Replace(
				regex.ReplaceAllString(expected, ""),
				"{UUID}",
				typedEvent.UUID(),
				-1,
			)

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					expectedWithUUID,
					"{PACKET_ACK_ERROR}",
					"unable to unescrow tokens, this may be caused by a malicious counterparty module or a bug: please open an issue on counterparty module: 3800201912basetcro is smaller than 200000002700066912basetcro: insufficient funds",
					-1,
				),
			))
		})
	})
})
