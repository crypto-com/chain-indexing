package github

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	appprojection "github.com/crypto-com/chain-indexing/projection"
)

type GithubMigrationHelper struct {
	Config Config

	MaybeSourceURL   *string
	MaybeDatabaseURL *string
}

type Config struct {
	appprojection.Config
	ConnString string
}

const MIGRATION_GITHUB_URL_FORMAT = "github://%s:%s@crypto-com/chain-indexing/%s"
const MIGRATION_DIRECTORY_FORMAT = "projection/%s/migrations"
const MIGRATION_TABLE_NAME_FORMAT = "%s_schema_migrations"

func NewGithubMigrationHelper(
	config Config,
	maybeSourceURL *string,
	maybeDatabaeURL *string,
) *GithubMigrationHelper {
	return &GithubMigrationHelper{
		Config: config,

		MaybeSourceURL:   maybeSourceURL,
		MaybeDatabaseURL: maybeDatabaeURL,
	}
}

// Implement MigrationHelper interface
func (gmh *GithubMigrationHelper) InitAndRunMigrate() {
	if gmh.MaybeSourceURL == nil {
		panic(fmt.Errorf("GithubMigrationHelper.MaybeSourceURL is nil when executing InitAndRunMigrate"))
	}
	if gmh.MaybeDatabaseURL == nil {
		panic(fmt.Errorf("GithubMigrationHelper.MaybeDatabaseURL is nil when executing InitAndRunMigrate"))
	}

	m, err := migrate.New(*gmh.MaybeSourceURL, *gmh.MaybeDatabaseURL)
	if err != nil {
		panic(fmt.Errorf("failed to init migration: %v", err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Errorf("failed to run migration: %v", err))
	}
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
