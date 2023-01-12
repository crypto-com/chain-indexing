package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CreateMsgSubmitProposal struct {
	msgCommonParams event.MsgCommonParams
	params          v1_model.MsgSubmitProposalParams
}

func NewCreateMsgSubmitProposal(
	msgCommonParams event.MsgCommonParams,
	params v1_model.MsgSubmitProposalParams,
) *CreateMsgSubmitProposal {
	return &CreateMsgSubmitProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitProposal) Name() string {
	return "/cosmos.gov..MsgSubmitProposal.Create"
}

// Version returns version of command
func (*CreateMsgSubmitProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
