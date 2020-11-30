package factory

import (
	"math/rand"

	random "github.com/brianvoe/gofakeit/v5"
)

func addGoTypesFuncLookup() {
	random.AddFuncLookup("+int", random.Info{
		Category:    "custom",
		Description: "Random positive int",
		Example:     "0",
		Output:      "int",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			// nolint:gosec
			return rand.Int(), nil
		},
	})

	random.AddFuncLookup("+int64", random.Info{
		Category:    "custom",
		Description: "Random positive int64",
		Example:     "0",
		Output:      "int64",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			// nolint:gosec
			return rand.Int63(), nil
		},
	})

	random.AddFuncLookup("+int32", random.Info{
		Category:    "custom",
		Description: "Random positive int32",
		Example:     "0",
		Output:      "int32",
		Call: func(m *map[string][]string, info *random.Info) (interface{}, error) {
			// nolint:gosec
			return rand.Int31(), nil
		},
	})
}
