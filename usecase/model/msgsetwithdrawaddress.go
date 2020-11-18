package model

type MsgSetWithdrawAddressParams struct {
	DelegatorAddress string `json:"delegatorAddress"`
	WithdrawAddress  string `json:"withdrawAddress"`
}
