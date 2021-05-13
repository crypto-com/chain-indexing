package model

type MsgNFTBurnNFTParams struct {
	DenomId string `json:"denomId"`
	TokenId string `json:"tokenId"`
	Sender  string `json:"sender"`
}
