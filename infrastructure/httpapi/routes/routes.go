package routes

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	searchHandler              *handlers.Search
	blocksHandler              *handlers.Blocks
	statusHandler              *handlers.StatusHandler
	transactionHandler         *handlers.Transactions
	blockEventHandler          *handlers.BlockEvents
	validatorsHandler          *handlers.Validators
	accountTransactionsHandler *handlers.AccountTransactions
	accountMessagesHandler     *handlers.AccountMessages
	accountsHandler            *handlers.Accounts
	proposalsHandler           *handlers.Proposals
	nftsHandler                *handlers.NFTs
}

func NewRoutesRegistry(
	searchHandler *handlers.Search,
	blocksHandler *handlers.Blocks,
	statusHandler *handlers.StatusHandler,
	transactionHandler *handlers.Transactions,
	blockEventHandler *handlers.BlockEvents,
	validatorsHandler *handlers.Validators,
	accountTransactionsHandler *handlers.AccountTransactions,
	accountMessagesHandler *handlers.AccountMessages,
	accountsHandler *handlers.Accounts,
	proposalsHandler *handlers.Proposals,
	nftsHandler *handlers.NFTs,
) *RouteRegistry {
	return &RouteRegistry{
		searchHandler,
		blocksHandler,
		statusHandler,
		transactionHandler,
		blockEventHandler,
		validatorsHandler,
		accountTransactionsHandler,
		accountMessagesHandler,
		accountsHandler,
		proposalsHandler,
		nftsHandler,
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
	server.GET(fmt.Sprintf("%s/api/v1/accounts", routePrefix), registry.accountsHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/accounts/{account}", routePrefix), registry.accountsHandler.FindBy)
	server.GET(fmt.Sprintf("%s/api/v1/accounts/{account}/transactions", routePrefix), registry.accountTransactionsHandler.ListByAccount)
	server.GET(fmt.Sprintf("%s/api/v1/accounts/{account}/messages", routePrefix), registry.accountMessagesHandler.ListByAccount)
	server.GET(fmt.Sprintf("%s/api/v1/blocks", routePrefix), registry.blocksHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height-or-hash}", routePrefix), registry.blocksHandler.FindBy)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height}/transactions", routePrefix), registry.blocksHandler.ListTransactionsByHeight)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height}/events", routePrefix), registry.blocksHandler.ListEventsByHeight)
	server.GET(fmt.Sprintf("%s/api/v1/blocks/{height}/commitments", routePrefix), registry.blocksHandler.ListCommitmentsByHeight)
	server.GET(fmt.Sprintf("%s/api/v1/events", routePrefix), registry.blockEventHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/events/{id}", routePrefix), registry.blockEventHandler.FindById)
	server.GET(fmt.Sprintf("%s/api/v1/proposals", routePrefix), registry.proposalsHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/proposals/{id}", routePrefix), registry.proposalsHandler.FindById)
	server.GET(fmt.Sprintf("%s/api/v1/proposals/{id}/votes", routePrefix), registry.proposalsHandler.ListVotesById)
	server.GET(fmt.Sprintf("%s/api/v1/proposals/{id}/depositors", routePrefix), registry.proposalsHandler.ListDepositorsById)
	server.GET(fmt.Sprintf("%s/api/v1/status", routePrefix), registry.statusHandler.GetStatus)
	server.GET(fmt.Sprintf("%s/api/v1/transactions", routePrefix), registry.transactionHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/transactions/{hash}", routePrefix), registry.transactionHandler.FindByHash)
	server.GET(fmt.Sprintf("%s/api/v1/validators", routePrefix), registry.validatorsHandler.List)
	server.GET(fmt.Sprintf("%s/api/v1/validators/active", routePrefix), registry.validatorsHandler.ListActive)
	server.GET(fmt.Sprintf("%s/api/v1/validators/{address}", routePrefix), registry.validatorsHandler.FindBy)
	server.GET(fmt.Sprintf("%s/api/v1/validators/{address}/activities", routePrefix), registry.validatorsHandler.ListActivities)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/transfers", routePrefix), registry.nftsHandler.ListTransfers)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/denoms", routePrefix), registry.nftsHandler.ListDenoms)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/denoms/{denomId}", routePrefix), registry.nftsHandler.FindDenomById)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/denoms/{denomId}/tokens", routePrefix), registry.nftsHandler.ListTokensByDenomId)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/denoms/{denomId}/tokens/{tokenId}", routePrefix), registry.nftsHandler.FindTokenById)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/denoms/{denomId}/tokens/{tokenId}/transfers", routePrefix), registry.nftsHandler.ListTransfersByToken)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/drops", routePrefix), registry.nftsHandler.ListDrops)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/drops/{drop}/tokens", routePrefix), registry.nftsHandler.ListTokensByDrop)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/accounts/{account}/tokens", routePrefix), registry.nftsHandler.ListTokensByAccount)
	server.GET(fmt.Sprintf("%s/api/v1/nfts/accounts/{account}/tokens/transfers", routePrefix), registry.nftsHandler.ListTransfersByAccount)

}
