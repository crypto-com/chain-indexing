package model

type MsgCreateVestingAccountParams struct {
	RawMsgCreateVestingAccount
}

type RawMsgCreateVestingAccount struct {
	Type        string                          `mapstructure:"@type" json:"@type"`
	FromAddress string                          `mapstructure:"from_address" json:"fromAddress"`
	ToAddress   string                          `mapstructure:"to_Address" json:"toAddress"`
	Amount      []MsgCreateVestingAccountAmount `mapstructure:"amount" json:"amount"`
	EndTime     int64                           `mapstructure:"end_time" json:"endTime"`
	Delayed     bool                            `mapstructure:"delayed" json:"delayed"`
}

type MsgCreateVestingAccountAmount struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}
