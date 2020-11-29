package factory

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"

	random "github.com/brianvoe/gofakeit/v5"
)

func addBlockFuncLookup() {
	random.AddFuncLookup("blockheight", random.Info{
		Category:    "custom",
		Description: "Random block height",
		Example:     "1",
		Output:      "int64",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			return RandomBlockHeight(), nil
		},
	})

	random.AddFuncLookup("blockhash", random.Info{
		Category:    "custom",
		Description: "Random block hash",
		Example:     "FCC8CE74C8B7B4B5212A57813F4025F16BEAA22D296B0421FACB7FE374A4CA2F",
		Output:      "string",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			return RandomBlockHash(), nil
		},
	})

	random.AddFuncLookup("apphash", random.Info{
		Category:    "custom",
		Description: "Random app hash",
		Example:     "8E32E5DECD0B9AAA675150BCB1EF2FE54BD246FD19BA7A6A7CE80C495CF6A4FD",
		Output:      "string",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			return RandomAppHash(), nil
		},
	})

	random.AddFuncLookup("validatoraddress", random.Info{
		Category:    "custom",
		Description: "Random validator address",
		Example:     "3D947104BCE9B9997333447E07AEAE52FFE02A2A",
		Output:      "string",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			return RandomValidatorAddress(), nil
		},
	})

	random.AddFuncLookup("commitsignature", random.Info{
		Category:    "custom",
		Description: "Random commit signature",
		Example:     "Wy8w7tD6yLWmT8bNoYNtkIXApbYvzzSMRUnkgZRgmmcgFSbCrU0fpgafVbKfMUuxJS2byiNqx6N2NE4jE+sFDA==",
		Output:      "string",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			return RandomCommitSignature(), nil
		},
	})

}

func RandomBlockHeight() int64 {
	// nolint:gosec
	return rand.Int63()
}

func RandomBlockHash() string {
	return strings.ToUpper(hex.EncodeToString(randomHex(20)))
}

func RandomAppHash() string {
	return strings.ToUpper(hex.EncodeToString(randomHex(32)))
}

func RandomValidatorAddress() string {
	return strings.ToUpper(hex.EncodeToString(randomHex(32)))
}

func RandomCommitSignature() string {
	return base64.StdEncoding.EncodeToString(randomHex(20))
}

func RandomTxHash() string {
	return strings.ToUpper(hex.EncodeToString(randomHex(32)))
}

func randomHex(n int) []byte {
	placeholder := make([]byte, n)
	// nolint:gosec
	_, err := rand.Read(placeholder)
	if err != nil {
		panic(err)
	}
	return placeholder
}
