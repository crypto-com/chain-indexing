package coin_test

import (
	"encoding/json"
	"errors"
	"math/big"

	usecase_coin "github.com/crypto-com/chain-indexing/usecase/coin"

	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coin", func() {
	Describe("NewCoin", func() {
		It("should return Error when the big.Int pointer is nil", func() {
			var anyNilBigInt *big.Int
			_, err := usecase_coin.NewCoin(anyNilBigInt)

			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, usecase_coin.ErrCoinNil)).To(BeTrue())
		})
	})

	Describe("NewCoinFromString", func() {
		It("should return Error when the string is invalid coin", func() {
			_, err := usecase_coin.NewCoinFromString("invalid")

			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, usecase_coin.ErrCoinInvalid)).To(BeTrue())
		})

		It("should return Error when the string has decimal place", func() {
			_, err := usecase_coin.NewCoinFromString("1.12345678")

			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, usecase_coin.ErrCoinInvalid)).To(BeTrue())

		})

		It("should return the Coin representation of the string", func() {
			anyCoin := "100000000"
			coin, err := usecase_coin.NewCoinFromString(anyCoin)

			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("100000000"))
		})
	})

	Describe("NewCoinFromCRO", func() {
		It("should return false when the CRO has more than 8 decimal places", func() {
			_, err := usecase_coin.NewCoinFromCRO("1.123456789")

			Expect(err).NotTo(BeNil())
			Expect(errors.Is(err, usecase_coin.ErrCoinInvalid)).To(BeTrue())
		})

		It("should return the Coin representation of a negative CRO", func() {
			var coin usecase_coin.Coin
			var err error

			coin, err = usecase_coin.NewCoinFromCRO("-99999999999.09999999")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("-9999999999909999999"))

			coin, err = usecase_coin.NewCoinFromCRO("-12345678")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("-1234567800000000"))

			coin, err = usecase_coin.NewCoinFromCRO("-0.00000884")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("-884"))
		})

		It("should return the Coin representation of a CRO integer", func() {
			coin, err := usecase_coin.NewCoinFromCRO("12345678")

			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("1234567800000000"))
		})

		It("should return the Coin representation of the CRO", func() {
			var coin usecase_coin.Coin
			var err error

			coin, err = usecase_coin.NewCoinFromCRO("0.12345678")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("12345678"))

			coin, err = usecase_coin.NewCoinFromCRO("0.00000884")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("884"))

			coin, err = usecase_coin.NewCoinFromCRO("0.00088400")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("88400"))
		})

		It("should return the Coin representation of a small CRO ", func() {
			coin, err := usecase_coin.NewCoinFromCRO("0.00000001")

			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("1"))
		})

		It("should return the Coin representation of a large CRO ", func() {
			var coin usecase_coin.Coin
			var err error

			coin, err = usecase_coin.NewCoinFromCRO("100000000000")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("10000000000000000000"))

			coin, err = usecase_coin.NewCoinFromCRO("99999999999.09999999")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("9999999999909999999"))

			coin, err = usecase_coin.NewCoinFromCRO("99999999999.00000884")
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("9999999999900000884"))
		})
	})

	Describe("WithinTotalSupply", func() {
		It("should return Error when the string exceeds maximum supply", func() {
			exceedTotalSupply := "10000000000000000001"
			result, err := usecase_coin.NewCoinFromString(exceedTotalSupply)

			Expect(err).To(BeNil())
			Expect(result.WithinTotalSupply()).To(BeFalse())
		})

		It("should return Error when the string negatively exceeds maximum supply", func() {
			exceedTotalSupply := "-10000000000000000001"
			result, err := usecase_coin.NewCoinFromString(exceedTotalSupply)

			Expect(err).To(BeNil())
			Expect(result.WithinTotalSupply()).To(BeFalse())
		})

		It("should return false when the CRO exceed maximum supply", func() {
			result, err := usecase_coin.NewCoinFromCRO("100000000000.1")

			Expect(err).To(BeNil())
			Expect(result.WithinTotalSupply()).To(BeFalse())
		})

		It("should return false when the CRO negatively exceed maximum supply", func() {
			result, err := usecase_coin.NewCoinFromCRO("-100000000000.1")

			Expect(err).To(BeNil())
			Expect(result.WithinTotalSupply()).To(BeFalse())
		})
		It("should return false when the add result exceeds total supply", func() {
			anyAmount := "10000000000000000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			oneUnit := "1"
			oneCoin, _ := usecase_coin.NewCoinFromString(oneUnit)

			result, err := coin.Add(oneCoin)
			Expect(err).To(BeNil())
			Expect(result.WithinTotalSupply()).To(BeFalse())
		})
	})

	Describe("Add", func() {
		It("should return the sum of the two coin values", func() {
			anyAmount := "10000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			oneUnit := "1"
			oneCoin, _ := usecase_coin.NewCoinFromString(oneUnit)

			actual, err := coin.Add(oneCoin)
			Expect(err).To(BeNil())
			Expect(actual.String()).To(Equal("10000001"))
		})

		It("should not update the original coin value", func() {
			anyAmount := "10000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			oneUnit := "1"
			oneCoin, _ := usecase_coin.NewCoinFromString(oneUnit)

			_, err := coin.Add(oneCoin)
			Expect(err).To(BeNil())
			Expect(coin.String()).To(Equal("10000000"))
		})
	})

	Describe("Negate", func() {
		It("should negate the underlying coin value when it is positive", func() {
			anyAmount := "10000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			actual := coin.Negate()
			Expect(actual.String()).To(Equal("-10000000"))
		})

		It("should convert negative coin to positive", func() {
			anyAmount := "-10000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			actual := coin.Negate()
			Expect(actual.String()).To(Equal("10000000"))
		})

		It("should not update the original coin value", func() {
			anyAmount := "10000000"
			coin, _ := usecase_coin.NewCoinFromString(anyAmount)

			_ = coin.Negate()
			Expect(coin.String()).To(Equal("10000000"))
		})
	})

	Describe("ToBigInt", func() {
		It("should return copy of the underlying big.Int", func() {
			anyCoin := "100000000"
			coin, _ := usecase_coin.NewCoinFromString(anyCoin)

			copiedCoin := coin.ToBigInt()
			_, _ = copiedCoin.SetString("12345678", 10)

			Expect(coin.String()).To(Equal("100000000"))
		})
	})

	Describe("jsoniter.Marshal", func() {
		It("should return null when the Coin value is nil", func() {
			coin := usecase_coin.Coin{}
			result, err := jsoniter.Marshal(coin)

			Expect(err).To(BeNil())
			Expect(result).To(Equal([]byte("null")))
		})

		It("should return the Coin amount in string", func() {
			coin, _ := usecase_coin.NewCoinFromString("12345678")
			result, err := jsoniter.Marshal(coin)

			Expect(err).To(BeNil())
			Expect(result).To(Equal([]byte("\"12345678\"")))
		})
	})

	Describe("json.Marshal", func() {
		It("should return null when the Coin value is nil", func() {
			coin := usecase_coin.Coin{}
			result, err := json.Marshal(coin)

			Expect(err).To(BeNil())
			Expect(result).To(Equal([]byte("null")))
		})

		It("should return the Coin amount in string", func() {
			coin, _ := usecase_coin.NewCoinFromString("12345678")
			result, err := json.Marshal(coin)

			Expect(err).To(BeNil())
			Expect(result).To(Equal([]byte("\"12345678\"")))
		})
	})
})
