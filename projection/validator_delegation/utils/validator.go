package utils

import (
	"errors"

	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func ValidatorIssueNewShares(
	v view.ValidatorRow,
	amount coin.Int,
) (coin.Dec, error) {
	var issuedShares coin.Dec

	if v.Shares.IsZero() {
		issuedShares = amount.ToDec()
	} else {

		if v.Tokens.IsZero() {
			return coin.Dec{}, errors.New("validator tokens is zero, cannot issue new shares")
		}

		issuedShares = v.Shares.MulInt(amount).QuoTruncate(v.Tokens.ToDec())
	}

	return issuedShares, nil
}

// Returns the shares of a delegation given a bond amount. It
// returns an error if the validator has no tokens.
func ValidatorSharesFromTokens(
	v view.ValidatorRow,
	amount coin.Int,
) (coin.Dec, error) {
	if v.Tokens.IsZero() {
		return coin.ZeroDec(), errors.New("validator insufficient shares")
	}

	return v.Shares.MulInt(amount).QuoInt(v.Tokens), nil
}

// Calculate the token worth of provided shares
func ValidatorTokensFromShares(
	v view.ValidatorRow,
	shares coin.Dec,
) coin.Dec {
	return (shares.MulInt(v.Tokens)).Quo(v.Shares)
}

// ValidatorRemoveDelShares removes delegator shares from a validator.
// NOTE: because token fractions are left in the valiadator,
//       the exchange rate of future shares of this validator can increase.
func ValidatorRemoveDelShares(
	v view.ValidatorRow,
	delShares coin.Dec,
) (view.ValidatorRow, coin.Int, error) {
	remainingShares := v.Shares.Sub(delShares)

	var issuedTokens coin.Int
	if remainingShares.IsZero() {
		// last delegation share gets any trimmings
		issuedTokens = v.Tokens
		v.Tokens = coin.ZeroInt()
	} else {
		// leave excess tokens in the validator
		// however fully use all the delegator shares
		issuedTokens = ValidatorTokensFromShares(v, delShares).TruncateInt()
		v.Tokens = v.Tokens.Sub(issuedTokens)

		if v.Tokens.IsNegative() {
			return view.ValidatorRow{}, coin.ZeroInt(), errors.New("attempting to remove more tokens than available in validator")
		}
	}

	v.Shares = remainingShares

	return v, issuedTokens, nil
}
