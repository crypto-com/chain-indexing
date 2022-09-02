package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgNFTIssueDenom struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgNFTIssueDenomParams
}

func NewCreateMsgNFTIssueDenom(
	msgCommonParams event.MsgCommonParams,
	params model.MsgNFTIssueDenomParams,
) *CreateMsgNFTIssueDenom {
	return &CreateMsgNFTIssueDenom{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgNFTIssueDenom) Name() string {
	return "/chainmain.nft.v1.MsgIssueDenom.Create"
}

// Version returns version of command
func (*CreateMsgNFTIssueDenom) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgNFTIssueDenom) Exec() (entity_event.Event, error) {
	event := event.NewMsgNFTIssueDenom(cmd.msgCommonParams, cmd.params)
	return event, nil
}
