package model

type MsgNFTTransferNFTParams struct {
	TokenId   string `json:"tokenId"`
	DenomId   string `json:"denomId"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}
