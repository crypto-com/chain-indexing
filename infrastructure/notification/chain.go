package notification

import (
	"github.com/crypto-com/chainindex/usecase/model"
)

type BlockNotification struct {
	Height   int64
	Block    *model.Block
	RawBlock *model.RawBlock
}

func NewBlockNotification(height int64, block *model.Block, rawBlock *model.RawBlock) *BlockNotification {
	return &BlockNotification{
		Height:   height,
		Block:    block,
		RawBlock: rawBlock,
	}
}
