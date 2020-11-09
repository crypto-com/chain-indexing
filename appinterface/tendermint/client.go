package tendermint

import (
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

type Client interface {
	Block(height int64) (*usecase_model.Block, *usecase_model.RawBlock, error)
	LatestBlockHeight() (int64, error)
}
