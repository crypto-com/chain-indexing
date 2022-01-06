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

	Describe("En/DecodeEvidence", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTendermintAddress := "D8C5B2B2A1B58FAC65A9E58F5A4CB22536C74961"
			anyInfractionHeight := int64(1001)
			event := event_usecase.NewEvidence(anyHeight, model.EvidenceParams{
				TendermintAddress: anyTendermintAddress,
				InfractionHeight:  anyInfractionHeight,
			})

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.EVIDENCE, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.Evidence)
			Expect(typedEvent.Name()).To(Equal(event_usecase.EVIDENCE))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TendermintAddress).To(Equal(anyTendermintAddress))
			Expect(typedEvent.InfractionHeight).To(Equal(anyInfractionHeight))
		})
	})
})
