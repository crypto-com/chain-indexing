package validator_delegation

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

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
		return coin.ZeroInt(), fmt.Errorf("error in unbondingDelegationsView.Upsert(): %v", err)
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
			return coin.ZeroInt(), fmt.Errorf("error in unbondingDelegationsView.Upsert(): %v", err)
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
			return coin.ZeroInt(), fmt.Errorf("error projection.unbond(): %v", err)
		}
	}

	return totalSlashAmount, nil
}
