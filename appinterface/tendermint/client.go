package tendermint

import (
	"github.com/crypto-com/chainindex/appinterface/tendermint/types"
	"github.com/crypto-com/chainindex/entity/model"
)

type Client interface {
	Block(height int64) (*types.Block, *model.RawBlock, error)
	LatestBlockHeight() (int64, error)
}
