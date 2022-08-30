package model

type Tx struct {
	Tx         CosmosTx   `json:"tx"`
	TxResponse TxResponse `json:"tx_response"`
}

type CosmosTx struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type Body struct {
	Messages                    []map[string]interface{} `json:"messages"`
	Memo                        string                   `json:"memo"`
	TimeoutHeight               string                   `json:"timeout_height"`
	ExtensionOptions            []interface{}            `json:"extension_options"`
	NonCriticalExtensionOptions []interface{}            `json:"non_critical_extension_options"`
}

type AuthInfo struct {
	SignerInfos []SignerInfo `json:"signer_infos"`
	Fee         Fee          `json:"fee"`
}

type Fee struct {
	Amount   []Amount `json:"amount"`
	GasLimit string   `json:"gas_limit"`
	Payer    string   `json:"payer"`
	Granter  string   `json:"granter"`
}

type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type SignerInfo struct {
	MaybePublicKey *SignerInfoPublicKey `json:"public_key"`
	ModeInfo       ModeInfo             `json:"mode_info"`
	Sequence       string               `json:"sequence"`
}

type SignerInfoPublicKey struct {
	Type            string      `json:"@type"`
	MaybeThreshold  *int        `json:"threshold,omitempty"`
	MaybePublicKeys []PublicKey `json:"public_keys,omitempty"`
	MaybeKey        *string     `json:"key,omitempty"`
}

type PublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type ModeInfo struct {
	MaybeSingle *Single `json:"single,omitempty"`
	MaybeMulti  *Multi  `json:"multi,omitempty"`
}

type Single struct {
	Mode string `json:"mode"`
}

type Multi struct {
	Bitarray  Bitarray         `json:"bitarray"`
	ModeInfos []SingleModeInfo `json:"mode_infos"`
}

type SingleModeInfo struct {
	Single Single `json:"single"`
}

type Bitarray struct {
	ExtraBitsStored int64  `json:"extra_bits_stored"`
	Elems           string `json:"elems"`
}

type TxResponse struct {
	Height    int64           `json:"height,omitempty"`
	TxHash    string          `json:"txhash,omitempty"`
	Codespace string          `json:"codespace,omitempty"`
	Code      uint32          `json:"code,omitempty"`
	Data      string          `json:"data,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
	Logs      []TxResponseLog `json:"logs"`
	Info      string          `json:"info,omitempty"`
	GasWanted int64           `json:"gas_wanted,omitempty"`
	GasUsed   int64           `json:"gas_used,omitempty"`
	Tx        TxResponseTx    `json:"tx,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`
	Events    []BlockEvent    `json:"events"`
}

type TxResponseTx struct {
	Type string `json:"@type"`
	CosmosTx
}

type TxResponseLog struct {
	MsgIndex uint32       `json:"msg_index,omitempty"`
	Log      string       `json:"log,omitempty"`
	Events   []BlockEvent `json:"events"`
}
