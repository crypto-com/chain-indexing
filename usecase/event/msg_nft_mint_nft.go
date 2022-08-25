package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_NFT_MINT_NFT = "/chainmain.nft.v1.MsgMintNFT"
const MSG_NFT_MINT_NFT_CREATED = "MsgMintNFTCreated"
const MSG_NFT_MINT_NFT_FAILED = "MsgMintNFTFailed"

type MsgNFTMintNFT struct {
	MsgBase

	DenomId   string `json:"denomId"`
	TokenId   string `json:"tokenId"`
	TokenName string `json:"tokenName"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

func NewMsgNFTMintNFT(
	msgCommonParams MsgCommonParams,
	params model.MsgNFTMintNFTParams,
) *MsgNFTMintNFT {
	return &MsgNFTMintNFT{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_NFT_MINT_NFT,
			MsgSuccess:      MSG_NFT_MINT_NFT_CREATED,
			MsgFailed:       MSG_NFT_MINT_NFT_FAILED,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params.DenomId,
		params.TokenId,
		params.TokenName,
		params.URI,
		params.Data,
		params.Sender,
		params.Recipient,
	}
}

func (event *MsgNFTMintNFT) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgNFTMintNFT) String() string {
	return render.Render(event)
}

func DecodeMsgNFTMintNFT(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgNFTMintNFT
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
