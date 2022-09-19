package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

type MsgExtensionOptionDynamicFeeTx struct {
	MsgBase

	Params model.MsgDynamicFeeTxParams `json:"params"`
}

func NewMsgExtensionOptionDynamicFeeTx(
	msgCommonParams MsgCommonParams,
	params model.MsgDynamicFeeTxParams,
) *MsgExtensionOptionDynamicFeeTx {
	return &MsgExtensionOptionDynamicFeeTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_ETHEREUM_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgExtensionOptionDynamicFeeTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgExtensionOptionDynamicFeeTx) String() string {
	return render.Render(event)
}

func DecodeMsgExtensionOptionDynamicFeeTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgExtensionOptionDynamicFeeTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
