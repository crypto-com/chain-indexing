package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func CliApp(args []string) error {
	cliApp := &cli.App{
		Name:      filepath.Base(args[0]),
		Usage:     "Crypto.com Chain Index",
		Version:   "v0.0.1",
		Copyright: "(c) 2020 Crypto.com",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config,c",
				Value: "./config/config.toml",
				Usage: "TOML `FILE` to load configuration from",
			},
			&cli.StringFlag{
				Name:    "tendermintURL",
				Usage:   "Tendermint HTTP RPC URL",
				EnvVars: []string{"TENDERMINT_URL"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if args := ctx.Args(); args.Len() > 0 {
				return fmt.Errorf("Unexpected arguments: %q", args.Get(0))
			}

			configPath := ctx.String("config")

			cliConfig := CLIConfig{
				TendermintHTTPRPCURL: ctx.String("tendermintURL"),
			}

			indexServerApp, err := NewIndexServer(configPath, &cliConfig)
			if err != nil {
				return fmt.Errorf("error creating server: %v", err)
			}

			if err = indexServerApp.Run(); err != nil {
				return fmt.Errorf("error when starting indexing server: %v", err)
			}

			return nil
		},
	}

	err := cliApp.Run(args)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "indexing").
			Msgf("Cannot start %s", "polling server")
	}

	return nil
}
