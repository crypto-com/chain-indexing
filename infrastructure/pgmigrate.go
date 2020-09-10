package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PgMigrate struct {
	*migrate.Migrate
}

func NewPgMigrate(config *PgConnConfig, sourceFolder string) (*PgMigrate, error) {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", sourceFolder),
		config.ToURL(),
	)
	if err != nil {
		return nil, err
	}

	return &PgMigrate{
		m,
	}, nil
}

func (m *PgMigrate) Up() error {
	if err := m.Migrate.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	return nil
}

func (m *PgMigrate) Down() error {
	if err := m.Migrate.Down(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	return nil
}

func (m *PgMigrate) Reset() error {
	if err := m.Force(0); err != nil {
		return err
	}
	if err := m.Drop(); err != nil {
		return err
	}

	return nil
}
