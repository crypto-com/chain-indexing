package json

import jsoniter "github.com/json-iterator/go"

func MustMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		panic(err)
	}

	return s
}

func MustMarshal(v interface{}) []byte {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}

func MustUnmarshalFromString(s string, v interface{}) {
	if err := jsoniter.UnmarshalFromString(s, v); err != nil {
		panic(err)
	}
}

func MustUnmarshal(b []byte, v interface{}) {
	if err := jsoniter.Unmarshal(b, v); err != nil {
		panic(err)
	}
}
