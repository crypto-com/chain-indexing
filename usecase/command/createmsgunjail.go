package command

import (
	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

// CreateMsgUnjail is a command to create MsgUnjail event
type CreateMsgUnjail struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgUnjailParams
}

// NewCreateMsgUnjail create a new instance of CreateMsgUnjail command
func NewCreateMsgUnjail(msgCommonParams event.MsgCommonParams, params model.MsgUnjailParams) *CreateMsgUnjail {
	return &CreateMsgUnjail{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgUnjail) Name() string {
	return "/cosmos.slashing.v1beta1.MsgUnjail.Create"
}

// Version returns version of command
func (*CreateMsgUnjail) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgUnjail) Exec() (entity_event.Event, error) {
	event := event.NewMsgUnjail(cmd.msgCommonParams, cmd.params)
	return event, nil
}
