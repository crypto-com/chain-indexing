package notification

import (
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type BlockNotification struct {
	Height       int64
	Block        *model.Block
	RawBlock     *model.RawBlock
	BlockResults *model.BlockResults
}

func NewBlockNotification(
	height int64,
	block *model.Block,
	rawBlock *model.RawBlock,
	blockResults *model.BlockResults,
) *BlockNotification {
	return &BlockNotification{
		Height:       height,
		Block:        block,
		RawBlock:     rawBlock,
		BlockResults: blockResults,
	}
}
