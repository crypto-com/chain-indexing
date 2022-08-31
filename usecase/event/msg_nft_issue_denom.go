package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_ISSUE_DENOM = "/chainmain.nft.v1.MsgIssueDenom"
const MSG_NFT_ISSUE_DENOM_CREATED = "/chainmain.nft.v1.MsgIssueDenom.Created"
const MSG_NFT_ISSUE_DENOM_FAILED = "/chainmain.nft.v1.MsgIssueDenom.Failed"

type MsgNFTIssueDenom struct {
	MsgBase

	DenomId   string `json:"denomId"`
	DenomName string `json:"denomName"`
	Schema    string `json:"schema"`
	Sender    string `json:"sender"`
}

func NewMsgNFTIssueDenom(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTIssueDenomParams,
) *MsgNFTIssueDenom {
	return &MsgNFTIssueDenom{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_ISSUE_DENOM,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.DenomName,
		params.Schema,
		params.Sender,
	}
}

func (event *MsgNFTIssueDenom) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTIssueDenom) String() string {
	return render.Render(event)
}

func DecodeMsgNFTIssueDenom(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTIssueDenom
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
