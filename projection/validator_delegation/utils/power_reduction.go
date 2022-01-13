package utils

import "github.com/crypto-com/chain-indexing/usecase/coin"

// TokensFromConsensusPower - convert input power to tokens
func TokensFromConsensusPower(power int64, powerReduction coin.Int) coin.Int {
	return coin.NewInt(power).Mul(powerReduction)
}
