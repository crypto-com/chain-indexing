package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

// CreateMsgBeginRedelegate is a command to create MsgBeginRedelegate event
type CreateMsgBeginRedelegate struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgBeginRedelegateParams
}

// NewCreateMsgBeginRedelegate create a new instance of CreateMsgBeginRedelegate command
func NewCreateMsgBeginRedelegate(msgCommonParams event.MsgCommonParams, params model.MsgBeginRedelegateParams) *CreateMsgBeginRedelegate {
	return &CreateMsgBeginRedelegate{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgBeginRedelegate) Name() string {
	return "CreateMsgBeginRedelegate"
}

// Version returns version of command
func (*CreateMsgBeginRedelegate) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgBeginRedelegate) Exec() (entity_event.Event, error) {
	event := event.NewMsgBeginRedelegate(cmd.msgCommonParams, cmd.params)
	return event, nil
}
