package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeBlockCommissioned", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyValidator := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyAmount := coin.MustParseDecCoins("123456basetcro,456789tcro")
			event := event_usecase.NewBlockCommissioned(anyHeight, anyValidator, anyAmount)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.BLOCK_COMMISSIONED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.BlockCommissioned)
			Expect(typedEvent.Name()).To(Equal(event_usecase.BLOCK_COMMISSIONED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.Validator).To(Equal(anyValidator))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
