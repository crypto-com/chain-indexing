package json

import (
	"fmt"
	"math/big"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

// NumericString is numeric value stored in string representation. It supports
// JSON unmarshalling from numeric and string.
type NumericString struct {
	v string
}

func NewNumericString(v string) (*NumericString, error) {
	_, ok := new(big.Int).SetString(v, 10)
	if !ok {
		return nil, fmt.Errorf("invalid number")
	}
	return &NumericString{v: v}, nil
}

func NewNumericStringFromUint64(v uint64) *NumericString {
	return &NumericString{v: strconv.FormatUint(v, 10)}
}

func (u NumericString) String() string {
	return u.v
}

func (u NumericString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

func (u *NumericString) UnmarshalJSON(data []byte) error {
	var untypedVal interface{}
	if err := jsoniter.Unmarshal(data, &untypedVal); err != nil {
		return err
	}

	if stringVal, ok := untypedVal.(string); ok {
		_, ok := new(big.Int).SetString(stringVal, 10)
		if !ok {
			return fmt.Errorf("invalid number")
		}
		u.v = stringVal

		return nil
	}

	var uint64Val uint64
	parseUint64ValErr := jsoniter.Unmarshal(data, &uint64Val)
	if parseUint64ValErr == nil {
		u.v = strconv.FormatUint(uint64Val, 10)
		return nil
	}

	var int64Val int64
	if err := jsoniter.Unmarshal(data, &int64Val); err != nil {
		return err
	}
	u.v = strconv.FormatInt(int64Val, 10)

	return nil
}
