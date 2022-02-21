package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/crypto-com/chain-indexing/bootstrap"
	configuration "github.com/crypto-com/chain-indexing/bootstrap/config"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/primptr"
	"github.com/crypto-com/chain-indexing/infrastructure"

	appconfig "github.com/crypto-com/chain-indexing/example/app/example-app/config"
	"github.com/crypto-com/chain-indexing/example/app/example-app/routes"
	"github.com/crypto-com/chain-indexing/example/internal/filereader/yaml"
)

func run(args []string) error {
	cliApp := &cli.App{
		Name:                 filepath.Base(args[0]),
		Usage:                "Crypto.org Chain Indexing Service",
		Version:              "v0.0.1",
		Copyright:            "(c) 2020 Crypto.com",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "./config/config.yaml",
				Usage: "YAML `FILE` to load configuration from",
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
			var config configuration.Config
			err := yaml.FromYAMLFile(configPath, &config)
			if err != nil {
				return fmt.Errorf("error config from yaml: %v", err)
			}

			var customConfig appconfig.CustomConfig
			err = yaml.FromYAMLFile(configPath, &customConfig)
			if err != nil {
				return fmt.Errorf("error custom config from yaml: %v", err)
			}

			cliConfig := appconfig.CLIConfig{
				LogLevel: ctx.String("logLevel"),

				DatabaseHost:     ctx.String("dbHost"),
				DatabaseUsername: ctx.String("dbUsername"),
				DatabasePassword: ctx.String("dbPassword"),
				DatabaseName:     ctx.String("dbName"),
				DatabaseSchema:   ctx.String("dbSchema"),

				TendermintHTTPRPCUrl: ctx.String("tendermintURL"),
				CosmosHTTPRPCUrl:     ctx.String("cosmosAppURL"),

				GithubAPIUsername: ctx.String("githubAPIUsername"),
				GithubAPIToken:    ctx.String("githubAPIToken"),
			}
			if ctx.IsSet("color") {
				cliConfig.LoggerColor = primptr.Bool(ctx.Bool("color"))
			}
			if ctx.IsSet("dbSSL") {
				cliConfig.DatabaseSSL = primptr.Bool(ctx.Bool("dbSSL"))
			}
			if ctx.IsSet("dbPort") {
				cliConfig.DatabasePort = primptr.Int32(int32(ctx.Int("dbPort")))
			}

			appconfig.OverrideByCLIConfig(&config, &cliConfig)

			// Create logger
			logLevel := parseLogLevel(config.Logger.Level)
			logger := infrastructure.NewZerologLogger(os.Stdout)
			logger.SetLogLevel(logLevel)

			app := bootstrap.NewApp(logger, &config)

			app.InitIndexService(
				initProjections(logger, app.GetRDbConn(), &config, &customConfig),
				initCronJobs(logger, app.GetRDbConn(), &config),
			)
			app.InitHTTPAPIServer(routes.InitRouteRegistry(logger, app.GetRDbConn(), &config, &customConfig))

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
