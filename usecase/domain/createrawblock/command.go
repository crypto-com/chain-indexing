package createrawblock

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

type CreateRawBlockCommand struct {
	rawBlock *usecase_model.RawBlock
}

func NewCommand(rawBlock *usecase_model.RawBlock) *CreateRawBlockCommand {
	return &CreateRawBlockCommand{
		rawBlock,
	}
}

func (_ *CreateRawBlockCommand) Name() string {
	return "CreateRawBlockCommand"
}

func (_ *CreateRawBlockCommand) Version() int {
	return 1
}

func (cmd *CreateRawBlockCommand) Exec() (entity_event.Event, error) {
	evt := NewEvent(cmd.rawBlock)
	return evt, nil
}
