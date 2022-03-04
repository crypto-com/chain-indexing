package json

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func MustMarshalToString(v interface{}) string {
	s, err := MarshalToString(v)
	if err != nil {
		panic(err)
	}

	return s
}

func MarshalToString(v interface{}) (string, error) {
	s, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(v)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(s, "\\u0000", ""), nil
}

func MustMarshal(v interface{}) []byte {
	b, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}

func MustUnmarshalFromString(s string, v interface{}) {
	if err := UnmarshalFromString(s, v); err != nil {
		panic(err)
	}
}

func UnmarshalFromString(str string, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.UnmarshalFromString(str, v)
}

func MustUnmarshal(b []byte, v interface{}) {
	if err := Unmarshal(b, v); err != nil {
		panic(err)
	}
}

func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}
