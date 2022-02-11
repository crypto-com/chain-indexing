package routes

import (
	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/tendermint"
	"github.com/crypto-com/chain-indexing/bootstrap"
	"github.com/crypto-com/chain-indexing/bootstrap/config"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	httpapi_handlers "github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	tendermint_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/tendermint"

	appconfig "github.com/crypto-com/chain-indexing/example/app/example-app/config"
	custom_httpapi_handlers "github.com/crypto-com/chain-indexing/example/httpapi/handlers"
)

func InitRouteRegistry(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *config.Config,
	customConfig *appconfig.CustomConfig,
) bootstrap.RouteRegistry {
	var cosmosAppClient cosmosapp.Client
	if config.CosmosApp.Insecure {
		cosmosAppClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
			config.CosmosApp.HTTPRPCUrl,
			config.Blockchain.BondingDenom,
		)
	} else {
		cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(
			config.CosmosApp.HTTPRPCUrl,
			config.Blockchain.BondingDenom,
		)
	}

	var tendermintClient tendermint.Client
	if config.TendermintApp.Insecure {
		tendermintClient = tendermint_infrastructure.NewInsecureHTTPClient(
			config.TendermintApp.HTTPRPCUrl,
			config.TendermintApp.StrictGenesisParsing,
		)
	} else {
		tendermintClient = tendermint_infrastructure.NewHTTPClient(
			config.TendermintApp.HTTPRPCUrl,
			config.TendermintApp.StrictGenesisParsing,
		)
	}

	accountAddressPrefix := config.Blockchain.AccountAddressPrefix
	validatorAddressPrefix := config.Blockchain.ValidatorAddressPrefix
	conNodeAddressPrefix := config.Blockchain.ConNodeAddressPrefix

	routes := make([]Route, 0)
	searchHandler := httpapi_handlers.NewSearch(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/search",
			handler: searchHandler.Search,
		},
	)

	blocksHandler := httpapi_handlers.NewBlocks(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/blocks",
			handler: blocksHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height-or-hash}",
			handler: blocksHandler.FindBy,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/transactions",
			handler: blocksHandler.ListTransactionsByHeight,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/events",
			handler: blocksHandler.ListEventsByHeight,
		},
		Route{
			Method:  GET,
			path:    "api/v1/blocks/{height}/commitments",
			handler: blocksHandler.ListCommitmentsByHeight,
		},
	)

	accountsHandlers := httpapi_handlers.NewAccounts(
		logger,
		rdbConn.ToHandle(),
		cosmosAppClient,
		config.Blockchain.ValidatorAddressPrefix,
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/accounts",
			handler: accountsHandlers.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/accounts/{account}",
			handler: accountsHandlers.FindBy,
		},
	)

	accountMessagesHandlers := httpapi_handlers.NewAccountMessages(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/accounts/{account}/messages",
			handler: accountMessagesHandlers.ListByAccount,
		},
	)

	accountTransactionsHandler := httpapi_handlers.NewAccountTransactions(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/accounts/{account}/transactions",
			handler: accountTransactionsHandler.ListByAccount,
		},
	)

	statusHandlers := httpapi_handlers.NewStatusHandler(logger, cosmosAppClient, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/status",
			handler: statusHandlers.GetStatus,
		},
	)

	transactionHandler := httpapi_handlers.NewTransactions(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/transactions",
			handler: transactionHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/transactions/{hash}",
			handler: transactionHandler.FindByHash,
		},
	)

	blockEventHandler := httpapi_handlers.NewBlockEvents(logger, rdbConn.ToHandle())
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/events",
			handler: blockEventHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/events/{id}",
			handler: blockEventHandler.FindById,
		},
	)

	proposalsHandler := httpapi_handlers.NewProposals(
		logger,
		rdbConn.ToHandle(),
		cosmosAppClient,
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/proposals",
			handler: proposalsHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/proposals/{id}",
			handler: proposalsHandler.FindById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/proposals/{id}/votes",
			handler: proposalsHandler.ListVotesById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/proposals/{id}/depositors",
			handler: proposalsHandler.ListDepositorsById,
		},
	)

	validatorsHandler := httpapi_handlers.NewValidators(
		logger,
		validatorAddressPrefix,
		conNodeAddressPrefix,
		cosmosAppClient,
		tendermintClient,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/validators",
			handler: validatorsHandler.List,
		},
		Route{
			Method:  GET,
			path:    "api/v1/validators/active",
			handler: validatorsHandler.ListActive,
		},
		Route{
			Method:  GET,
			path:    "api/v1/validators/{address}",
			handler: validatorsHandler.FindBy,
		},
		Route{
			Method:  GET,
			path:    "api/v1/validators/{address}/activities",
			handler: validatorsHandler.ListActivities,
		},
	)

	nftsHandler := httpapi_handlers.NewNFTs(
		logger,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/nfts/messages",
			handler: nftsHandler.ListMessages,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denom/name/{denomName}",
			handler: nftsHandler.FindDenomByName,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denom/id/{denomId}",
			handler: nftsHandler.FindDenomById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denom",
			handler: nftsHandler.ListDenoms,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/tokens",
			handler: nftsHandler.ListTokens,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}",
			handler: nftsHandler.FindDenomById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}/messages",
			handler: nftsHandler.ListMessagesByDenom,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}/tokens",
			handler: nftsHandler.ListTokensByDenomId,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}/tokens/{tokenId}",
			handler: nftsHandler.FindTokenById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}/tokens/{tokenId}/transfers",
			handler: nftsHandler.ListTransfersByToken,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/denoms/{denomId}/tokens/{tokenId}/messages",
			handler: nftsHandler.ListMessagesByToken,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/drops",
			handler: nftsHandler.ListDrops,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/drops/{drop}/tokens",
			handler: nftsHandler.ListTokensByDrop,
		},
		Route{
			Method:  GET,
			path:    "api/v1/nfts/accounts/{account}/tokens",
			handler: nftsHandler.ListTokensByAccount,
		},
	)

	ibcChannelHandler := httpapi_handlers.NewIBCChannel(
		logger,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/ibc/channels",
			handler: ibcChannelHandler.ListChannels,
		},
		Route{
			Method:  GET,
			path:    "api/v1/ibc/channels/{channelId}",
			handler: ibcChannelHandler.FindChannelById,
		},
		Route{
			Method:  GET,
			path:    "api/v1/ibc/denom-hash-mappings",
			handler: ibcChannelHandler.ListAllDenomHashMapping,
		},
	)

	ibcChannelMessageHandler := httpapi_handlers.NewIBCChannelMessage(
		logger,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/ibc/channels/{channelId}/messages",
			handler: ibcChannelMessageHandler.ListByChannelID,
		},
	)

	bridgesHandler := httpapi_handlers.NewBridges(
		logger,
		rdbConn.ToHandle(),
		accountAddressPrefix,
		appconfig.ParseBridgesConfig(&customConfig.BridgeAPI),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/bridges/txs/{hash}",
			handler: bridgesHandler.FindByTransactionHash,
		},
		Route{
			Method:  GET,
			path:    "api/v1/bridges/activities",
			handler: bridgesHandler.ListActivities,
		},
		Route{
			Method:  GET,
			path:    "api/v1/bridges/{network}/account/{account}/activities",
			handler: bridgesHandler.ListActivitiesByNetwork,
		},
	)

	exampleHandler := custom_httpapi_handlers.NewExample(
		logger,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/v1/examples",
			handler: exampleHandler.List,
		},
	)

	validatorDelegationHandler := httpapi_handlers.NewValidatorDelegation(
		logger,
		config.Blockchain.AccountAddressPrefix,
		config.Blockchain.ValidatorAddressPrefix,
		config.Blockchain.ConNodeAddressPrefix,
		rdbConn.ToHandle(),
	)
	routes = append(routes,
		Route{
			Method:  GET,
			path:    "api/test/validators",
			handler: validatorDelegationHandler.ListValidator,
		},
		Route{
			Method:  GET,
			path:    "api/test/validators/{address}",
			handler: validatorDelegationHandler.FindValidatorByAddress,
		},
		Route{
			Method:  GET,
			path:    "api/test/validators/{address}/delegations",
			handler: validatorDelegationHandler.ListDelegationByValidator,
		},
		Route{
			Method:  GET,
			path:    "api/test/validators/{address}/unbonding_delegations",
			handler: validatorDelegationHandler.ListUnbondingDelegationByValidator,
		},
		Route{
			Method:  GET,
			path:    "api/test/validators/{srcValAddress}/redelegations",
			handler: validatorDelegationHandler.ListRedelegationBySrcValidator,
		},
		Route{
			Method:  GET,
			path:    "api/test/delegators/{address}/delegations",
			handler: validatorDelegationHandler.ListDelegationByDelegator,
		},
	)

	return &RouteRegistry{routes: routes}
}
