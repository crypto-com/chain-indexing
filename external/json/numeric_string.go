package json

import (
	"fmt"
	"math/big"

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

func NewNumericStringFromBigInt(v *big.Int) *NumericString {
	return &NumericString{v: v.String()}
}

func (u NumericString) String() string {
	return u.v
}

func (u NumericString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

func (u NumericString) UnmarshalJSON(data []byte) error {
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

	var bigIntVal big.Int
	parseBigIntValErr := jsoniter.Unmarshal(data, &bigIntVal)
	if parseBigIntValErr == nil {
		u.v = bigIntVal.String()
		return nil
	}

	return nil
}
