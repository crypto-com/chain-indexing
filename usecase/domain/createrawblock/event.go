package createrawblock

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const EVENT_NAME = "RawBlockCreatedEvent"

type RawBlockCreatedEvent struct {
	entity_event.Base

	RawBlock *usecase_model.RawBlock `json:"rawBlock"`
}

func NewEvent(rawBlock *usecase_model.RawBlock) *RawBlockCreatedEvent {
	blockHeight, err := strconv.ParseInt(rawBlock.Block.Header.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Missing block blockHeight in raw block: %v", err))
	}

	return &RawBlockCreatedEvent{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        EVENT_NAME,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		rawBlock,
	}
}

func (event *RawBlockCreatedEvent) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *RawBlockCreatedEvent) String() string {
	return render.Render(event)
}

func DecodeEvent(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawBlockCreatedEvent
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
