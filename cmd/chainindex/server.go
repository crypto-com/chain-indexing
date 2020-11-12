package main

import (
	"fmt"
	"os"

	"github.com/crypto-com/chainindex/usecase/parser"

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
	logger := infrastructure.NewZerologLoggerWithColor(os.Stderr)
	logger.Debug("TODO: should start server here")
	logger.Debug("TODO: should load module accounts configuration")

	moduleAccoutns := &parser.ModuleAccounts{
		FeeCollector:        "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
		Mint:                "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
		Distribution:        "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
		Gov:                 "tcro10d07y265gmmuvt4z0w9aw880jnsr700jvvjc2n",
		BondedTokensPool:    "tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h",
		NotBondedTokensPool: "tcro1tygms3xhhs3yv487phx3dw4a95jn7t7lh45rnr",
	}

	blockFeed := NewSyncManager(
		logger,
		s.TendermintHTTPRPCURL,
		340400,
		moduleAccoutns,
	)
	err := blockFeed.Run()
	if err != nil {
		return fmt.Errorf("error start running block polling service: %v", err)
	}

	return nil
}
