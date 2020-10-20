package typeconv_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/internal/typeconv"
)

var _ = Describe("MustAtoi32", func() {
	It("should panic when the string is not an integer value", func() {
		Expect(func() {
			typeconv.MustAtoi32("invalid")
		}).To(Panic())
	})

	It("should panic when the integer value exceed int32", func() {
		Expect(func() {
			typeconv.MustAtoi32("2147483648")
		}).To(Panic())
	})

	It("should return int32 representation of the intever value", func() {
		Expect(typeconv.MustAtoi32("1001")).To(Equal(int32(1001)))
	})
})

var _ = Describe("MustAtou32", func() {
	It("should panic when the string is not an integer value", func() {
		Expect(func() {
			typeconv.MustAtou32("invalid")
		}).To(Panic())
	})

	It("should panic when the integer value exceed uint32", func() {
		Expect(func() {
			typeconv.MustAtou32("4294967296")
		}).To(Panic())
	})

	It("should panic when the intever value is negative", func() {
		Expect(func() {
			typeconv.MustAtou32("-1001")
		}).To(Panic())
	})

	It("should return uint32 representation of the intever value", func() {
		Expect(typeconv.MustAtou32("1001")).To(Equal(uint32(1001)))
	})
})
