package pg

import (
	"errors"
	"fmt"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	appprojection "github.com/crypto-com/chain-indexing/projection"
)

type Migrate struct {
	*gomigrate.Migrate

	config       *ConnConfig
	sourceFolder string

	reCreateClientOnReset bool
}

func MustNewMigrate(config *ConnConfig, sourceFolder string) *Migrate {
	pgMigrate, err := NewMigrate(config, sourceFolder)
	if err != nil {
		panic(err)
	}

	return pgMigrate
}

func NewMigrate(config *ConnConfig, sourceFolder string) (*Migrate, error) {
	m, err := _newMigrate(config, sourceFolder)
	if err != nil {
		return nil, err
	}

	reCreateClientOnReset := true
	return &Migrate{
		m,

		config,
		sourceFolder,

		reCreateClientOnReset,
	}, nil
}

func _newMigrate(config *ConnConfig, sourceFolder string) (*gomigrate.Migrate, error) {
	return gomigrate.New(
		fmt.Sprintf("file://%s", sourceFolder),
		config.ToURL(),
	)
}

func (m *Migrate) SetReCreateOnReset(shouldReset bool) {
	m.reCreateClientOnReset = shouldReset
}

func (m *Migrate) MustUp() {
	if err := m.Up(); err != nil {
		panic(err)
	}
}

func (m *Migrate) Up() error {
	if err := m.Migrate.Up(); err != nil {
		if errors.Is(err, gomigrate.ErrNoChange) {
			return nil
		}

		return err
	}

	return nil
}

func (m *Migrate) MustDown() {
	if err := m.Down(); err != nil {
		panic(err)
	}
}

func (m *Migrate) Down() error {
	if err := m.Migrate.Down(); err != nil {
		if errors.Is(err, gomigrate.ErrNoChange) {
			return nil
		}

		return err
	}

	return nil
}

func (m *Migrate) MustReset() {
	if err := m.Reset(); err != nil {
		panic(err)
	}
}

func (m *Migrate) Reset() error {
	if err := m.Force(0); err != nil {
		return err
	}
	if err := m.Drop(); err != nil {
		return err
	}

	// https://github.com/golang-migrate/migrate/issues/226#issuecomment-494496734
	// There is a known bug of "golang-migrate" that after `Drop()`, "schema_migrations" won't be
	// create by `Up()`. The solution is to re-create the "golang-migrate" client
	if m.reCreateClientOnReset {
		migrate, err := _newMigrate(m.config, m.sourceFolder)
		if err != nil {
			return err
		}
		m.Migrate = migrate
	}

	return nil
}

// Get the migration source files from Github, init the migrate and then run it.
func InitAndRunMigrateFromGithub(
	rdbConn rdb.Conn,
	migrationGithubTarget string,
	config *appprojection.Config,
	migrationDirectory string,
	migrationTableName string,
) error {
	migrationSourceURL := GithubMigrationSourceURL(
		migrationGithubTarget,
		config.GithubAPIUser,
		config.GithubAPIToken,
		migrationDirectory,
		config.MigrationRepoRef,
	)
	databaseURL := MigrationDBConnString(
		rdbConn.(*PgxConn).ConnString(),
		migrationTableName,
	)
	if err := InitAndRunMigrate(migrationSourceURL, databaseURL); err != nil {
		return err
	}
	return nil
}

func InitAndRunMigrate(
	migrationSourceURL string,
	databaseURL string,
) error {
	m, err := gomigrate.New(migrationSourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to init migration: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("failed to run migration: %v", err)
	}

	return nil
}

// Generate the Github URL to migration source files
func GithubMigrationSourceURL(
	format string,
	apiUser string,
	apiToken string,
	migrationDirectory string,
	migrationRepoRef string,
) string {
	ref := ""
	if migrationRepoRef != "" {
		ref = "#" + migrationRepoRef
	}

	return fmt.Sprintf(format, apiUser, apiToken, migrationDirectory+ref)
}

// Generate DB conn string with customized migration table name
func MigrationDBConnString(
	connString string,
	migrationTableName string,
) string {
	if connString[len(connString)-1:] == "?" {
		return connString + "x-migrations-table=" + migrationTableName
	} else {
		return connString + "&x-migrations-table=" + migrationTableName
	}
}
