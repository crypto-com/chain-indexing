package main

import (
	"fmt"

	"github.com/crypto-com/chainindex/infrastructure/httpapi/routes"

	"github.com/crypto-com/chainindex/infrastructure/httpapi/handlers"

	"github.com/crypto-com/chainindex/appinterface/rdb"
	"github.com/crypto-com/chainindex/infrastructure/httpapi"
	applogger "github.com/crypto-com/chainindex/internal/logger"
)

type HTTPAPIServer struct {
	logger  applogger.Logger
	rdbConn rdb.Conn

	listeningAddress string
}

// NewIndexService creates a new server instance for polling and indexing
func NewHTTPAPIServer(logger applogger.Logger, rdbConn rdb.Conn, config *FileConfig) *HTTPAPIServer {
	return &HTTPAPIServer{
		logger:  logger,
		rdbConn: rdbConn,

		listeningAddress: config.HTTP.ListeningAddress,
	}
}

// Run function runs the polling server to index the data from Tendermint
func (server *HTTPAPIServer) Run() error {
	httpServer := httpapi.NewServer(
		server.listeningAddress,
	).WithLogger(
		server.logger,
	)

	blocksHandler := handlers.NewBlocks(server.logger, server.rdbConn.ToHandle())
	statusHandler := handlers.NewStatusHandler(server.logger, server.rdbConn.ToHandle())
	transactionsHandler := handlers.NewTransactions(server.logger, server.rdbConn.ToHandle())
	blockEventsHandler := handlers.NewBlockEvents(server.logger, server.rdbConn.ToHandle())
	validatorsHandler := handlers.NewValidators(server.logger, server.rdbConn.ToHandle())
	routeRegistry := routes.NewRoutesRegistry(
		blocksHandler,
		statusHandler,
		transactionsHandler,
		blockEventsHandler,
		validatorsHandler,
	)
	routeRegistry.Register(httpServer)

	server.logger.Infof("server start listening on: %s", server.listeningAddress)
	if err := httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error listening and serving HTTP API server: %v", err)
	}

	return nil
}
