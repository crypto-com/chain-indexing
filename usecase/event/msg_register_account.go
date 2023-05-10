package event

import (
	"bytes"

	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_REGISTER_ACCOUNT = "/chainmain.icaauth.v1.MsgRegisterAccount"
const MSG_REGISTER_ACCOUNT_CREATED = "/chainmain.icaauth.v1.MsgRegisterAccount.Created"
const MSG_REGISTER_ACCOUNT_FAILED = "/chainmain.icaauth.v1.MsgRegisterAccount.Failed"

type MsgRegisterAccount struct {
	MsgBase

	Params icaauthmodel.MsgRegisterAccountParams `json:"params"`
}

func NewMsgRegisterAccount(
	msgCommonParams MsgCommonParams,
	params icaauthmodel.MsgRegisterAccountParams,
) *MsgRegisterAccount {
	return &MsgRegisterAccount{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_REGISTER_ACCOUNT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgRegisterAccount) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgRegisterAccount) String() string {
	return render.Render(event)
}

func DecodeMsgRegisterAccount(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgRegisterAccount
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
