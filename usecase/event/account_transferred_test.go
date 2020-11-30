package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeAccountTransferred", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anySender := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyRecipient := "tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := model.AccountTransferParams{
				Sender:    anySender,
				Recipient: anyRecipient,
				Amount:    anyAmount,
			}
			event := event_usecase.NewAccountTransferred(anyHeight, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.ACCOUNT_TRANSFERRED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.AccountTransferred)
			Expect(typedEvent.Name()).To(Equal(event_usecase.ACCOUNT_TRANSFERRED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.Sender).To(Equal(anySender))
			Expect(typedEvent.Recipient).To(Equal(anyRecipient))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
