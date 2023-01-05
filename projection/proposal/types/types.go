package types

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

type CommunityPoolSpendData struct {
	RecipientAddress string     `json:"recipient"`
	Amount           coin.Coins `json:"amount"`
}

type MsgSoftwareUpgradeData struct {
	Type      string                          `json:"@type"`
	Authority string                          `json:"authority"`
	Plan      v1_model.MsgSoftwareUpgradePlan `json:"plan"`
}

type MsgCancelUpgradeData struct {
	Type      string `json:"@type"`
	Authority string `json:"authority"`
}
