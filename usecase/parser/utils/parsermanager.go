package utils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/command"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CosmosParserManager struct {
	store  map[CosmosParserKey]BlockHeightToCosmosParserMap
	logger applogger.Logger
	config CosmosParserManagerConfig
}

type CosmosParserKey string

type ParserBlockHeight uint64

type BlockHeightToCosmosParserMap map[ParserBlockHeight]CosmosParser

type CosmosParserManagerConfig struct {
	CosmosVersionBlockHeight
}

type CosmosVersionBlockHeight struct {
	V0_43_0 ParserBlockHeight
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
		store:  make(map[CosmosParserKey]BlockHeightToCosmosParserMap),
		logger: params.Logger,
		config: params.Config,
	}

	return cpm
}

func (cpm *CosmosParserManager) GetCosmosV0_43_0BlockHeight() ParserBlockHeight {
	return cpm.config.CosmosVersionBlockHeight.V0_43_0
}

func (cpm *CosmosParserManager) RegisterParser(name CosmosParserKey, fromHeight ParserBlockHeight, parser CosmosParser) {
	if cpm.store[name] == nil {
		cpm.store[name] = make(map[ParserBlockHeight]CosmosParser)
	}
	cpm.store[name][fromHeight] = parser
}

func (cpm *CosmosParserManager) GetParser(name CosmosParserKey, blockHeight ParserBlockHeight) CosmosParser {
	parserVersions, ok := cpm.store[name]
	if !ok {
		panic(fmt.Sprintf("Requesting invalid parser :%s", name))
	}
	resultBlockHeight := ParserBlockHeight(0)
	resultParser := parserVersions[resultBlockHeight]
	for fromBlockHeight, parser := range parserVersions {
		if fromBlockHeight <= blockHeight && fromBlockHeight > resultBlockHeight {
			resultBlockHeight = fromBlockHeight
			resultParser = parser
		}
	}

	return resultParser
}
