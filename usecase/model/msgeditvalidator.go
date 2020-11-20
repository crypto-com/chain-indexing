package model

type MsgEditValidatorParams struct {
	ValidatorAddress  string `json:"validatorAddress"`
	CommissionRate    string `json:"commissionRate"`
	MinSelfDelegation string `json:"minSelfDelegation"`
}
