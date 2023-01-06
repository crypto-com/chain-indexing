package event

import (
	"bytes"

	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SUBMIT_TX = "/chainmain.icaauth.v1.MsgSubmitTx"
const MSG_SUBMIT_TX_CREATED = "/chainmain.icaauth.v1.MsgSubmitTx.Created"
const MSG_SUBMIT_TX_FAILED = "/chainmain.icaauth.v1.MsgSubmitTx.Failed"

type MsgSubmitTx struct {
	MsgBase

	Params icaauthmodel.MsgSubmitTxParams `json:"params"`
}

func NewMsgSubmitTx(
	msgCommonParams MsgCommonParams,
	params icaauthmodel.MsgSubmitTxParams,
) *MsgSubmitTx {
	return &MsgSubmitTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgSubmitTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSubmitTx) String() string {
	return render.Render(event)
}

func DecodeMsgSubmitTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSubmitTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
