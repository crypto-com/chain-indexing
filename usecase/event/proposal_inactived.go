package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const PROPOSAL_INACTIVED = "ProposalInactived"

type ProposalInactived struct {
	event_entity.Base

	ProposalId string `json:"proposalId"`
	Result     string `json:"result"`
}

func NewProposalInactived(blockHeight int64, proposalId string, result string) *ProposalInactived {
	return &ProposalInactived{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        PROPOSAL_INACTIVED,
			Version:     1,
			BlockHeight: blockHeight,
		}),
		proposalId,
		result,
	}

}
func (event *ProposalInactived) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ProposalInactived) String() string {
	return render.Render(event)
}

func DecodeProposalInactived(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ProposalInactived
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
