package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

// CreateMsgUndelegate is a command to create MsgUndelegate event
type CreateMsgUndelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgUndelegateParams
}

// NewCreateMsgUndelegate create a new instance of CreateMsgUndelegate command
func NewCreateMsgUndelegate(msgCommonParams event.MsgCommonParams, params model.MsgUndelegateParams) *CreateMsgUndelegate {
	return &CreateMsgUndelegate{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgUndelegate) Name() string {
	return "CreateMsgUndelegate"
}

// Version returns version of command
func (*CreateMsgUndelegate) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgUndelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgUndelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
