package rdb

import (
	"errors"
	"fmt"
)

var (
	ErrPrepare = errors.New("error preparing statement")
	ErrOpen    = errors.New("error opening database connection")
	ErrQuery   = errors.New("error querying from table")
	ErrWrite   = errors.New("error writing to table")

	ErrBuildSQLStmt = fmt.Errorf("error building SQL statement: %w", ErrPrepare)
	ErrTypeConv     = fmt.Errorf("error converting between types: %w", ErrPrepare)

	// when trying to scan a null row
	ErrNoRows = errors.New("no rows in result set")
)
