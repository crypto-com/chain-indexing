package json_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/external/json"
)

var _ = Describe("Json", func() {
	Describe("MarshalToString", func() {
		It("Should work", func() {
			Expect(json.MarshalToString(map[string]interface{}{
				"foo": 0,
				"bar": "baz",
			})).To(Equal("{\"bar\":\"baz\",\"foo\":0}"))
		})

		It("Should escape \u0000 to empty (NUL)", func() {
			Expect(json.MarshalToString(map[string]interface{}{
				"foo": 0,
				"bar": "\u0000baz\u0000",
			})).To(Equal("{\"bar\":\"baz\",\"foo\":0}"))
		})
	})
})
