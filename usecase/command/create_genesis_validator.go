package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"
)

type CreateGenesisValidator struct {
	params genesis.CreateGenesisValidatorParams
}

func NewCreateGenesisValidator(
	params genesis.CreateGenesisValidatorParams,
) *CreateGenesisValidator {
	return &CreateGenesisValidator{
		params,
	}
}

func (*CreateGenesisValidator) Name() string {
	return "CreateGenesisValidator"
}

func (*CreateGenesisValidator) Version() int {
	return 1
}

func (cmd *CreateGenesisValidator) Exec() (entity_event.Event, error) {
	event := event.NewCreateGenesisValidator(cmd.params)
	return event, nil
}
