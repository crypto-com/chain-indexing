package hex

import "encoding/hex"

func MustDecodeFromString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}
