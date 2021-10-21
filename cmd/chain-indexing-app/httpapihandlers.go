package main

import (
	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	httpapi_handlers "github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

func initHTTPAPIHanlders(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *bootstrap.Config,
) []bootstrap.Handler {
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

	handlers := make([]bootstrap.Handler, 0)
	handlers = append(handlers, httpapi_handlers.NewSearch(logger, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewBlocks(logger, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewAccounts(
		logger,
		rdbConn.ToHandle(),
		cosmosAppClient,
		config.Blockchain.ValidatorAddressPrefix,
	))
	handlers = append(handlers, httpapi_handlers.NewAccountMessages(logger, rdbConn.ToHandle()))

	return handlers
}
