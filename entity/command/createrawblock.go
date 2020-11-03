package command

import (
	"github.com/crypto-com/chainindex/ddd"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/model"
)

type CreateRawBlockCommand struct {
	rawBlock *model.RawBlock
}

func NewCreateRawBlockCommand(rawBlock *model.RawBlock) CreateRawBlockCommand {
	return CreateRawBlockCommand{
		rawBlock,
	}
}

func (_ *CreateRawBlockCommand) Name() string {
	return "CreateRawBlock"
}

func (_ *CreateRawBlockCommand) Version() int {
	return 1
}

func (cmd *CreateRawBlockCommand) Exec() (ddd.Event, error) {
	evt := event.NewRawBlockCreatedEvent(cmd.rawBlock)
	return evt, nil
}
