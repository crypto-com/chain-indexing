package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_VOTE_V1 = "/cosmos.gov.v1.MsgVote"
const MSG_VOTE_V1_CREATED = "/cosmos.gov.v1.MsgVote.Created"
const MSG_VOTE_V1_FAILED = "/cosmos.gov.v1.MsgVote.Failed"

type MsgVoteV1 struct {
	MsgBase

	ProposalId string `json:"proposalId"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
	Metadata   string `json:"metadata"`
}

func NewMsgVoteV1(msgCommonParams MsgCommonParams, params model.MsgVoteParams) *MsgVoteV1 {
	return &MsgVoteV1{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_VOTE_V1,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Voter,
		params.Option,
		params.Metadata,
	}
}

func (event *MsgVoteV1) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgVoteV1) String() string {
	return render.Render(event)
}

func DecodeMsgVoteV1(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgVoteV1
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
