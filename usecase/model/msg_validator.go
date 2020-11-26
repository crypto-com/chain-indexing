package model

type MsgValidatorDescription struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"securityContact"`
	Details         string `json:"details"`
}

type MsgValidatorCommission struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"maxRate"`
	MaxChangeRate string `json:"maxChangeRate"`
}
