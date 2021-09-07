package utils

import (
	"github.com/crypto-com/chain-indexing/entity/command"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CosmosParserManager struct {
	store  map[string]CosmosParserBlockHeight
	logger applogger.Logger
	config CosmosParserManagerConfig
}

type CosmosParserBlockHeight map[uint64]CosmosParser

type CosmosParserManagerConfig struct {
	CosmosVersionBLockHeight
}

type CosmosVersionBLockHeight struct {
	V0_43_0 uint64
}

type CosmosParserManagerParams struct {
	Logger applogger.Logger
	Config CosmosParserManagerConfig
}

type CosmosParser func(
	CosmosParserParams,
) []command.Command

type CosmosParserParams struct {
	MsgCommonParams event.MsgCommonParams
	TxsResult       model.BlockResultsTxsResult
	MsgIndex        int
	Msg             map[string]interface{}
}

func NewCosmosParserManager(params CosmosParserManagerParams) *CosmosParserManager {
	cpm := &CosmosParserManager{
		store:  make(map[string]CosmosParserBlockHeight),
		logger: params.Logger,
		config: params.Config,
	}

	return cpm
}

func (cpm *CosmosParserManager) GetConfig() CosmosParserManagerConfig {
	return cpm.config
}

func (cpm *CosmosParserManager) RegisterParser(name string, fromHeight uint64, parser CosmosParser) {
	if cpm.store[name] == nil {
		cpm.store[name] = make(map[uint64]CosmosParser)
	}
	cpm.store[name][fromHeight] = parser
}

func (cpm *CosmosParserManager) GetParser(name string, blockHeight uint64) CosmosParser {
	parserVersions, ok := cpm.store[name]
	if !ok {
		cpm.logger.Errorf("Requesting invalid parser :%s", name)
	}
	resultBlockHeight := uint64(0)
	resultParser := parserVersions[resultBlockHeight]
	for fromBlockHeight, parser := range parserVersions {
		if fromBlockHeight <= blockHeight && fromBlockHeight > resultBlockHeight {
			resultParser = parser
		}
	}

	return resultParser
}
