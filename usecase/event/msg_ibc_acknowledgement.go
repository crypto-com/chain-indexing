package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_ACKNOWLEDGEMENT = "/ibc.core.channel.v1.MsgAcknowledgement"
const MSG_IBC_ACKNOWLEDGEMENT_CREATED = "MsgAcknowledgementCreated"
const MSG_IBC_ACKNOWLEDGEMENT_FAILED = "MsgAcknowledgementFailed"

type MsgIBCAcknowledgement struct {
	MsgBase

	Params ibc_model.MsgAcknowledgementParams `json:"params"`
}

func NewMsgIBCAcknowledgement(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgAcknowledgementParams,
) *MsgIBCAcknowledgement {
	return &MsgIBCAcknowledgement{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_ACKNOWLEDGEMENT,
			MsgSuccess:      MSG_IBC_ACKNOWLEDGEMENT_CREATED,
			MsgFailed:       MSG_IBC_ACKNOWLEDGEMENT_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCAcknowledgement) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCAcknowledgement) String() string {
	return render.Render(event)
}

func DecodeMsgIBCAcknowledgement(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCAcknowledgement
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}

const MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT = "MsgAlreadyRelayedAcknowledgement"
const MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_CREATED = "MsgAlreadyRelayedAcknowledgementCreated"
const MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_FAILED = "MsgAlreadyRelayedAcknowledgementFailed"

type MsgAlreadyRelayedIBCAcknowledgement struct {
	MsgBase

	Params ibc_model.MsgAlreadyRelayedAcknowledgementParams `json:"params"`
}

func NewMsgAlreadyRelayedIBCAcknowledgement(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgAlreadyRelayedAcknowledgementParams,
) *MsgAlreadyRelayedIBCAcknowledgement {
	return &MsgAlreadyRelayedIBCAcknowledgement{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT,
			MsgSuccess:      MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_CREATED,
			MsgFailed:       MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgAlreadyRelayedIBCAcknowledgement) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgAlreadyRelayedIBCAcknowledgement) String() string {
	return render.Render(event)
}

func DecodeMsgAlreadyRelayedIBCAcknowledgement(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgAlreadyRelayedIBCAcknowledgement
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
