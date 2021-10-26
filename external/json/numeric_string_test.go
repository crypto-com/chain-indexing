package json_test

import (
	json2 "github.com/crypto-com/chain-indexing/external/json"
	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const MAX_UINT64 = "18446744073709551615"
const MAX_INT64 = "9223372036854775807"
const MAX_NEGATIVE_INT64 = "-9223372036854775808"

var _ = Describe("NumericString", func() {
	Describe("JSON en/decoding", func() {
		It("should decode to nil pointer from null value", func() {
			var actual *json2.NumericString
			err := jsoniter.Unmarshal([]byte("null"), &actual)

			Expect(err).To(BeNil())
			Expect(actual).To(BeNil())
		})

		It("should decode to 0 from null value", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte("null"), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("0"))
		})

		It("should encode nil pointer to null value", func() {
			var v *json2.NumericString = nil
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("null")))
		})

		It("should encode number which exceeds int64 to string", func() {
			v, _ := json2.NewNumericString("9223372036854775807123456")
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"9223372036854775807123456\"")))
		})

		It("should encode max uint64 to the string representation", func() {
			v, _ := json2.NewNumericString(MAX_UINT64)
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"18446744073709551615\"")))
		})

		It("should encode max int64 to the string representation", func() {
			v, _ := json2.NewNumericString(MAX_INT64)
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"9223372036854775807\"")))
		})

		It("should encode max negative int64 to the string representation", func() {
			v, _ := json2.NewNumericString(MAX_NEGATIVE_INT64)
			actual, err := jsoniter.Marshal(v)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"-9223372036854775808\"")))
		})

		It("should decode from number which exceeds int64 in string", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte("\"9223372036854775807123456\""), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("9223372036854775807123456"))
		})

		It("should decode from negative number which exceeds int64 in string", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte("\"-9223372036854775807123456\""), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("-9223372036854775807123456"))
		})

		It("should decode from max uint64 in string", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte("\"18446744073709551615\""), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("18446744073709551615"))
		})

		It("should decode from max uint64 in number", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte(MAX_UINT64), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("18446744073709551615"))
		})

		It("should decode from max int64 in number", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte(MAX_INT64), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("9223372036854775807"))
		})

		It("should decode from max negative int64 in number", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte(MAX_NEGATIVE_INT64), &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("-9223372036854775808"))
		})

		It("should return error when decoding invalid number in string", func() {
			var actual json2.NumericString
			err := jsoniter.Unmarshal([]byte("\"invalid\""), &actual)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("invalid number"))
		})

		It("should be able to encode and decode to/from json", func() {
			expected, _ := json2.NewNumericString("18446744073709551615")
			encoded, _ := jsoniter.Marshal(expected)

			var actual json2.Uint64
			err := jsoniter.Unmarshal(encoded, &actual)

			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal(expected.String()))
		})
	})
})
