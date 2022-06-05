package parser

import (
	"errors"
	"fmt"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"
)

func ParseGenesisCommands(
	rawGenesis *genesis.Genesis,
	accountAddressPrefix string,
) ([]command.Command, error) {
	commands := []command.Command{
		command_usecase.NewCreateGenesis(*rawGenesis),
	}
	for _, genTx := range rawGenesis.AppState.Genutil.GenTxs {
		for _, message := range genTx.Body.Messages {
			if message["@type"] == "/cosmos.staking.v1beta1.MsgCreateValidator" {
				commands = append(commands, parseGenesisGenTxsMsgCreateValidator(message)...)
			}
		}
	}
	for _, validator := range rawGenesis.AppState.Staking.Validators {
		var status constants.Status
		switch validator.Status {
		case "BOND_STATUS_UNSPECIFIED":
			status = constants.INACTIVE
		case "BOND_STATUS_BONDED":
			status = constants.BONDED
		case "BOND_STATUS_UNBONDING":
			status = constants.UNBONDING
		case "BOND_STATUS_UNBONDED":
			status = constants.UNBONDED
		}
		amountInt, parseAmountIntOk := coin.NewIntFromString(validator.Tokens)
		if !parseAmountIntOk {
			return nil, errors.New("error parsing genesis validator amount")
		}

		amount, parseAmountErr := coin.NewCoin(rawGenesis.AppState.Staking.Params.BondDenom, amountInt)
		if parseAmountErr != nil {
			return nil, fmt.Errorf("error parsing genesis validator amount: %v", parseAmountErr)
		}
		commands = append(commands, command_usecase.NewCreateGenesisValidator(
			genesis.CreateGenesisValidatorParams{
				Status: status,
				Description: model.ValidatorDescription{
					Moniker:         validator.Description.Moniker,
					Identity:        validator.Description.Identity,
					Website:         validator.Description.Website,
					SecurityContact: validator.Description.SecurityContact,
					Details:         validator.Description.Details,
				},
				Commission: model.ValidatorCommission{
					Rate:          validator.Commission.CommissionRates.Rate,
					MaxRate:       validator.Commission.CommissionRates.MaxRate,
					MaxChangeRate: validator.Commission.CommissionRates.MaxChangeRate,
				},
				MinSelfDelegation: validator.MinSelfDelegation,
				DelegatorAddress: tmcosmosutils.MustAccountAddressFromValidatorAddress(
					accountAddressPrefix,
					validator.OperatorAddress,
				),
				ValidatorAddress: validator.OperatorAddress,
				TendermintPubkey: validator.ConsensusPubkey.Key,
				Amount:           amount,
				Jailed:           validator.Jailed,
			},
		))
	}
	return commands, nil
}
