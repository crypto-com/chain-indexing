package main

import (
	"fmt"

	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi/routes"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

type HTTPAPIServer struct {
	logger          applogger.Logger
	rdbConn         rdb.Conn
	cosmosAppClient cosmosapp.Client

	validatorAddressPrefix string
	conNodeAddressPrefix   string

	listeningAddress string
	routePrefix      string
}

// NewIndexService creates a new server instance for polling and indexing
func NewHTTPAPIServer(logger applogger.Logger, rdbConn rdb.Conn, config *Config) *HTTPAPIServer {
	return &HTTPAPIServer{
		logger:          logger,
		rdbConn:         rdbConn,
		cosmosAppClient: cosmosapp_infrastructure.NewHTTPClient(config.CosmosApp.HTTPRPCUL),

		validatorAddressPrefix: config.Blockchain.ValidatorAddressPrefix,
		conNodeAddressPrefix:   config.Blockchain.ConNodeAddressPrefix,
		listeningAddress:       config.HTTP.ListeningAddress,
		routePrefix:            config.HTTP.RoutePrefix,
	}
}

// Run function runs the polling server to index the data from Tendermint
func (server *HTTPAPIServer) Run() error {
	httpServer := httpapi.NewServer(
		server.listeningAddress,
	).WithLogger(
		server.logger,
	)

	searchHandler := handlers.NewSearch(server.logger, server.rdbConn.ToHandle())
	blocksHandler := handlers.NewBlocks(server.logger, server.rdbConn.ToHandle())
	statusHandler := handlers.NewStatusHandler(server.logger, server.rdbConn.ToHandle())
	transactionsHandler := handlers.NewTransactions(server.logger, server.rdbConn.ToHandle())
	blockEventsHandler := handlers.NewBlockEvents(server.logger, server.rdbConn.ToHandle())
	validatorsHandler := handlers.NewValidators(
		server.logger,
		server.validatorAddressPrefix,
		server.conNodeAddressPrefix,
		server.cosmosAppClient,
		server.rdbConn.ToHandle(),
	)
	accountsHandler := handlers.NewAccounts(server.logger, server.rdbConn.ToHandle())

	routeRegistry := routes.NewRoutesRegistry(
		searchHandler,
		blocksHandler,
		statusHandler,
		transactionsHandler,
		blockEventsHandler,
		validatorsHandler,
		accountsHandler,
	)
	routeRegistry.Register(httpServer, server.routePrefix)

	server.logger.Infof("server start listening on: %s", server.listeningAddress)
	if err := httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error listening and serving HTTP API server: %v", err)
	}

	return nil
}
