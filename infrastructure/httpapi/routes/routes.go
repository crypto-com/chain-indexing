package routes

import (
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	"github.com/crypto-com/chainindex/infrastructure/httpapi/handlers"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	blocksHandler *handlers.Blocks
}

func NewRoutesRegistry(
	blocksHandler *handlers.Blocks,
) *RouteRegistry {
	return &RouteRegistry{
		blocksHandler,
	}
}

func (registry *RouteRegistry) Register(server *httpapi.Server) {
	server.GET("/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})
	server.GET("/api/v1/blocks", registry.blocksHandler.List)
}
