package filesystem

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

type FilesystemMigrationHelper struct {
	sourceURL   string
	databaseURL string
}

func NewFilesystemMigrationHelper(sourceURL, databaseURL string) *FilesystemMigrationHelper {
	if sourceURL == "" {
		panic(fmt.Errorf("FilesystemMigrationHelper.SourceURL is empty when executing Migrate()"))
	}
	if databaseURL == "" {
		panic(fmt.Errorf("FilesystemMigrationHelper.DatabaseURL is empty when executing Migrate()"))
	}

	return &FilesystemMigrationHelper{
		sourceURL:   sourceURL,
		databaseURL: databaseURL,
	}
}

// Implement MigrationHelper interface
func (fmh *FilesystemMigrationHelper) Migrate() {
	m, err := migrate.New(fmh.sourceURL, fmh.databaseURL)
	if err != nil {
		panic(fmt.Errorf("failed to init migration: %v", err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Errorf("failed to run migration: %v", err))
	}
}

// Generate the source URL to migration source files.
// `sourceFolder` could be absolute or relative path.
func GenerateSourceURL(sourceFolder string) string {
	return fmt.Sprintf("file://%s", sourceFolder)
}
