package tmcosmosutils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func SumAmount(baseDenom string, amounts []Amount) (coin.Coin, error) {
	var err error

	sum := coin.Zero()
	for _, amount := range amounts {
		if amount.Denom != baseDenom {
			return coin.Coin{}, fmt.Errorf("unrecognized denom %s when parsing amount", amount.Denom)
		}

		var amountCoin coin.Coin
		amountCoin, err = coin.NewCoinFromString(amount.Amount)
		if err != nil {
			return coin.Coin{}, fmt.Errorf("error parsing amount %s to coin: %v", amount.Amount, err)
		}
		sum, err = sum.Add(amountCoin)
		if err != nil {
			return coin.Coin{}, fmt.Errorf("error adding amount togetehr: %v", err)
		}
	}

	return sum, nil
}

type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
