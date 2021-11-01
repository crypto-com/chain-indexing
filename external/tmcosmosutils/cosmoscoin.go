package tmcosmosutils

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/coin"
)

// MustNewCoinFromAmountInterface returns a Coin from the amount map in the form of
// map[string]interface{}. It panics when the amount is invalid.
func MustNewCoinFromAmountInterface(amount map[string]interface{}) coin.Coin {
	result, err := NewCoinFromAmountInterface(amount)
	if err != nil {
		panic(err)
	}

	return result
}

// NewCoinFromAmountInterface returns a Coin from the amount in the form of
// map[string]interface{}. It returns error when the amount is invalid.
func NewCoinFromAmountInterface(amount map[string]interface{}) (coin.Coin, error) {
	result, err := coin.NewCoinFromString(amount["denom"].(string), amount["amount"].(string))
	if err != nil {
		return coin.Coin{}, err
	}

	return result, nil
}

// NewCoinsFromAmountInterface returns Coins from the list of amount in the
// from of []interface{}. It behaves the same as NewCoinsFromAmountInterface
// except it panics on any error.
func MustNewCoinsFromAmountInterface(amounts []interface{}) coin.Coins {
	result, err := NewCoinsFromAmountInterface(amounts)
	if err != nil {
		panic(err)
	}

	return result
}

// NewCoinsFromAmountInterface returns Coins from the list of amount in the
// from of []interface{}. It returns error when any of the amount inside is
// invalid
func NewCoinsFromAmountInterface(amounts []interface{}) (coin.Coins, error) {
	coins := coin.NewCoins()
	for _, rawAmount := range amounts {
		amount, ok := rawAmount.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid coin in the amount list: %v", rawAmount)
		}
		coinUnit, err := NewCoinFromAmountInterface(amount)
		if err != nil {
			return nil, err
		}

		coins = coins.Add(coinUnit)
	}

	return coins, nil
}
