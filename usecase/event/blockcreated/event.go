package blockcreated

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const NAME = "BlockCreated"

type BlockCreated struct {
	entity_event.Base

	Block *usecase_model.Block `json:"block"`
}

func New(block *usecase_model.Block) *BlockCreated {
	return &BlockCreated{
		entity_event.NewBase(block.Height),

		block,
	}
}

func (_ *BlockCreated) Name() string {
	return NAME
}

func (_ *BlockCreated) Version() int {
	return 1
}

func (event *BlockCreated) String() string {
	return render.Render(event)
}

func Decode(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
