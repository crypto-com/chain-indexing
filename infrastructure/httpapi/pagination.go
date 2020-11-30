package httpapi

import (
	"strconv"

	"github.com/valyala/fasthttp"

	pagination_interface "github.com/crypto-com/chain-indexing/appinterface/pagination"
)

func ParsePagination(ctx *fasthttp.RequestCtx) (*pagination_interface.Pagination, error) {
	var err error

	var pagination pagination_interface.PaginationType
	var page, limit int64

	queryArgs := NewQueryArgs(ctx.QueryArgs())

	pagination = queryArgs.Get("pagination")
	if pagination == "" {
		pagination = pagination_interface.PAGINATION_OFFSET
	}
	if pagination != pagination_interface.PAGINATION_OFFSET {
		return nil, ErrInvalidPagination
	}

	pageQuery := queryArgs.Get("page")
	if pageQuery == "" {
		page = int64(1)
	} else {
		page, err = strconv.ParseInt(pageQuery, 10, 64)
		if err != nil {
			return nil, ErrInvalidPage
		}
		if page == 0 {
			return nil, ErrInvalidPage
		}
	}

	limitQuery := queryArgs.Get("limit")
	if limitQuery == "" {
		limit = int64(20)
	} else {
		limit, err = strconv.ParseInt(limitQuery, 10, 64)
		if err != nil {
			return nil, ErrInvalidPage
		}
	}

	return pagination_interface.NewOffsetPagination(page, limit), nil
}
