package bootstrap

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/metric/prometheus"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/golang-migrate/migrate/v4"
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

	if config.System.Mode != SYSTEM_MODE_API_ONLY {
		ref := ""
		if config.GithubAPI.MigrationRepoRef != "" {
			ref = "#" + config.GithubAPI.MigrationRepoRef
		}

		m, err := migrate.New(
			fmt.Sprintf(MIGRATION_GITHUB_TARGET, config.GithubAPI.Username, config.GithubAPI.Token, ref),
			migrationDBConnString(rdbConn),
		)
		if err != nil {
			logger.Panicf("failed to init migration: %v", err)
		}

		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			logger.Panicf("failed to run migration: %v", err)
		}
	}

	return &app{
		logger:  logger,
		config:  config,
		rdbConn: rdbConn,
	}
}

const (
	MIGRATION_TABLE_NAME    = "schema_migrations"
	MIGRATION_GITHUB_TARGET = "github://%s:%s@crypto-com/chain-indexing/migrations%s"
)

func migrationDBConnString(conn rdb.Conn) string {
	connString := conn.(*pg.PgxConn).ConnString()
	if connString[len(connString)-1:] == "?" {
		return connString + "x-migrations-table=" + MIGRATION_TABLE_NAME
	} else {
		return connString + "&x-migrations-table=" + MIGRATION_TABLE_NAME
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

	if a.config.Prometheus.Enable {
		go func() {
			if runErr := prometheus.Run(a.config.Prometheus.ExportPath, a.config.Prometheus.Port); runErr != nil {
				a.logger.Panicf("%v", runErr)
			}
		}()
	}

	select {}
}
