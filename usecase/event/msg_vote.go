package event

import (
	"bytes"

	"github.com/crypto-com/chainindex/usecase/model"

	entity_event "github.com/crypto-com/chainindex/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_VOTE = "MsgVote"
const MSG_VOTE_CREATED = "MsgVoteCreated"
const MSG_VOTE_FAILED = "MsgVoteFailed"

type MsgVote struct {
	MsgBase

	ProposalId string `json:"proposalId"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
}

func NewMsgVote(msgCommonParams MsgCommonParams, params model.MsgVoteParams) *MsgVote {
	return &MsgVote{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_VOTE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.ProposalId,
		params.Voter,
		params.Option,
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
