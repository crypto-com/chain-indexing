package model

type MsgEthereumTxParams struct {
	RawMsgEthereumTx
}

type RawMsgEthereumTx struct {
	Type string   `mapstructure:"@type" json:"@type"`
	Size int      `mapstructure:"size" json:"size"`
	Data LegacyTx `mapstructure:"data" json:"data"`
	From string   `mapstructure:"from" json:"from"`
	Hash string   `mapstructure:"hash" json:"hash"`
}

type MsgEthereumTx struct {
	Denom  string `mapstructure:"denom" json:"denom,omitempty"`
	Amount string `mapstructure:"amount" json:"amount"`
}

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
