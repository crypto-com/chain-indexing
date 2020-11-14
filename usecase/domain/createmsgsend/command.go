package createmsgsend

import entity_event "github.com/crypto-com/chainindex/entity/event"

type CreateMsgSendCommand struct {
	blockHeight int64
	txHash      string
	msgIndex    int
	params      Params
}

func NewCommand(blockHeight int64, txHash string, msgIndex int, params Params) *CreateMsgSendCommand {
	return &CreateMsgSendCommand{
		blockHeight,
		txHash,
		msgIndex,
		params,
	}
}

func (_ *CreateMsgSendCommand) Name() string {
	return "CreateMsgSendCommand"
}

func (_ *CreateMsgSendCommand) Version() int {
	return 1
}

func (cmd *CreateMsgSendCommand) Exec() (entity_event.Event, error) {
	event := NewEvent(cmd.blockHeight, cmd.txHash, cmd.msgIndex, cmd.params)
	return event, nil
}
