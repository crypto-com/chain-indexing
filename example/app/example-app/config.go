package main

import (
	"github.com/crypto-com/chain-indexing/bootstrap"
)

type CustomConfig struct {
	ServerGithubAPI ServerGithubAPIConfig `toml:"server_github_api"`
}

type ServerGithubAPIConfig struct {
	MigrationRepoRef string `toml:"migration_repo_ref"`
}

type CLIConfig struct {
	LoggerColor *bool
	LogLevel    string

	DatabaseSSL      *bool
	DatabaseHost     string
	DatabasePort     *int32
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabaseSchema   string

	TendermintHTTPRPCUrl string
	CosmosHTTPRPCUrl     string

	GithubAPIUsername string
	GithubAPIToken    string
}

func OverrideByCLIConfig(config *bootstrap.Config, cliConfig *CLIConfig) {
	if cliConfig.LogLevel != "" {
		config.Logger.Level = cliConfig.LogLevel
	}
	if cliConfig.LoggerColor != nil {
		config.Logger.Color = *cliConfig.LoggerColor
	}
	if cliConfig.DatabaseSSL != nil {
		config.Database.SSL = *cliConfig.DatabaseSSL
	}
	if cliConfig.DatabaseHost != "" {
		config.Database.Host = cliConfig.DatabaseHost
	}
	if cliConfig.DatabasePort != nil {
		config.Database.Port = *cliConfig.DatabasePort
	}
	if cliConfig.DatabaseUsername != "" {
		config.Database.Username = cliConfig.DatabaseUsername
	}
	// Always overwrite database password with CLI (ENV)
	config.Database.Password = cliConfig.DatabasePassword
	if cliConfig.DatabaseName != "" {
		config.Database.Name = cliConfig.DatabaseName
	}
	if cliConfig.DatabaseSchema != "" {
		config.Database.Schema = cliConfig.DatabaseSchema
	}
	if cliConfig.TendermintHTTPRPCUrl != "" {
		config.Tendermint.HTTPRPCUrl = cliConfig.TendermintHTTPRPCUrl
	}
	if cliConfig.CosmosHTTPRPCUrl != "" {
		config.CosmosApp.HTTPRPCUrl = cliConfig.CosmosHTTPRPCUrl
	}
	if cliConfig.GithubAPIUsername != "" {
		config.GithubAPI.Username = cliConfig.GithubAPIUsername
	}
	if cliConfig.GithubAPIToken != "" {
		config.GithubAPI.Token = cliConfig.GithubAPIToken
	}
}

