package main

import (
	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/appinterface/tendermint"
	"github.com/crypto-com/chain-indexing/bootstrap"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	httpapi_handlers "github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	tendermint_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/tendermint"
)

func initHTTPAPIHandlers(
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

	var tendermintClient tendermint.Client
	if config.Tendermint.Insecure {
		tendermintClient = tendermint_infrastructure.NewInsecureHTTPClient(
			config.Tendermint.HTTPRPCUrl,
			config.Tendermint.StrictGenesisParsing,
		)
	} else {
		tendermintClient = tendermint_infrastructure.NewHTTPClient(
			config.Tendermint.HTTPRPCUrl,
			config.Tendermint.StrictGenesisParsing,
		)
	}

	accountAddressPrefix := config.Blockchain.AccountAddressPrefix
	validatorAddressPrefix := config.Blockchain.ValidatorAddressPrefix
	conNodeAddressPrefix := config.Blockchain.ConNodeAddressPrefix

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
	handlers = append(handlers, httpapi_handlers.NewStatusHandler(logger, cosmosAppClient, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewTransactions(logger, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewBlockEvents(logger, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewValidators(
		logger,
		validatorAddressPrefix,
		conNodeAddressPrefix,
		cosmosAppClient,
		tendermintClient,
		rdbConn.ToHandle(),
	))
	handlers = append(handlers, httpapi_handlers.NewAccountTransactions(logger, rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewProposals(
		logger,
		rdbConn.ToHandle(),
		cosmosAppClient,
	))
	handlers = append(handlers, httpapi_handlers.NewNFTs(
		logger,
		rdbConn.ToHandle(),
	))
	handlers = append(handlers, httpapi_handlers.NewIBCChannel(
		logger,
		rdbConn.ToHandle(),
	))
	handlers = append(handlers, httpapi_handlers.NewIBCChannelMessage(
		logger,
		rdbConn.ToHandle(),
	))
	handlers = append(handlers, httpapi_handlers.NewBridges(
		logger,
		rdbConn.ToHandle(),
		accountAddressPrefix,
	))

	return handlers
}
