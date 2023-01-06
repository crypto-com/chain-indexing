package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_VOTE_WEIGHTED = "/cosmos.gov.v1.MsgVoteWeighted"
const MSG_VOTE_WEIGHTED_CREATED = "/cosmos.gov.v1.MsgVoteWeighted.Created"
const MSG_VOTE_WEIGHTED_FAILED = "/cosmos.gov.v1.MsgVoteWeighted.Failed"

type MsgVoteWeighted struct {
	MsgBase

	ProposalId string           `json:"proposalId"`
	Voter      string           `json:"voter"`
	VoteOption model.VoteOption `json:"vote_option"`
	Metadata   string           `json:"metadata"`
}

func NewMsgVoteWeighted(msgCommonParams MsgCommonParams, params model.MsgVoteWeightedParams) *MsgVoteWeighted {
	return &MsgVoteWeighted{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_VOTE_WEIGHTED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Voter,
		params.VoteOption,
		params.Metadata,
	}
}

func (event *MsgVoteWeighted) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgVoteWeighted) String() string {
	return render.Render(event)
}

func DecodeMsgVoteWeighted(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgVoteWeighted
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
