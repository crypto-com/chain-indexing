package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_CLEAR_POSITION = "/chainmain.tieredrewards.v1.MsgClearPosition"
const MSG_TIER_CLEAR_POSITION_CREATED = "/chainmain.tieredrewards.v1.MsgClearPosition.Created"
const MSG_TIER_CLEAR_POSITION_FAILED = "/chainmain.tieredrewards.v1.MsgClearPosition.Failed"

type MsgTierClearPosition struct {
	MsgBase
	Owner      string `json:"owner"`
	PositionId string `json:"position_id"`
}

func NewMsgTierClearPosition(msgCommonParams MsgCommonParams, params model.MsgClearPositionParams) *MsgTierClearPosition {
	return &MsgTierClearPosition{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_CLEAR_POSITION,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Owner,
		params.PositionId,
	}
}

func (event *MsgTierClearPosition) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierClearPosition) String() string {
	return render.Render(event)
}

func DecodeMsgTierClearPosition(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierClearPosition
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
