package handlers

import (
	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chainindex/appinterface/rdb"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type Info struct {
	logger applogger.Logger
}

func NewInfo(logger applogger.Logger, rdbHandle *rdb.Handle) *Info {
	return &Info{
		logger.WithFields(applogger.LogFields{
			"module": "BlocksHandler",
		}),
	}
}

func (handler *Info) GetLatestHeight(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("1"))
}
