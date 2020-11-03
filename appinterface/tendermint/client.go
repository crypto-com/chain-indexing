package tendermint

import "github.com/crypto-com/chainindex/appinterface/tendermint/types"

type Client interface {
	Block(height int64) (*types.Block, error)
	LatestBlockHeight() (int64, error)
}
