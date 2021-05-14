package model

type MsgNFTMintNFTParams struct {
	DenomId   string `json:"denomId"`
	TokenId   string `json:"tokenId"`
	TokenName string `json:"tokenName"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}
