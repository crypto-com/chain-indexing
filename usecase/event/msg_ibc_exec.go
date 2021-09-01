package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_EXEC = "MsgExec"
const MSG_IBC_EXEC_CREATED = "MsgExecCreated"
const MSG_IBC_EXEC_FAILED = "MsgExecFailed"

type MsgIBCExec struct {
	MsgBase

	Params ibc_model.MsgExecParams `json:"params"`
}

func NewMsgIBCExec(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgExecParams,
) *MsgIBCExec {
	return &MsgIBCExec{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_EXEC,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCExec) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCExec) String() string {
	return render.Render(event)
}

func DecodeMsgIBCExec(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCExec
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
