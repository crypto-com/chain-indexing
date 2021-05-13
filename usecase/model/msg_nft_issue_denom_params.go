package model

type MsgNFTIssueDenomParams struct {
	DenomId   string `json:"denomId"`
	DenomName string `json:"denomName"`
	Schema    string `json:"schema"`
	Sender    string `json:"sender"`
}
