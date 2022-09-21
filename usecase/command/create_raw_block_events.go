package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateBlockRawEvent struct {
	blockHeight int64
	params      usecase_model.CreateBlockRawEventParams
}

func NewCreateBlockRawEvent(blockHeight int64, params usecase_model.CreateBlockRawEventParams) *CreateBlockRawEvent {
	return &CreateBlockRawEvent{
		blockHeight,
		params,
	}
}

func (_ *CreateBlockRawEvent) Name() string {
	return "CreateBlockRawEvent"
}

func (_ *CreateBlockRawEvent) Version() int {
	return 1
}

func (cmd *CreateBlockRawEvent) Exec() (entity_event.Event, error) {
	params := usecase_model.CreateBlockRawEventParams{
		BlockHash:  cmd.params.BlockHash,
		BlockTime:  cmd.params.BlockTime,
		FromResult: cmd.params.FromResult,
		Data: usecase_model.DataParams{
			Type:    cmd.params.Data.Type,
			Content: cmd.params.Data.Content,
		},
	}
	event := event.NewBlockRawEventCreated(cmd.blockHeight, &params)
	return event, nil
}
