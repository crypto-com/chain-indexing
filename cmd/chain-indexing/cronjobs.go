package main

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/projection"
)

func initCronJobs(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *Config,
) []projection_entity.CronJob {
	cronJobs := make([]projection_entity.CronJob, 0, len(config.CronJob.Enables))
	initParams := projection.InitCronJobParams{
		Logger:  logger,
		RdbConn: rdbConn,
	}

	for _, cronJobName := range config.CronJob.Enables {
		cronJob := projection.InitCronJob(
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
