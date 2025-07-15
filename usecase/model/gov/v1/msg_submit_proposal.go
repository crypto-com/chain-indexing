package model

import (
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type RawMsgSubmitProposal struct {
	Type           string        `json:"@type"`
	Messages       []interface{} `json:"messages"`
	Proposer       string        `json:"proposer"`
	InitialDeposit []interface{} `mapstructure:"initial_deposit" json:"initial_deposit"`
	Metadata       string        `json:"metadata"`
	Title          string        `json:"title"`
	Summary        string        `json:"summary"`
}
type MsgSubmitProposalMsg struct {
	Type string `mapstructure:"@type" json:"@type"`
}

type MsgSubmitProposalParams struct {
	MaybeProposalId *string       `json:"proposalId"`
	Messages        []interface{} `json:"messages"`
	InitialDeposit  coin.Coins    `json:"initial_deposit"`
	Proposer        string        `json:"proposer"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Metadata        string        `json:"metadata"`
}
