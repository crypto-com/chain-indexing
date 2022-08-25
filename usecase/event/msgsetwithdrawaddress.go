package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SET_WITHDRAW_ADDRESS = "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress"
const MSG_SET_WITHDRAW_ADDRESS_CREATED = "MsgSetWithdrawAddressCreated"
const MSG_SET_WITHDRAW_ADDRESS_FAILED = "MsgSetWithdrawAddressFailed"

type MsgSetWithdrawAddress struct {
	MsgBase

	model.MsgSetWithdrawAddressParams
}

func NewMsgSetWithdrawAddress(
	msgCommonParams MsgCommonParams,
	params model.MsgSetWithdrawAddressParams,
) *MsgSetWithdrawAddress {
	return &MsgSetWithdrawAddress{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SET_WITHDRAW_ADDRESS,
			MsgSuccess:      MSG_SET_WITHDRAW_ADDRESS_CREATED,
			MsgFailed:       MSG_SET_WITHDRAW_ADDRESS_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSetWithdrawAddress) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSetWithdrawAddress) String() string {
	return render.Render(event)
}

func DecodeMsgSetWithdrawAddress(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSetWithdrawAddress
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
