package event

import (
	"github.com/luci/go-render/render"

	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

const BLOCK_CREATED = "BlockCreated"

type BlockCreated struct {
	event.Base

	Block *model.Block `json:"block"`
}

func NewBlockCreated(block *model.Block) *BlockCreated {
	return &BlockCreated{
		event.NewBase(block.Height),

		block,
	}
}

func (event *BlockCreated) Name() string {
	return BLOCK_CREATED
}

func (event *BlockCreated) Version() int {
	return 1
}

func (evt *BlockCreated) String() string {
	return render.Render(evt)
}
