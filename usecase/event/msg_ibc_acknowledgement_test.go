package event_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/internal/json"

	"github.com/crypto-com/chain-indexing/internal/must"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCAcknowledgement", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"
			anyApplication := "transfer"
			anyMessageType := "MsgTransfer"

			var anyRawValue map[string]interface{}
			var anyRawMsg ibc_model.RawMsgAcknowledgement
			msgRecvPacketDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsg,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgAcknowledgement",
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
  "acknowledgement": "eyJyZXN1bHQiOiJBUT09In0=",
  "proof_acked": "CscCCsQCCjJhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtMC9zZXF1ZW5jZXMvMRIgCPdVftUYJv4Y2EUSvyTsdQAe268hI6R333KgqfNkCnwaCwgBGAEgASoDAAI0IisIARIEAgQ0IBohICsdSIdU+v81lhDQ/dswq/j+y5pEx0RcBqDNNvGBURKNIisIARIEBAg0IBohIO5n1rXg/PKXipcxX5Uy+XyTZ3jGwpIB/RDNehpbIU0jIisIARIEBhA0IBohIIymA52HzmLEaKECJT3n3bOHLJhBuDBeQclE+lp93vOyIisIARIECBo0IBohIHQvJirT50kRamjMV0YCUFCu6329YWyKcegokMjXHxBQIisIARIECjI0IBohIHRaqiQwujUalYQJUdxvggy1gNyGKCYldtX1dq+fAjtwCtMBCtABCgNpYmMSIHdgqTmy3PkcFIK3GAEKyMa1kQJniPJbk/Hl3o6gyztLGgkIARgBIAEqAQAiJQgBEiEBRrIPWyoPk3eSgkugPkJ6eniFTnDkezXSqYrL5TiW+RMiJQgBEiEBb6XKyCWaqTgyGMhPJZ7TKQAjel7eQsEcAP3XGSsSBg4iJQgBEiEBpzj0MzNAMsi156BdbnVm9mYPLOVKKnxZmEfyj3pKw5kiJwgBEgEBGiA7xilGseIAH2wYkx1DUIexthQg/0UX9+RjCYJ6M0kL5Q==",
  "proof_height": {
    "revision_number": "2",
    "revision_height": "27"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(msgRecvPacketDecoder.Decode(anyRawValue))

			var anyFungibleTokenPacketData ibc_model.FungibleTokenPacketData
			fungibleTokenPacketDataDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyFungibleTokenPacketData,
			})

			json.MustUnmarshalFromString(`
{
  "amount":"1234",
  "denom":"basecro",
  "receiver":"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "sender":"cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(fungibleTokenPacketDataDecoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgAcknowledgementParams{
				RawMsgAcknowledgement: anyRawMsg,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
					FungibleTokenPacketData: anyFungibleTokenPacketData,
					Success:                 true,
					Acknowledgement:         "{0xc0038ae7a0}",
				},

				PacketSequence:  anyPacketSequence,
				ChannelOrdering: anyChannelOrdering,
				ConnectionID:    anyConnectionId,
			}

			event := event_usecase.NewMsgIBCAcknowledgement(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCAcknowledgement)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_ACKNOWLEDGEMENT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.RawMsgAcknowledgement).To(Equal(anyParams.RawMsgAcknowledgement))
			Expect(typedEvent.Params.MaybeFungibleTokenPacketData).To(Equal(anyParams.MaybeFungibleTokenPacketData))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyPacketSequence := uint64(1)
			anyChannelOrdering := "ORDER_UNORDERED"
			anyConnectionId := "connection-0"
			anyApplication := "transfer"
			anyMessageType := "MsgTransfer"

			var anyRawValue map[string]interface{}
			var anyRawMsg ibc_model.RawMsgAcknowledgement
			msgRecvPacketDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsg,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgAcknowledgement",
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
  "acknowledgement": "eyJyZXN1bHQiOiJBUT09In0=",
  "proof_acked": "CscCCsQCCjJhY2tzL3BvcnRzL3RyYW5zZmVyL2NoYW5uZWxzL2NoYW5uZWwtMC9zZXF1ZW5jZXMvMRIgCPdVftUYJv4Y2EUSvyTsdQAe268hI6R333KgqfNkCnwaCwgBGAEgASoDAAI0IisIARIEAgQ0IBohICsdSIdU+v81lhDQ/dswq/j+y5pEx0RcBqDNNvGBURKNIisIARIEBAg0IBohIO5n1rXg/PKXipcxX5Uy+XyTZ3jGwpIB/RDNehpbIU0jIisIARIEBhA0IBohIIymA52HzmLEaKECJT3n3bOHLJhBuDBeQclE+lp93vOyIisIARIECBo0IBohIHQvJirT50kRamjMV0YCUFCu6329YWyKcegokMjXHxBQIisIARIECjI0IBohIHRaqiQwujUalYQJUdxvggy1gNyGKCYldtX1dq+fAjtwCtMBCtABCgNpYmMSIHdgqTmy3PkcFIK3GAEKyMa1kQJniPJbk/Hl3o6gyztLGgkIARgBIAEqAQAiJQgBEiEBRrIPWyoPk3eSgkugPkJ6eniFTnDkezXSqYrL5TiW+RMiJQgBEiEBb6XKyCWaqTgyGMhPJZ7TKQAjel7eQsEcAP3XGSsSBg4iJQgBEiEBpzj0MzNAMsi156BdbnVm9mYPLOVKKnxZmEfyj3pKw5kiJwgBEgEBGiA7xilGseIAH2wYkx1DUIexthQg/0UX9+RjCYJ6M0kL5Q==",
  "proof_height": {
    "revision_number": "2",
    "revision_height": "27"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(msgRecvPacketDecoder.Decode(anyRawValue))

			var anyFungibleTokenPacketData ibc_model.FungibleTokenPacketData
			fungibleTokenPacketDataDecoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					ibcmsg.StringToDurationHookFunc(),
					ibcmsg.StringToByteSliceHookFunc(),
				),
				Result: &anyFungibleTokenPacketData,
			})

			json.MustUnmarshalFromString(`
{
  "amount":"1234",
  "denom":"basecro",
  "receiver":"cro1dulwqgcdpemn8c34sjd92fxepz5p0sqpeevw7f",
  "sender":"cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(fungibleTokenPacketDataDecoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgAcknowledgementParams{
				RawMsgAcknowledgement: anyRawMsg,

				Application: anyApplication,
				MessageType: anyMessageType,
				MaybeFungibleTokenPacketData: &ibc_model.MsgAcknowledgementFungibleTokenPacketData{
					FungibleTokenPacketData: anyFungibleTokenPacketData,
					Success:                 true,
					Acknowledgement:         "{0xc0038ae7a0}",
				},

				PacketSequence:  anyPacketSequence,
				ChannelOrdering: anyChannelOrdering,
				ConnectionID:    anyConnectionId,
			}

			event := event_usecase.NewMsgIBCAcknowledgement(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_ACKNOWLEDGEMENT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCAcknowledgement)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_ACKNOWLEDGEMENT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.RawMsgAcknowledgement).To(Equal(anyParams.RawMsgAcknowledgement))
			Expect(typedEvent.Params.MaybeFungibleTokenPacketData).To(Equal(anyParams.MaybeFungibleTokenPacketData))
			Expect(typedEvent.Params.Application).To(Equal(anyParams.Application))
			Expect(typedEvent.Params.MessageType).To(Equal(anyParams.MessageType))

			Expect(typedEvent.Params.PacketSequence).To(Equal(anyParams.PacketSequence))
			Expect(typedEvent.Params.ChannelOrdering).To(Equal(anyParams.ChannelOrdering))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})
	})
})
