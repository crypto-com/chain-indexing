package parser_test

import (
	"regexp"
	"strings"

	"github.com/hashicorp/go-version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/crypto-com/chain-indexing/usecase/parser/test"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgIBCRecvPacket", func() {
		Describe("MsgTransfer", func() {
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
      "success": {SUCCESS},
      "maybeDenominationTrace": {
        "hash": "6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
        "denom": "ibc/6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86"
	  }
    },

    "packetSequence": "1",
    "channelOrdering": "ORDER_UNORDERED",
    "connectionId": "connection-0",
    "packetAck": {
      "result": "AQ=="
    }
  }
}
`
			It("should parse Msg commands when there is MsgIBCRecvPacket in the transaction", func() {
				txDecoder := utils.NewTxDecoder()
				block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
				))
				blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
				))

				accountAddressPrefix := "cro"
				stakingDenom := "basecro"
				v0_42_7 := version.Must(version.NewVersion("v0.42.7"))
				cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
					txDecoder,
					block,
					blockResults,
					accountAddressPrefix,
					stakingDenom,
					v0_42_7,
				)
				Expect(err).To(BeNil())
				Expect(cmds).To(HaveLen(2))
				cmd := cmds[1]
				Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

				untypedEvent, _ := cmd.Exec()
				typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

				regex, _ := regexp.Compile("\n?\r?\\s?")

				Expect(json.MustMarshalToString(typedEvent)).To(Equal(
					strings.ReplaceAll(
						strings.ReplaceAll(
							regex.ReplaceAllString(expected, ""),
							"{UUID}",
							typedEvent.UUID(),
						),
						"{SUCCESS}",
						"false",
					),
				))
			})

			It("should reverse success value when Cosmos SDK version is before v0.42.7", func() {
				txDecoder := utils.NewTxDecoder()
				block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
				))
				blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
				))

				accountAddressPrefix := "cro"
				stakingDenom := "basecro"
				v0_42 := version.Must(version.NewVersion("v0.42"))
				cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
					txDecoder,
					block,
					blockResults,
					accountAddressPrefix,
					stakingDenom,
					v0_42,
				)
				Expect(err).To(BeNil())
				Expect(cmds).To(HaveLen(2))
				cmd := cmds[1]
				Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

				untypedEvent, _ := cmd.Exec()
				typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

				regex, _ := regexp.Compile("\n?\r?\\s?")

				Expect(json.MustMarshalToString(typedEvent)).To(Equal(
					strings.ReplaceAll(
						strings.ReplaceAll(
							regex.ReplaceAllString(expected, ""),
							"{UUID}",
							typedEvent.UUID(),
						),
						"{SUCCESS}",
						"true",
					),
				))
			})

			It("should parse success value using the on-chain value when Cosmos SDK version is v0.42.7", func() {
				txDecoder := utils.NewTxDecoder()
				block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
				))
				blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
				))

				accountAddressPrefix := "cro"
				stakingDenom := "basecro"
				v0_42_7 := version.Must(version.NewVersion("v0.42.7"))
				cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
					txDecoder,
					block,
					blockResults,
					accountAddressPrefix,
					stakingDenom,
					v0_42_7,
				)
				Expect(err).To(BeNil())
				Expect(cmds).To(HaveLen(2))
				cmd := cmds[1]
				Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

				untypedEvent, _ := cmd.Exec()
				typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

				regex, _ := regexp.Compile("\n?\r?\\s?")

				Expect(json.MustMarshalToString(typedEvent)).To(Equal(
					strings.ReplaceAll(
						strings.ReplaceAll(
							regex.ReplaceAllString(expected, ""),
							"{UUID}",
							typedEvent.UUID(),
						),
						"{SUCCESS}",
						"false",
					),
				))
			})

			It("should parse success value using the on-chain value when Cosmos SDK version is v0.43", func() {
				txDecoder := utils.NewTxDecoder()
				block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESP,
				))
				blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
					usecase_parser_test.TX_MSG_RECV_PACKET_BLOCK_RESULTS_RESP,
				))

				accountAddressPrefix := "cro"
				stakingDenom := "basecro"
				v0_42_7 := version.Must(version.NewVersion("v0.42.7"))
				cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
					txDecoder,
					block,
					blockResults,
					accountAddressPrefix,
					stakingDenom,
					v0_42_7,
				)
				Expect(err).To(BeNil())
				Expect(cmds).To(HaveLen(2))
				cmd := cmds[1]
				Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

				untypedEvent, _ := cmd.Exec()
				typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

				regex, _ := regexp.Compile("\n?\r?\\s?")

				Expect(json.MustMarshalToString(typedEvent)).To(Equal(
					strings.ReplaceAll(
						strings.ReplaceAll(
							regex.ReplaceAllString(expected, ""),
							"{UUID}",
							typedEvent.UUID(),
						),
						"{SUCCESS}",
						"false",
					),
				))
			})
		})

		It("should parse Msg commands when there is SoloMachine MsgIBCRecvPacket in the transaction", func() {
			expected := `{
  "name": "MsgRecvPacketCreated",
  "version": 1,
  "height": 14803,
  "uuid": "{UUID}",
  "msgName": "MsgRecvPacket",
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
      "timeoutHeight": {
    	"revisionNumber": "4",
    	"revisionHeight": "14812"
      },
      "timeoutTimestamp": "0"
    },
    "proofCommitment": "CkQKQhJAqfGjUQ5IBMpw/u/sAm+xxKztwsL9zJHGs/GObfCQPyd2Yi489Y8BbB0I9nQiOXymYa5Tu/lTuo+tJMQSLzCr+BCqsPyIBg==",
    "proofHeight": {
      "revisionNumber": "0",
      "revisionHeight": "5"
    },
    "signer": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",

    "application": "transfer",
    "messageType": "MsgTransfer",
    "maybeMsgTransfer": {
      "sender": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
      "receiver": "tcro14wku4hr74m0m4tvexs4f6jvuy6vnu2x2dg7hsy",
      "denom": "solotoken",
      "amount": "20",
      "success": true,
      "maybeDenominationTrace": {
        "hash": "1A35E932DCE61466ED9F72D0B436628C388FB1BC60CB23A055039B7DE54883CC",
        "denom": "ibc/1A35E932DCE61466ED9F72D0B436628C388FB1BC60CB23A055039B7DE54883CC"
	  }
    },

    "packetSequence": "1",
    "channelOrdering": "ORDER_UNORDERED",
    "connectionId": "connection-0",
    "packetAck": {
      "result": "AQ=="
    }
  }
}
`
			txDecoder := utils.NewTxDecoder()
			block, _, _ := tendermint.ParseBlockResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_SOLO_MACHINE_BLOCK_RESP,
			))
			blockResults, _ := tendermint.ParseBlockResultsResp(strings.NewReader(
				usecase_parser_test.TX_MSG_RECV_PACKET_SOLO_MACHINE_BLOCK_RESULTS_RESP,
			))

			accountAddressPrefix := "cro"
			stakingDenom := "basecro"
			v0_43 := version.Must(version.NewVersion("v0.43"))
			cmds, err := parser.ParseBlockResultsTxsMsgToCommands(
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				stakingDenom,
				v0_43,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			cmd := cmds[0]
			Expect(cmd.Name()).To(Equal("CreateMsgIBCRecvPacket"))

			untypedEvent, _ := cmd.Exec()
			typedEvent := untypedEvent.(*event.MsgIBCRecvPacket)

			regex, _ := regexp.Compile("\n?\r?\\s?")

			Expect(json.MustMarshalToString(typedEvent)).To(Equal(
				strings.ReplaceAll(
					regex.ReplaceAllString(expected, ""),
					"{UUID}",
					typedEvent.UUID(),
				),
			))
		})

	})
})
