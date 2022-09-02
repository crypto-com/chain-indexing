package event_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/json"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/must"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	mapstructure_utils "github.com/crypto-com/chain-indexing/usecase/parser/utils/mapstructure"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCRecvPacket", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"
			anyPacketAck := ibc_model.MsgRecvPacketPacketAck{MaybeResult: base64.MustDecodeString("AQ==")}
			anyApplication := "transfer"
			anyMessageType := "/ibc.applications.transfer.v1.MsgTransfer"

			var anyRawValue map[string]interface{}
			var anyRawMsgRecvPacket ibc_model.RawMsgRecvPacket
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgRecvPacket",
  "packet": {
    "sequence": "1",
    "source_port": "transfer",
    "source_channel": "channel-0",
    "destination_port": "transfer",
    "destination_channel": "channel-0",
    "data": "eyJhbW91bnQiOiIxMjM0IiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjcm8xZHVsd3FnY2RwZW1uOGMzNHNqZDkyZnhlcHo1cDBzcXBlZXZ3N2YiLCJzZW5kZXIiOiJjcm8xMHNuaGx2a3B1YzR4aHE4MnV5ZzVleDJlZXptbWY1ZWQ1dG1xc3YifQ==",
    "timeout_height": {
   	  "revision_number": "2",
   	  "revision_height": "1023"
    },
    "timeout_timestamp": "0"
  },
  "proof_commitment": "CpsCCpgCCjljb21taXRtZW50cy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTAvc2VxdWVuY2VzLzESIOL9pxFSPp6ZejEK+FtCDChW0bJhdl57P/d6zx+hVhgpGgsIARgBIAEqAwACMCIpCAESJQIEMCDj88ZEbojOy3mqsQvLWRQf8tOSvTpdOPsZA8AECkOeziAiKQgBEiUECDAglzLiPEahZdTWY2pN870XzhUMewaTClTsjv3dhk1e6+ggIisIARIECBYwIBohIMytK1VpzSkLoEFzfe4606li/jw6NhJDwdjFKOR/YVqfIikIARIlCiYwIMdfhADZpHkpq3IfEwlxQT8z0N9wSA4ASWUuYkFML8CzIArTAQrQAQoDaWJjEiAbbEgt1HZqlOsM1hT5uoYHf+4r+1wCNMmALauIR+pv+xoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAad0DU0QldmQc2csoMIRyWEWnRpubv1IwniOeASCc9f+IiUIARIhARFr3/h1b3vKOoGq7PWsyM0e7UKp3Wd9nNus5KlIGgiZIicIARIBARogXnlUV86r9evx40joRXDS41kvFZuFFW8KFgSUfQW/uRQ=",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "25"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(anyRawValue, &anyRawMsgRecvPacket)

			var anyRawMsgRecvPacketFungibleTokenPacketData ibc_model.FungibleTokenPacketData
			json.MustUnmarshalFromString(`
{
  "amount":"1234",
  "denom":"basecro",
  "receiver":"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "sender":"cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(
				anyRawValue,
				&anyRawMsgRecvPacketFungibleTokenPacketData,
			)

			anyParams := ibc_model.MsgRecvPacketParams{
				RawMsgRecvPacket: anyRawMsgRecvPacket,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
					FungibleTokenPacketData: anyRawMsgRecvPacketFungibleTokenPacketData,
					Success:                 true,
					MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
						Hash:  "6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
						Denom: "ibc/6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
					},
				},

				PacketSequence:  anyPacketSequence,
				ChannelOrdering: anyChannelOrdering,
				ConnectionID:    anyConnectionId,
				PacketAck:       anyPacketAck,
			}

			event := event_usecase.NewMsgIBCRecvPacket(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_RECV_PACKET_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCRecvPacket)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_RECV_PACKET_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.RawMsgRecvPacket).To(Equal(anyParams.RawMsgRecvPacket))
			Expect(typedEvent.Params.MaybeFungibleTokenPacketData).To(Equal(anyParams.MaybeFungibleTokenPacketData))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.PacketAck).To(Equal(anyParams.PacketAck))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"
			anyPacketAck := ibc_model.MsgRecvPacketPacketAck{MaybeResult: base64.MustDecodeString("AQ==")}
			anyApplication := "transfer"
			anyMessageType := "/ibc.applications.transfer.v1.MsgTransfer"

			var anyRawValue map[string]interface{}
			var anyRawMsgRecvPacket ibc_model.RawMsgRecvPacket
			msgRecvPacketDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					mapstructure_utils.StringToDurationHookFunc(),
					mapstructure_utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgRecvPacket,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgRecvPacket",
  "packet": {
    "sequence": "1",
    "source_port": "transfer",
    "source_channel": "channel-0",
    "destination_port": "transfer",
    "destination_channel": "channel-0",
    "data": "eyJhbW91bnQiOiIxMjM0IiwiZGVub20iOiJiYXNlY3JvIiwicmVjZWl2ZXIiOiJjcm8xZHVsd3FnY2RwZW1uOGMzNHNqZDkyZnhlcHo1cDBzcXBlZXZ3N2YiLCJzZW5kZXIiOiJjcm8xMHNuaGx2a3B1YzR4aHE4MnV5ZzVleDJlZXptbWY1ZWQ1dG1xc3YifQ==",
    "timeout_height": {
   	  "revision_number": "2",
   	  "revision_height": "1023"
    },
    "timeout_timestamp": "0"
  },
  "proof_commitment": "CpsCCpgCCjljb21taXRtZW50cy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTAvc2VxdWVuY2VzLzESIOL9pxFSPp6ZejEK+FtCDChW0bJhdl57P/d6zx+hVhgpGgsIARgBIAEqAwACMCIpCAESJQIEMCDj88ZEbojOy3mqsQvLWRQf8tOSvTpdOPsZA8AECkOeziAiKQgBEiUECDAglzLiPEahZdTWY2pN870XzhUMewaTClTsjv3dhk1e6+ggIisIARIECBYwIBohIMytK1VpzSkLoEFzfe4606li/jw6NhJDwdjFKOR/YVqfIikIARIlCiYwIMdfhADZpHkpq3IfEwlxQT8z0N9wSA4ASWUuYkFML8CzIArTAQrQAQoDaWJjEiAbbEgt1HZqlOsM1hT5uoYHf+4r+1wCNMmALauIR+pv+xoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAad0DU0QldmQc2csoMIRyWEWnRpubv1IwniOeASCc9f+IiUIARIhARFr3/h1b3vKOoGq7PWsyM0e7UKp3Wd9nNus5KlIGgiZIicIARIBARogXnlUV86r9evx40joRXDS41kvFZuFFW8KFgSUfQW/uRQ=",
  "proof_height": {
    "revision_number": "1",
    "revision_height": "25"
  },
  "signer": "cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f"
}
`, &anyRawValue)
			must.Do(msgRecvPacketDecoder.Decode(anyRawValue))

			var anyRawMsgRecvPacketFungibleTokenPacketData ibc_model.FungibleTokenPacketData
			json.MustUnmarshalFromString(`
{
  "amount":"1234",
  "denom":"basecro",
  "receiver":"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "sender":"cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			mapstructure_utils.DefaultMapstructureDecoder.MustDecode(
				anyRawValue,
				&anyRawMsgRecvPacketFungibleTokenPacketData,
			)

			anyParams := ibc_model.MsgRecvPacketParams{
				RawMsgRecvPacket: anyRawMsgRecvPacket,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeFungibleTokenPacketData: &ibc_model.MsgRecvPacketFungibleTokenPacketData{
					FungibleTokenPacketData: anyRawMsgRecvPacketFungibleTokenPacketData,
					Success:                 true,
					MaybeDenominationTrace: &ibc_model.MsgRecvPacketFungibleTokenDenominationTrace{
						Hash:  "6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
						Denom: "ibc/6411AE2ADA1E73DB59DB151A8988F9B7D5E7E233D8414DB6817F8F1A01611F86",
					},
				},

				PacketSequence:  anyPacketSequence,
				ChannelOrdering: anyChannelOrdering,
				ConnectionID:    anyConnectionId,
				PacketAck:       anyPacketAck,
			}

			event := event_usecase.NewMsgIBCRecvPacket(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_RECV_PACKET_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCRecvPacket)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_RECV_PACKET_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))

			Expect(typedEvent.Params.RawMsgRecvPacket).To(Equal(anyParams.RawMsgRecvPacket))
			Expect(typedEvent.Params.MaybeFungibleTokenPacketData).To(Equal(anyParams.MaybeFungibleTokenPacketData))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
			Expect(typedEvent.Params.PacketAck).To(Equal(anyParams.PacketAck))
		})
	})
})
