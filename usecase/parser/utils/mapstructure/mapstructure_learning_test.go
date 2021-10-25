package mapstructure_test

import (
	"github.com/mitchellh/mapstructure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mapstructure library learning test", func() {
	Context("when weakTypedInput is enabled", func() {
		It("should decode within bound number to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": 100,
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(100)))
		})

		It("should decode max uint32 number to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": uint32(4294967295),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(4294967295)))
		})

		It("should decode max uint16 number to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": uint16(65535),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(65535)))
		})

		It("should decode max uint16 number to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": 4294967295,
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(4294967295)))
		})

		It("should decode max uint64 number to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": uint64(18446744073709551615),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(18446744073709551615)))
		})

		It("should decode max uint64 string to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": "18446744073709551615",
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint64(18446744073709551615)))
		})

		It("should return error when decoding exceed max uint64 string to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": "18446744073709551615123456",
			})
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("value out of range"))
		})

		It("should decode max int64 number to int64", func() {
			type t struct {
				V int64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": int64(9223372036854775807),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(int64(9223372036854775807)))
		})

		It("should decode max int64 string to int64", func() {
			type t struct {
				V int64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": "9223372036854775807",
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(int64(9223372036854775807)))
		})

		It("should decode max int64 string to int64", func() {
			type t struct {
				V string `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": int64(9223372036854775807),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal("9223372036854775807"))
		})

		It("should return error when decoding exceed max uint64 string to uint64", func() {
			type t struct {
				V uint64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": "18446744073709551615123456",
			})
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("value out of range"))
		})

		It("should trim the number to max int64 when decoding max uint64 to int64", func() {
			type t struct {
				V int64 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": int64(9223372036854775807),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(int64(9223372036854775807)))

		})

		It("should decode max uint64 to string", func() {
			type t struct {
				V string `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": uint64(18446744073709551615),
			})
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal("18446744073709551615"))
		})

		It("should trim to max uint32 when decoding max uint64 number to uint32", func() {
			type t struct {
				V uint32 `mapstructure:"v"`
			}
			var v t

			decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				DecodeHook:       nil,
				Result:           &v,
			})

			err := decoder.Decode(map[string]interface{}{
				"v": uint64(18446744073709551615),
			})
			//Expect(err).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(v.V).To(Equal(uint32(4294967295)))
		})
	})

	It("should trim to max uint32 when decoding max uint64 number to uint32", func() {
		type t struct {
			V uint32 `mapstructure:"v"`
		}
		var v t

		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			WeaklyTypedInput: false,
			DecodeHook:       nil,
			Result:           &v,
		})

		err := decoder.Decode(map[string]interface{}{
			"v": uint64(18446744073709551615),
		})
		Expect(err).To(BeNil())
		Expect(v.V).To(Equal(uint32(4294967295)))
	})
})
