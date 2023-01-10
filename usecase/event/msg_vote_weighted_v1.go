package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_VOTE_WEIGHTED_V1 = "/cosmos.gov.v1.MsgVoteWeighted"
const MSG_VOTE_WEIGHTED_V1_CREATED = "/cosmos.gov.v1.MsgVoteWeighted.Created"
const MSG_VOTE_WEIGHTED_V1_FAILED = "/cosmos.gov.v1.MsgVoteWeighted.Failed"

type MsgVoteWeightedV1 struct {
	MsgBase

	ProposalId  string             `json:"proposalId"`
	Voter       string             `json:"voter"`
	VoteOptions []model.VoteOption `json:"vote_options"`
	Metadata    string             `json:"metadata"`
}

func NewMsgVoteWeightedV1(msgCommonParams MsgCommonParams, params model.MsgVoteWeightedParams) *MsgVoteWeightedV1 {
	return &MsgVoteWeightedV1{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_VOTE_WEIGHTED_V1,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Voter,
		params.VoteOptions,
		params.Metadata,
	}
}

func (event *MsgVoteWeightedV1) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgVoteWeightedV1) String() string {
	return render.Render(event)
}

func DecodeMsgVoteWeightedV1(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgVoteWeightedV1
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
