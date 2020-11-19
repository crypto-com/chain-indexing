package command

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

// CreateMsgDelegate is a command to create MsgDelegate event
type CreateMsgDelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgDelegateParams
}

// NewCreateMsgDelegate create a new instance of CreateMsgDelegate command
func NewCreateMsgDelegate(msgCommonParams event.MsgCommonParams, params model.MsgDelegateParams) *CreateMsgDelegate {
	return &CreateMsgDelegate{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgDelegate) Name() string {
	return "CreateMsgDelegate"
}

// Version returns version of command
func (*CreateMsgDelegate) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgDelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgDelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
