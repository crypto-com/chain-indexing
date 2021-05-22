package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgNFTMintNFT struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgNFTMintNFTParams
}

func NewCreateMsgNFTMintNFT(
	msgCommonParams event.MsgCommonParams,
	params model.MsgNFTMintNFTParams,
) *CreateMsgNFTMintNFT {
	return &CreateMsgNFTMintNFT{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgNFTMintNFT) Name() string {
	return "CreateMsgNFTMintNFT"
}

// Version returns version of command
func (*CreateMsgNFTMintNFT) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgNFTMintNFT) Exec() (entity_event.Event, error) {
	event := event.NewMsgNFTMintNFT(cmd.msgCommonParams, cmd.params)
	return event, nil
}
