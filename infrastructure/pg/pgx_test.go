package pg_test

import (
	"time"

	"github.com/crypto-com/chain-indexing/external/primptr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crypto-com/chain-indexing/infrastructure/pg"
)

var _ = Describe("PgxConnPoolConfig", func() {
	Describe("ToURL", func() {
		Context("When no config is provided", func() {
			It("should return Postgres connection string without auth when no username is provided", func() {
				config := PgxConnPoolConfig{
					ConnConfig: ConnConfig{
						Host:          "127.0.0.1",
						Port:          5432,
						MaybeUsername: nil,
						MaybePassword: nil,
						Database:      "chain-indexing",
						SSL:           false,
					},
					MaybeMaxConns:          nil,
					MaybeMinConns:          nil,
					MaybeMaxConnLifeTime:   nil,
					MaybeMaxConnIdleTime:   nil,
					MaybeHealthCheckPeriod: nil,
				}

				Expect(config.ToURL()).To(Equal("postgres://127.0.0.1:5432/chain-indexing?sslmode=disable"))
			})

			It("should return Postgres connection string without auth when username and password is provided", func() {
				config := PgxConnPoolConfig{
					ConnConfig: ConnConfig{
						Host:          "127.0.0.1",
						Port:          5432,
						MaybeUsername: primptr.String("user"),
						MaybePassword: primptr.String("password"),
						Database:      "chain-indexing",
						SSL:           false,
					},
					MaybeMaxConns:          nil,
					MaybeMinConns:          nil,
					MaybeMaxConnLifeTime:   nil,
					MaybeMaxConnIdleTime:   nil,
					MaybeHealthCheckPeriod: nil,
				}

				Expect(config.ToURL()).To(Equal("postgres://user:password@127.0.0.1:5432/chain-indexing?sslmode=disable"))
			})

			It("should return Postgres connection string when SSL is enabled", func() {
				config := PgxConnPoolConfig{
					ConnConfig: ConnConfig{
						Host:          "127.0.0.1",
						Port:          5432,
						MaybeUsername: primptr.String("user"),
						MaybePassword: primptr.String("password"),
						Database:      "chain-indexing",
						SSL:           true,
					},
					MaybeMaxConns:          nil,
					MaybeMinConns:          nil,
					MaybeMaxConnLifeTime:   nil,
					MaybeMaxConnIdleTime:   nil,
					MaybeHealthCheckPeriod: nil,
				}

				Expect(config.ToURL()).To(Equal("postgres://user:password@127.0.0.1:5432/chain-indexing"))
			})
		})

		Context("When config is provided", func() {
			It("should render the config in the connection string", func() {
				config := PgxConnPoolConfig{
					ConnConfig: ConnConfig{
						Host:          "127.0.0.1",
						Port:          5432,
						MaybeUsername: primptr.String("user"),
						MaybePassword: primptr.String("password"),
						Database:      "chain-indexing",
						SSL:           true,
					},
					MaybeMaxConns:          primptr.Int32(50),
					MaybeMinConns:          primptr.Int32(1),
					MaybeMaxConnLifeTime:   primptr.Duration(300 * time.Second),
					MaybeMaxConnIdleTime:   primptr.Duration(60 * time.Second),
					MaybeHealthCheckPeriod: primptr.Duration(30 * time.Second),
				}

				Expect(config.ToURL()).To(Equal("postgres://user:password@127.0.0.1:5432/chain-indexing?pool_health_check_period=30s&pool_max_conn_idle_time=1m0s&pool_max_conn_lifetime=5m0s&pool_max_conns=50&pool_min_conns=1"))
			})
		})
	})
})
