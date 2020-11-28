package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const VALIDATOR_SLASHED = "ValidatorSlashed"

type ValidatorSlashed struct {
	event_entity.Base

	ConsensusNodeAddress string `json:"consensusNodeAddress"`
	SlashedPower         string `json:"slashedPower"`
	Reason               string `json:"reason"`
}

func NewValidatorSlashed(blockHeight int64, params model.SlashValidatorParams) *ValidatorSlashed {
	return &ValidatorSlashed{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        VALIDATOR_SLASHED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params.ConsensusNodeAddress,
		params.SlashedPower,
		params.Reason,
	}

}
func (event *ValidatorSlashed) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ValidatorSlashed) String() string {
	return render.Render(event)
}

func DecodeValidatorSlashed(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ValidatorSlashed
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
