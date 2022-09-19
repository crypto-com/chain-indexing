package event

import (
	"bytes"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/utctime"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

const BLOCK_RAW_EVENT_CREATED = "BlockRawEventCreated"

type BlockRawEventCreated struct {
	entity_event.Base

	BlockHash  string                   `json:"blockHash"`
	BlockTime  utctime.UTCTime          `json:"blockTime"`
	FromResult string                   `json:"fromResult"`
	Data       usecase_model.DataParams `json:"data"`
}

func NewBlockRawEventCreated(blockHeight int64, params *usecase_model.CreateBlockRawEventParams) *BlockRawEventCreated {
	return &BlockRawEventCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        BLOCK_RAW_EVENT_CREATED,
			Version:     1,
			BlockHeight: blockHeight,
		}),
		params.BlockHash,
		params.BlockTime,
		params.FromResult,
		params.Data,
	}
}

func (event *BlockRawEventCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockRawEventCreated) String() string {
	return render.Render(event)
}

func DecodeBlockRawEventCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockRawEventCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
