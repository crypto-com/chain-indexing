package utils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/command"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
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
	V0_42_7 ParserBlockHeight
}

type CosmosParserManagerParams struct {
	Logger applogger.Logger
	Config CosmosParserManagerConfig
}

type CosmosParser func(
	CosmosParserParams,
) ([]command.Command, string)

type CosmosParserParams struct {
	AddressPrefix   string
	StakingDenom    string
	TxsResult       model.BlockResultsTxsResult
	MsgCommonParams event.MsgCommonParams
	MsgIndex        int
	Msg             map[string]interface{}
	ParserManager   *CosmosParserManager
}

func NewCosmosParserManager(params CosmosParserManagerParams) *CosmosParserManager {
	cpm := &CosmosParserManager{
		store:  make(map[CosmosParserKey]BlockHeightToCosmosParserMap),
		logger: params.Logger,
		config: params.Config,
	}

	return cpm
}

// GetCosmosV0_42_7BlockHeight return height of the first block with cosmos sdk v0.42.7
func (cpm *CosmosParserManager) GetCosmosV0_42_7BlockHeight() ParserBlockHeight {
	return cpm.config.CosmosVersionBlockHeight.V0_42_7
}

// RegisterParser register a cosmos message parser by a given key and a starting block height
func (cpm *CosmosParserManager) RegisterParser(name CosmosParserKey, fromHeight ParserBlockHeight, parser CosmosParser) {
	if cpm.store[name] == nil {
		cpm.store[name] = make(map[ParserBlockHeight]CosmosParser)
	}
	cpm.store[name][fromHeight] = parser
}

// GetParser return a cosmos message parser from a registered key and a specific block height.
// Panic if the key is not found in the registered store
func (cpm *CosmosParserManager) GetParser(name CosmosParserKey, blockHeight ParserBlockHeight) CosmosParser {
	parserVersions, ok := cpm.store[name]
	if !ok {
		panic(fmt.Sprintf("Requesting invalid parser :%s", name))
	}
	enabledBlockHeight := ParserBlockHeight(0)
	parser := parserVersions[enabledBlockHeight]
	for curEnabledBlockHeight, curParser := range parserVersions {
		if isEnabledBlock(curEnabledBlockHeight, blockHeight) && isLaterVersion(curEnabledBlockHeight, enabledBlockHeight) {
			enabledBlockHeight = curEnabledBlockHeight
			parser = curParser
		}
	}

	return parser
}

func isEnabledBlock(enabledBlockHeight ParserBlockHeight, blockHeight ParserBlockHeight) bool {
	return enabledBlockHeight <= blockHeight
}

func isLaterVersion(enabledBlockHeight ParserBlockHeight, existingEnableBlockHeight ParserBlockHeight) bool {
	return enabledBlockHeight > existingEnableBlockHeight
}
