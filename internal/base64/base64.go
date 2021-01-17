package base64

import "encoding/base64"

func MustDecodeString(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}
