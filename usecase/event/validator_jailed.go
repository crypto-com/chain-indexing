package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const VALIDATOR_JAILED = "ValidatorJailed"

type ValidatorJailed struct {
	event_entity.Base

	ConsensusNodeAddress string `json:"consensusNodeAddress"`
	Reason               string `json:"reason"`
}

func NewValidatorJailed(blockHeight int64, consensusNodeAddress string, reason string) *ValidatorJailed {
	return &ValidatorJailed{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        VALIDATOR_JAILED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		consensusNodeAddress,
		reason,
	}

}
func (event *ValidatorJailed) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ValidatorJailed) String() string {
	return render.Render(event)
}

func DecodeValidatorJailed(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ValidatorJailed
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
