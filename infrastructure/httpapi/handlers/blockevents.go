package handlers

import (
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type BlockEvents struct {
	logger applogger.Logger

	blockEventsView *view.BlockEvents
}

func NewBlockEvents(logger applogger.Logger, rdbHandle *rdb.Handle) *BlockEvents {
	return &BlockEvents{
		logger.WithFields(applogger.LogFields{
			"module": "BlockEventsHandler",
		}),

		view.NewBlockEvents(rdbHandle),
	}
}

func (handler *BlockEvents) FindById(ctx *fasthttp.RequestCtx) {
	idParam := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid event id"))
		return
	}
	blockEvents, err := handler.blockEventsView.FindById(id)
	if err != nil {
		handler.logger.Errorf("error finding event by id: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, blockEvents)
}

func (handler *BlockEvents) List(ctx *fasthttp.RequestCtx) {
	var err error

	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	blockEvents, paginationResult, err := handler.blockEventsView.List(view.BlockEventsListFilter{
		MaybeBlockHeight: nil,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blockEvents, paginationResult)
}
