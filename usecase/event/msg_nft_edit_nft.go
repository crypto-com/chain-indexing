package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_EDIT_NFT = "/chainmain.nft.v1.MsgEditNFT"
const MSG_NFT_EDIT_NFT_CREATED = "/chainmain.nft.v1.MsgEditNFT.Created"
const MSG_NFT_EDIT_NFT_FAILED = "/chainmain.nft.v1.MsgEditNFT.Failed"

type MsgNFTEditNFT struct {
	MsgBase

	DenomId   string `json:"denomId"`
	TokenId   string `json:"tokenId"`
	TokenName string `json:"tokenName"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Sender    string `json:"sender"`
}

func NewMsgNFTEditNFT(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTEditNFTParams,
) *MsgNFTEditNFT {
	return &MsgNFTEditNFT{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_EDIT_NFT,
			Version:         1,

			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.TokenId,
		params.TokenName,
		params.URI,
		params.Data,
		params.Sender,
	}
}

func (event *MsgNFTEditNFT) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTEditNFT) String() string {
	return render.Render(event)
}

func DecodeMsgNFTEditNFT(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTEditNFT
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
