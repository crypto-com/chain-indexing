package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_TRANSFER_TRANSFER = "MsgTransfer"
const MSG_IBC_TRANSFER_TRANSFER_CREATED = "MsgTransferCreated"
const MSG_IBC_TRANSFER_TRANSFER_FAILED = "MsgTransferFailed"

type MsgIBCTransferTransfer struct {
	MsgBase

	Params ibc_model.MsgTransferParams `json:"params"`
}

func NewMsgIBCTransferTransfer(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgTransferParams,
) *MsgIBCTransferTransfer {
	return &MsgIBCTransferTransfer{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_TRANSFER_TRANSFER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCTransferTransfer) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCTransferTransfer) String() string {
	return render.Render(event)
}

func DecodeMsgIBCTransferTransfer(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCTransferTransfer
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
