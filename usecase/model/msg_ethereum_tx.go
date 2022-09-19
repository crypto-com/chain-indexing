package model

type MsgEthereumTxParams struct {
	RawMsgEthereumTx
}

type RawMsgEthereumTx struct {
	Type string `mapstructure:"@type" json:"@type"`
	Size int    `mapstructure:"size" json:"size"`
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
	Data LegacyTx `mapstructure:"data" json:"data"`
	From string   `mapstructure:"from" json:"from"`
	Hash string   `mapstructure:"hash" json:"hash"`
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

type MsgDynamicFeeTxParams struct {
	RawDynamicFeeTx
}
type RawDynamicFeeTx struct {
	Type string `mapstructure:"@type" json:"@type"`
	Size int    `mapstructure:"size" json:"size"`
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/730
	Data DynamicFeeTx `mapstructure:"data" json:"data"`
	From string       `mapstructure:"from" json:"from"`
	Hash string       `mapstructure:"hash" json:"hash"`
}
type DynamicFeeTx struct {
	Type      string `mapstructure:"@type" json:"@type"`
	ChainId   string `mapstructure:"chain_id" json:"chainId"`
	Nonce     string `mapstructure:"nonce" json:"nonce"`
	GasTipCap string `mapstructure:"gas_tip_cap" json:"gasTipCap"`
	GasFeeCap string `mapstructure:"Gas_fee_cap" json:"gasFeeCap"`
	Gas       string `mapstructure:"gas" json:"gas"`
	To        string `mapstructure:"to" json:"to"`
	Value     string `mapstructure:"value" json:"value"`
	Data      string `mapstructure:"data" json:"data"`
	V         string `mapstructure:"v" json:"v"`
	R         string `mapstructure:"r" json:"r"`
	S         string `mapstructure:"s" json:"s"`
}
