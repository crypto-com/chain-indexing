package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_EXEC = "/cosmos.authz.v1beta1.MsgExec"
const MSG_EXEC_CREATED = "/cosmos.authz.v1beta1.MsgExec.Created"
const MSG_EXEC_FAILED = "/cosmos.authz.v1beta1.MsgExec.Failed"

type MsgExec struct {
	MsgBase

	Params model.MsgExecParams `json:"params"`
}

func NewMsgExec(
	msgCommonParams MsgCommonParams,
	params model.MsgExecParams,
) *MsgExec {
	return &MsgExec{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_EXEC,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgExec) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgExec) String() string {
	return render.Render(event)
}

func DecodeMsgExec(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgExec
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
