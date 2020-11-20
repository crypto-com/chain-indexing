package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

type CreateMsgEditValidator struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgEditValidatorParams
}

func NewCreateMsgEditValidator(msgCommonParams event.MsgCommonParams, params model.MsgEditValidatorParams) *CreateMsgEditValidator {
	return &CreateMsgEditValidator{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgEditValidator) Name() string {
	return "CreateMsgEditValidator"
}

func (*CreateMsgEditValidator) Version() int {
	return 1
}

func (cmd *CreateMsgEditValidator) Exec() (entity_event.Event, error) {
	event := event.NewMsgEditValidator(cmd.msgCommonParams, cmd.params)
	return event, nil
}
