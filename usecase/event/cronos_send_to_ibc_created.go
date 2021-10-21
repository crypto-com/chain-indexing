package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	"github.com/luci/go-render/render"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
)

const CRONOS_SEND_TO_IBC_CREATED = "CronosSendToIBCCreated"

type CronosSendToIBCCreated struct {
	event_entity.Base

	Params model.CronosSendToIBCParams `json:"params"`
}

func NewCronosSendToIBCCreated(
	blockHeight int64,
	params model.CronosSendToIBCParams,
) *CronosSendToIBCCreated {
	return &CronosSendToIBCCreated{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        CRONOS_SEND_TO_IBC_CREATED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *CronosSendToIBCCreated) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *CronosSendToIBCCreated) String() string {
	return render.Render(event)
}

func DecodeCronosSendToIBCCreated(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *CronosSendToIBCCreated
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
