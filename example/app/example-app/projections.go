package main

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	configuration "github.com/crypto-com/chain-indexing/bootstrap/config"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	custom_projection "github.com/crypto-com/chain-indexing/example/projection"
	"github.com/crypto-com/chain-indexing/example/projection/example"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	github_migrationhelper "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	"github.com/crypto-com/chain-indexing/projection/account"
	"github.com/crypto-com/chain-indexing/projection/account_message"
	"github.com/crypto-com/chain-indexing/projection/account_transaction"
	"github.com/crypto-com/chain-indexing/projection/block"
	"github.com/crypto-com/chain-indexing/projection/blockevent"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_pending_activity"
	"github.com/crypto-com/chain-indexing/projection/chainstats"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel"
	"github.com/crypto-com/chain-indexing/projection/ibc_channel_message"
	"github.com/crypto-com/chain-indexing/projection/nft"
	"github.com/crypto-com/chain-indexing/projection/proposal"
	"github.com/crypto-com/chain-indexing/projection/transaction"
	"github.com/crypto-com/chain-indexing/projection/validator"
	"github.com/crypto-com/chain-indexing/projection/validatorstats"
	"github.com/ettle/strcase"
)

func initProjections(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *configuration.Config,
	customConfig *CustomConfig,
) []projection_entity.Projection {
	if !config.IndexService.Enable {
		return []projection_entity.Projection{}
	}

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

	projections := make([]projection_entity.Projection, 0, len(config.IndexService.Projection.Enables))
	initParams := InitProjectionParams{
		Logger:  logger,
		RdbConn: rdbConn,

		ExtraConfigs: config.IndexService.Projection.ExtraConfigs,

		CosmosAppClient:       cosmosAppClient,
		AccountAddressPrefix:  config.Blockchain.AccountAddressPrefix,
		ConsNodeAddressPrefix: config.Blockchain.ConNodeAddressPrefix,

		GithubAPIUser:    config.IndexService.GithubAPI.Username,
		GithubAPIToken:   config.IndexService.GithubAPI.Token,
		MigrationRepoRef: config.IndexService.GithubAPI.MigrationRepoRef,

		ServerMigrationRepoRef: customConfig.ServerGithubAPI.MigrationRepoRef,
	}
	for _, projectionName := range config.IndexService.Projection.Enables {
		projection := InitProjection(
			projectionName, initParams,
		)
		if projection == nil {
			continue
		} else {
			if onInitErr := projection.OnInit(); onInitErr != nil {
				logger.Errorf(
					"error initializing projection %s, system will attempt to initialize the projection again on next restart: %v",
					projection.Id(), onInitErr,
				)
				continue
			}
		}
		projections = append(projections, projection)
	}

	// Append additional projection
	for _, projectionName := range config.IndexService.Projection.Enables {
		projection := InitAdditionalProjection(
			projectionName, initParams,
		)
		if projection == nil {
			continue
		} else {
			if onInitErr := projection.OnInit(); onInitErr != nil {
				logger.Errorf(
					"error initializing additional projection %s, system will attempt to initialize the projection again on next restart: %v",
					projection.Id(), onInitErr,
				)
				continue
			}
		}
		projections = append(projections, projection)
	}

	logger.Infof("Enabled the follow projection: [%s]", strings.Join(config.IndexService.Projection.Enables, ", "))

	return projections
}

func InitAdditionalProjection(name string, params InitProjectionParams) projection_entity.Projection {
	connString := params.RdbConn.(*pg.PgxConn).ConnString()

	githubMigrationHelperConfigForCustomProjection := github_migrationhelper.Config{
		GithubAPIUser:    params.GithubAPIUser,
		GithubAPIToken:   params.GithubAPIToken,
		MigrationRepoRef: params.ServerMigrationRepoRef,
		ConnString:       connString,
	}

	switch name {
	case "Example":
		sourceURL := generateGithubMigrationSrouceURLForCustomProjection(name, githubMigrationHelperConfigForCustomProjection)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return example.NewAdditionalProjection(params.Logger, params.RdbConn, migrationHelper)
	}

	return nil
}

func InitProjection(name string, params InitProjectionParams) projection_entity.Projection {
	connString := params.RdbConn.(*pg.PgxConn).ConnString()

	githubMigrationHelperConfig := github_migrationhelper.Config{
		GithubAPIUser:    params.GithubAPIUser,
		GithubAPIToken:   params.GithubAPIToken,
		MigrationRepoRef: params.MigrationRepoRef,
		ConnString:       connString,
	}

	switch name {
	case "Account":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return account.NewAccount(params.Logger, params.RdbConn, params.CosmosAppClient, migrationHelper)
	case "AccountTransaction":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return account_transaction.NewAccountTransaction(params.Logger, params.RdbConn, params.AccountAddressPrefix, migrationHelper)
	case "AccountMessage":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return account_message.NewAccountMessage(params.Logger, params.RdbConn, params.AccountAddressPrefix, migrationHelper)
	case "Block":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return block.NewBlock(params.Logger, params.RdbConn, migrationHelper)
	case "BlockEvent":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			blockevent.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return blockevent.NewBlockEvent(params.Logger, params.RdbConn, migrationHelper)
	case "ChainStats":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			chainstats.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return chainstats.NewChainStats(params.Logger, params.RdbConn, migrationHelper)
	case "Proposal":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return proposal.NewProposal(params.Logger, params.RdbConn, params.ConsNodeAddressPrefix, migrationHelper)
	case "Transaction":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return transaction.NewTransaction(params.Logger, params.RdbConn, migrationHelper)
	case "Validator":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return validator.NewValidator(params.Logger, params.RdbConn, params.ConsNodeAddressPrefix, migrationHelper)
	case "ValidatorStats":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			validatorstats.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDatabaseURL(
			connString,
			validatorstats.MIGRATION_TABLE_NAME,
		)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return validatorstats.NewValidatorStats(params.Logger, params.RdbConn, migrationHelper)
	case "NFT":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return nft.NewNFT(
			params.Logger,
			params.RdbConn,
			&nft.Config{
				EnableDrop:       false,
				DropDataAccessor: "",
			},
			migrationHelper,
		)
	case "CryptoComNFT":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			nft.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDatabaseURL(
			connString,
			nft.MIGRATION_TABLE_NAME,
		)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return nft.NewNFT(
			params.Logger,
			params.RdbConn,
			&nft.Config{
				EnableDrop:       true,
				DropDataAccessor: "dropId",
			},
			migrationHelper,
		)
	case "IBCChannel":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return ibc_channel.NewIBCChannel(
			params.Logger,
			params.RdbConn,
			&ibc_channel.Config{
				EnableTxMsgTrace: false,
			},
			migrationHelper,
		)
	case "IBCChannelTxMsgTrace":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			ibc_channel.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDatabaseURL(
			connString,
			ibc_channel.MIGRATION_TABLE_NAME,
		)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return ibc_channel.NewIBCChannel(
			params.Logger,
			params.RdbConn,
			&ibc_channel.Config{
				EnableTxMsgTrace: true,
			},
			migrationHelper,
		)
	case "IBCChannelMessage":
		sourceURL := github_migrationhelper.GenerateDefaultSourceURL(name, githubMigrationHelperConfig)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return ibc_channel_message.NewIBCChannelMessage(params.Logger, params.RdbConn, migrationHelper)
	case "BridgePendingActivity":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			bridge_pending_activity.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		config, err := bridge_pending_activity.ConfigFromInterface(params.ExtraConfigs[name])
		if err != nil {
			params.Logger.Panicf(err.Error())
		}

		return bridge_pending_activity.NewBridgePendingActivity(params.Logger, params.RdbConn, migrationHelper, config)
	}

	return nil
}

func generateGithubMigrationSrouceURLForCustomProjection(
	projectionId string,
	config github_migrationhelper.Config,
) string {
	projectionSnakeId := strcase.ToSnake(projectionId)
	migrationDirectory := fmt.Sprintf(github_migrationhelper.MIGRATION_DIRECTORY_FORMAT, projectionSnakeId)

	return github_migrationhelper.GenerateSourceURL(
		custom_projection.MIGRATION_GITHUB_URL_FORMAT,
		config.GithubAPIUser,
		config.GithubAPIToken,
		migrationDirectory,
		config.MigrationRepoRef,
	)
}

type InitProjectionParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn

	ExtraConfigs map[string]interface{}

	CosmosAppClient       cosmosapp.Client
	AccountAddressPrefix  string
	ConsNodeAddressPrefix string

	GithubAPIUser    string
	GithubAPIToken   string
	MigrationRepoRef string

	ServerMigrationRepoRef string
}
