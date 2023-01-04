package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/event"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_VOTE = "/cosmos.gov.v1.MsgVote"
const MSG_VOTE_CREATED = "/cosmos.gov.v1.MsgVote.Created"
const MSG_VOTE_FAILED = "/cosmos.gov.v1.MsgVote.Failed"

type MsgVote struct {
	event.MsgBase

	ProposalId string `json:"proposalId"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
	Metadata   string `json:"metadata"`
}

func NewMsgVote(msgCommonParams event.MsgCommonParams, params model.MsgVoteParams) *MsgVote {
	return &MsgVote{
		event.NewMsgBase(event.MsgBaseParams{
			MsgName:         MSG_VOTE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Voter,
		params.Option,
		params.Metadata,
	}
}

func (event *MsgVote) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgVote) String() string {
	return render.Render(event)
}

func DecodeMsgVote(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgVote
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
