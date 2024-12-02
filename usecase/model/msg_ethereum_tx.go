package model

type MsgEthereumTxParams struct {
	RawMsgEthereumTx
}

type RawMsgEthereumTx struct {
	Type string `mapstructure:"@type" json:"@type"`
	Size int    `mapstructure:"size" json:"size"`
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
	Data       *LegacyTx   `mapstructure:"data" json:"data"`
	From       string      `mapstructure:"from" json:"from"`
	Hash       string      `mapstructure:"hash" json:"hash"`
	Raw        string      `mapstructure:"raw" json:"raw"`
	DecodedRaw *DecodedRaw `mapstructure:"decodedRaw,omitempty" json:"decodedRaw,omitempty"`
}

type MsgEthereumTx struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
type LegacyTx struct {
	Type     string `mapstructure:"@type" json:"@type"`
	Nonce    string `mapstructure:"nonce" json:"nonce"`
	GasPrice string `mapstructure:"gas_price" json:"gasPrice"`
	Gas      string `mapstructure:"gas" json:"gas"`
	To       string `mapstructure:"to" json:"to"`
	Value    string `mapstructure:"value" json:"value"`
	Data     string `mapstructure:"data" json:"data"`
	V        string `mapstructure:"v" json:"v"`
	R        string `mapstructure:"r" json:"r"`
	S        string `mapstructure:"s" json:"s"`
}

type DecodedRaw struct {
	Type                 string `mapstructure:"type" json:"type"`
	Hash                 string `mapstructure:"hash" json:"hash"`
	Nonce                string `mapstructure:"nonce" json:"nonce"`
	GasLimit             string `mapstructure:"gasLimit" json:"gasLimit"`
	GasPrice             string `mapstructure:"gasPrice" json:"gasPrice"`
	MaxFeePerGas         string `mapstructure:"maxFeePerGas" json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `mapstructure:"maxPriorityFeePerGas" json:"maxPriorityFeePerGas"`
	From                 string `mapstructure:"from" json:"from"`
	To                   string `mapstructure:"to" json:"to"`
	PublicKey            string `mapstructure:"publicKey" json:"publicKey"`
	V                    string `mapstructure:"v" json:"v"`
	R                    string `mapstructure:"r" json:"r"`
	S                    string `mapstructure:"s" json:"s"`
	Value                string `mapstructure:"value" json:"value"`
	Input                string `mapstructure:"input" json:"input"`
	FunctionHash         string `mapstructure:"functionHash" json:"functionHash"`
}
