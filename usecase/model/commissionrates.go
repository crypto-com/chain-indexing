package model

type CommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"maxRate"`
	MaxChangeRate string `json:"maxChangeRate"`
}
