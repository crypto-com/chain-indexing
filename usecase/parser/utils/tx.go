package utils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func SumAmount(amounts []model.CosmosTxAuthInfoFeeAmount) (coin.Coins, error) {
	var err error

	coins := coin.NewEmptyCoins()
	for _, amount := range amounts {
		var amountCoin coin.Coin
		amountCoin, err = coin.NewCoinFromString(amount.Denom, amount.Amount)
		if err != nil {
			return nil, fmt.Errorf("error parsing amount %s to Coin: %v", amount.Amount, err)
		}
		coins = coins.Add(amountCoin)
	}

	return coins, nil
}
