package v1_model

import "github.com/crypto-com/chain-indexing/usecase/coin"

type MsgDepositParams struct {
	ProposalId string     `json:"proposalId"`
	Depositor  string     `json:"depositor"`
	Amount     coin.Coins `json:"amount"`
}
