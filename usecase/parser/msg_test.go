package parser_test

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chainindex/usecase/model/genesis"

	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	"github.com/crypto-com/chainindex/usecase/model"
)

func mustParseBlockResp(rawResp string) (*model.Block, *model.RawBlock) {
	block, rawBlock, err := tendermint.ParseBlockResp(strings.NewReader(rawResp))
	if err != nil {
		panic(fmt.Sprintf("error parsing block response: %v", err))
	}

	return block, rawBlock
}

func mustParseBlockResultsResp(rawResp string) *model.BlockResults {
	blockResults, err := tendermint.ParseBlockResultsResp(strings.NewReader(rawResp))

	if err != nil {
		panic(fmt.Sprintf("error parsing block results response: %v", err))
	}

	return blockResults
}

func mustParseGenesisResp(rawResp string) *genesis.Genesis {
	genesis, err := tendermint.ParseGenesisResp(strings.NewReader(rawResp))

	if err != nil {
		panic(fmt.Sprintf("error parsing block genesis response: %v", err))
	}

	return genesis
}
