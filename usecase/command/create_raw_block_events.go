package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateRawBlockEvent struct {
	blockHeight int64
	params      usecase_model.CreateRawBlockEventParams
}

func NewCreateRawBlockEvent(blockHeight int64, params usecase_model.CreateRawBlockEventParams) *CreateRawBlockEvent {
	return &CreateRawBlockEvent{
		blockHeight,
		params,
	}
}

func (_ *CreateRawBlockEvent) Name() string {
	return "CreateRawBlockEvent"
}

func (_ *CreateRawBlockEvent) Version() int {
	return 1
}

func (cmd *CreateRawBlockEvent) Exec() (entity_event.Event, error) {
	params := usecase_model.CreateRawBlockEventParams{
		BlockHash:  cmd.params.BlockHash,
		BlockTime:  cmd.params.BlockTime,
		FromResult: cmd.params.FromResult,
		RawData: usecase_model.RawDataParams{
			Type:    cmd.params.RawData.Type,
			Content: cmd.params.RawData.Content,
		},
	}
	event := event.NewRawBlockEventCreated(cmd.blockHeight, &params)
	return event, nil
}
