package bootstrap

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/bootstrap/config"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
)

func SetupRDbConn(config *config.Config, logger applogger.Logger) (rdb.Conn, error) {
	var pgxConnPool *pg.PgxConn
	var err error

	// GetFee duration strings to duration
	maxConnLifeTime, err := time.ParseDuration(config.Postgres.PoolMaxConnLifeTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing MaxConnLifeTime string to duration %v", err)
	}
	maxConnIdleTime, err := time.ParseDuration(config.Postgres.PoolMaxConnIdleTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing MaxConnIdleTime string to duration %v", err)
	}
	healthCheckInterval, err := time.ParseDuration(config.Postgres.PoolHealthCheckInterval)
	if err != nil {
		return nil, fmt.Errorf("error parsing HealthCheckInterval string to duration %v", err)
	}

	for pgxConnPool == nil {
		pgxConnPool, err = pg.NewPgxConnPool(&pg.PgxConnPoolConfig{
			ConnConfig: pg.ConnConfig{
				Host:          config.Postgres.Host,
				Port:          config.Postgres.Port,
				MaybeUsername: &config.Postgres.Username,
				MaybePassword: &config.Postgres.Password,
				Database:      config.Postgres.Name,
				SSL:           config.Postgres.SSL,
			},
			MaybeMaxConns:          &config.Postgres.PoolMaxConns,
			MaybeMinConns:          &config.Postgres.PoolMinConns,
			MaybeMaxConnLifeTime:   &maxConnLifeTime,
			MaybeMaxConnIdleTime:   &maxConnIdleTime,
			MaybeHealthCheckPeriod: &healthCheckInterval,
		}, logger)

		if err != nil {
			logger.Errorf("error setting up connection to database, will retry in 5 seconds: %v", err)
			<-time.After(5 * time.Second)
		}
	}

	logger.Info("successfully setup database connection")
	return pgxConnPool, nil
}
