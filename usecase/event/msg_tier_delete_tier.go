package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_DELETE_TIER = "/chainmain.tieredrewards.v1.MsgDeleteTier"
const MSG_TIER_DELETE_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgDeleteTier.Created"
const MSG_TIER_DELETE_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgDeleteTier.Failed"

type MsgTierDeleteTier struct {
	MsgBase
	Authority string `json:"authority"`
	Id        uint32 `json:"id"`
}

func NewMsgTierDeleteTier(msgCommonParams MsgCommonParams, params model.MsgDeleteTierParams) *MsgTierDeleteTier {
	return &MsgTierDeleteTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_DELETE_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Authority,
		params.Id,
	}
}

func (event *MsgTierDeleteTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierDeleteTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierDeleteTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierDeleteTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
