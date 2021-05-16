package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/projection/nft/utils"
)

var _ = Describe("Utils", func() {
	Describe("GetValueFromJSONData", func() {
		It("should return error if rawJSON is invalid JSON", func() {
			invalidJSON := "{{{"

			_, err := utils.GetValueFromJSONData([]byte(invalidJSON), "a.b")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid character"))
		})

		It("should return buk if accessor does not exist", func() {
			rawJSON := "{\"a\": {\"b\": \"c\"}}"

			value, err := utils.GetValueFromJSONData([]byte(rawJSON), "x.y")
			Expect(err).To(BeNil())
			Expect(value).To(BeNil())
		})

		It("should return the value if accessor exists", func() {
			rawJSON := "{\"a\": {\"b\": \"c\"}}"

			value, err := utils.GetValueFromJSONData([]byte(rawJSON), "a.b")
			Expect(err).To(BeNil())
			Expect(value).To(Equal("c"))
		})
	})
})
