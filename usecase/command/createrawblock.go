package command

import (
	"github.com/crypto-com/chainindex/entity/event"
	usecase_event "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateRawBlock struct {
	rawBlock *model.RawBlock
}

func NewCreateRawBlock(rawBlock *model.RawBlock) CreateRawBlock {
	return CreateRawBlock{
		rawBlock,
	}
}

func (_ *CreateRawBlock) Name() string {
	return "CreateRawBlock"
}

func (_ *CreateRawBlock) Version() int {
	return 1
}

func (cmd *CreateRawBlock) Exec() (event.Event, error) {
	evt := usecase_event.NewRawBlockCreated(cmd.rawBlock)
	return evt, nil
}
