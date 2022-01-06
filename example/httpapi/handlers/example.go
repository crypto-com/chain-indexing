package handlers

import (
	"github.com/crypto-com/chain-indexing/appinterface/projection/view"
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	example_view "github.com/crypto-com/chain-indexing/example/projection/example/view"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
)

type Example struct {
	logger applogger.Logger

	exampleView example_view.Examples
}

func NewExample(logger applogger.Logger, rdbHandle *rdb.Handle) *Example {
	return &Example{
		logger.WithFields(applogger.LogFields{
			"module": "ExampleModule",
		}),

		example_view.NewExamplesView(rdbHandle),
	}
}

func (handler *Example) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	addressOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "address.desc" {
			addressOrder = view.ORDER_DESC
		}
	}

	exampleRows, paginationResult, err := handler.exampleView.List(example_view.ExamplesListOrder{
		Address: addressOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing addresses: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, exampleRows, paginationResult)
}
