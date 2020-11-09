package tendermint

import (
	"github.com/crypto-com/chainindex/usecase/model"
)

type Client interface {
	Block(height int64) (*model.Block, *model.RawBlock, error)
	BlockResults(height int64) (*model.BlockResults, error)
	LatestBlockHeight() (int64, error)
}
