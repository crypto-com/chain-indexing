package event

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const BLOCK_CREATED = "BlockCreated"

type BlockCreated struct {
	entity_event.Base

	Block *usecase_model.Block `json:"block"`
}

func NewBlockCreated(block *usecase_model.Block) *BlockCreated {
	return &BlockCreated{
		entity_event.NewBase(block.Height),

		block,
	}
}

func (event *BlockCreated) Name() string {
	return BLOCK_CREATED
}

func (event *BlockCreated) Version() int {
	return 1
}

func (event *BlockCreated) ToJSON() string {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		panic(fmt.Sprintf("error encoding BlockCreated event to JSON: %v", err))
	}

	return string(encoded)
}

func (evt *BlockCreated) String() string {
	return render.Render(evt)
}
