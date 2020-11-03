package pg

import (
	"fmt"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		if err == gomigrate.ErrNoChange {
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
		if err == gomigrate.ErrNoChange {
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
