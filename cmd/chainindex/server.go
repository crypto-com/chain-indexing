package main

import (
	"github.com/crypto-com/chainindex/infrastructure"
	toml "github.com/crypto-com/chainindex/internal/filereader/toml"
	"os"
)

type IndexServer struct {
	TendermintHTTPRPCURL string
}

// NewIndexServer creates a new server instance for polling and indexing
func NewIndexServer(configPath string, config *CLIConfig) (*IndexServer, error) {
	configReader, configFileErr := toml.FromFile(configPath)
	if configFileErr != nil {
		return nil, configFileErr
	}

	var fileConfig FileConfig
	readConfigErr := configReader.Read(&fileConfig)
	if readConfigErr != nil {
		return nil, readConfigErr
	}

	return &IndexServer{
		TendermintHTTPRPCURL: fileConfig.Tendermint.HTTPRPCURL,
	}, nil
}

// Run function runs the polling server to index the data from Tendermint
func (s *IndexServer) Run() error {
	logger := infrastructure.NewZerologLoggerWithColor(os.Stderr)

	syncPollingManager := NewPollingManager(logger, s.TendermintHTTPRPCURL, 367000)
	syncPollingManager.Run()

	return nil
}
