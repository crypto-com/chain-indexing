package handlers

import (
	"fmt"
	applogger "github.com/crypto-com/chain-indexing/external/logger"

	"github.com/valyala/fasthttp"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
)

func URLValueGuard(ctx *fasthttp.RequestCtx, logger applogger.Logger, key string) (string, bool) {
	userValue := ctx.UserValue(key)
	if userValue == nil {
		logger.Errorf("missing URL value: %s, make sure it is registered in route definition", key)
		httpapi.InternalServerError(ctx)
		return "", false
	}

	value, castOk := userValue.(string)
	if !castOk {
		logger.Errorf("error casting URL value %s to string", key)
		httpapi.InternalServerError(ctx)
		return "", false
	}

	if value == "" {
		httpapi.BadRequest(ctx, fmt.Errorf("missing %s in URL", key))
		return "", false
	}

	return value, true
}
