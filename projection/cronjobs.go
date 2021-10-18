package projection

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/projection/bridge_activity/bridge_activity_matcher"
)

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
