package httpapi

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")

	ErrInvalidPagination = errors.New("invalid pagination type")
	ErrInvalidPage       = errors.New("invalid page number")
	ErrInvalidLimit      = errors.New("invalid page limit")

	ErrInvalidQuery = errors.New("invalid query parameter")
)
