package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const EVIDENCE_SUBMITTED = "EvidenceSubmitted"

type EvidenceSubmitted struct {
	event_entity.Base

	model.EvidenceParams
}

func NewEvidenceSubmitted(blockHeight int64, params model.EvidenceParams) *EvidenceSubmitted {
	return &EvidenceSubmitted{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        EVIDENCE_SUBMITTED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params,
	}

}
func (event *EvidenceSubmitted) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *EvidenceSubmitted) String() string {
	return render.Render(event)
}

func DecodeEvidenceSubmitted(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *EvidenceSubmitted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
