package event

import (
	"github.com/luci/go-render/render"

	"github.com/crypto-com/chainindex/ddd/event"
	"github.com/crypto-com/chainindex/entity/model"
)

const BLOCK_CREATED = "BlockCreated"

type BlockCreated struct {
	event.Base

	Block *model.Block
}

func NewBlockCreated(block *model.Block) *BlockCreated {
	return &BlockCreated{
		event.NewBase(),

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

func (event *BlockCreated) String() string {
	return render.Render(event)
}
