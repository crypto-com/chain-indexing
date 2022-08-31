package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgNFTEditNFT struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgNFTEditNFTParams
}

func NewCreateMsgNFTEditNFT(
	msgCommonParams event.MsgCommonParams,
	params model.MsgNFTEditNFTParams,
) *CreateMsgNFTEditNFT {
	return &CreateMsgNFTEditNFT{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgNFTEditNFT) Name() string {
	return "/chainmain.nft.v1.MsgEditNFT.Create"
}

// Version returns version of command
func (*CreateMsgNFTEditNFT) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgNFTEditNFT) Exec() (entity_event.Event, error) {
	event := event.NewMsgNFTEditNFT(cmd.msgCommonParams, cmd.params)
	return event, nil
}
