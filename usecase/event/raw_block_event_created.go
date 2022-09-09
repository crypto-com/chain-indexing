package event

import (
	"bytes"

	"github.com/luci/go-render/render"

	jsoniter "github.com/json-iterator/go"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/utctime"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

const RAW_BLOCK_EVENT_CREATED = "RawBlockEventCreated"

type RawBlockEventCreated struct {
	entity_event.Base

	BlockHash  string                          `json:"blockHash"`
	BlockTime  utctime.UTCTime                 `json:"blockTime"`
	FromResult string                          `json:"fromResult"`
	RawData    usecase_model.BlockResultsEvent `json:"rawData"`
}

func NewRawBlockEvent(blockHeight int64, params *usecase_model.CreateRawBlockEventParams) *RawBlockEventCreated {
	return &RawBlockEventCreated{
		entity_event.NewBase(entity_event.BaseParams{
			Name:        RAW_BLOCK_EVENT_CREATED,
			Version:     1,
			BlockHeight: blockHeight,
		}),
		params.BlockHash,
		params.BlockTime,
		params.FromResult,
		params.RawData,
	}
}

func (event *RawBlockEventCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *RawBlockEventCreated) String() string {
	return render.Render(event)
}

func DecodeRawBlockEventCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *RawBlockEventCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
