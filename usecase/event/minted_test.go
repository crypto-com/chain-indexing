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

	Describe("En/DecodeMinted", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyBondedRatio := "0.000004999957742178"
			anyInflation := "0.013000953646781450"
			anyAnnualProvisions := coin.MustParseCoinsNormalized("104008508208850609.312053700576913050basetcro")
			anyAmount := coin.MustParseCoinsNormalized("16479153707basetcro")

			anyParams := model.MintParams{
				BondedRatio:      anyBondedRatio,
				Inflation:        anyInflation,
				AnnualProvisions: anyAnnualProvisions,
				Amount:           anyAmount,
			}
			event := event_usecase.NewMinted(anyHeight, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MINTED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.Minted)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MINTED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.Height()).To(Equal(anyHeight))

			Expect(typedEvent.BondedRatio).To(Equal(anyBondedRatio))
			Expect(typedEvent.Inflation).To(Equal(anyInflation))
			Expect(typedEvent.AnnualProvisions).To(Equal(anyAnnualProvisions))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
