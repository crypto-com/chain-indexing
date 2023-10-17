package parser_test

import (
	"fmt"
	"strings"

	tendermint_interface "github.com/crypto-com/chain-indexing/appinterface/tendermint"

	"github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"
)

func mustParseBlockResp(rawResp string) (*model.Block, *model.RawBlock) {
	block, rawBlock, err := tendermint.ParseBlockResp(strings.NewReader(rawResp))
	if err != nil {
		panic(fmt.Sprintf("error parsing block response: %v", err))
	}

	return block, rawBlock
}

func mustParseBlockResultsResp(rawResp string, decoder tendermint_interface.BlockResultEventAttributeDecoder) *model.BlockResults {
	blockResults, err := tendermint.ParseBlockResultsResp(strings.NewReader(rawResp), decoder)

	if err != nil {
		panic(fmt.Sprintf("error parsing block results response: %v", err))
	}

	return blockResults
}

func mustParseGenesisResp(rawResp string, strictParsing bool) *genesis.Genesis {
	genesis, err := tendermint.ParseGenesisResp(strings.NewReader(rawResp), strictParsing)

	if err != nil {
		panic(fmt.Sprintf("error parsing block genesis response: %v", err))
	}

	return genesis
}

func MustParseTxsResp(rawResp string) *model.CosmosTxWithHash {
	tx, err := cosmosapp.ParseTxsResp(strings.NewReader(rawResp))

	if err != nil {
		panic(fmt.Sprintf("error parsing block results response: %v", err))
	}

	return &model.CosmosTxWithHash{
		Tx:   tx.Tx,
		Hash: tx.TxResponse.TxHash,
	}
}
