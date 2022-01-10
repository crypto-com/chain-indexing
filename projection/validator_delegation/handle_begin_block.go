package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/utils"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func (projection *ValidatorDelegation) handleEvidence(
	rdbTxHandle *rdb.Handle,
	height int64,
	tendermintAddress string,
	infractionHeight int64,
	rawEvidence model.BlockEvidence,
) error {

	evidencesView := NewEvidences(rdbTxHandle)

	evidence := view.EvidenceRow{
		Height:            height,
		TendermintAddress: tendermintAddress,
		InfractionHeight:  infractionHeight,
		RawEvidence:       rawEvidence,
	}

	if err := evidencesView.Insert(evidence); err != nil {
		return fmt.Errorf("error in evidencesView.Insert(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleValidatorJailed(
	rdbTxHandle *rdb.Handle,
	height int64,
	consensusNodeAddress string,
) error {

	validatorsView := NewValidators(rdbTxHandle)

	// Get the validator
	validator, found, err := validatorsView.FindByConsensusNodeAddr(
		consensusNodeAddress,
		height,
	)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByConsensusNodeAddr(): %v", err)
	}
	if !found {
		return fmt.Errorf("error validator not found, conNodeAddr: %v", consensusNodeAddress)
	}

	if err := projection.jailValidator(rdbTxHandle, validator); err != nil {
		return fmt.Errorf("error in projection.jailValidator(): %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) handleSlash(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	consensusNodeAddress string,
	distributionHeight int64,
	power int64,
	slashFactor coin.Dec,
) error {
	validatorsView := NewValidators(rdbTxHandle)
	unbondingDelegationsView := NewUnbondingDelegations(rdbTxHandle)
	redelegationsView := NewRedelegations(rdbTxHandle)

	if slashFactor.IsNegative() {
		return fmt.Errorf("attempted to slash with a negative slash factor: %v", slashFactor)
	}

	amount := utils.TokensFromConsensusPower(power, projection.config.defaultPowerReduction)
	slashAmountDec := amount.ToDec().Mul(slashFactor)
	slashAmount := slashAmountDec.TruncateInt()

	validator, found, err := validatorsView.FindByConsensusNodeAddr(consensusNodeAddress, height)
	if err != nil {
		return fmt.Errorf("error validatorsView.FindByConsensusNodeAddr(): %v", err)
	}
	if !found {
		// If not found, the validator must have been overslashed and removed - so we don't need to do anything
		// NOTE:  Correctness dependent on invariant that unbonding delegations / redelegations must also have been completely
		//        slashed in this case - which we don't explicitly check, but should be true.
		// https://github.com/cosmos/cosmos-sdk/blob/0cb7fd081e05317ed7a2f13b0e142349a163fe4d/x/staking/keeper/slash.go#L38
		return nil
	}

	if validator.IsUnbonded() {
		return fmt.Errorf("should not be slashing unbonded validator: %s", validator.OperatorAddress)
	}

	// Track remaining slash amount for the validator
	// This will decrease when we slash unbondings and
	// redelegations, as that stake has since unbonded
	remainingSlashAmount := slashAmount

	switch {
	case distributionHeight > height:
		// Can't slash infractions in the future
		return fmt.Errorf("impossible attempt to slash future infraction at height %d but we are at height %d",
			distributionHeight, height)

	case distributionHeight == height:
		// Special-case slash at current height for efficiency - we don't need to
		// look through unbonding delegations or redelegations.
		return nil

	case distributionHeight < height:
		// Iterate through unbonding delegations from slashed validator
		unbondingDelegations, err := unbondingDelegationsView.ListByValidator(validator.OperatorAddress, height)
		if err != nil {
			return fmt.Errorf("error in unbondingDelegationsView.ListByValidator(): %v", err)
		}

		for _, unbondingDelegation := range unbondingDelegations {
			amountSlashed, slashErr := projection.slashUnbondingDelegation(
				rdbTxHandle,
				blockTime,
				unbondingDelegation,
				distributionHeight,
				slashFactor,
			)
			if slashErr != nil {
				return fmt.Errorf("error in projection.slashUnbondingDelegation(): %v", slashErr)
			}
			if amountSlashed.IsZero() {
				continue
			}

			remainingSlashAmount = remainingSlashAmount.Sub(amountSlashed)
		}

		// Iterate through redelegations from slashed source validator
		redelegations, err := redelegationsView.ListBySrcValidator(validator.OperatorAddress, height)
		if err != nil {
			return fmt.Errorf("error in redelegationsView.ListBySrcValidator(): %v", err)
		}
		for _, redelegation := range redelegations {
			amountSlashed, slashErr := projection.slashRedelegation(
				rdbTxHandle,
				height,
				blockTime,
				redelegation,
				distributionHeight,
				slashFactor,
			)
			if slashErr != nil {
				return fmt.Errorf("error in projection.slashRedelegation(): %v", slashErr)
			}
			if amountSlashed.IsZero() {
				continue
			}

			remainingSlashAmount = remainingSlashAmount.Sub(amountSlashed)
		}
	}

	// cannot decrease balance below zero
	tokensToBurn := coin.MinInt(remainingSlashAmount, validator.Tokens)
	tokensToBurn = coin.MaxInt(tokensToBurn, coin.ZeroInt()) // defensive.

	// Deduct from validator's bonded tokens and update the validator.
	if err := projection.removeValidatorTokens(rdbTxHandle, validator, tokensToBurn); err != nil {
		return fmt.Errorf("error in projection.removeValidatorTokens(): %v", err)
	}

	return nil
}
