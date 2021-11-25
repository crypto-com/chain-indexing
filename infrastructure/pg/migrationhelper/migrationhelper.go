package migrationhelper

import (
	"fmt"

	"github.com/ettle/strcase"
)

type MigrationHelper interface {
	Migrate()
}

const MIGRATION_TABLE_NAME_FORMAT = "%s_schema_migrations"

// Generate PostgreSQL DB conn string with customized migration table name
func GenerateDatabaseURL(
	connString string,
	migrationTableName string,
) string {
	if connString[len(connString)-1:] == "?" {
		return connString + "x-migrations-table=" + migrationTableName
	} else {
		return connString + "&x-migrations-table=" + migrationTableName
	}
}

// Generate default PostgreSQL DB conn string with customized migration table name
func GenerateDefaultDatabaseURL(
	projectionId string,
	connString string,
) string {
	projectionSnakeId := strcase.ToSnake(projectionId)
	migrationTableName := fmt.Sprintf(MIGRATION_TABLE_NAME_FORMAT, projectionSnakeId)

	return GenerateDatabaseURL(connString, migrationTableName)
}
