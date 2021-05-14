package genesis

import (
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type CreateGenesisValidatorParams struct {
	Status            constants.Status           `json:"status"`
	Description       model.ValidatorDescription `json:"description"`
	Commission        model.ValidatorCommission  `json:"commission"`
	MinSelfDelegation string                     `json:"minSelfDelegation"`
	DelegatorAddress  string                     `json:"delegatorAddress"`
	ValidatorAddress  string                     `json:"validatorAddress"`
	TendermintPubkey  string                     `json:"tendermintPubkey"`
	Amount            coin.Coin                  `json:"amount"`
	Jailed            bool                       `json:"jailed"`
}
