package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeProposalEnded", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyProposalId := "2"
			anyResult := "proposal_rejected"
			event := event_usecase.NewProposalEnded(anyHeight, anyProposalId, anyResult)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.PROPOSAL_ENDED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.ProposalEnded)
			Expect(typedEvent.Name()).To(Equal(event_usecase.PROPOSAL_ENDED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.Height()).To(Equal(anyHeight))

			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Result).To(Equal(anyResult))
		})
	})
})
