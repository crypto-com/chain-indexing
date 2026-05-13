package model

type MsgExitTierWithDelegationParams struct {
	RawMsgExitTierWithDelegation

	TransferredAmount string `json:"transferredAmount"`
	TransferredShares string `json:"transferredShares"`
	FullExit          bool   `json:"fullExit"`
}

type RawMsgExitTierWithDelegation struct {
	Type       string `mapstructure:"@type" json:"@type"`
	Owner      string `mapstructure:"owner" json:"owner"`
	PositionId string `mapstructure:"position_id" json:"positionId"`
	Amount     string `mapstructure:"amount" json:"amount"`
}
