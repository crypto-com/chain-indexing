package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_GRANT = "MsgGrant"
const MSG_GRANT_CREATED = "MsgGrantCreated"
const MSG_GRANT_FAILED = "MsgGrantFailed"

type MsgGrant struct {
	MsgBase

	Params ibc_model.MsgGrantParams `json:"params"`
}

func NewMsgGrant(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgGrantParams,
) *MsgGrant {
	return &MsgGrant{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_GRANT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgGrant) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgGrant) String() string {
	return render.Render(event)
}

func DecodeMsgGrant(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgGrant
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
