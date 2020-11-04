package command

import (
	dddevent "github.com/crypto-com/chainindex/ddd/event"
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

func (cmd *CreateRawBlockCommand) Exec() (dddevent.Event, error) {
	evt := event.NewRawBlockCreatedEvent(cmd.rawBlock)
	return evt, nil
}
