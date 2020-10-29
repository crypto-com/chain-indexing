package appinterface

import (
	"errors"
	"fmt"
)

var (
	ErrRepoPrepare = errors.New("error preparing repository operation")
	ErrRepoOpen    = errors.New("error opening repository")
	ErrRepoQuery   = errors.New("error querying from repository")
	ErrRepoWrite   = errors.New("error writing to repository")

	ErrBuildSQLStmt = fmt.Errorf("error building SQL statement: %w", ErrRepoPrepare)
	ErrTypeConv     = fmt.Errorf("error converting between types: %w", ErrRepoPrepare)

	// when trying to scan a null row
	ErrNoRows = errors.New("no rows in result set")
)
