package github

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	appprojection "github.com/crypto-com/chain-indexing/projection"
)

type GithubMigrationHelper struct {
	Config Config

	SourceURL   string
	DatabaseURL string
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
	sourceURL string,
	databaeURL string,
) *GithubMigrationHelper {
	return &GithubMigrationHelper{
		Config: config,

		SourceURL:   sourceURL,
		DatabaseURL: databaeURL,
	}
}

// Implement MigrationHelper interface
func (gmh *GithubMigrationHelper) Migrate() {
	if gmh.SourceURL == "" {
		panic(fmt.Errorf("GithubMigrationHelper.SourceURL is empty when executing Migrate()"))
	}
	if gmh.DatabaseURL == "" {
		panic(fmt.Errorf("GithubMigrationHelper.DatabaseURL is empty when executing Migrate()"))
	}

	m, err := migrate.New(gmh.SourceURL, gmh.DatabaseURL)
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
