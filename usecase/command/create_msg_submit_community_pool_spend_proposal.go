package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgSubmitCommunityPoolSpendProposal struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSubmitCommunityPoolSpendProposalParams
}

func NewCreateMsgSubmitCommunityPoolSpendProposal(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSubmitCommunityPoolSpendProposalParams,
) *CreateMsgSubmitCommunityPoolSpendProposal {
	return &CreateMsgSubmitCommunityPoolSpendProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitCommunityPoolSpendProposal) Name() string {
	return "CreateMsgSubmitCommunityPoolSpendProposal"
}

// Version returns version of command
func (*CreateMsgSubmitCommunityPoolSpendProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitCommunityPoolSpendProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitCommunityPoolSpendProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
