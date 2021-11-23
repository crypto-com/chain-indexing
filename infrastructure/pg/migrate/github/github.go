package github

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func InitAndRunMigrate(
	migrationSourceURL string,
	databaseURL string,
) error {
	m, err := migrate.New(migrationSourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to init migration: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migration: %v", err)
	}

	return nil
}

// Generate the Github URL to migration source files
func GenerateSourceURL(
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
func GenerateDBConnString(
	connString string,
	migrationTableName string,
) string {
	if connString[len(connString)-1:] == "?" {
		return connString + "x-migrations-table=" + migrationTableName
	} else {
		return connString + "&x-migrations-table=" + migrationTableName
	}
}
