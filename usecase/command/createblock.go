package command

import (
	"github.com/crypto-com/chainindex/entity/event"
	usecase_event "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateBlock struct {
	block *model.Block
}

func NewCreateBlock(block *model.Block) CreateBlock {
	return CreateBlock{
		block,
	}
}

func (_ *CreateBlock) Name() string {
	return "CreateBlock"
}

func (_ *CreateBlock) Version() int {
	return 1
}

func (cmd *CreateBlock) Exec() (event.Event, error) {
	evt := usecase_event.NewBlockCreated(cmd.block)
	return evt, nil
}
