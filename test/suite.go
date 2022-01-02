package test

import (
	"fmt"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	. "github.com/crypto-com/chain-indexing/external/logger/test"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/typeconv"
	"os"
	"path"
	"runtime"
)

var DEFAULT_MIGRATIONS_TABLE_NAME = "schema_migrations"
var DEFAULT_MIGRATIONS_FILE_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../../migrations")
}()

func WithTestPgConnConfig(body func(*pg.ConnConfig)) bool {
	testDB := os.Getenv("TEST_DB")
	if testDB == "" || testDB == "0" {
		fmt.Println("Skipping DB testing")
		return true
	}

	host := os.Getenv("TEST_POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	ssl := true
	if os.Getenv("TEST_POSTGRES_SSL") == "0" {
		ssl = false
	}
	config := pg.ConnConfig{
		Host:          host,
		Port:          typeconv.MustAtoi32(os.Getenv("TEST_POSTGRES_PORT")),
		MaybeUsername: primptr.String(os.Getenv("TEST_POSTGRES_USERNAME")),
		MaybePassword: primptr.String(os.Getenv("TEST_POSTGRES_PASSWORD")),
		Database:      os.Getenv("TEST_POSTGRES_DATABASE"),
		SSL:           ssl,
	}

	body(&config)

	return true
}

func WithTestPgxConn(body func(*pg.PgxConn, rdb.Migrate)) bool {
	return WithTestPgConnConfig(func(config *pg.ConnConfig) {
		fakeLogger := NewFakeLogger()
		conn := pg.MustNewPgxConn(config, fakeLogger)

		migrate := pg.MustNewMigrate(config, DEFAULT_MIGRATIONS_FILE_PATH)

		body(conn, migrate)
	})
}

func WithProjectionTestEnv(body func(ProjectionTestEnv)) bool {
	return WithTestPgConnConfig(func(config *pg.ConnConfig) {
		fakeLogger := NewFakeLogger()
		conn := pg.MustNewPgxConn(config, fakeLogger)

		migrateCreator := func(tableName, filePath string) rdb.Migrate {
			migrateConfig := config.DeepClone()
			migrateConfig.MaybeXMigrationsTable = &tableName
			return pg.MustNewMigrate(&migrateConfig, filePath)
		}

		body(ProjectionTestEnv{
			Conn: conn,
			RootMigrate: migrateCreator(
				DEFAULT_MIGRATIONS_TABLE_NAME,
				DEFAULT_MIGRATIONS_FILE_PATH,
			),
			MigrateCreator: migrateCreator,
		})
	})
}

type MigrateCreator func(tableName, filePath string) rdb.Migrate

type ProjectionTestEnv struct {
	Conn           *pg.PgxConn
	RootMigrate    rdb.Migrate
	MigrateCreator MigrateCreator
}
