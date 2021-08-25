package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/internal/slice"
)

var _ = Describe("SLice", func() {
	Describe("String", func() {
		It("Contained string", func() {
			strArr := []string{
				"A", "B",
			}

			result := slice.ContainString(strArr, "A")
			Expect(result).To(Equal(true))
		})

		It("Not contained string ", func() {
			strArr := []string{
				"A", "B",
			}

			result := slice.ContainString(strArr, "C")
			Expect(result).To(Equal(false))
		})
	})
})
