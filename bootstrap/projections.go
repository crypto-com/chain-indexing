package bootstrap

import (
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection"
)

func initProjections(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *Config,
) []projection_entity.Projection {
	var cosmosAppClient cosmosapp.Client
	if config.CosmosApp.Insecure {
		cosmosAppClient = cosmosapp_infrastructure.NewInsecureHTTPClient(
			config.CosmosApp.HTTPRPCUrl, config.Blockchain.BondingDenom,
		)
	} else {
		cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(
			config.CosmosApp.HTTPRPCUrl, config.Blockchain.BondingDenom,
		)
	}

	projections := make([]projection_entity.Projection, 0, len(config.Projection.Enables))
	initParams := projection.InitProjectionParams{
		Logger:  logger,
		RdbConn: rdbConn,

		CosmosAppClient:       cosmosAppClient,
		AccountAddressPrefix:  config.Blockchain.AccountAddressPrefix,
		ConsNodeAddressPrefix: config.Blockchain.ConNodeAddressPrefix,
	}
	for _, projectionName := range config.Projection.Enables {
		projection := projection.InitProjection(
			projectionName, initParams,
		)
		if onInitErr := projection.OnInit(); onInitErr != nil {
			logger.Errorf(
				"error initializing projection %s, system will attempt to initialize the projection again on next restart: %v",
				projection.Id(), onInitErr,
			)
			continue
		}
		projections = append(projections, projection)
	}

	logger.Infof("Enabled the follow projections: [%s]", strings.Join(config.Projection.Enables, ", "))

	return projections
}
