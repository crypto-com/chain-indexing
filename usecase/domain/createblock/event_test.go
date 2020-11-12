package createblock_test

import (
	"encoding/json"

	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/domain/createblock"
	"github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("BlockCreated", func() {
	Describe("En/DecodeEvent", func() {
		It("should able to encode and decode to the same Event", func() {
			var block model.Block
			random.Struct(&block)
			event := createblock.NewEvent(&block)
			encoded, _ := json.Marshal(event)

			actual, err := createblock.DecodeEvent(encoded)
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(event))
		})
	})
})
