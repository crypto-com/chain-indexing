package test

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/internal/logger/test"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
)

func WithTestPgConnConfig(body func(*pg.ConnConfig)) bool {
	testDB := os.Getenv("TEST_DB")
	if testDB == "" || testDB == "0" {
		fmt.Println("Skipping DB testing")
		return true
	}

	ssl := true
	if os.Getenv("TEST_POSTGRES_SSL") == "0" {
		ssl = false
	}
	config := pg.ConnConfig{
		Host:          "127.0.0.1",
		Port:          typeconv.MustAtoi32(os.Getenv("TEST_POSTGRES_PORT")),
		MaybeUsername: primptr.String(os.Getenv("TEST_POSTGRES_USERNAME")),
		MaybePassword: primptr.String(os.Getenv("TEST_POSTGRES_PASSWORD")),
		Database:      os.Getenv("TEST_POSTGRES_DATABASE"),
		SSL:           ssl,
	}

	body(&config)

	return true
}

func WithTestPgxConn(body func(*pg.PgxConn, *pg.Migrate)) bool {
	testDB := os.Getenv("TEST_DB")
	if testDB == "" || testDB == "0" {
		fmt.Println("Skipping DB testing")
		return true
	}

	ssl := true
	if os.Getenv("TEST_POSTGRES_SSL") == "0" {
		ssl = false
	}
	config := &pg.ConnConfig{
		Host:          "127.0.0.1",
		Port:          typeconv.MustAtoi32(os.Getenv("TEST_POSTGRES_PORT")),
		MaybeUsername: primptr.String(os.Getenv("TEST_POSTGRES_USERNAME")),
		MaybePassword: primptr.String(os.Getenv("TEST_POSTGRES_PASSWORD")),
		Database:      os.Getenv("TEST_POSTGRES_DATABASE"),
		SSL:           ssl,
	}
	fakeLogger := NewFakeLogger()
	conn := pg.MustNewPgxConn(config, fakeLogger)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}
	migrate := pg.MustNewMigrate(config, path.Join(filename, "../../migrations"))

	body(conn, migrate)

	return true
}
