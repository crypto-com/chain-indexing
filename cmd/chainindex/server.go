package main

import (
	"fmt"
	"os"

	"github.com/crypto-com/chainindex/infrastructure"
	toml "github.com/crypto-com/chainindex/internal/filereader/toml"
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
	fmt.Println("TODO: will start server here")

	logger := infrastructure.NewZerologLoggerWithColor(os.Stderr)
	blockFeed := NewSyncManager(logger, s.TendermintHTTPRPCURL, 340400)
	err := blockFeed.Run()
	if err != nil {
		return fmt.Errorf("error start running block polling service: %v", err)
	}

	return nil
}
