package routes

import (
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	"github.com/crypto-com/chainindex/infrastructure/httpapi/handlers"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	blocksHandler      *handlers.Blocks
	transactionHandler *handlers.Transactions
	infoHandler        *handlers.Info
}

func NewRoutesRegistry(
	blocksHandler *handlers.Blocks,
	transactionHandler *handlers.Transactions,
	infoHandler *handlers.Info,
) *RouteRegistry {
	return &RouteRegistry{
		blocksHandler,
		transactionHandler,
		infoHandler,
	}
}

func (registry *RouteRegistry) Register(server *httpapi.Server) {
	server.GET("/api/v1/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})
	server.GET("/api/v1/getlatestheight", registry.infoHandler.GetLatestHeight)
	server.GET("/api/v1/blocks", registry.blocksHandler.List)
	server.GET("/api/v1/blocks/{height}/transactions", registry.blocksHandler.ListTransactionsByHeight)
	server.GET("/api/v1/transactions", registry.transactionHandler.List)
}
