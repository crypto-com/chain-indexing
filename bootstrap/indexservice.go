package bootstrap

import (
	"fmt"
	"time"

	event_interface "github.com/crypto-com/chain-indexing/appinterface/event"
	eventhandler_interface "github.com/crypto-com/chain-indexing/appinterface/eventhandler"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap/config"
	"github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
)

type IndexService struct {
	logger      applogger.Logger
	rdbConn     rdb.Conn
	projections []projection_entity.Projection
	cronJobs    []projection_entity.CronJob

	mode                     string
	accountAddressPrefix     string
	consNodeAddressPrefix    string
	bondingDenom             string
	windowSize               int
	tendermintHTTPRPCURL     string
	cosmosAppHTTPRPCURL      string
	insecureTendermintClient bool
	insecureCosmosAppClient  bool
	strictGenesisParsing     bool
	startingBlockHeight      int64

	cosmosVersionBlockHeight utils.CosmosVersionBlockHeight

	GithubAPIUser  string
	GithubAPIToken string
}

// NewIndexService creates a new server instance for polling and indexing
func NewIndexService(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *config.Config,
	projections []projection_entity.Projection,
	cronJobs []projection_entity.CronJob,
) *IndexService {
	return &IndexService{
		logger:      logger,
		rdbConn:     rdbConn,
		projections: projections,
		cronJobs:    cronJobs,

		mode:                     config.IndexService.Mode,
		consNodeAddressPrefix:    config.Blockchain.ConNodeAddressPrefix,
		accountAddressPrefix:     config.Blockchain.AccountAddressPrefix,
		bondingDenom:             config.Blockchain.BondingDenom,
		windowSize:               config.IndexService.WindowSize,
		tendermintHTTPRPCURL:     config.TendermintApp.HTTPRPCUrl,
		cosmosAppHTTPRPCURL:      config.CosmosApp.HTTPRPCUrl,
		insecureTendermintClient: config.TendermintApp.Insecure,
		insecureCosmosAppClient:  config.CosmosApp.Insecure,
		strictGenesisParsing:     config.TendermintApp.StrictGenesisParsing,
		startingBlockHeight:      config.IndexService.StartingBlockHeight,
		cosmosVersionBlockHeight: utils.CosmosVersionBlockHeight{
			V0_42_7: utils.ParserBlockHeight(config.IndexService.CosmosVersionEnabledHeight.V0_42_7),
		},
		GithubAPIUser:  config.IndexService.GithubAPI.Username,
		GithubAPIToken: config.IndexService.GithubAPI.Token,
	}
}

func (service *IndexService) Run() error {
	// run polling tendermint manager, update view tables directly
	infoManager := NewInfoManager(
		service.logger,
		service.rdbConn,
		service.tendermintHTTPRPCURL,
		service.insecureTendermintClient,
		service.strictGenesisParsing,
	)

	switch service.mode {
	case config.SYSTEM_MODE_EVENT_STORE:
		infoManager.Run()
		service.runCronJobs()
		return service.RunEventStoreMode()
	case config.SYSTEM_MODE_TENDERMINT_DIRECT:
		infoManager.Run()
		service.runCronJobs()
		return service.RunTendermintDirectMode()
	default:
		return fmt.Errorf("unsupported system mode: %s", service.mode)
	}
}

func (service *IndexService) runCronJobs() {
	for i := range service.cronJobs {
		cronJobClosure := service.cronJobs[i]
		go func() {
			logger := service.logger.WithFields(applogger.LogFields{
				"module": cronJobClosure.Id(),
			})
			for {
				if cronJobErr := cronJobClosure.Exec(); cronJobErr != nil {
					logger.Errorf("error executing cron job: %v", cronJobErr)
				}
				logger.Infof("successfully executed cron job, going to execute again in %s", cronJobClosure.Interval())
				<-time.After(cronJobClosure.Interval())
			}
		}()
	}
}

func (service *IndexService) RunEventStoreMode() error {
	eventRegistry := event.NewRegistry()
	event_usecase.RegisterEvents(eventRegistry)
	eventStore := event_interface.NewRDbStore(service.rdbConn.ToHandle(), eventRegistry)

	projectionManager := projection_entity.NewStoreBasedManager(service.logger, eventStore)

	for _, projection := range service.projections {
		if err := projectionManager.RegisterProjection(projection); err != nil {
			return fmt.Errorf("error registering projection `%s` to manager %v", projection.Id(), err)
		}
	}
	projectionManager.RunInBackground()

	eventStoreHandler := eventhandler_interface.NewRDbEventStoreHandler(
		service.logger,
		service.rdbConn,
		eventRegistry,
	)
	syncManager := NewSyncManager(
		SyncManagerParams{
			Logger:  service.logger,
			RDbConn: service.rdbConn,
			Config: SyncManagerConfig{
				WindowSize:               service.windowSize,
				TendermintRPCUrl:         service.tendermintHTTPRPCURL,
				CosmosAppHTTPRPCURL:      service.cosmosAppHTTPRPCURL,
				InsecureTendermintClient: service.insecureTendermintClient,
				InsecureCosmosAppClient:  service.insecureCosmosAppClient,
				StrictGenesisParsing:     service.strictGenesisParsing,
				AccountAddressPrefix:     service.accountAddressPrefix,
				StakingDenom:             service.bondingDenom,
				StartingBlockHeight:      service.startingBlockHeight,
			},
		},
		utils.NewCosmosParserManager(
			utils.CosmosParserManagerParams{
				Logger: service.logger,
				Config: utils.CosmosParserManagerConfig{
					CosmosVersionBlockHeight: service.cosmosVersionBlockHeight,
				},
			},
		),
		eventStoreHandler,
	)
	if err := syncManager.Run(); err != nil {
		return fmt.Errorf("error running sync manager %v", err)
	}

	return nil
}

func (service *IndexService) RunTendermintDirectMode() error {

	for i := range service.projections {
		go func(projection projection_entity.Projection) {
			syncManager := NewSyncManager(
				SyncManagerParams{
					Logger: service.logger.WithFields(applogger.LogFields{
						"projection": projection.Id(),
					}),
					RDbConn: service.rdbConn,
					Config: SyncManagerConfig{
						WindowSize:               service.windowSize,
						TendermintRPCUrl:         service.tendermintHTTPRPCURL,
						CosmosAppHTTPRPCURL:      service.cosmosAppHTTPRPCURL,
						InsecureTendermintClient: service.insecureTendermintClient,
						InsecureCosmosAppClient:  service.insecureCosmosAppClient,
						AccountAddressPrefix:     service.accountAddressPrefix,
						StakingDenom:             service.bondingDenom,
						StartingBlockHeight:      service.startingBlockHeight,
					},
				},
				utils.NewCosmosParserManager(
					utils.CosmosParserManagerParams{
						Logger: service.logger.WithFields(applogger.LogFields{
							"projection": projection.Id(),
						}),
						Config: utils.CosmosParserManagerConfig{
							CosmosVersionBlockHeight: service.cosmosVersionBlockHeight,
						},
					},
				),
				eventhandler_interface.NewProjectionHandler(service.logger, projection),
			)
			if err := syncManager.Run(); err != nil {
				panic(fmt.Sprintf("error running sync manager %v", err))
			}
		}(service.projections[i])
	}
	select {}
}
