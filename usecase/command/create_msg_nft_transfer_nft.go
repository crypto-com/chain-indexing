package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateMsgNFTTransferNFT struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgNFTTransferNFTParams
}

func NewCreateMsgNFTTransferNFT(
	msgCommonParams event.MsgCommonParams,
	params model.MsgNFTTransferNFTParams,
) *CreateMsgNFTTransferNFT {
	return &CreateMsgNFTTransferNFT{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgNFTTransferNFT) Name() string {
	return "CreateMsgNFTTransferNFT"
}

// Version returns version of command
func (*CreateMsgNFTTransferNFT) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgNFTTransferNFT) Exec() (entity_event.Event, error) {
	event := event.NewMsgNFTTransferNFT(cmd.msgCommonParams, cmd.params)
	return event, nil
}
