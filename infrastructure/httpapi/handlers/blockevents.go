package handlers

import (
	"errors"
	"strconv"

	"github.com/crypto-com/chainindex/appinterface/projection/view"

	"github.com/valyala/fasthttp"

	blockevents_view "github.com/crypto-com/chainindex/appinterface/projection/blockevent/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type BlockEvents struct {
	logger applogger.Logger

	blockEventsView *blockevents_view.BlockEvents
}

func NewBlockEvents(logger applogger.Logger, rdbHandle *rdb.Handle) *BlockEvents {
	return &BlockEvents{
		logger.WithFields(applogger.LogFields{
			"module": "BlockEventsHandler",
		}),

		blockevents_view.NewBlockEvents(rdbHandle),
	}
}

func (handler *BlockEvents) FindById(ctx *fasthttp.RequestCtx) {
	idParam, _ := ctx.UserValue("id").(string)
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		httpapi.BadRequest(ctx, errors.New("invalid event id"))
		return
	}
	blockEvents, err := handler.blockEventsView.FindById(id)
	if err != nil {
		if err == rdb.ErrNoRows {
			httpapi.NotFound(ctx)
			return
		}
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

	heightOrder := view.ORDER_ASC
	queryArgs := ctx.QueryArgs()
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			heightOrder = view.ORDER_DESC
		}
	}

	blockEvents, paginationResult, err := handler.blockEventsView.List(blockevents_view.BlockEventsListFilter{
		MaybeBlockHeight: nil,
	}, blockevents_view.BlockEventsListOrder{
		Height: heightOrder,
	}, pagination)
	if err != nil {
		handler.logger.Errorf("error listing events: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blockEvents, paginationResult)
}
