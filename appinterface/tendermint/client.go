package tendermint

import "github.com/crypto-com/chainindex/appinterface/tendermint/types"

type Client interface {
	Block(height uint64) (*types.Block, error)
}
