package appinterface

import "errors"

var (
	ErrRepoPrepare = errors.New("error preparing repository operation")
	ErrRepoOpen    = errors.New("error opening repository")
	ErrRepoQuery   = errors.New("error querying from repository")
	ErrRepoWrite   = errors.New("error writing to repository")
)
