package model

type MsgEditValidatorParams struct {
	Description            ValidatorDescription `json:"description"`
	ValidatorAddress       string               `json:"validatorAddress"`
	MaybeCommissionRate    *string              `json:"commissionRate"`
	MaybeMinSelfDelegation *string              `json:"minSelfDelegation"`
}
