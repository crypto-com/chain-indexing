package routes

import (
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	"github.com/crypto-com/chainindex/infrastructure/httpapi/handlers"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	blocksHandler      *handlers.Blocks
	statusHandler      *handlers.StatusHandler
	transactionHandler *handlers.Transactions
	blockEventHandler  *handlers.BlockEvents
	validatorsHandler  *handlers.Validators
}

func NewRoutesRegistry(
	blocksHandler *handlers.Blocks,
	statusHandler *handlers.StatusHandler,
	transactionHandler *handlers.Transactions,
	blockEventHandler *handlers.BlockEvents,
	validatorsHandler *handlers.Validators,
) *RouteRegistry {
	return &RouteRegistry{
		blocksHandler,
		statusHandler,
		transactionHandler,
		blockEventHandler,
		validatorsHandler,
	}
}

func (registry *RouteRegistry) Register(server *httpapi.Server) {
	server.GET("/api/v1/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})
	server.GET("/api/v1/blocks", registry.blocksHandler.List)
	server.GET("/api/v1/blocks/{height-or-hash}", registry.blocksHandler.FindBy)
	server.GET("/api/v1/blocks/{height}/transactions", registry.blocksHandler.ListTransactionsByHeight)
	server.GET("/api/v1/blocks/{height}/events", registry.blocksHandler.ListEventsByHeight)
	server.GET("/api/v1/status", registry.statusHandler.GetStatus)
	server.GET("/api/v1/transactions", registry.transactionHandler.List)
	server.GET("/api/v1/transactions/{hash}", registry.transactionHandler.FindByHash)
	server.GET("/api/v1/events", registry.blockEventHandler.List)
	server.GET("/api/v1/events/{id}", registry.blockEventHandler.FindById)
	server.GET("/api/v1/validators", registry.validatorsHandler.List)
	server.GET("/api/v1/validators/{operator_address}", registry.validatorsHandler.FindBy)
}
