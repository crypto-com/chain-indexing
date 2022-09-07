package model

type Tx struct {
	Tx         CosmosTx   `json:"tx"`
	TxResponse TxResponse `json:"tx_response"`
}

type CosmosTx struct {
	Body       CosmosTxBody     `json:"body"`
	AuthInfo   CosmosTxAuthInfo `json:"auth_info"`
	Signatures []string         `json:"signatures"`
}

type CosmosTxBody struct {
	Messages                    []map[string]interface{} `json:"messages"`
	Memo                        string                   `json:"memo"`
	TimeoutHeight               string                   `json:"timeout_height"`
	ExtensionOptions            []interface{}            `json:"extension_options"`
	NonCriticalExtensionOptions []interface{}            `json:"non_critical_extension_options"`
}

type CosmosTxAuthInfo struct {
	SignerInfos []CosmosTxSignerInfo `json:"signer_infos"`
	Fee         CosmosTxAuthInfoFee  `json:"fee"`
}

type CosmosTxAuthInfoFee struct {
	Amount   []CosmosTxAuthInfoFeeAmount `json:"amount"`
	GasLimit string                      `json:"gas_limit"`
	Payer    string                      `json:"payer"`
	Granter  string                      `json:"granter"`
}

type CosmosTxAuthInfoFeeAmount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type CosmosTxSignerInfo struct {
	MaybePublicKey *CosmosTxSignerInfoPublicKey `json:"public_key"`
	ModeInfo       CosmosTxSignerInfoModeInfo   `json:"mode_info"`
	Sequence       string                       `json:"sequence"`
}

type CosmosTxSignerInfoPublicKey struct {
	Type            string                                 `json:"@type"`
	MaybeThreshold  *int                                   `json:"threshold,omitempty"`
	MaybePublicKeys []CosmosTxSignerInfoPublicKeyPublicKey `json:"public_keys,omitempty"`
	MaybeKey        *string                                `json:"key,omitempty"`
}

type CosmosTxSignerInfoPublicKeyPublicKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type CosmosTxSignerInfoModeInfo struct {
	MaybeSingle *CosmosTxSignerInfoModeInfoSingle `json:"single,omitempty"`
	MaybeMulti  *CosmosTxSignerInfoModeInfoMulti  `json:"multi,omitempty"`
}

type CosmosTxSignerInfoModeInfoSingle struct {
	Mode string `json:"mode"`
}

type CosmosTxSignerInfoModeInfoMulti struct {
	Bitarray  CosmosTxSignerInfoModeInfoMultiBitarray         `json:"bitarray"`
	ModeInfos []CosmosTxSignerInfoModeInfoMultiSingleModeInfo `json:"mode_infos"`
}

type CosmosTxSignerInfoModeInfoMultiSingleModeInfo struct {
	Single CosmosTxSignerInfoModeInfoSingle `json:"single"`
}

type CosmosTxSignerInfoModeInfoMultiBitarray struct {
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
