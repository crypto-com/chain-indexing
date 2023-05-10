package handlers

import (
	applogger "github.com/crypto-com/chain-indexing/external/logger"

	status_polling "github.com/crypto-com/chain-indexing/appinterface/polling"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/valyala/fasthttp"
)

type HealthHandler struct {
	logger applogger.Logger

	statusView *status_polling.Status
}

func NewHealthHandler(logger applogger.Logger, rdbHandle *rdb.Handle) *HealthHandler {
	return &HealthHandler{
		logger.WithFields(applogger.LogFields{
			"module": "HealthHandler",
		}),
		status_polling.NewStatus(rdbHandle),
	}
}

func (handler *HealthHandler) HealthCheck(ctx *fasthttp.RequestCtx) {
	_, err := handler.statusView.FindBy("LatestHeight")
	if err != nil {
		handler.logger.Errorf("error fetching latest height: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.Success(ctx, "OK")
}
