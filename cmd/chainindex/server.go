package main

import (
	"fmt"
	"os"
	"time"

	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/infrastructure/pg"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"github.com/crypto-com/chainindex/usecase/parser"
)

type IndexServer struct {
	TendermintHTTPRPCURL string
	RdbConn              *pg.PgxConn
	logger               applogger.Logger
}

// NewIndexServer creates a new server instance for polling and indexing
func NewIndexServer(config *FileConfig) (*IndexServer, error) {
	logger := infrastructure.NewZerologLoggerWithColor(os.Stdout)

	pgxConnPool, err := SetupRdbConn(config, logger)
	if err != nil {
		return nil, fmt.Errorf("error setting up DB connection %v", err)
	}

	return &IndexServer{
		TendermintHTTPRPCURL: config.Tendermint.HTTPRPCURL,
		RdbConn:              pgxConnPool,
		logger:               logger,
	}, nil
}

func SetupRdbConn(config *FileConfig, logger applogger.Logger) (*pg.PgxConn, error) {
	var pgxConnPool *pg.PgxConn
	var err error

	// Parse duration strings to duration
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

// Run function runs the polling server to index the data from Tendermint
func (s *IndexServer) Run() error {
	s.logger.Debug("TODO: should load module accounts configuration")

	moduleAccounts := &parser.ModuleAccounts{
		FeeCollector:        "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
		Mint:                "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
		Distribution:        "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
		Gov:                 "tcro10d07y265gmmuvt4z0w9aw880jnsr700jvvjc2n",
		BondedTokensPool:    "tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h",
		NotBondedTokensPool: "tcro1tygms3xhhs3yv487phx3dw4a95jn7t7lh45rnr",
	}

	syncManager := NewSyncManager(
		s.logger,
		s.TendermintHTTPRPCURL,
		s.RdbConn,
		moduleAccounts,
	)
	if err := syncManager.Run(); err != nil {
		return fmt.Errorf("error running sync manager %v", err)
	}

	return nil
}
