package event

import (
	"bytes"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

const MSG_TIER_UPDATE_TIER = "/chainmain.tieredrewards.v1.MsgUpdateTier"
const MSG_TIER_UPDATE_TIER_CREATED = "/chainmain.tieredrewards.v1.MsgUpdateTier.Created"
const MSG_TIER_UPDATE_TIER_FAILED = "/chainmain.tieredrewards.v1.MsgUpdateTier.Failed"

type MsgTierUpdateTier struct {
	MsgBase
	Authority string `json:"authority"`
	Tier      string `json:"tier"`
}

func NewMsgTierUpdateTier(msgCommonParams MsgCommonParams, params model.MsgUpdateTierParams) *MsgTierUpdateTier {
	return &MsgTierUpdateTier{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_TIER_UPDATE_TIER,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.Authority,
		params.Tier,
	}
}

func (event *MsgTierUpdateTier) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func (event *MsgTierUpdateTier) String() string {
	return render.Render(event)
}

func DecodeMsgTierUpdateTier(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()
	var event *MsgTierUpdateTier
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
