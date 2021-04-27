package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodePowerChanged", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTendermintPubkey := "tWY6qzpOg/6HFj2X3a8+tzIAehW7k2MWOgrjotcWCuI="
			anyPower := "123456"
			event := event_usecase.NewPowerChanged(anyHeight, model.PowerChangeParams{
				TendermintPubkey: anyTendermintPubkey,
				Power:            anyPower,
			})

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.POWER_CHANGED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.PowerChanged)
			Expect(typedEvent.Name()).To(Equal(event_usecase.POWER_CHANGED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TendermintPubkey).To(Equal(anyTendermintPubkey))
			Expect(typedEvent.Power).To(Equal(anyPower))
		})
	})
})
