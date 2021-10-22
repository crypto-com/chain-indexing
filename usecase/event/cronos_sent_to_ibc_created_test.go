package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/internal/json"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeCronosSendToIBCCreated", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			event := event_usecase.NewCronosSendToIBCCreated(anyHeight, model.CronosSendToIBCParams{
				TxHash:         "E0DFC9314BAFC8A95991D49764CAFCDDE233FEB3892E2A053C7AA6BB1708241E",
				EthereumTxHash: "0x33c03179001d0c780416fbdbfcbab2e7778dc8692ab5202699225c9e2c592242",
				SourcePort:     "transfer",
				SourceChannel:  "channel-0",
				Token: model.CronosSendToIBCToken{
					Amount: json.NewUint64(1000000000),
					Denom:  "transfer/channel-0/basetcro",
				},
				Sender:   "tcrc13yux6z8mh6w5t3v4uq7clewnh35znrgdgye0k2",
				Receiver: "tcro1tzhdkuc328cgh2hycyfddtdpqfwwu42ywyfvkj",
				TimeoutHeight: model.CronosSendToIBCHeight{
					RevisionNumber: uint64(0),
					RevisionHeight: uint64(0),
				},
				TimeoutTimestamp:   "1633722759099392805",
				PacketDataHex:      "7b22616d6f756e74223a2231303030303030303030222c2264656e6f6d223a227472616e736665722f6368616e6e656c2d302f626173657463726f222c227265636569766572223a227463726f31747a68646b756333323863676832687963796664647464707166777775343279777966766b6a222c2273656e646572223a22746372633133797578367a386d6836773574337634757137636c65776e6833357a6e726764677965306b32227d\"",
				PacketSequence:     28,
				DestinationPort:    "transfer",
				DestinationChannel: "channel-129",
				ChannelOrdering:    "ORDER_UNORDERED",
				ConnectionID:       "connection-0",
			})

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.CRONOS_SEND_TO_IBC_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.CronosSendToIBCCreated)
			Expect(typedEvent.Name()).To(Equal(event_usecase.CRONOS_SEND_TO_IBC_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.Params.TxHash).To(Equal("E0DFC9314BAFC8A95991D49764CAFCDDE233FEB3892E2A053C7AA6BB1708241E"))
			Expect(typedEvent.Params.EthereumTxHash).To(Equal("0x33c03179001d0c780416fbdbfcbab2e7778dc8692ab5202699225c9e2c592242"))
		})
	})
})
