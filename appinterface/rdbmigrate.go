package appinterface

type RDbMigrate interface {
	// Up() looks at the currently active migration version and migrate all the
	// way up.
	Up() error
	// Down() looks at the currently active migration version and migrate all
	// the way down.
	Down() error
	// Steps() looks at the currently active migration version and migrate up
	// or down according to the steps delta.
	// If the delta exceed the maximum number of up/down migrations that can be
	// executed, it will execute all the migrations and return an error.
	Steps(delta int) error
	// Reset() deletes everything in the database (including migration
	// histories).
	Reset() error
	// Version() returns the current active migration version and whether the
	// database is at dirty state
	Version() (version uint, dirty bool, err error)
}
