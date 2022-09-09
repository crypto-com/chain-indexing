package command

import (
	"fmt"

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
		RawData:    cmd.params.RawData,
	}
	event := event.NewRawBlockEvent(cmd.blockHeight, &params)
	fmt.Println("===> exec:", event)
	return event, nil
}
