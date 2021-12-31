package utils

import (
	"errors"

	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

func ValidatorIssueNewShares(
	validator view.ValidatorRow,
	amount coin.Int,
) (coin.Dec, error) {
	var issuedShares coin.Dec

	if validator.Shares.IsZero() {
		issuedShares = amount.ToDec()
	} else {

		if validator.Tokens.IsZero() {
			return coin.Dec{}, errors.New("validator tokens is zero, cannot issue new shares")
		}

		// newShares = existingShares * newTokens / existingTokens
		issuedShares = validator.Shares.MulInt(amount).QuoTruncate(validator.Tokens.ToDec())
	}

	return issuedShares, nil
}
