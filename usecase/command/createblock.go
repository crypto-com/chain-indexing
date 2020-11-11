package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event/blockcreated"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
)

type CreateBlock struct {
	block *usecase_model.Block
}

func NewCreateBlock(block *usecase_model.Block) *CreateBlock {
	return &CreateBlock{
		block,
	}
}

func (_ *CreateBlock) Name() string {
	return "CreateBlock"
}

func (_ *CreateBlock) Version() int {
	return 0
}

func (cmd *CreateBlock) Exec() (entity_event.Event, error) {
	evt := blockcreated.New(cmd.block)
	return evt, nil
}
