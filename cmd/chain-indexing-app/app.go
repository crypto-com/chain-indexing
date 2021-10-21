package main

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	httpapi_handlers "github.com/crypto-com/chain-indexing/infrastructure/httpapi/handlers"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection"
)

type app struct {
	logger applogger.Logger
	config *bootstrap.Config

	rdbConn       rdb.Conn
	httpAPIServer *bootstrap.HTTPAPIServer
	indexService  *bootstrap.IndexService
}

func NewApp(logger applogger.Logger, config *bootstrap.Config) *app {
	return &app{
		logger: logger,
		config: config,
	}
}

func (a *app) Init() {
	rdbConn, err := bootstrap.SetupRDbConn(a.config, a.logger)
	if err != nil {
		a.logger.Panicf("error setting up RDb connection: %v", err)
	}
	a.rdbConn = rdbConn
}

func (a *app) InitHTTPAPIServer() {
	var cosmosAppClient cosmosapp.Client
	if a.config.CosmosApp.Insecure {
		cosmosAppClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
			a.config.CosmosApp.HTTPRPCUrl,
			a.config.Blockchain.BondingDenom,
		)
	} else {
		cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(
			a.config.CosmosApp.HTTPRPCUrl,
			a.config.Blockchain.BondingDenom,
		)
	}

	a.httpAPIServer = bootstrap.NewHTTPAPIServer(a.logger, a.config)

	handlers := make([]bootstrap.Handler, 0)
	handlers = append(handlers, httpapi_handlers.NewSearch(a.logger, a.rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewBlocks(a.logger, a.rdbConn.ToHandle()))
	handlers = append(handlers, httpapi_handlers.NewAccounts(
		a.logger,
		a.rdbConn.ToHandle(),
		cosmosAppClient,
		a.config.Blockchain.ValidatorAddressPrefix,
	))
	handlers = append(handlers, httpapi_handlers.NewAccountMessages(a.logger, a.rdbConn.ToHandle()))

	a.httpAPIServer.RegisterHandlers(handlers)
}

func (a *app) InitIndexService() {
	var cosmosAppClient cosmosapp.Client
	if a.config.CosmosApp.Insecure {
		cosmosAppClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
			a.config.CosmosApp.HTTPRPCUrl, a.config.Blockchain.BondingDenom,
		)
	} else {
		cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(
			a.config.CosmosApp.HTTPRPCUrl, a.config.Blockchain.BondingDenom,
		)
	}

	projections := make([]projection_entity.Projection, 0, len(a.config.Projection.Enables))
	initProjectionParams := projection.InitProjectionParams{
		Logger:  a.logger,
		RdbConn: a.rdbConn,

		CosmosAppClient:       cosmosAppClient,
		AccountAddressPrefix:  a.config.Blockchain.AccountAddressPrefix,
		ConsNodeAddressPrefix: a.config.Blockchain.ConNodeAddressPrefix,
	}
	for _, projectionName := range a.config.Projection.Enables {
		projection := projection.InitProjection(
			projectionName, initProjectionParams,
		)
		if onInitErr := projection.OnInit(); onInitErr != nil {
			a.logger.Errorf(
				"error initializing projection %s, system will attempt to initialize the projection again on next restart: %v",
				projection.Id(), onInitErr,
			)
			continue
		}
		projections = append(projections, projection)
	}

	a.logger.Infof("Enabled the follow projections: [%s]", strings.Join(a.config.Projection.Enables, ", "))

	cronJobs := make([]projection_entity.CronJob, 0, len(a.config.CronJob.Enables))
	initCronJobParams := projection.InitCronJobParams{
		Logger:  a.logger,
		RdbConn: a.rdbConn,
	}

	for _, cronJobName := range a.config.CronJob.Enables {
		cronJob := projection.InitCronJob(
			cronJobName, initCronJobParams,
		)
		if onInitErr := cronJob.OnInit(); onInitErr != nil {
			panic(fmt.Errorf(
				"error initializing cron job %s: %v",
				cronJob.Id(), onInitErr,
			))
		}
		cronJobs = append(cronJobs, cronJob)
	}

	a.indexService = bootstrap.NewIndexService(a.logger, a.rdbConn, a.config, projections, cronJobs)
}

func (a *app) Run() {
	switch a.config.System.Mode {
	case bootstrap.SYSTEM_MODE_EVENT_STORE,
		bootstrap.SYSTEM_MODE_TENDERMINT_DIRECT,
		bootstrap.SYSTEM_MODE_API_ONLY:
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
