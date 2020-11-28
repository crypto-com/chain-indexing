package command

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_usecase "github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateBlock struct {
	block *model_usecase.Block
}

func NewCreateBLock(block *model_usecase.Block) *CreateBlock {
	return &CreateBlock{
		block,
	}
}

func (_ *CreateBlock) Name() string {
	return "CreateBlock"
}

func (_ *CreateBlock) Version() int {
	return 1
}

func (cmd *CreateBlock) Exec() (event_entity.Event, error) {
	event := event_usecase.NewBlockCreated(cmd.block)
	return event, nil
}
