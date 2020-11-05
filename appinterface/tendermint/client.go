package tendermint

import (
	"github.com/crypto-com/chainindex/usecase/model"
)

type Client interface {
	Block(height int64) (*model.Block, *model.RawBlock, error)
	LatestBlockHeight() (int64, error)
}
