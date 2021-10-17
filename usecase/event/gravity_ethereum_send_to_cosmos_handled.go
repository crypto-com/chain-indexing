package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED = "GravityEthereumSendToCosmosHandled"

type GravityEthereumSendToCosmosHandled struct {
	event_entity.Base

	Params model.GravityEthereumSendToCosmosHandledEventParams `json:"params"`
}

func NewGravityEthereumSendToCosmosHandled(blockHeight int64, params model.GravityEthereumSendToCosmosHandledEventParams) *GravityEthereumSendToCosmosHandled {
	return &GravityEthereumSendToCosmosHandled{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params,
	}

}
func (event *GravityEthereumSendToCosmosHandled) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *GravityEthereumSendToCosmosHandled) String() string {
	return render.Render(event)
}

func DecodeGravityEthereumSendToCosmosHandled(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ValidatorJailed
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
