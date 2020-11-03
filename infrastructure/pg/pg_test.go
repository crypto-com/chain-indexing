package pg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chainindex/infrastructure/pg"
	"github.com/crypto-com/chainindex/internal/primptr"
)

var _ = Describe("ConnConfig", func() {
	Describe("ToURL", func() {
		It("should return Postgres connection string without auth when no username is provided", func() {
			config := ConnConfig{
				Host:          "127.0.0.1",
				Port:          5432,
				MaybeUsername: nil,
				MaybePassword: nil,
				Database:      "chainindex",
				SSL:           false,
			}

			Expect(config.ToURL()).To(Equal("postgres://127.0.0.1:5432/chainindex?sslmode=disable"))
		})

		It("should return Postgres connection string without auth when username and password is provided", func() {
			config := ConnConfig{
				Host:          "127.0.0.1",
				Port:          5432,
				MaybeUsername: primptr.String("user"),
				MaybePassword: primptr.String("password"),
				Database:      "chainindex",
				SSL:           false,
			}

			Expect(config.ToURL()).To(Equal("postgres://user:password@127.0.0.1:5432/chainindex?sslmode=disable"))
		})

		It("should return Postgres connection string when SSL is enabled", func() {
			config := ConnConfig{
				Host:          "127.0.0.1",
				Port:          5432,
				MaybeUsername: primptr.String("user"),
				MaybePassword: primptr.String("password"),
				Database:      "chainindex",
				SSL:           true,
			}

			Expect(config.ToURL()).To(Equal("postgres://user:password@127.0.0.1:5432/chainindex"))
		})
	})
})
