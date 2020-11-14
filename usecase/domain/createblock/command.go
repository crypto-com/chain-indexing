package createblock

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

type CreateBlockCommand struct {
	block *usecase_model.Block
}

func NewCommand(block *usecase_model.Block) *CreateBlockCommand {
	return &CreateBlockCommand{
		block,
	}
}

func (_ *CreateBlockCommand) Name() string {
	return "CreateBlockCommand"
}

func (_ *CreateBlockCommand) Version() int {
	return 1
}

func (cmd *CreateBlockCommand) Exec() (entity_event.Event, error) {
	event := NewEvent(cmd.block)
	return event, nil
}
