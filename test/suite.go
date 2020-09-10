package test

import (
	"os"

	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/internal/primptr"
	"github.com/crypto-com/chainindex/internal/typeconv"
)

func WithTestPgConnConfig(body func(*infrastructure.PgConnConfig)) bool {
	ssl := true
	if os.Getenv("TEST_POSTGRES_SSL") == "0" {
		ssl = false
	}
	config := infrastructure.PgConnConfig{
		Host:          "127.0.0.1",
		Port:          typeconv.MustAtou32(os.Getenv("TEST_POSTGRES_PORT")),
		MaybeUsername: primptr.String(os.Getenv("TEST_POSTGRES_USERNAME")),
		MaybePassword: primptr.String(os.Getenv("TEST_POSTGRES_PASSWORD")),
		Database:      os.Getenv("TEST_POSTGRES_DATABASE"),
		SSL:           ssl,
	}

	body(&config)

	return true
}

func WithTestPgx(body func(infrastructure.PgConnConfig)) bool {
	ssl := true
	if os.Getenv("TEST_POSTGRES_SSL") == "0" {
		ssl = false
	}
	config := infrastructure.PgConnConfig{
		Host:          "127.0.0.1",
		Port:          typeconv.MustAtou32(os.Getenv("TEST_POSTGRES_PORT")),
		MaybeUsername: primptr.String(os.Getenv("TEST_POSTGRES_USERNAME")),
		MaybePassword: primptr.String(os.Getenv("TEST_POSTGRES_PASSWORD")),
		Database:      os.Getenv("TEST_POSTGRES_DATABASE"),
		SSL:           ssl,
	}

	body(config)

	return true
}
