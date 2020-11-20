package model

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"securityContact"`
	Details         string `json:"details"`
}
