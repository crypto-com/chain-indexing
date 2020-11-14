package createtransaction

import entity_event "github.com/crypto-com/chainindex/entity/event"

type CreateTransactionCommand struct {
	blockHeight int64
	params      Params
}

func NewCommand(blockHeight int64, params Params) *CreateTransactionCommand {
	return &CreateTransactionCommand{
		blockHeight,
		params,
	}
}

func (_ *CreateTransactionCommand) Name() string {
	return "CreateTransactionCommand"
}

func (_ *CreateTransactionCommand) Version() int {
	return 1
}

func (cmd *CreateTransactionCommand) Exec() (entity_event.Event, error) {
	event := NewEvent(cmd.blockHeight, cmd.params)
	return event, nil
}
