package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CORE_RECV_PACKET = "MsgRecvPacket"
const MSG_IBC_CORE_RECV_PACKET_CREATED = "MsgRecvPacketCreated"
const MSG_IBC_CORE_RECV_PACKET_FAILED = "MsgRecvPacketFailed"

type MsgIBCCoreRecvPacket struct {
	MsgBase

	Params ibc_model.MsgRecvPacketParams `json:"params"`
}

func NewMsgIBCCoreRecvPacket(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgRecvPacketParams,
) *MsgIBCCoreRecvPacket {
	return &MsgIBCCoreRecvPacket{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CORE_RECV_PACKET,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCCoreRecvPacket) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCCoreRecvPacket) String() string {
	return render.Render(event)
}

func DecodeMsgIBCCoreRecvPacket(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCCoreRecvPacket
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
