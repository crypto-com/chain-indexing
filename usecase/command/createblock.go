package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	usecase_event "github.com/crypto-com/chainindex/usecase/event"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

type CreateBlock struct {
	block *usecase_model.Block
}

func NewCreateBlock(block *usecase_model.Block) CreateBlock {
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

func (cmd *CreateBlock) Exec() (entity_event.Event, error) {
	evt := usecase_event.NewBlockCreated(cmd.block)
	return evt, nil
}
