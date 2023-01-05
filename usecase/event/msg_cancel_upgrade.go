package event

import (
	"bytes"

	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_CANCEL_UPGRADE = "/cosmos.upgrade.v1beta1.MsgCancelUpgrade"
const MSG_CANCEL_UPGRADE_CREATED = "/cosmos.upgrade.v1beta1.MsgCancelUpgrade.Created"
const MSG_CANCEL_UPGRADE_FAILED = "/cosmos.upgrade.v1beta1.MsgCancelUpgrade.Failed"

type MsgCancelUpgrade struct {
	MsgBase

	v1_model.MsgCancelUpgradeParams
}

func NewMsgCancelUpgrade(
	msgCommonParams MsgCommonParams,
	params v1_model.MsgCancelUpgradeParams,
) *MsgCancelUpgrade {
	return &MsgCancelUpgrade{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgCancelUpgrade) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgCancelUpgrade) String() string {
	return render.Render(event)
}

func DecodeMsgCancelUpgrade(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgCancelUpgrade
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
