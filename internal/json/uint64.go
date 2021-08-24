package json

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

type Uint64 struct {
	v uint64
}

func NewUint64(v uint64) Uint64 {
	return Uint64{v: v}
}

func NewUint64FromString(v string) (*Uint64, error) {
	typedVal, parseErr := strconv.ParseUint(v, 10, 64)
	if parseErr != nil {
		return nil, parseErr
	}

	return &Uint64{v: typedVal}, nil
}

func (u Uint64) Uint64() uint64 {
	return u.v
}

func (u Uint64) String() string {
	return strconv.FormatUint(u.v, 10)
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

func (u *Uint64) UnmarshalJSON(data []byte) error {
	var untypedVal interface{}
	if err := jsoniter.Unmarshal(data, &untypedVal); err != nil {
		return err
	}

	if stringVal, ok := untypedVal.(string); ok {
		typedVal, parseErr := strconv.ParseUint(stringVal, 10, 64)
		if parseErr != nil {
			return parseErr
		}
		u.v = typedVal

		return nil
	}

	var uint64Val uint64
	if err := jsoniter.Unmarshal(data, &uint64Val); err != nil {
		return err
	}
	u.v = uint64Val

	return nil
}
