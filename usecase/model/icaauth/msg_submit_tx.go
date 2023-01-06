package icaauth

import (
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

type MsgSubmitTxParams struct {
	RawMsgSubmitTx

	ibc_model.Packet
}

type RawMsgSubmitTx struct {
	Owner           string           `mapstructure:"owner" json:"owner"`
	ConnectionId    string           `mapstructure:"connectionId" json:"connectionId"`
	Msgs            []MsgSubmitTxMsg `mapstructure:"msgs" json:"msgs"`
	TimeoutDuration string           `mapstructure:"timeoutDuration" json:"timeoutDuration"`
}

type MsgSubmitTxMsg struct {
	Type string `mapstructure:"@type" json:"@type"`
}
