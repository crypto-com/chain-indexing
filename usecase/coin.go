package usecase

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// Coin is a wrapper of CRO amount in basic unit
type Coin struct {
	value *big.Int
}

const (
	TOTAL_COIN_SUPPLY_STR = "100_000_000_000_0000_0000"
	MAX_COIN_DECIMALS_STR = "1_0000_0000"
)

var (
	reSign                      = `-?`
	reInt                       = `[[:digit:]]+`
	reDecimalPoint              = `\.?`
	reDecimal                   = `[[:digit:]]*`
	reCRO                       = regexp.MustCompile(fmt.Sprintf(`^(%s)(%s)%s(%s)$`, reSign, reInt, reDecimalPoint, reDecimal))
	TOTAL_COIN_SUPPLY_BIGINT, _ = new(big.Int).SetString(TOTAL_COIN_SUPPLY_STR, 0)
	MAX_COIN_DECIMALS_BIGINT, _ = new(big.Int).SetString(MAX_COIN_DECIMALS_STR, 0)
)

var (
	ErrCoinNil          = errors.New("big.Int has nil value")
	ErrCoinInvalid      = errors.New("invalid coin")
	ErrCoinExceedSupply = errors.New("exceed total supply")
)

// MustNewCoin() returns a new Coin from the provided big.Int. Any error would
// panic immediately
func MustNewCoin(value *big.Int) Coin {
	coin, err := NewCoin(value)
	if err != nil {
		panic(err)
	}

	return coin
}

// NewCoin() returns a new Coin from the provided big.Int
func NewCoin(value *big.Int) (Coin, error) {
	if value == nil {
		return Coin{}, ErrCoinNil
	}
	if value.CmpAbs(TOTAL_COIN_SUPPLY_BIGINT) == 1 {
		return Coin{}, ErrCoinExceedSupply
	}

	return Coin{
		value,
	}, nil
}

// MustNewCoinFromInt() accepts and parse a int64 coin amount to Coin and
// returns it. Any error would panic immediately
func MustNewCoinFromInt(value int64) Coin {
	var err error

	coin, err := NewCoinFromInt(value)
	if err != nil {
		panic(err)
	}
	return coin
}

// NewCoinFromInt() accepts and parse a int64 coin amount to Coin and returns
// it
func NewCoinFromInt(value int64) (Coin, error) {
	var err error

	b := big.NewInt(value)
	coin, err := NewCoin(b)
	if err != nil {
		return Coin{}, err
	}
	return coin, nil
}

// MustNewCoinFromString() accepts and parse a string to Coin and returns it.
// Any error would panic immediately
func MustNewCoinFromString(value string) Coin {
	var err error

	coin, err := NewCoinFromString(value)
	if err != nil {
		panic(err)
	}
	return coin
}

// NewCoinFromString() accepts and parse a string to Coin and returns it
func NewCoinFromString(value string) (Coin, error) {
	var err error

	b, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return Coin{}, ErrCoinInvalid
	}

	coin, err := NewCoin(b)
	if err != nil {
		return Coin{}, err
	}
	return coin, nil
}

// NewCoinFromCRO() accepts and parse CRO amount in string to Coin and
// returns it
func NewCoinFromCRO(value string) (Coin, error) {
	var err error

	matches := reCRO.FindStringSubmatch(value)
	if matches == nil {
		return Coin{}, fmt.Errorf("invalid CRO expression %s: %w", value, ErrCoinInvalid)
	}

	sign, intPart, decimalPart := matches[1], matches[2], matches[3]
	intCoin, ok := new(big.Int).SetString(intPart, 0)
	if !ok {
		return Coin{}, fmt.Errorf("invalid integer part in CRO expression %s: %w", value, ErrCoinInvalid)
	}

	unit := intCoin.Mul(intCoin, MAX_COIN_DECIMALS_BIGINT)

	totalDecimalLen := len(decimalPart)
	if totalDecimalLen > 8 {
		return Coin{}, fmt.Errorf("more than 8 decimal places in CRO expression %s: %w", value, ErrCoinInvalid)
	}
	decimalPart = strings.TrimLeft(decimalPart, "0")
	trimmedDecimalLen := len(decimalPart)

	if trimmedDecimalLen > 0 {
		decimalCoin, ok := new(big.Int).SetString(decimalPart, 0)
		if !ok {
			return Coin{}, fmt.Errorf("invalid decimal part in CRO expression %s: %w", value, ErrCoinInvalid)
		}
		if totalDecimalLen < 8 {
			// 0.00884 == 0.00884000 -> 884000
			decimalCoin = decimalCoin.Mul(
				decimalCoin,
				new(big.Int).Exp(
					big.NewInt(1),
					big.NewInt(8-int64(trimmedDecimalLen)),
					nil,
				),
			)
		}

		unit = intCoin.Add(intCoin, decimalCoin)
	}

	coin, err := NewCoin(unit)
	if err != nil {
		return Coin{}, err
	}

	if sign == "-" {
		coin = coin.Negate()
	}

	return coin, nil
}

// Add() add y to the coin receiver and returns a new coin value
func (coin *Coin) Add(y Coin) (Coin, error) {
	z := new(big.Int).Add(coin.value, y.value)
	if z.CmpAbs(TOTAL_COIN_SUPPLY_BIGINT) > 0 {
		return Coin{}, ErrCoinExceedSupply
	}

	return MustNewCoin(z), nil
}

// Negate() negates the coin value and returns a new coin value
func (coin *Coin) Negate() Coin {
	z := new(big.Int).Neg(coin.value)
	return MustNewCoin(z)
}

// String() returns the string representation of the Coin
func (coin *Coin) String() string {
	return coin.value.String()
}

// ToBigInt() returns a copy of the underlying big.Int
func (coin *Coin) ToBigInt() *big.Int {
	copy := new(big.Int).Set(coin.value)
	return copy
}

func (coin Coin) MarshalJSON() ([]byte, error) {
	if coin.value == nil {
		return []byte("null"), nil
	}

	s := coin.String()

	return []byte("\"" + s + "\""), nil
}

func (coin Coin) UnmarshalJSON(encoded []byte) ([]byte, error) {
	if coin.value == nil {
		return []byte("null"), nil
	}

	s := coin.String()

	return []byte("\"" + s + "\""), nil
}
