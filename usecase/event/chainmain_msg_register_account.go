package event

import (
	"bytes"

	icaauthmodel "github.com/crypto-com/chain-indexing/usecase/model/icaauth"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const CHAINMAIN_MSG_REGISTER_ACCOUNT = "/chainmain.icaauth.v1.MsgRegisterAccount"
const CHAINMAIN_MSG_REGISTER_ACCOUNT_CREATED = "/chainmain.icaauth.v1.MsgRegisterAccount.Created"
const CHAINMAIN_MSG_REGISTER_ACCOUNT_FAILED = "/chainmain.icaauth.v1.MsgRegisterAccount.Failed"

type ChainmainMsgRegisterAccount struct {
	MsgBase

	Params icaauthmodel.ChainmainMsgRegisterAccountParams `json:"params"`
}

func NewChainmainMsgRegisterAccount(
	msgCommonParams MsgCommonParams,
	params icaauthmodel.ChainmainMsgRegisterAccountParams,
) *ChainmainMsgRegisterAccount {
	return &ChainmainMsgRegisterAccount{
		NewMsgBase(MsgBaseParams{
			MsgName:         CHAINMAIN_MSG_REGISTER_ACCOUNT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *ChainmainMsgRegisterAccount) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *ChainmainMsgRegisterAccount) String() string {
	return render.Render(event)
}

func DecodeChainmainMsgRegisterAccount(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *ChainmainMsgRegisterAccount
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
