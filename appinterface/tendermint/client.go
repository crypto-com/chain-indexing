package tendermint

import (
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

type Client interface {
	Block(height int64) (*usecase_model.Block, *usecase_model.RawBlock, error)
	BlockResults(height int64) (*usecase_model.BlockResults, error)
	LatestBlockHeight() (int64, error)
}
