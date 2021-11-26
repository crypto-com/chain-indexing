package github

import (
	"errors"
	"fmt"

	"github.com/ettle/strcase"
	"github.com/golang-migrate/migrate/v4"
)

type GithubMigrationHelper struct {
	sourceURL   string
	databaseURL string
}

type Config struct {
	GithubAPIUser    string
	GithubAPIToken   string
	MigrationRepoRef string
	ConnString       string
}

const MIGRATION_GITHUB_URL_FORMAT = "github://%s:%s@crypto-com/chain-indexing/%s"
const MIGRATION_DIRECTORY_FORMAT = "projection/%s/migrations"

func NewGithubMigrationHelper(sourceURL string, databaseURL string) *GithubMigrationHelper {
	if sourceURL == "" {
		panic(fmt.Errorf("GithubMigrationHelper.SourceURL is empty when executing Migrate()"))
	}
	if databaseURL == "" {
		panic(fmt.Errorf("GithubMigrationHelper.DatabaseURL is empty when executing Migrate()"))
	}

	return &GithubMigrationHelper{
		sourceURL:   sourceURL,
		databaseURL: databaseURL,
	}
}

// Implement MigrationHelper interface
func (gmh *GithubMigrationHelper) Migrate() {
	m, err := migrate.New(gmh.sourceURL, gmh.databaseURL)
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

// Generate the default Github URL to migration source files
func GenerateDefaultSourceURL(
	projectionId string,
	config Config,
) string {
	projectionSnakeId := strcase.ToSnake(projectionId)
	migrationDirectory := fmt.Sprintf(MIGRATION_DIRECTORY_FORMAT, projectionSnakeId)

	return GenerateSourceURL(
		MIGRATION_GITHUB_URL_FORMAT,
		config.GithubAPIUser,
		config.GithubAPIToken,
		migrationDirectory,
		config.MigrationRepoRef,
	)
}
