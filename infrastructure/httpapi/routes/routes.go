package routes

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	searchHandler      *handlers.Search
	blocksHandler      *handlers.Blocks
	statusHandler      *handlers.StatusHandler
	transactionHandler *handlers.Transactions
	blockEventHandler  *handlers.BlockEvents
	validatorsHandler  *handlers.Validators
	accountsHandler    *handlers.Accounts
}

func NewRoutesRegistry(
	searchHandler *handlers.Search,
	blocksHandler *handlers.Blocks,
	statusHandler *handlers.StatusHandler,
	transactionHandler *handlers.Transactions,
	blockEventHandler *handlers.BlockEvents,
	validatorsHandler *handlers.Validators,
	accountsHandler *handlers.Accounts,
) *RouteRegistry {
	return &RouteRegistry{
		searchHandler,
		blocksHandler,
		statusHandler,
		transactionHandler,
		blockEventHandler,
		validatorsHandler,
		accountsHandler,
	}
}

func (registry *RouteRegistry) Register(server *httpapi.Server, routePrefix string) {
	if routePrefix == "/" {
		routePrefix = ""
	}

	server.GET(fmt.Sprintf("%s/api/v1/health", routePrefix), func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})
	server.GET(fmt.Sprintf("%s/api/v1/search", routePrefix), registry.searchHandler.Search)
	server.GET(fmt.Sprintf("%s/api/v1/blocks", routePrefix), registry.blocksHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height-or-hash}", routePrefix), registry.blocksHandler.FindBy)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height}/transactions", routePrefix), registry.blocksHandler.ListTransactionsByHeight)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height}/events", routePrefix), registry.blocksHandler.ListEventsByHeight)
	server.GET(fmt.Sprintf("%s/api/v1/status", routePrefix), registry.statusHandler.GetStatus)
	server.GET(fmt.Sprintf("%s/api/v1/transactions", routePrefix), registry.transactionHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/transactions/{hash}", routePrefix), registry.transactionHandler.FindByHash)
	server.GET(fmt.Sprintf("%s/api/v1/events", routePrefix), registry.blockEventHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/events/{id}", routePrefix), registry.blockEventHandler.FindById)
	server.GET(fmt.Sprintf("%s/api/v1/validators", routePrefix), registry.validatorsHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/validators/active", routePrefix), registry.validatorsHandler.ListActive)
	server.GET(fmt.Sprintf("%s/api/v1/validators/{address}", routePrefix), registry.validatorsHandler.FindBy)
	server.GET(fmt.Sprintf("%s/api/v1/validators/{address}/activities", routePrefix), registry.validatorsHandler.ListActivities)
	server.GET(fmt.Sprintf("%s/api/v1/accounts", routePrefix), registry.accountsHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/accounts/{address}", routePrefix), registry.accountsHandler.FindBy)

}
