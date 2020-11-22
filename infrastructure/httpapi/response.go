package httpapi

import (
	pagination_interface "github.com/crypto-com/chainindex/appinterface/pagination"
	"github.com/valyala/fasthttp"

	jsoniter "github.com/json-iterator/go"
)

func Success(ctx *fasthttp.RequestCtx, result interface{}) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	err := jsoniter.NewEncoder(ctx.Response.BodyWriter()).Encode(Response{
		Result: result,
		Err:    "",
	})
	if err != nil {
		InternalServerError(ctx)
	}
}

func SuccessWithPagination(
	ctx *fasthttp.RequestCtx,
	result interface{},
	paginationResult *pagination_interface.PaginationResult,
) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	err := jsoniter.NewEncoder(ctx.Response.BodyWriter()).Encode(PagedResponse{
		Response: Response{
			Result: result,
			Err:    "",
		},
		OffsetPagination: OptPaginationOffsetResponseFromResult(paginationResult.OffsetResult()),
	})
	if err != nil {
		InternalServerError(ctx)
	}
}

func NotFound(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	message, err := jsoniter.Marshal(Response{
		Err: "Record not found",
	})
	if err != nil {
		InternalServerError(ctx)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusNotFound)
	ctx.Write(message)
}

func BadRequest(ctx *fasthttp.RequestCtx, errResp error) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	message, err := jsoniter.Marshal(Response{
		Err: errResp.Error(),
	})
	if err != nil {
		InternalServerError(ctx)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.Write(message)
}

func InternalServerError(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	message, _ := jsoniter.Marshal(Response{
		Err: ErrInternalServerError.Error(),
	})

	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.Write(message)
}

type PagedResponse struct {
	Response

	OffsetPagination *PaginationOffsetResponse `json:"pagination,omitempty"`
}

type Response struct {
	Result interface{} `json:"result"`
	Err    string      `json:"error,omitempty"`
}

type PaginationOffsetResponse struct {
	TotalRecord uint64 `json:"total_record"`
	TotalPage   uint64 `json:"total_page"`
	CurrentPage uint64 `json:"current_page"`
	Limit       uint64 `json:"limit"`
}

func OptPaginationOffsetResponseFromResult(
	offsetResult *pagination_interface.PaginationOffsetResult,
) *PaginationOffsetResponse {
	if offsetResult == nil {
		return nil
	}

	return &PaginationOffsetResponse{
		TotalRecord: offsetResult.TotalRecord,
		TotalPage:   offsetResult.TotalPage(),
		CurrentPage: offsetResult.CurrentPage,
		Limit:       offsetResult.Limit,
	}
}
