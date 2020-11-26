package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/luci/go-render/render"
)

const PROPOSAL_ENDED = "ProposalEnded"

type ProposalEnded struct {
	event_entity.Base

	ProposalId string `json:"proposalId"`
	Result     string `json:"result"`
}

func NewProposalEnded(blockHeight int64, proposalId string, result string) *ProposalEnded {
	return &ProposalEnded{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        PROPOSAL_ENDED,
			Version:     1,
			BlockHeight: blockHeight,
		}),
		proposalId,
		result,
	}

}
func (event *ProposalEnded) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ProposalEnded) String() string {
	return render.Render(event)
}

func DecodeProposalEnded(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ProposalEnded
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
