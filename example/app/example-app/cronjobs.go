package main

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	github_migrationhelper "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_activity_matcher"
)

func initCronJobs(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *bootstrap.Config,
) []projection_entity.CronJob {
	cronJobs := make([]projection_entity.CronJob, 0, len(config.CronJob.Enables))
	initParams := InitCronJobParams{
		Logger:  logger,
		RdbConn: rdbConn,
	}

	for _, cronJobName := range config.CronJob.Enables {
		cronJob := InitCronJob(
			cronJobName, initParams,
		)
		if onInitErr := cronJob.OnInit(); onInitErr != nil {
			panic(fmt.Errorf(
				"error initializing cron job %s: %v",
				cronJob.Id(), onInitErr,
			))
		}
		cronJobs = append(cronJobs, cronJob)
	}

	logger.Infof("Enabled the follow cron jobs: [%s]", strings.Join(config.CronJob.Enables, ", "))

	return cronJobs
}

func InitCronJob(name string, params InitCronJobParams) projection_entity.CronJob {
	connString := params.RdbConn.(*pg.PgxConn).ConnString()

	switch name {
	case "BridgeActivityMatcher":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			bridge_activity_matcher.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		return bridge_activity_matcher.New(params.Logger, params.RdbConn, migrationHelper)
	// register more cronjobs here
	default:
		panic(fmt.Sprintf("Unrecognized cron job: %s", name))
	}
}

type InitCronJobParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn

	GithubAPIUser    string
	GithubAPIToken   string
	MigrationRepoRef string
}
