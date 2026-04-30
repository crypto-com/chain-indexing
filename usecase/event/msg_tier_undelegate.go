package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_UNDELEGATE = "/chainmain.tieredrewards.v1.MsgTierUndelegate"
const MSG_TIER_UNDELEGATE_CREATED = "/chainmain.tieredrewards.v1.MsgTierUndelegate.Created"
const MSG_TIER_UNDELEGATE_FAILED = "/chainmain.tieredrewards.v1.MsgTierUndelegate.Failed"

type MsgTierUndelegate struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"position_id"`
}

func NewMsgTierUndelegate(msgCommonParams MsgCommonParams, params model.MsgTierUndelegateParams) *MsgTierUndelegate {
	return &MsgTierUndelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_UNDELEGATE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
	}
}

func (event *MsgTierUndelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierUndelegate) String() string {
	return render.Render(event)
}

func DecodeMsgTierUndelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierUndelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
