package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgSubmitUnknownProposal struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSubmitUnknownProposalParams
}

func NewCreateMsgSubmitUnknownProposal(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSubmitUnknownProposalParams,
) *CreateMsgSubmitUnknownProposal {
	return &CreateMsgSubmitUnknownProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitUnknownProposal) Name() string {
	return "CreateMsgSubmitUnknownProposal"
}

// Version returns version of command
func (*CreateMsgSubmitUnknownProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitUnknownProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitUnknownProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
