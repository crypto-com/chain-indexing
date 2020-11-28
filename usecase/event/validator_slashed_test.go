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

	Describe("En/DecodeValidatorSlashed", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyConsensusNodeAddress := "crocnclcons1548f5hydddg0ea4sdgxse7t7j4jn84zp6y954e"
			anySlashedPower := "123"
			anyReason := "missing_signature"
			event := event_usecase.NewValidatorSlashed(anyHeight, model.SlashValidatorParams{
				ConsensusNodeAddress: anyConsensusNodeAddress,
				SlashedPower:         anySlashedPower,
				Reason:               anyReason,
			})

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.VALIDATOR_SLASHED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.ValidatorSlashed)
			Expect(typedEvent.Name()).To(Equal(event_usecase.VALIDATOR_SLASHED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.ConsensusNodeAddress).To(Equal(anyConsensusNodeAddress))
			Expect(typedEvent.SlashedPower).To(Equal(anySlashedPower))
			Expect(typedEvent.Reason).To(Equal(anyReason))
		})
	})
})
