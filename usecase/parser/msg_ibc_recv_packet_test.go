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
	Describe("MsgIBCRecvPacket", func() {
		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction", func() {
			expected := `{
  "name": "/ibc.core.channel.v1.MsgRecvPacket.Created",
  "version": 1,
  "height": 26,
  "uuid": "{UUID}",
  "msgName": "/ibc.core.channel.v1.MsgRecvPacket",
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
    "messageType": "/ibc.applications.transfer.v1.MsgTransfer",
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
      "error": null
    }
  }
}
`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_RECV_PACKET_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgRecvPacket.Create"))

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
			Expect(possibleSignerAddresses[1]).To(Equal("cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"))
		})

		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction and MsgIBCRecvPacket.PacketAck.Error is NOT empty string", func() {
			expected := `{
	"name": "/ibc.core.channel.v1.MsgRecvPacket.Created",
	"version": 1,
	"height": 130418,
	"uuid": "{UUID}",
	"msgName": "/ibc.core.channel.v1.MsgRecvPacket",
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
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
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

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_PACKET_ACK_ERROR_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_PACKET_ACK_ERROR_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_RECV_PACKET_PACKET_ACK_ERROR_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgRecvPacket.Create"))

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
			Expect(possibleSignerAddresses[1]).To(Equal("tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6"))
		})

		It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction and missing fungible_token_packet event in TxsResult log", func() {
			expected := `{
	"name": "MsgAlreadyRelayedRecvPacket.Created",
	"version": 1,
	"height": 317994,
	"uuid": "{UUID}",
	"msgName": "MsgAlreadyRelayedRecvPacket",
	"txHash": "D4BBE348AB746FCED63D2028D8886B9091312336142AC040C1C2841E7BF78C9C",
	"msgIndex": 1,
	"params": {
		"packet": {
			"sequence": "45",
			"sourcePort": "transfer",
			"sourceChannel": "channel-0",
			"destinationPort": "transfer",
			"destinationChannel": "channel-3",
			"data": "eyJhbW91bnQiOiIxMDAwMDAwMDAwMDAwIiwiZGVub20iOiJ0cmFuc2Zlci9jaGFubmVsLTAvYmFzZXRjcm8iLCJyZWNlaXZlciI6InRjcm8xZjZxY3ZwMzNkYzc5eHpwdXdsbDdtbG41bG5lcHVxdjhkN2xlZDkiLCJzZW5kZXIiOiJldGgxbXRjbjI1MDVrMzdtbHp0eXdmOGVnOHNwdjBrcG5zcWFtMnpzMDIifQ==",
			"timeoutHeight": { "revisionNumber": "4", "revisionHeight": "318973" },
			"timeoutTimestamp": "1631044281305171742"
		},
		"proofCommitment": "Ct4DCtsDCjpjb21taXRtZW50cy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTAvc2VxdWVuY2VzLzQ1EiAwXmZ+E3z/y5/o578BC3kosoTiz+6+AojpYblmSh5W8xoNCAEYASABKgUAAr7ZECItCAESBgIEvtkQIBohIPbTrk6EJ8X5DxhymRoZgwtWIdCWUEDvERLzfc0ZAqsGIisIARInBAi+2RAgWaR5S/iXR6U3mJrwh2+coIl4NuGr7NmYE6Q9BaRg5OUgIisIARInBg6+2RAggfivNe1TRsR5vKy2Q/8+PFs+4aH1/SXUSQrz8S5OQzogIisIARInCBq+2RAgxdZ4NqbPmgBmrQjrBjseWC4e8qfIh9HB2p3l5PFGuKogIisIARInCji+2RAgSbJl1ZFAyv2EfWKcK7c0EX9BYSWC8Qqe15sVwfznQKggIisIARInDHK+2RAgNP1NCTsx4mWJXKq6JA28CLfe7AKb+NlRzt1Iz8ap7E0gIi4IARIHDp4BvtkQIBohIDevlIZff9mKJUxPfds1+iJkBdYITPcR4K2lLtT9rtM2IiwIARIoEo4EvtkQIC1OxSRQNX8wkn4yu5jUjYisJqp/Octv4mrIo7cZ65riIAr+AQr7AQoDaWJjEiB81eIohS3kImn0slDchX3TsOpmg1KtoBJjeoS0XamZ3xoJCAEYASABKgEAIicIARIBARog+VbefwQZr0EJzBl04fE3Iwq9K4y59Sd3XuzKGogXDyIiJQgBEiEBHhDMIutpz1N/dqAD7CKxN8Le8+xw92XPht9+FkWcIJgiJwgBEgEBGiD/qVFKOoGhKJbUxy/btGutELgMJU2cxTOhZsIRJkfjNiIlCAESIQHmnMo6MTcTPxTquoyKFdT4/zdYaQSWXpWpi+bivu8T3SInCAESAQEaIHVgcAqfN4KuqeY9pK7xJf8ARSjBPNtY2ZeCttQCEJOO",
		"proofHeight": { "revisionNumber": "1", "revisionHeight": "136800" },
		"signer": "tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6",
		"application": "transfer",
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
		"maybeMsgTransfer": {
			"sender": "eth1mtcn2505k37mlztywf8eg8spv0kpnsqam2zs02",
			"receiver": "tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9",
			"denom": "transfer/channel-0/basetcro",
			"amount": "1000000000000",
			"maybeDenominationTrace": null
		},
		"packetSequence": "45",
		"channelOrdering": "",
		"connectionId": "",
		"packetAck": {
			"result": null,
			"error": null
		}
	}
}
`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_MISSING_FUNGIBLE_TOKEN_PACKET_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_MISSING_FUNGIBLE_TOKEN_PACKET_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_RECV_PACKET_MISSING_FUNGIBLE_TOKEN_PACKET_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("CreateMsgAlreadyRelayedIBCRecvPacket"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgAlreadyRelayedIBCRecvPacket)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.Replace(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
					-1,
				),
			))
			Expect(possibleSignerAddresses[1]).To(Equal("tcro18mcwp6vtlvpgxy62eledk3chhjguw636x8n7h6"))
		})

		It("should parse Msg commands when there is SoloMachine MsgIBCRecvPacket in the transaction", func() {
			expected := `{
	"name": "/ibc.core.channel.v1.MsgRecvPacket.Created",
	"version": 1,
	"height": 14803,
	"uuid": "{UUID}",
	"msgName": "/ibc.core.channel.v1.MsgRecvPacket",
	"txHash": "0696B4561D093E0AF784D6CC5701C4FB0645E47BE425C47108737E23BB4FBDEA",
	"msgIndex": 0,
	"params": {
		"packet": {
			"sequence": "1",
			"sourcePort": "transfer",
			"sourceChannel": "channel-VSAv",
			"destinationPort": "transfer",
			"destinationChannel": "channel-0",
			"data": "eyJkZW5vbSI6InNvbG90b2tlbiIsImFtb3VudCI6MjAsInNlbmRlciI6InRjcm8xNHdrdTRocjc0bTBtNHR2ZXhzNGY2anZ1eTZ2bnUyeDJkZzdoc3kiLCJyZWNlaXZlciI6InRjcm8xNHdrdTRocjc0bTBtNHR2ZXhzNGY2anZ1eTZ2bnUyeDJkZzdoc3kifQ==",
			"timeoutHeight": { "revisionNumber": "4", "revisionHeight": "14812" },
			"timeoutTimestamp": "0"
		},
		"proofCommitment": "CkQKQhJAqfGjUQ5IBMpw/u/sAm+xxKztwsL9zJHGs/GObfCQPyd2Yi489Y8BbB0I9nQiOXymYa5Tu/lTuo+tJMQSLzCr+BCqsPyIBg==",
		"proofHeight": { "revisionNumber": "0", "revisionHeight": "5" },
		"signer": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
		"application": "transfer",
		"messageType": "/ibc.applications.transfer.v1.MsgTransfer",
		"maybeMsgTransfer": {
			"sender": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
			"receiver": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
			"denom": "solotoken",
			"amount": "20",
			"success": false,
			"maybeDenominationTrace": {
				"hash": "1A35E932DCE61466ED9F72D0B436628C388FB1BC60CB23A055039B7DE54883CC",
				"denom": "ibc/1A35E932DCE61466ED9F72D0B436628C388FB1BC60CB23A055039B7DE54883CC"
			}
		},
		"packetSequence": "1",
		"channelOrdering": "ORDER_UNORDERED",
		"connectionId": "connection-0",
		"packetAck": {
			"result": "AQ==",
			"error": null
		}
	}
}
`

			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_SOLO_MACHINE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_SOLO_MACHINE_BLOCK_RESULTS_RESP,
			))

			tx := mustParseTxsResp(usecase_parser_test.TX_MSG_RECV_PACKET_SOLO_MACHINE_TXS_RESP)
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
			Expect(cmd.Name()).To(Equal("/ibc.core.channel.v1.MsgRecvPacket.Create"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			expectedWithUUID := strings.Replace(
				regex.ReplaceAllString(expected, ""),
				"{UUID}",
				typedEvent.UUID(),
				-1,
			)

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(expectedWithUUID))
			Expect(possibleSignerAddresses[0]).To(Equal("tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy"))
		})
	})
})
