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

	Describe("En/DecodeCoinMint", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyAddress := "tcrocncl1sruzd529lhjju6hfcwd2fxp3v0e7p0vqqtme76"
			anyAmount := coin.MustParseCoinsNormalized("123456basetcro,456789tcro")
			event := event_usecase.NewCoinMint(anyHeight, model.CoinMintParams{
				Address: anyAddress,
				Amount:  anyAmount,
			})

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.COIN_MINT, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.CoinMint)
			Expect(typedEvent.Name()).To(Equal(event_usecase.COIN_MINT))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.Address).To(Equal(anyAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
