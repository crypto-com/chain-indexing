package bootstrap

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
)

type app struct {
	logger applogger.Logger
	config *Config

	rdbConn       rdb.Conn
	httpAPIServer *HTTPAPIServer
	indexService  *IndexService
}

func NewApp(logger applogger.Logger, config *Config) *app {
	rdbConn, err := SetupRDbConn(config, logger)
	if err != nil {
		logger.Panicf("error setting up RDb connection: %v", err)
	}

	return &app{
		logger:  logger,
		config:  config,
		rdbConn: rdbConn,
	}
}

func (a *app) GetRDbConn() rdb.Conn {
	return a.rdbConn
}

func (a *app) InitHTTPAPIServer(registry RouteRegistry) {
	a.httpAPIServer = NewHTTPAPIServer(a.logger, a.config)
	a.httpAPIServer.RegisterRoutes(registry)
}

func (a *app) InitIndexService(projections []projection_entity.Projection, cronJobs []projection_entity.CronJob) {
	a.indexService = NewIndexService(a.logger, a.rdbConn, a.config, projections, cronJobs)
}

func (a *app) Run() {
	switch a.config.System.Mode {
	case SYSTEM_MODE_EVENT_STORE,
		SYSTEM_MODE_TENDERMINT_DIRECT,
		SYSTEM_MODE_API_ONLY:
	default:
		a.logger.Panicf("unrecognized system mode: %s", a.config.System.Mode)
	}

	if a.httpAPIServer != nil {
		go func() {
			if runErr := a.httpAPIServer.Run(); runErr != nil {
				a.logger.Panicf("%v", runErr)
			}
		}()
	}

	if a.indexService != nil {
		go func() {
			if runErr := a.indexService.Run(); runErr != nil {
				a.logger.Panicf("%v", runErr)
			}
		}()
	}

	select {}
}
