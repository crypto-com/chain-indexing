package main

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
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
	switch name {
	case "BridgeActivityMatcher":
		return bridge_activity_matcher.New(params.Logger, params.RdbConn)
	// register more cronjobs here
	default:
		panic(fmt.Sprintf("Unrecognized cron job: %s", name))
	}
}

type InitCronJobParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn
}
