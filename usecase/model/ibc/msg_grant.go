package ibc

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
	Authorization MaybeSendAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string                 `mapstructure:"expiration" json:"expiration"`
}

type StakeGrant struct {
	Authorization MaybeStakeAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string                  `mapstructure:"expiration" json:"expiration"`
}

type GenericGrant struct {
	Authorization MaybeGenericAuthorization `mapstructure:"authorization" json:"authorization"`
	Expiration    string                    `mapstructure:"expiration" json:"expiration"`
}

type MsgGrantAuthorization struct {
	Type                      string                     `mapstructure:"@type" json:"@type"`
	MaybeSendAuthorization    *MaybeSendAuthorization    `json:"maybeSendAuthorization"`
	MaybeStakeAuthorization   *MaybeStakeAuthorization   `json:"maybeStakeAuthorization"`
	MaybeGenericAuthorization *MaybeGenericAuthorization `json:"maybeGenericAuthorization"`
}

type MaybeSendAuthorization struct {
	SpendLimit []MsgGrantSpendLimit `mapstructure:"spend_limit" json:"spendLimit"`
}

type MsgGrantSpendLimit struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type MaybeStakeAuthorization struct {
	MaxTokens         MsgGrantMaxTokens `mapstructure:"max_tokens" json:"maxTokens,omitempty"`
	AllowList         Validators        `mapstructure:"allow_list" json:"allowList,omitempty"`
	DenyList          Validators        `mapstructure:"deny_list" json:"denyList,omitempty"`
	AuthorizationType int32             `mapstructure:"authorization_type" json:"authorization_type,omitempty"`
}

type MsgGrantMaxTokens struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

type Validators struct {
	Address []string `mapstructure:"address" json:"address,omitempty"`
}

type MaybeGenericAuthorization struct {
	Msg string `mapstructure:"msg" json:"msg,omitempty"`
}
