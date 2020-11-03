package command

import (
	"github.com/crypto-com/chainindex/ddd"
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/entity/model"
)

type CreateBlockCommand struct {
	block *model.Block
}

func NewCreateBlockCommand(block *model.Block) CreateBlockCommand {
	return CreateBlockCommand{
		block,
	}
}

func (_ *CreateBlockCommand) Name() string {
	return "CreateBlock"
}

func (_ *CreateBlockCommand) Version() int {
	return 1
}

func (cmd *CreateBlockCommand) Exec() (ddd.Event, error) {
	evt := event.NewBlockCreated(cmd.block)
	return evt, nil
}
