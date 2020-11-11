package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event/rawblockcreated"
	usecase_model "github.com/crypto-com/chainindex/usecase/model"
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
	return 0
}

func (cmd *CreateRawBlock) Exec() (entity_event.Event, error) {
	evt := rawblockcreated.New(cmd.rawBlock)
	return evt, nil
}
