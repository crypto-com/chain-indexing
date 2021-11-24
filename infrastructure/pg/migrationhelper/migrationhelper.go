package migrationhelper

type MigrationHelper interface {
	InitAndRunMigrate()
}

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
