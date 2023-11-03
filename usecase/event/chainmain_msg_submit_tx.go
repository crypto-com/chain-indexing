package event

import (
	"bytes"

	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const CHAINMAIN_MSG_SUBMIT_TX = "/chainmain.icaauth.v1.MsgSubmitTx"
const CHAINMAIN_MSG_SUBMIT_TX_CREATED = "/chainmain.icaauth.v1.MsgSubmitTx.Created"
const CHAINMAIN_MSG_SUBMIT_TX_FAILED = "/chainmain.icaauth.v1.MsgSubmitTx.Failed"

type ChainmainMsgSubmitTx struct {
	MsgBase

	Params icaauthmodel.MsgSubmitTxParams `json:"params"`
}

func NewChainmainMsgSubmitTx(
	msgCommonParams MsgCommonParams,
	params icaauthmodel.MsgSubmitTxParams,
) *ChainmainMsgSubmitTx {
	return &ChainmainMsgSubmitTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         CHAINMAIN_MSG_SUBMIT_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *ChainmainMsgSubmitTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ChainmainMsgSubmitTx) String() string {
	return render.Render(event)
}

func DecodeChainmainMsgSubmitTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ChainmainMsgSubmitTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
