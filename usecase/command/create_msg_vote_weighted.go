package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgVoteWeighted struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgVoteWeightedParams
}

func NewCreateMsgVoteWeightedV1(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgVoteWeightedParams,
) *CreateMsgVoteWeighted {
	return &CreateMsgVoteWeighted{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgVoteWeighted) Name() string {
	return "/cosmos.gov.v1.MsgVoteWeighted.Create"
}

// Version returns version of command
func (*CreateMsgVoteWeighted) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgVoteWeighted) Exec() (entity_event.Event, error) {
	event := event.NewMsgVoteWeighted(cmd.msgCommonParams, cmd.params)
	return event, nil
}
