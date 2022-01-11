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
		return fmt.Errorf("error in inserting evidence record: %v", err)
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
		return fmt.Errorf("error finding validator by consensusNodeAddr: %v", err)
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
			return fmt.Errorf("error finding all unbonding delegations under the validator: %v", err)
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
				return fmt.Errorf("error slashing an unbonding delegation: %v", slashErr)
			}
			if amountSlashed.IsZero() {
				continue
			}

			remainingSlashAmount = remainingSlashAmount.Sub(amountSlashed)
		}

		// Iterate through redelegations from slashed source validator
		redelegations, err := redelegationsView.ListBySrcValidator(validator.OperatorAddress, height)
		if err != nil {
			return fmt.Errorf("error getting all redelegations under the src validator: %v", err)
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
				return fmt.Errorf("error slashing an redelegation: %v", slashErr)
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
		return fmt.Errorf("error removing tokens from validator: %v", err)
	}

	return nil
}

func (projection *ValidatorDelegation) slashUnbondingDelegation(
	rdbTxHandle *rdb.Handle,
	blockTime utctime.UTCTime,
	unbondingDelegation view.UnbondingDelegationRow,
	distributionHeight int64,
	slashFactor coin.Dec,
) (coin.Int, error) {

	totalSlashAmount := coin.ZeroInt()

	// perform slashing on all entries within the unbonding delegation
	for i, entry := range unbondingDelegation.Entries {
		// If unbonding started before this height, stake didn't contribute to infraction
		if entry.CreationHeight < distributionHeight {
			continue
		}

		if entry.IsMature(blockTime) {
			// Unbonding delegation no longer eligible for slashing, skip it
			continue
		}

		// Calculate slash amount proportional to stake contributing to infraction
		slashAmountDec := slashFactor.MulInt(entry.InitialBalance)
		slashAmount := slashAmountDec.TruncateInt()
		totalSlashAmount = totalSlashAmount.Add(slashAmount)

		// Don't slash more tokens than held
		// Possible since the unbonding delegation may already
		// have been slashed, and slash amounts are calculated
		// according to stake held at time of infraction
		unbondingSlashAmount := coin.MinInt(slashAmount, entry.Balance)

		// Update unbonding delegation if necessary
		if unbondingSlashAmount.IsZero() {
			continue
		}

		entry.Balance = entry.Balance.Sub(unbondingSlashAmount)
		unbondingDelegation.Entries[i] = entry
	}

	unbondingDelegationsView := NewUnbondingDelegations(rdbTxHandle)
	if err := unbondingDelegationsView.Upsert(unbondingDelegation); err != nil {
		return coin.ZeroInt(), fmt.Errorf("error upserting unbonding delegation record: %v", err)
	}

	return totalSlashAmount, nil
}

func (projection *ValidatorDelegation) slashRedelegation(
	rdbTxHandle *rdb.Handle,
	height int64,
	blockTime utctime.UTCTime,
	redelegation view.RedelegationRow,
	distributionHeight int64,
	slashFactor coin.Dec,
) (coin.Int, error) {

	delegationsView := NewDelegations(rdbTxHandle)

	totalSlashAmount := coin.ZeroInt()

	// perform slashing on all entries within the redelegation
	for _, entry := range redelegation.Entries {
		// If redelegation started before this height, stake didn't contribute to infraction
		if entry.CreationHeight < distributionHeight {
			continue
		}

		if entry.IsMature(blockTime) {
			// Redelegation no longer eligible for slashing, skip it
			continue
		}

		// Calculate slash amount proportional to stake contributing to infraction
		slashAmountDec := slashFactor.MulInt(entry.InitialBalance)
		slashAmount := slashAmountDec.TruncateInt()
		totalSlashAmount = totalSlashAmount.Add(slashAmount)

		// Unbond from target validator
		sharesToUnbond := slashFactor.Mul(entry.SharesDst)
		if sharesToUnbond.IsZero() {
			continue
		}

		valDstAddr := redelegation.ValidatorDstAddress
		delegatorAddress := redelegation.DelegatorAddress

		delegation, found, err := delegationsView.FindBy(valDstAddr, delegatorAddress, height)
		if err != nil {
			return coin.ZeroInt(), fmt.Errorf("error finding delegation: %v", err)
		}
		if !found {
			// If deleted, delegation has zero shares, and we can't unbond any more
			continue
		}

		if sharesToUnbond.GT(delegation.Shares) {
			sharesToUnbond = delegation.Shares
		}

		_, err = projection.unbond(
			rdbTxHandle,
			height,
			valDstAddr,
			delegatorAddress,
			sharesToUnbond,
		)
		if err != nil {
			return coin.ZeroInt(), fmt.Errorf("error unbonding shares from delegation and dst validator: %v", err)
		}
	}

	return totalSlashAmount, nil
}
