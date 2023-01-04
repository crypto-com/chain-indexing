package event

import (
	"bytes"

	model "github.com/crypto-com/chain-indexing/usecase/model/v1"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_SOFTWARE_UPGRADE = "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade"
const MSG_SOFTWARE_UPGRADE_CREATED = "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade.Created"
const MSG_SOFTWARE_UPGRADE_FAILED = "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade.Failed"

type MsgSoftwareUpgrade struct {
	MsgBase

	model.MsgSoftwareUpgradeParams
}

func NewMsgSoftwareUpgrade(
	msgCommonParams MsgCommonParams,
	params model.MsgSoftwareUpgradeParams,
) *MsgSoftwareUpgrade {
	return &MsgSoftwareUpgrade{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_SOFTWARE_UPGRADE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

func (event *MsgSoftwareUpgrade) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgSoftwareUpgrade) String() string {
	return render.Render(event)
}

func DecodeMsgSoftwareUpgrade(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgSoftwareUpgrade
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
