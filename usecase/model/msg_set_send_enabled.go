package model

type MsgSetSendEnabled struct {
	Authority     string        `json:"authority"`
	SendEnabled   []SendEnabled `json:"send_enabled"`
	UseDefaultFor []string      `json:"use_default_for"`
}

type SendEnabled struct {
	Denom   string `json:"denom"`
	Enabled bool   `json:"enabled"`
}
