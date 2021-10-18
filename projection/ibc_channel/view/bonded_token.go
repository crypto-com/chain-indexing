package view

import "github.com/crypto-com/chain-indexing/usecase/coin"

// OnThisChain:         keep track of tokens from counterparty chain transferring to this chain
// OnCounterpartyChain: keep track of tokens from this chain transferring to counterparty chain
type BondedTokens struct {
	OnThisChain         []BondedToken `json:"onThisChain"`
	OnCounterpartyChain []BondedToken `json:"onCounterpartyChain"`
}

func NewEmptyBondedTokens() *BondedTokens {
	return &BondedTokens{
		OnThisChain:         []BondedToken{},
		OnCounterpartyChain: []BondedToken{},
	}
}

type BondedToken struct {
	Denom  string   `json:"denom"`
	Amount coin.Int `json:"amount"`
}

func NewBondedToken(denom string, amount coin.Int) *BondedToken {
	return &BondedToken{
		Denom:  denom,
		Amount: amount,
	}
}
