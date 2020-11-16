package event_test

import (
	"encoding/json"

	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("BlockCreated", func() {
	Describe("En/DecodeBlockCreated", func() {
		It("should able to encode and decode to the same Event", func() {
			var block model.Block
			random.Struct(&block)
			event := event_usecase.NewBlockCreated(&block)
			encoded, _ := json.Marshal(event)

			actual, err := event_usecase.DecodeBlockCreated(encoded)
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(event))
		})
	})
})
