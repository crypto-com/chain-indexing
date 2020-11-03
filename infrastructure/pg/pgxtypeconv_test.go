package pg_test

import (
	"math/big"

	"github.com/jackc/pgtype"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/infrastructure/pg"
	"github.com/crypto-com/chainindex/internal/utctime"
)

var _ = Describe("PgxTypeConv", func() {
	var pgxTypeConv pg.PgxTypeConv

	BeforeEach(func() {
		//nolint:staticcheck
		pgxTypeConv = pg.PgxTypeConv{}
	})

	Describe("Bton", func() {
		It("should return numeric null when bigInt is nil", func() {
			var expected pgtype.Numeric
			_ = expected.Set(nil)

			actual := pgxTypeConv.Bton(nil)
			Expect(actual).To(Equal(expected))
		})

		It("should return numeric of the bigInt value", func() {
			var expected pgtype.Numeric
			_ = expected.Scan("10")

			actual := pgxTypeConv.Bton(big.NewInt(10))
			Expect(actual).To(Equal(expected))
		})

		It("should work for negative number", func() {
			var expected pgtype.Numeric
			_ = expected.Scan("-10")

			actual := pgxTypeConv.Bton(big.NewInt(-10))
			Expect(actual).To(Equal(expected))
		})
	})

	Describe("Iton", func() {
		It("should return numeric representation of the int", func() {
			var expected pgtype.Numeric
			_ = expected.Set(10)

			Expect(pgxTypeConv.Iton(10)).To(Equal(expected))
		})

		It("should work for negative number", func() {
			var expected pgtype.Numeric
			_ = expected.Set(-10)

			Expect(pgxTypeConv.Iton(-10)).To(Equal(expected))
		})
	})

	Describe("NtobReader", func() {
		It("should return Error for decimal number", func() {
			var n pgtype.Numeric
			_ = n.Set(1.23456)

			ntobReader := pgxTypeConv.NtobReader()
			arg, _ := ntobReader.ScannableArg().(*pgtype.Numeric)
			*arg = n

			_, err := ntobReader.Parse()
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("cannot convert 123456e-5 to bigInt"))
		})

		It("should return nil when numeric is null", func() {
			var n pgtype.Numeric
			_ = n.Set(nil)

			ntobReader := pgxTypeConv.NtobReader()
			arg, _ := ntobReader.ScannableArg().(*pgtype.Numeric)
			*arg = n

			actual, err := ntobReader.Parse()
			Expect(err).To(BeNil())
			Expect(actual).To(BeNil())
		})

		It("should return big.Int of the numeric value", func() {
			var n pgtype.Numeric
			_ = n.Set(10)

			expected := big.NewInt(10)

			ntobReader := pgxTypeConv.NtobReader()
			arg, _ := ntobReader.ScannableArg().(*pgtype.Numeric)
			*arg = n

			actual, err := ntobReader.Parse()
			Expect(err).To(BeNil())
			Expect(*actual).To(Equal(*expected))
		})

		It("should work for negative number", func() {
			var n pgtype.Numeric
			_ = n.Set(-10)

			expected := big.NewInt(-10)

			ntobReader := pgxTypeConv.NtobReader()
			arg, _ := ntobReader.ScannableArg().(*pgtype.Numeric)
			*arg = n

			actual, err := ntobReader.Parse()
			Expect(err).To(BeNil())
			Expect(*actual).To(Equal(*expected))
		})
	})

	Describe("Tton", func() {
		It("should return nil when time is nil", func() {
			Expect(pgxTypeConv.Tton(nil)).To(BeNil())
		})

		It("should convert time to native time type", func() {
			expected := int64(707810766000000000)
			maybeAnyTime := utctime.FromUnixNano(expected)

			actual := pgxTypeConv.Tton(&maybeAnyTime)
			Expect(actual).To(Equal(expected))
		})
	})

	Describe("NtotReader", func() {
		It("should return nil when time is null", func() {
			reader := pgxTypeConv.NtotReader()

			arg := reader.ScannableArg().(**int64)
			*arg = nil

			actual, err := reader.Parse()
			Expect(err).To(BeNil())
			Expect(actual).To(BeNil())
		})

		It("should return utctime.UTCTime of the time value", func() {
			expected := utctime.FromUnixNano(707810766000000000)

			reader := pgxTypeConv.NtotReader()
			arg := reader.ScannableArg().(**int64)
			**arg = 707810766000000000

			actual, err := reader.Parse()
			Expect(err).To(BeNil())
			Expect(*actual).To(Equal(expected))
		})
	})
})
