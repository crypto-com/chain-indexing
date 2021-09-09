package model

type MsgGrantAllowanceParams struct {
	MaybeBasicAllowance      *RawMsgGrantBasicAllowance      `json:"maybeBasicAllowance"`
	MaybePeriodicAllowance   *RawMsgGrantPeriodicAllowance   `json:"maybePeriodicAllowance"`
	MaybeAllowedMsgAllowance *RawMsgGrantAllowedMsgAllowance `json:"maybeAllowedMsgAllowance"`
}

type RawMsgGrantBasicAllowance struct {
	Type      string         `mapstructure:"@type" json:"@type"`
	Granter   string         `mapstructure:"granter" json:"granter"`
	Grantee   string         `mapstructure:"grantee" json:"grantee"`
	Allowance BasicAllowance `mapstructure:"allowance" json:"allowance"`
}

type RawMsgGrantPeriodicAllowance struct {
	Type      string            `mapstructure:"@type" json:"@type"`
	Granter   string            `mapstructure:"granter" json:"granter"`
	Grantee   string            `mapstructure:"grantee" json:"grantee"`
	Allowance PeriodicAllowance `mapstructure:"allowance" json:"allowance"`
}

type RawMsgGrantAllowedMsgAllowance struct {
	Type      string              `mapstructure:"@type" json:"@type"`
	Granter   string              `mapstructure:"granter" json:"granter"`
	Grantee   string              `mapstructure:"grantee" json:"grantee"`
	Allowance AllowedMsgAllowance `mapstructure:"allowance" json:"allowance"`
}

type Allowance struct {
	Type string `mapstructure:"@type" json:"@type"`
}

type BasicAllowance struct {
	Type       string                        `mapstructure:"@type" json:"@type"`
	SpendLimit []MsgGrantAllowanceSpendLimit `mapstructure:"spend_limit" json:"spendLimit"`
	Expiration string                        `mapstructure:"expiration" json:"expiration"`
}

type MsgGrantAllowanceSpendLimit struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type PeriodicAllowance struct {
	Type             string                              `mapstructure:"@type" json:"@type"`
	Basic            BasicAllowance                      `mapstructure:"basic" json:"basic"`
	Period           Duration                            `mapstructure:"period" json:"period"`
	PeriodSpendLimit []MsgGrantAllowancePeriodSpendLimit `mapstructure:"period_spend_limit" json:"periodSpendLimit"`
	PeriodCanSpend   []MsgGrantAllowancePeriodCanSpend   `mapstructure:"period_can_spend" json:"periodCanSpend"`
	PeriodReset      string                              `mapstructure:"period_reset" json:"periodReset"`
}

type MsgGrantAllowancePeriodSpendLimit struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type MsgGrantAllowancePeriodCanSpend struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type AllowedMsgAllowance struct {
	Type            string    `mapstructure:"@type" json:"@type"`
	Allowance       Allowance `mapstructure:"allowance" json:"allowance,omitempty"`
	AllowedMessages []string  `mapstructure:"allowed_messages" json:"allowedMessages,omitempty"`
}
