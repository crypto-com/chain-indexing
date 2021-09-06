package event_test

import (
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	"time"

	"github.com/crypto-com/chain-indexing/internal/must"

	"github.com/crypto-com/chain-indexing/internal/json"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCChannelOpenAck", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyPortId := "transfer"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelOpenAck ibc_model.RawMsgChannelOpenAck
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenAck,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenAck",
  "port_id": "transfer",
  "channel_id": "channel-0",
  "counterparty_channel_id": "channel-0",
  "counterparty_version": "ics20-1",
  "proof_try": "CqcCCqQCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASMggCEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0wKgdpY3MyMC0xGgsIARgBIAEqAwACICIrCAESBAIEICAaISCKn2oMLG+m7x551nh06PSUqF3nF6zuDAUZzHWN3Y7R6iIrCAESBAQGICAaISAl+qVfRZQA8J2FNbAhThANLKQJekTVQlahuZE83dRMGyIrCAESBAgQICAaISAnLE8YAY/vNnafNqIgxC5dnti4svEmGWT97i5E1czeVSIrCAESBAokICAaISD6l2Xp3knzZRXOdJhJ8ToqY0qLt6QgDUruuoeIyeXg6wrTAQrQAQoDaWJjEiDpdjgQlYa7NrP+ylMayPUCN/n48pTXiof8JywRXiIjxxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAaj29iVS1fIRuqzuZBN/XoI11+IAJTrJPqBHC1ZQFHsbIiUIARIhAdmJCKPzDC4kc5/SoZhd1InuuJbVmFMLbq5ugt5dhByhIicIARIBARogEqy+GTuqKMTr9SDspw2rkC/9N0sLQcUcY4c9336zn3M=",
  "proof_height": {
    "revision_number": "2",
    "revision_height": "17"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelOpenAckParams{
				RawMsgChannelOpenAck: anyRawMsgChannelOpenAck,

				CounterpartyPortID: anyCounterpartyPortId,
				ConnectionID:       anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenAck(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenAck)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.CounterpartyVersion).To(Equal(anyParams.CounterpartyVersion))
			Expect(typedEvent.Params.ProofTry).To(Equal(anyParams.ProofTry))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyCounterpartyPortId := "transfer"
			anyConnectionId := "connection-0"

			var anyRawValue map[string]interface{}
			var anyRawMsgChannelOpenAck ibc_model.RawMsgChannelOpenAck
			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
					utils.StringToDurationHookFunc(),
					utils.StringToByteSliceHookFunc(),
				),
				Result: &anyRawMsgChannelOpenAck,
			})
			json.MustUnmarshalFromString(`
{
  "@type": "/ibc.core.channel.v1.MsgChannelOpenAck",
  "port_id": "transfer",
  "channel_id": "channel-0",
  "counterparty_channel_id": "channel-0",
  "counterparty_version": "ics20-1",
  "proof_try": "CqcCCqQCCi1jaGFubmVsRW5kcy9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9jaGFubmVsLTASMggCEAEaFQoIdHJhbnNmZXISCWNoYW5uZWwtMCIMY29ubmVjdGlvbi0wKgdpY3MyMC0xGgsIARgBIAEqAwACICIrCAESBAIEICAaISCKn2oMLG+m7x551nh06PSUqF3nF6zuDAUZzHWN3Y7R6iIrCAESBAQGICAaISAl+qVfRZQA8J2FNbAhThANLKQJekTVQlahuZE83dRMGyIrCAESBAgQICAaISAnLE8YAY/vNnafNqIgxC5dnti4svEmGWT97i5E1czeVSIrCAESBAokICAaISD6l2Xp3knzZRXOdJhJ8ToqY0qLt6QgDUruuoeIyeXg6wrTAQrQAQoDaWJjEiDpdjgQlYa7NrP+ylMayPUCN/n48pTXiof8JywRXiIjxxoJCAEYASABKgEAIiUIARIhAUayD1sqD5N3koJLoD5Cenp4hU5w5Hs10qmKy+U4lvkTIiUIARIhAaj29iVS1fIRuqzuZBN/XoI11+IAJTrJPqBHC1ZQFHsbIiUIARIhAdmJCKPzDC4kc5/SoZhd1InuuJbVmFMLbq5ugt5dhByhIicIARIBARogEqy+GTuqKMTr9SDspw2rkC/9N0sLQcUcY4c9336zn3M=",
  "proof_height": {
    "revision_number": "2",
    "revision_height": "17"
  },
  "signer": "cro10snhlvkpuc4xhq82uyg5ex2eezmmf5ed5tmqsv"
}
`, &anyRawValue)
			must.Do(decoder.Decode(anyRawValue))

			anyParams := ibc_model.MsgChannelOpenAckParams{
				RawMsgChannelOpenAck: anyRawMsgChannelOpenAck,

				CounterpartyPortID: anyCounterpartyPortId,
				ConnectionID:       anyConnectionId,
			}

			event := event_usecase.NewMsgIBCChannelOpenAck(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCChannelOpenAck)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CHANNEL_OPEN_ACK_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))

			Expect(typedEvent.Params.PortID).To(Equal(anyParams.PortID))
			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
			Expect(typedEvent.Params.CounterpartyChannelID).To(Equal(anyParams.CounterpartyChannelID))
			Expect(typedEvent.Params.CounterpartyVersion).To(Equal(anyParams.CounterpartyVersion))
			Expect(typedEvent.Params.ProofTry).To(Equal(anyParams.ProofTry))
			Expect(typedEvent.Params.ProofHeight).To(Equal(anyParams.ProofHeight))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))

			Expect(typedEvent.Params.ChannelID).To(Equal(anyParams.ChannelID))
		})
	})
})
