package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateRawBlock struct {
	rawBlock *usecase_model.RawBlock
}

func NewCreateRawBlock(rawBlock *usecase_model.RawBlock) *CreateRawBlock {
	return &CreateRawBlock{
		rawBlock,
	}
}

func (_ *CreateRawBlock) Name() string {
	return "CreateRawBlock"
}

func (_ *CreateRawBlock) Version() int {
	return 1
}

func (cmd *CreateRawBlock) Exec() (entity_event.Event, error) {
	evt := event.NewEvent(cmd.rawBlock)
	return evt, nil
}
