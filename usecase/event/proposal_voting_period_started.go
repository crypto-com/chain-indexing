package event

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const PROPOSAL_VOTING_PERIOD_STARTED = "ProposalVotingPeriodStarted"

type ProposalVotingPeriodStarted struct {
	event_entity.Base

	ProposalId string `json:"proposalId"`
}

func NewProposalVotingPeriodStarted(blockHeight int64, proposalId string) *ProposalVotingPeriodStarted {
	return &ProposalVotingPeriodStarted{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        PROPOSAL_VOTING_PERIOD_STARTED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		proposalId,
	}

}
func (event *ProposalVotingPeriodStarted) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ProposalVotingPeriodStarted) String() string {
	return render.Render(event)
}

func DecodeProposalVotingPeriodStarted(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ProposalVotingPeriodStarted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
