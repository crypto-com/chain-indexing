package main

import (
	"github.com/crypto-com/chain-indexing/bootstrap/config"
)

type CustomConfig struct {
	ServerGithubAPI ServerGithubAPIConfig `yaml:"server_github_api"`
}

type ServerGithubAPIConfig struct {
	MigrationRepoRef string `yaml:"migration_repo_ref"`
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

func OverrideByCLIConfig(config *config.Config, cliConfig *CLIConfig) {
	if cliConfig.LogLevel != "" {
		config.Logger.Level = cliConfig.LogLevel
	}
	if cliConfig.LoggerColor != nil {
		config.Logger.Color = *cliConfig.LoggerColor
	}
	if cliConfig.DatabaseSSL != nil {
		config.Postgres.SSL = *cliConfig.DatabaseSSL
	}
	if cliConfig.DatabaseHost != "" {
		config.Postgres.Host = cliConfig.DatabaseHost
	}
	if cliConfig.DatabasePort != nil {
		config.Postgres.Port = *cliConfig.DatabasePort
	}
	if cliConfig.DatabaseUsername != "" {
		config.Postgres.Username = cliConfig.DatabaseUsername
	}
	// Always overwrite database password with CLI (ENV)
	config.Postgres.Password = cliConfig.DatabasePassword
	if cliConfig.DatabaseName != "" {
		config.Postgres.Name = cliConfig.DatabaseName
	}
	if cliConfig.DatabaseSchema != "" {
		config.Postgres.Schema = cliConfig.DatabaseSchema
	}
	if cliConfig.TendermintHTTPRPCUrl != "" {
		config.TendermintApp.HTTPRPCUrl = cliConfig.TendermintHTTPRPCUrl
	}
	if cliConfig.CosmosHTTPRPCUrl != "" {
		config.CosmosApp.HTTPRPCUrl = cliConfig.CosmosHTTPRPCUrl
	}
	if cliConfig.GithubAPIUsername != "" {
		config.IndexService.GithubAPI.Username = cliConfig.GithubAPIUsername
	}
	if cliConfig.GithubAPIToken != "" {
		config.IndexService.GithubAPI.Token = cliConfig.GithubAPIToken
	}
}
