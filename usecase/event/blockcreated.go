package event

import (
	"bytes"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

const BLOCK_CREATED = "BlockCreated"

type BlockCreated struct {
	entity_event.Base

	Block *usecase_model.Block `json:"block"`
}

func NewBlockCreated(block *usecase_model.Block) *BlockCreated {
	return &BlockCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        BLOCK_CREATED,
			Version:     1,
			BlockHeight: block.Height,
		}),

		block,
	}
}

func (event *BlockCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockCreated) String() string {
	return render.Render(event)
}

func DecodeBlockCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
