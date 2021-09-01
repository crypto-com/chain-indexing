package ibc

import "github.com/crypto-com/chain-indexing/usecase/coin"

type MsgCreateVestingAccount struct {
	RawMsgCreateVestingAccount
}

type RawMsgCreateVestingAccount struct {
	Type        string     `mapstructure:"@type" json:"@type"`
	FromAddress string     `mapstructure:"from_address" json:"fromAddress"`
	ToAddress   string     `mapstructure:"to_Address" json:"toAddress"`
	Amount      coin.Coins `mapstructure:"amount" json:"amount"`
	EndTime     int64      `mapstructure:"end_time" json:"endTime"`
	Delayed     bool       `mapstructure:"delayed" json:"delayed"`
}
