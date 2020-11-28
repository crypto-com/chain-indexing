package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const POWER_CHANGED = "PowerChanged"

type PowerChanged struct {
	event_entity.Base

	TendermintPubkey string `json:"tendermintPubkey"`
	Power            string `json:"power"`
}

func NewPowerChanged(blockHeight int64, params model.PowerChangeParams) *PowerChanged {
	return &PowerChanged{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        POWER_CHANGED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.TendermintPubkey,
		params.Power,
	}

}
func (event *PowerChanged) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *PowerChanged) String() string {
	return render.Render(event)
}

func DecodePowerChanged(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *PowerChanged
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
