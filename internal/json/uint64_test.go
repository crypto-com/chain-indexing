package json_test

import (
	"github.com/crypto-com/chain-indexing/internal/json"
	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Uint64", func() {
	Describe("JSON en/decoding", func() {
		It("should decode to nil pointer from null value", func() {
			var actual *json.Uint64
			err := jsoniter.Unmarshal([]byte("null"), &actual)

			Expect(err).To(BeNil())
			Expect(actual).To(BeNil())
		})

		It("should decode to 0 from null value", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("null"), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("0"))
		})

		It("should encode nil pointer to null value", func() {
			var v *json.Uint64 = nil
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("null")))
		})

		It("should encode to the string representation", func() {
			v := json.NewUint64(18446744073709551615)
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"18446744073709551615\"")))
		})

		It("should decode from string", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("\"18446744073709551615\""), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("18446744073709551615"))
		})

		It("should decode from number", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("18446744073709551615"), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("18446744073709551615"))
		})

		It("should return error if number overflows uint64", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("18446744073709551616"), &actual)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("unmarshalerDecoder: readUint64: overflow"))
		})

		It("should return error if number is negative", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("-18446744073709551615"), &actual)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("unmarshalerDecoder: readUint64: unexpected character"))
		})

		It("should return error if string is negative number", func() {
			var actual json.Uint64
			err := jsoniter.Unmarshal([]byte("\"-18446744073709551615\""), &actual)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("unmarshalerDecoder: strconv.ParseUint"))
		})

		It("should be able to encode and decode to/from json", func() {
			expected := json.NewUint64(18446744073709551615)
			encoded, _ := jsoniter.Marshal(expected)

			var actual json.Uint64
			err := jsoniter.Unmarshal(encoded, &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal(expected.String()))
		})
	})
})
