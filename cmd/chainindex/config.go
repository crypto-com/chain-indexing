package main

// CLIConfig defines struct of CLI input configs
type CLIConfig struct {
	TendermintHTTPRPCURL string
}

// FileConfig is the struct matches config.toml
type FileConfig struct {
	Tendermint      TendermintConfig
}

type TendermintConfig struct {
	HTTPRPCURL string `toml:"http_rpc_url"`
}