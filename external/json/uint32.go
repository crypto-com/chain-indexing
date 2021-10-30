package json

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

type Uint32 struct {
	v uint32
}

func NewUint32(v uint32) Uint32 {
	return Uint32{v: v}
}

func (u Uint32) Uint32() uint32 {
	return u.v
}

func (u Uint32) String() string {
	return strconv.Itoa(int(u.v))
}

func (u Uint32) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *Uint32) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	var untypedVal interface{}
	if err := jsoniter.Unmarshal(data, &untypedVal); err != nil {
		return err
	}

	if stringVal, ok := untypedVal.(string); ok {
		typedVal, parseErr := strconv.ParseUint(stringVal, 10, 32)
		if parseErr != nil {
			return parseErr
		}
		u.v = uint32(typedVal)

		return nil
	}

	var uint32Val uint32
	if err := jsoniter.Unmarshal(data, &uint32Val); err != nil {
		return err
	}
	u.v = uint32Val

	return nil
}
