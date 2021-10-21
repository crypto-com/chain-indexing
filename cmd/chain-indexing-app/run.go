package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crypto-com/chain-indexing/bootstrap"
	"github.com/urfave/cli/v2"

	"github.com/crypto-com/chain-indexing/infrastructure"
	"github.com/crypto-com/chain-indexing/internal/filereader/toml"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/primptr"
	projection_usecase "github.com/crypto-com/chain-indexing/usecase/projection"
)

func run(args []string) error {
	cliApp := &cli.App{
		Name:                 filepath.Base(args[0]),
		Usage:                "Crypto.com Chain Indexing Service",
		Version:              "v0.0.1",
		Copyright:            "(c) 2020 Crypto.com",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "./config/config.toml",
				Usage: "TOML `FILE` to load configuration from",
			},

			&cli.StringFlag{
				Name:  "logLevel,l",
				Usage: "Log level (Allowed values: fatal,error,info,debug)",
			},
			&cli.BoolFlag{
				Name:  "color",
				Usage: "Display colored log",
			},

			&cli.BoolFlag{
				Name:    "dbSSL",
				Usage:   "Enable Postgres SSL mode",
				EnvVars: []string{"DB_SSL"},
			},
			&cli.StringFlag{
				Name:    "dbHost",
				Usage:   "Postgres database hostname",
				EnvVars: []string{"DB_HOST"},
			},
			&cli.UintFlag{
				Name:    "dbPort",
				Usage:   "Postgres database port",
				EnvVars: []string{"DB_PORT"},
			},
			&cli.StringFlag{
				Name:    "dbUsername",
				Usage:   "Postgres username",
				EnvVars: []string{"DB_USERNAME"},
			},
			&cli.StringFlag{
				Name:     "dbPassword",
				Usage:    "Postgres password",
				EnvVars:  []string{"DB_PASSWORD"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "dbName",
				Usage:   "Postgres database name",
				EnvVars: []string{"DB_NAME"},
			},
			&cli.StringFlag{
				Name:    "dbSchema",
				Usage:   "Postgres schema name",
				EnvVars: []string{"DB_SCHEMA"},
			},

			&cli.StringFlag{
				Name:    "tendermintURL",
				Usage:   "Tendermint HTTP RPC URL",
				EnvVars: []string{"TENDERMINT_URL"},
			},
			&cli.StringFlag{
				Name:    "cosmosAppURL",
				Usage:   " Cosmos App RPC URL",
				EnvVars: []string{"COSMOSAPP_URL"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if args := ctx.Args(); args.Len() > 0 {
				return fmt.Errorf("Unexpected arguments: %q", args.Get(0))
			}

			// Prepare FileConfig
			configPath := ctx.String("config")
			configReader, configFileErr := toml.FromFile(configPath)
			if configFileErr != nil {
				return fmt.Errorf("error creating Toml config reader: %v", configFileErr)
			}
			var fileConfig bootstrap.FileConfig
			if readConfigErr := configReader.Read(&fileConfig); readConfigErr != nil {
				return fmt.Errorf("error reading Toml config: %v", readConfigErr)
			}

			projectionConfigReader, projectionConfigFileErr := toml.FromFile(configPath)
			if projectionConfigFileErr != nil {
				return fmt.Errorf("error creating projection Toml config reader: %v", configFileErr)
			}
			if readProjectionConfigErr := projectionConfigReader.Read(&projection_usecase.GlobalConfig); readProjectionConfigErr != nil {
				return fmt.Errorf("error reading global projection Toml config: %v", readProjectionConfigErr)
			}

			cliConfig := bootstrap.CLIConfig{
				LogLevel: ctx.String("logLevel"),

				DatabaseHost:     ctx.String("dbHost"),
				DatabaseUsername: ctx.String("dbUsername"),
				DatabasePassword: ctx.String("dbPassword"),
				DatabaseName:     ctx.String("dbName"),
				DatabaseSchema:   ctx.String("dbSchema"),

				TendermintHTTPRPCUrl: ctx.String("tendermintURL"),
				CosmosHTTPRPCUrl:     ctx.String("cosmosAppURL"),
			}
			if ctx.IsSet("color") {
				cliConfig.LoggerColor = primptr.Bool(ctx.Bool("color"))
			}
			if ctx.IsSet("dbSSL") {
				cliConfig.DatabaseSSL = primptr.Bool(ctx.Bool("dbSSL"))
			}
			if ctx.IsSet("dgPort") {
				cliConfig.DatabasePort = primptr.Int32(int32(ctx.Int("dbPort")))
			}

			config := bootstrap.Config{
				FileConfig: fileConfig,
			}
			config.OverrideByCLIConfig(&cliConfig)

			// Create logger
			logLevel := parseLogLevel(config.Logger.Level)
			logger := infrastructure.NewZerologLogger(os.Stdout)
			logger.SetLogLevel(logLevel)

			app := bootstrap.NewApp(logger, &config)

			app.InitIndexService(
				initProjections(logger, app.GetRDbConn(), &config),
				initCronJobs(logger, app.GetRDbConn(), &config),
			)
			app.InitHTTPAPIServer(initHTTPAPIHandlers(logger, app.GetRDbConn(), &config))

			app.Run()

			return nil
		},
	}

	err := cliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}

func parseLogLevel(level string) applogger.LogLevel {
	switch level {
	case "panic":
		return applogger.LOG_LEVEL_PANIC
	case "error":
		return applogger.LOG_LEVEL_ERROR
	case "info":
		return applogger.LOG_LEVEL_INFO
	case "debug":
		return applogger.LOG_LEVEL_DEBUG
	default:
		panic("Unsupported log level: " + level)
	}
}
