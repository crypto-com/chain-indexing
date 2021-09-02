package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_IBC_CREATE_VESTING_ACCOUNT = "MsgCreateVestingAccount"
const MSG_IBC_CREATE_VESTING_ACCOUNT_CREATED = "MsgCreateVestingAccountCreated"
const MSG_IBC_CREATE_VESTING_ACCOUNT_FAILED = "MsgCreateVestingAccountFailed"

type MsgIBCCreateVestingAccount struct {
	MsgBase

	Params ibc_model.MsgCreateVestingAccountParams `json:"params"`
}

func NewMsgIBCCreateVestingAccount(
	msgCommonParams MsgCommonParams,
	params ibc_model.MsgCreateVestingAccountParams,
) *MsgIBCCreateVestingAccount {
	return &MsgIBCCreateVestingAccount{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_IBC_CREATE_VESTING_ACCOUNT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgIBCCreateVestingAccount) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgIBCCreateVestingAccount) String() string {
	return render.Render(event)
}

func DecodeMsgIBCCreateVestingAccount(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgIBCCreateVestingAccount
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
