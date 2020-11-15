package event

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

const RAW_BLOCK_CREATED_NAME = "RawBlockCreated"

type RawBlockCreated struct {
	entity_event.Base

	RawBlock *usecase_model.RawBlock `json:"rawBlock"`
}

func NewEvent(rawBlock *usecase_model.RawBlock) *RawBlockCreated {
	blockHeight, err := strconv.ParseInt(rawBlock.Block.Header.Height, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Missing block blockHeight in raw block: %v", err))
	}

	return &RawBlockCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        RAW_BLOCK_CREATED_NAME,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		rawBlock,
	}
}

func (event *RawBlockCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *RawBlockCreated) String() string {
	return render.Render(event)
}

func DecodeRawBlockCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawBlockCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
