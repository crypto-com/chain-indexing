package event_test

import (
	"encoding/json"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"

	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

var _ = Describe("RawBlockEventCreated", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeRawBlockEventCreated", func() {
		It("should able to encode and decode to the same Event", func() {
			var rawBlockEvent model.CreateRawBlockEventParams
			var blockheight = random.Int64()
			random.Struct(&rawBlockEvent)
			event := event_usecase.NewRawBlockEventCreated(blockheight, &rawBlockEvent)
			encoded, _ := json.Marshal(event)

			decodedEvent, err := registry.DecodeByType(
				event_usecase.RAW_BLOCK_EVENT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
		})
	})
})
