package main

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection"
	"strings"
)

func initProjections(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *Config,
) []projection_entity.Projection {
	var cosmosAppClient = cosmosapp_infrastructure.NewHTTPClient(config.CosmosApp.HTTPRPCUL)

	projections := make([]projection_entity.Projection, 0, len(config.Projection.Enables))
	initParams := projection.InitParams{
		Logger:                logger,
		RdbConn:               rdbConn,
		CosmosAppClient:       cosmosAppClient,
		ConsNodeAddressPrefix: config.Blockchain.ConNodeAddressPrefix,

		ValidatorAddressPrefix: config.Blockchain.ValidatorAddressPrefix,
		AccountAddressPrefix:   config.Blockchain.ValidatorAddressPrefix,

		PhaseOneStartTime:        config.Crossfire.PhaseOneStartTime,
		PhaseTwoStartTime:        config.Crossfire.PhaseTwoStartTime,
		PhaseThreeStartTime:      config.Crossfire.PhaseThreeStartTime,
		CompetitionEndTime:       config.Crossfire.CompetitionEndTime,
		AdminAddress:             config.Crossfire.AdminAddress,
		NetworkUpgradeProposalID: config.Crossfire.NetworkUpgradeProposalID,
		ParticipantsListURL:      config.Crossfire.ParticipantsListURL,
	}
	for _, projectionName := range config.Projection.Enables {
		projections = append(projections, projection.InitProjection(
			projectionName, initParams,
		))
	}

	logger.Infof("Enabled the follow projections: [%s]", strings.Join(config.Projection.Enables, ", "))

	return projections
}
