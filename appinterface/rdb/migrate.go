package rdb

type Migrate interface {
	// Up() looks at the currently active migration version and migrate all the
	// way up.
	Up() error
	MustUp()
	// Down() looks at the currently active migration version and migrate all
	// the way down.
	Down() error
	MustDown()
	// Reset() deletes everything in the database (including migration
	// histories).
	Reset() error
	MustReset()
	// Version() returns the current active migration version and whether the
	// database is at dirty state
	Version() (version uint, dirty bool, err error)
}
