package utctime_test

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/external/utctime"
)

var _ = Describe("UTCTime", func() {
	Describe("Add", func() {
		It("should work", func() {
			t := utctime.FromUnixNano(int64(100000))
			result := t.Add(123 * time.Second)
			Expect(result.UnixNano()).To(Equal(int64(123000100000)))
		})
	})

	Describe("JSON en/decoding", func() {
		It("should encode to the unix nano time", func() {
			t := utctime.FromUnixNano(1000000)
			actual, err := jsoniter.Marshal(t)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal([]byte("\"1970-01-01T00:00:00.001Z\"")))
		})

		It("should decode from time.Time string", func() {
			var err error
			t, _ := jsoniter.Marshal(time.Unix(0, 1000000).UTC())

			actual, err := utctime.UnmarshalJSON(t)
			Expect(err).To(BeNil())
			Expect(actual.UnixNano()).To(Equal(int64(1000000)))
		})

		It("should be able to encode and decode to/from json", func() {
			expected := utctime.FromUnixNano(1000000)
			encoded, _ := jsoniter.Marshal(expected)

			actual, err := utctime.UnmarshalJSON(encoded)

			Expect(err).To(BeNil())
			Expect(actual.UnixNano()).To(Equal(expected.UnixNano()))
		})
	})
})
