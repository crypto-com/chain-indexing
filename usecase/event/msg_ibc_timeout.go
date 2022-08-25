package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_TIMEOUT = "/ibc.core.channel.v1.MsgTimeout"
const MSG_IBC_TIMEOUT_CREATED = "MsgTimeoutCreated"
const MSG_IBC_TIMEOUT_FAILED = "MsgTimeoutFailed"

type MsgIBCTimeout struct {
	MsgBase

	Params ibc_model.MsgTimeoutParams `json:"params"`
}

func NewMsgIBCTimeout(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgTimeoutParams,
) *MsgIBCTimeout {
	return &MsgIBCTimeout{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_TIMEOUT,
			MsgSuccess:      MSG_IBC_TIMEOUT_CREATED,
			MsgFailed:       MSG_IBC_TIMEOUT_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCTimeout) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCTimeout) String() string {
	return render.Render(event)
}

func DecodeMsgIBCTimeout(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCTimeout
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}

const MSG_ALREADY_RELAYED_IBC_TIMEOUT = "MsgAlreadyRelayedTimeout"
const MSG_ALREADY_RELAYED_IBC_TIMEOUT_CREATED = "MsgAlreadyRelayedTimeoutCreated"
const MSG_ALREADY_RELAYED_IBC_TIMEOUT_FAILED = "MsgAlreadyRelayedTimeoutFailed"

type MsgAlreadyRelayedIBCTimeout struct {
	MsgBase

	Params ibc_model.MsgTimeoutParams `json:"params"`
}

func NewMsgAlreadyRelayedIBCTimeout(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgTimeoutParams,
) *MsgAlreadyRelayedIBCTimeout {
	return &MsgAlreadyRelayedIBCTimeout{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_ALREADY_RELAYED_IBC_TIMEOUT,
			MsgSuccess:      MSG_ALREADY_RELAYED_IBC_TIMEOUT_CREATED,
			MsgFailed:       MSG_ALREADY_RELAYED_IBC_TIMEOUT_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgAlreadyRelayedIBCTimeout) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgAlreadyRelayedIBCTimeout) String() string {
	return render.Render(event)
}

func DecodeMsgAlreadyRelayedIBCTimeout(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgAlreadyRelayedIBCTimeout
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
