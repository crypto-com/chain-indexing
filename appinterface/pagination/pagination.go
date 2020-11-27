package pagination

import "math"

// Pagination stores pagination request data
type Pagination struct {
	t PaginationType

	offsetParams PaginationOffsetParams
	// TODO: Cursor pagination
	// cursorParams *PaginationCursorParams
}

func NewOffsetPagination(page int64, limit int64) *Pagination {
	return &Pagination{
		t: PAGINATION_OFFSET,

		offsetParams: PaginationOffsetParams{
			Page:  page,
			Limit: limit,
		},
	}
}

func (pagination *Pagination) Type() PaginationType {
	return pagination.t
}

func (pagination *Pagination) OffsetParams() *PaginationOffsetParams {
	if pagination.Type() != PAGINATION_OFFSET {
		return nil
	}
	return &pagination.offsetParams
}

// TODO: Cursor pagination
// func (pagination *Pagination) CursorParams() *PaginationCursorParams {
// 	return &pagination.cursorParams
// }

func (pagination *Pagination) OffsetResult(totalRecord int64) *PaginationResult {
	if pagination.Type() != PAGINATION_OFFSET {
		return nil
	}
	return NewOffsetPaginationResult(
		totalRecord,
		pagination.offsetParams.Page,
		pagination.offsetParams.Limit,
	)
}

type PaginationOffsetParams struct {
	Page  int64
	Limit int64
}

func (params *PaginationOffsetParams) Offset() int64 {
	return params.Limit * (params.Page - 1)
}

// TODO: Cursor pagination
// type PaginationCursorParams struct {
// 	Direction PaginationCursorDirection
// 	CursorId  interface{}
// }

// type PaginationCursorDirection = int8

// const (
// 	PAGINATION_CURSOR_NEXT = 0
// 	PAGINATION_CURSOR_BACK = 1
// )

type PaginationResult struct {
	t PaginationType

	offsetResult PaginationOffsetResult
}

func NewOffsetPaginationResult(totalRecord int64, currentPage int64, limit int64) *PaginationResult {
	return &PaginationResult{
		t: PAGINATION_OFFSET,

		offsetResult: PaginationOffsetResult{
			TotalRecord: totalRecord,
			CurrentPage: currentPage,
			Limit:       limit,
		},
	}
}

func (result *PaginationResult) Type() PaginationType {
	return result.t
}

func (result *PaginationResult) OffsetResult() *PaginationOffsetResult {
	if result.Type() != PAGINATION_OFFSET {
		return nil
	}
	return &result.offsetResult
}

type PaginationOffsetResult struct {
	TotalRecord int64
	CurrentPage int64
	Limit       int64
}

func (result *PaginationOffsetResult) TotalPage() int64 {
	return int64(math.Ceil(float64(result.TotalRecord) / float64(result.Limit)))
}

type PaginationType = string

const (
	PAGINATION_OFFSET PaginationType = "offset"
	// PAGINATION_CURSOR
)
