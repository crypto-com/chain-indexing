package event

import "github.com/crypto-com/chainindex/entity/model"

const BLOCK_CREATED = "BlockCreated"

type BlockCreated struct {
	Block *model.Block
}

func NewBlockCreated(block *model.Block) *BlockCreated {
	return &BlockCreated{
		block,
	}
}

func (event *BlockCreated) Name() string {
	return BLOCK_CREATED
}

func (event *BlockCreated) Version() int {
	return 1
}

func (event *BlockCreated) Payload() interface{} {
	return event.Block
}
