package main

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
)

func SetupRDbConn(config *Config, logger applogger.Logger) (rdb.Conn, error) {
	var pgxConnPool *pg.PgxConn
	var err error

	// GetFee duration strings to duration
	maxConnLifeTime, err := time.ParseDuration(config.Postgres.MaxConnLifeTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing MaxConnLifeTime string to duration %v", err)
	}
	maxConnIdleTime, err := time.ParseDuration(config.Postgres.MaxConnIdleTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing MaxConnIdleTime string to duration %v", err)
	}
	healthCheckInterval, err := time.ParseDuration(config.Postgres.HealthCheckInterval)
	if err != nil {
		return nil, fmt.Errorf("error parsing HealthCheckInterval string to duration %v", err)
	}

	for pgxConnPool == nil {
		pgxConnPool, err = pg.NewPgxConnPool(&pg.PgxConnPoolConfig{
			ConnConfig: pg.ConnConfig{
				Host:          config.Database.Host,
				Port:          config.Database.Port,
				MaybeUsername: &config.Database.Username,
				MaybePassword: &config.Database.Password,
				Database:      config.Database.Name,
				SSL:           config.Database.SSL,
			},
			MaybeMaxConns:          &config.Postgres.MaxConns,
			MaybeMinConns:          &config.Postgres.MinConns,
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
