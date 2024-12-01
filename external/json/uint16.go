package json

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

type Uint16 struct {
	v uint16
}

func NewUint16(v uint16) Uint16 {
	return Uint16{v: v}
}

func (u Uint16) Uint16() uint16 {
	return u.v
}

func (u Uint16) String() string {
	return strconv.Itoa(int(u.v))
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	// uint16 number can be handled by all languages as number
	return []byte(u.String()), nil
}

func (u *Uint16) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	var untypedVal interface{}
	if err := jsoniter.Unmarshal(data, &untypedVal); err != nil {
		return err
	}

	if stringVal, ok := untypedVal.(string); ok {
		typedVal, parseErr := strconv.ParseUint(stringVal, 10, 16)
		if parseErr != nil {
			return parseErr
		}
		u.v = uint16(typedVal)

		return nil
	}

	var uint16Val uint16
	if err := jsoniter.Unmarshal(data, &uint16Val); err != nil {
		return err
	}
	u.v = uint16Val

	return nil
}
