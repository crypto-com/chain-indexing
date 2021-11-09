package model

type MsgGrantParams struct {
	MaybeSendGrant    *RawMsgSendGrant    `json:"maybeSendGrant"`
	MaybeStakeGrant   *RawMsgStakeGrant   `json:"maybeStakeGrant"`
	MaybeGenericGrant *RawMsgGenericGrant `json:"maybeGenericGrant"`
}

type RawMsgSendGrant struct {
	Type    string    `mapstructure:"@type" json:"@type"`
	Granter string    `mapstructure:"granter" json:"granter"`
	Grantee string    `mapstructure:"grantee" json:"grantee"`
	Grant   SendGrant `mapstructure:"grant" json:"grant"`
}

type RawMsgStakeGrant struct {
	Type    string     `mapstructure:"@type" json:"@type"`
	Granter string     `mapstructure:"granter" json:"granter"`
	Grantee string     `mapstructure:"grantee" json:"grantee"`
	Grant   StakeGrant `mapstructure:"grant" json:"grant"`
}

type RawMsgGenericGrant struct {
	Type    string       `mapstructure:"@type" json:"@type"`
	Granter string       `mapstructure:"granter" json:"granter"`
	Grantee string       `mapstructure:"grantee" json:"grantee"`
	Grant   GenericGrant `mapstructure:"grant" json:"grant"`
}

type SendGrant struct {
	Authorization SendAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string            `mapstructure:"expiration" json:"expiration"`
}

type StakeGrant struct {
	Authorization StakeAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string             `mapstructure:"expiration" json:"expiration"`
}

type GenericGrant struct {
	Authorization GenericAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string               `mapstructure:"expiration" json:"expiration"`
}

type SendAuthorization struct {
	SpendLimit []MsgGrantSpendLimit `mapstructure:"spend_limit" json:"spendLimit"`
}

type MsgGrantSpendLimit struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type StakeAuthorization struct {
	MaxTokens         MsgGrantMaxTokens `mapstructure:"max_tokens" json:"maxTokens,omitempty"`
	MaybeAllowList    *Validators       `mapstructure:"allow_list" json:"allowList,omitempty"`
	MaybeDenyList     *Validators       `mapstructure:"deny_list" json:"denyList,omitempty"`
	AuthorizationType string            `mapstructure:"authorization_type" json:"authorizationType,omitempty"`
}

type MsgGrantMaxTokens struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type Validators struct {
	Address []string `mapstructure:"address" json:"address,omitempty"`
}

type GenericAuthorization struct {
	Msg string `mapstructure:"msg" json:"msg,omitempty"`
}
