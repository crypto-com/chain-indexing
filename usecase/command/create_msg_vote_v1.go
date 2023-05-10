package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
)

type CreateMsgVotV1 struct {
	msgCommonParams event.MsgCommonParams
	params          model_gov_v1.MsgVoteParams
}

func NewCreateMsgVoteV1(
	msgCommonParams event.MsgCommonParams,
	params model_gov_v1.MsgVoteParams,
) *CreateMsgVotV1 {
	return &CreateMsgVotV1{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgVotV1) Name() string {
	return "/cosmos.gov.v1.MsgVote.Create"
}

// Version returns version of command
func (*CreateMsgVotV1) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgVotV1) Exec() (entity_event.Event, error) {
	event := event.NewMsgVoteV1(cmd.msgCommonParams, cmd.params)
	return event, nil
}
