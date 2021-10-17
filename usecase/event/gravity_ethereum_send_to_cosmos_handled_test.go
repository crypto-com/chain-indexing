package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeGravityEthereumSendToCosmosHandled", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)

			anyParams := model.GravityEthereumSendToCosmosHandledEventParams{
				Module:                    "Gravity",
				Sender:                    "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
				Receiver:                  "tcrc1j6xhntmfqauwn0rqhqc89yr9a2xqe8vxthedhp",
				Amount:                    coin.NewCoins(coin.MustNewCoinFromString("basetcro", "1")),
				BridgeChainId:             338,
				EthereumTokenContract:     "0x564a1c3af089d02d0b6c311c650ea3768424cbfa",
				Nonce:                     2,
				EthereumEventVoteRecordId: "vote-record-id",
			}
			event := event_usecase.NewGravityEthereumSendToCosmosHandled(anyHeight, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.GravityEthereumSendToCosmosHandled)
			Expect(typedEvent.Name()).To(Equal(event_usecase.GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.Params).To(Equal(anyParams))
		})
	})
})
