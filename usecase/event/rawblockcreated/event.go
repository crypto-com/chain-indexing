package rawblockcreated

import (
	"bytes"
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const NAME = "RawBlockCreated"

type RawBlockCreated struct {
	entity_event.Base

	RawBlock *usecase_model.RawBlock `json:"rawBlock"`
}

func New(rawBlock *usecase_model.RawBlock) *RawBlockCreated {
	height, err := strconv.ParseInt(rawBlock.Block.Header.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Missing block height in raw block: %v", err))
	}

	return &RawBlockCreated{
		entity_event.NewBase(height),

		rawBlock,
	}
}

func (_ *RawBlockCreated) Name() string {
	return NAME
}

func (_ *RawBlockCreated) Version() int {
	return 1
}

func (event *RawBlockCreated) ToJSON() string {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		panic(fmt.Sprintf("error encoding RawBlockCreated event to JSON: %v", err))
	}

	return string(encoded)
}

func (evt *RawBlockCreated) String() string {
	return render.Render(evt)
}

func Decode(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawBlockCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
