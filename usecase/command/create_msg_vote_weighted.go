package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgVoteWeightedV1 struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgVoteWeightedParams
}

func NewCreateMsgVoteWeightedV1(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgVoteWeightedParams,
) *CreateMsgVoteWeightedV1 {
	return &CreateMsgVoteWeightedV1{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgVoteWeightedV1) Name() string {
	return "/cosmos.gov.v1.MsgVoteWeighted.Create"
}

// Version returns version of command
func (*CreateMsgVoteWeightedV1) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgVoteWeightedV1) Exec() (entity_event.Event, error) {
	event := event.NewMsgVoteWeightedV1(cmd.msgCommonParams, cmd.params)
	return event, nil
}
