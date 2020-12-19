package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgSubmitTextProposal struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSubmitTextProposalParams
}

func NewCreateMsgSubmitTextProposal(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSubmitTextProposalParams,
) *CreateMsgSubmitTextProposal {
	return &CreateMsgSubmitTextProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitTextProposal) Name() string {
	return "CreateMsgSubmitTextProposal"
}

// Version returns version of command
func (*CreateMsgSubmitTextProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitTextProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitTextProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
