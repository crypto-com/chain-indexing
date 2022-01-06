package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const EVIDENCE = "Evidence"

type Evidence struct {
	event_entity.Base

	model.EvidenceParams
}

func NewEvidence(blockHeight int64, params model.EvidenceParams) *Evidence {
	return &Evidence{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        EVIDENCE,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		params,
	}

}
func (event *Evidence) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *Evidence) String() string {
	return render.Render(event)
}

func DecodeEvidence(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *Evidence
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
