package pg

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/jackc/pgtype"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
)

var (
	BIGINT_INT0  = big.NewInt(0)
	BIGINT_INT10 = big.NewInt(10)
)

type PgxTypeConv struct{}

func (conv *PgxTypeConv) Bton(b *big.Int) interface{} {
	var err error

	var nilValue pgtype.Numeric
	_ = nilValue.Set(nil)
	if b == nil {
		return nilValue
	}

	var num pgtype.Numeric
	if err = num.Set(b.String()); err != nil {
		panic(fmt.Sprintf("cannot convert %v to numeric: %v", b, err))
	}

	return num
}
func (conv *PgxTypeConv) Iton(i int) interface{} {
	var num pgtype.Numeric
	if err := num.Set(i); err != nil {
		panic(fmt.Sprintf("cannot convert %v to numeric: %v", i, err))
	}

	return num
}
func (conv *PgxTypeConv) NtobReader() rdb.NtobReader {
	return new(PgxNtobReader)
}

type PgxNtobReader struct {
	num pgtype.Numeric
}

func (reader *PgxNtobReader) ScannableArg() interface{} {
	return &reader.num
}
func (reader *PgxNtobReader) Parse() (*big.Int, error) {
	var err error

	value, err := reader.num.Value()
	if err != nil {
		return nil, err
	}
	switch str := value.(type) {
	case string:
		// pgtype.Numeric.Amount() returns scientific notation e.g. "1000e0".
		i, err := sciToBigIntPtr(str)
		if err != nil {
			return nil, err
		}
		return i, nil
	case nil:
		return nil, nil
	default:
		return nil, errors.New("unknown pgtype value")
	}
}
func sciToBigIntPtr(sci string) (*big.Int, error) {
	var err error
	var ok bool

	parts := strings.Split(sci, "e")
	if len(parts) != 2 {
		return nil, errors.New("non-scientific notation value error")
	}

	intval, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, errors.New("integer part to bigInt error")
	}
	if parts[1] == "0" {
		return intval, nil
	}

	exp, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("exponent part to bigInt error: %v", err)
	}

	num := &big.Int{}
	num.Set(intval)
	if exp > 0 {
		mul := &big.Int{}
		mul.Exp(BIGINT_INT10, big.NewInt(exp), nil)
		num.Mul(num, mul)
		return num, nil
	}

	div := &big.Int{}
	div.Exp(BIGINT_INT10, big.NewInt(int64(-exp)), nil)
	remainder := &big.Int{}
	num.DivMod(num, div, remainder)
	if remainder.Cmp(BIGINT_INT0) != 0 {
		return nil, fmt.Errorf("cannot convert %v to bigInt", sci)
	}
	return num, nil
}

func (conv *PgxTypeConv) Tton(t *utctime.UTCTime) interface{} {
	if t == nil {
		return nil
	}
	return t.UnixNano()
}

func (conv *PgxTypeConv) NtotReader() rdb.NtotReader {
	return NewPgxNtotReader()
}

type PgxNtotReader struct {
	unixNano *int64
}

func NewPgxNtotReader() *PgxNtotReader {
	var i int64
	return &PgxNtotReader{
		unixNano: &i,
	}
}

func (reader *PgxNtotReader) ScannableArg() interface{} {
	return &reader.unixNano
}
func (reader *PgxNtotReader) Parse() (*utctime.UTCTime, error) {
	if reader.unixNano == nil {
		return nil, nil
	}
	t := utctime.FromUnixNano(*reader.unixNano)
	return &t, nil
}
