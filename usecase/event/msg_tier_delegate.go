package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_TIER_DELEGATE = "/chainmain.tieredrewards.v1.MsgTierDelegate"
const MSG_TIER_DELEGATE_CREATED = "/chainmain.tieredrewards.v1.MsgTierDelegate.Created"
const MSG_TIER_DELEGATE_FAILED = "/chainmain.tieredrewards.v1.MsgTierDelegate.Failed"

type MsgTierDelegate struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"positionId"`
	Validator  string `json:"validator"`
}

func NewMsgTierDelegate(msgCommonParams MsgCommonParams, params model.MsgTierDelegateParams) *MsgTierDelegate {
	return &MsgTierDelegate{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_DELEGATE,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
		params.Validator,
	}
}

func (event *MsgTierDelegate) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierDelegate) String() string {
	return render.Render(event)
}

func DecodeMsgTierDelegate(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierDelegate
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
