package filesystem

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

type FilesystemMigrationHelper struct {
	SourceURL   string
	DatabaseURL string
}

func NewFilesystemMigrationHelper(sourceURL, databaeURL string) *FilesystemMigrationHelper {
	return &FilesystemMigrationHelper{
		SourceURL:   sourceURL,
		DatabaseURL: databaeURL,
	}
}

// Implement MigrationHelper interface
func (fmh *FilesystemMigrationHelper) InitAndRunMigrate() {
	m, err := migrate.New(fmh.SourceURL, fmh.DatabaseURL)
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
