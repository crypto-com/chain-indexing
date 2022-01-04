package validator_delegation

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &ValidatorDelegation{}

var (
	NewValidators                = view.NewValidatorsView
	NewUnbondingValidators       = view.NewUnbondingValidatorsView
	NewDelegations               = view.NewDelegationsView
	NewUnbondingDelegations      = view.NewUnbondingDelegationsView
	NewUBDQueue                  = view.NewUBDQueueView
	UpdateLastHandledEventHeight = (*ValidatorDelegation).UpdateLastHandledEventHeight
)

type ValidatorDelegation struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	accountAddressPrefix   string
	validatorAddressPrefix string
	conNodeAddressPrefix   string
	UnbondingTime          time.Duration

	migrationHelper migrationhelper.MigrationHelper
}

func NewValidatorDelegation(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	validatorAddressPrefix string,
	conNodeAddressPrefix string,
	UnbondingTime time.Duration,
	migrationHelper migrationhelper.MigrationHelper,
) *ValidatorDelegation {
	return &ValidatorDelegation{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"ValidatorDelegation",
		),

		rdbConn,
		logger,

		accountAddressPrefix,
		validatorAddressPrefix,
		conNodeAddressPrefix,
		UnbondingTime,

		migrationHelper,
	}
}

func (_ *ValidatorDelegation) GetEventsToListen() []string {
	return []string{
		event_usecase.GENESIS_CREATED,
		event_usecase.BLOCK_CREATED,

		// Genesis
		event_usecase.GENESIS_VALIDATOR_CREATED, // parsed from genesis.

		// BeginBlock
		event_usecase.VALIDATOR_SLASHED, // parsed from BlockResult.BeginBlockEvents, emitted in BeginBlock.
		event_usecase.VALIDATOR_JAILED,  // generated along with `slash` event, introduced by ourselves.

		// Tx
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_CREATED,
		event_usecase.MSG_DELEGATE_CREATED,
		event_usecase.MSG_BEGIN_REDELEGATE_CREATED,
		event_usecase.MSG_UNDELEGATE_CREATED,
		event_usecase.MSG_UNJAIL_CREATED,

		// EndBlock
		event_usecase.POWER_CHANGED, // parsed from BlockResult.ValidatorUpdates, emitted in EndBlock.
	}
}

func (projection *ValidatorDelegation) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

	return nil
}

func (projection *ValidatorDelegation) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()

	// Get block time
	var blockTime utctime.UTCTime
	for _, event := range events {
		if genesisEvent, ok := event.(*event_usecase.GenesisCreated); ok {
			blockTime = utctime.MustParse(time.RFC3339, genesisEvent.Genesis.GenesisTime)
		} else if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	// Genesis block height is 0
	// If this is NOT a Genesis block, then always clone ValidatorRow, DelegationRow, UnbondingDelegationRow in previous height.
	// This is to keep track historical states in all height.
	if height > 0 {
		validatorsView := NewValidators(rdbTxHandle)
		if err := validatorsView.Clone(height-1, height); err != nil {
			return fmt.Errorf("error in validatorsView.Clone(): %v", err)
		}

		delegationsView := NewDelegations(rdbTxHandle)
		if err := delegationsView.Clone(height-1, height); err != nil {
			return fmt.Errorf("error in delegationsView.Clone(): %v", err)
		}

		unbondingDelegationsView := NewUnbondingDelegations(rdbTxHandle)
		if err := unbondingDelegationsView.Clone(height-1, height); err != nil {
			return fmt.Errorf("error in unbondingDelegationsView.Clone(): %v", err)
		}
	}

	// Handle events in Genesis block
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.CreateGenesisValidator); ok {

			if err := projection.handleCreateNewValidator(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddress,
				typedEvent.DelegatorAddress,
				typedEvent.TendermintPubkey,
				typedEvent.Amount.Amount,
				typedEvent.MinSelfDelegation,
			); err != nil {
				return fmt.Errorf("error handleCreateNewValidator when CreateGenesisValidator: %v", err)
			}

		}
	}

	// Handle events in BeginBlock.
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.ValidatorJailed); ok {

			if err := projection.handleValidatorJailed(
				rdbTxHandle,
				height,
				typedEvent.ConsensusNodeAddress,
			); err != nil {
				return fmt.Errorf("error handleValidatorJailed: %v", err)
			}

			// } else if _, ok := event.(*event_usecase.ValidatorSlashed); ok {

			// 	// TODO: the real devil...

		}
	}

	// Handle events and msgs in Tx
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {

			if err := projection.handleCreateNewValidator(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddress,
				typedEvent.DelegatorAddress,
				typedEvent.TendermintPubkey,
				typedEvent.Amount.Amount,
				typedEvent.MinSelfDelegation,
			); err != nil {
				return fmt.Errorf("error handleCreateNewValidator when MsgCreateValidator: %v", err)
			}

		} else if typedEvent, ok := event.(*event_usecase.MsgEditValidator); ok {

			if typedEvent.MaybeMinSelfDelegation != nil {

				if err := projection.handleEditValidator(
					rdbTxHandle,
					height,
					typedEvent.ValidatorAddress,
					*typedEvent.MaybeMinSelfDelegation,
				); err != nil {
					return fmt.Errorf("error handleEditValidator: %v", err)
				}

			}

		} else if typedEvent, ok := event.(*event_usecase.MsgDelegate); ok {

			if err := projection.handleDelegate(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddress,
				typedEvent.DelegatorAddress,
				typedEvent.Amount.Amount,
			); err != nil {
				return fmt.Errorf("error handleDelegate: %v", err)
			}

		} else if typedEvent, ok := event.(*event_usecase.MsgUndelegate); ok {

			// Successful tx with MsgUndelegate always has unbonding completion in the Msg.
			if typedEvent.MaybeUnbondCompleteAt != nil {

				if err := projection.handleUndelegate(
					rdbTxHandle,
					height,
					typedEvent.ValidatorAddress,
					typedEvent.DelegatorAddress,
					typedEvent.Amount.Amount,
					*typedEvent.MaybeUnbondCompleteAt,
				); err != nil {
					return fmt.Errorf("error handleUndelegate: %v", err)
				}

			}

			// } else if _, ok := event.(*event_usecase.MsgBeginRedelegate); ok {

			// 	// TODO

		} else if typedEvent, ok := event.(*event_usecase.MsgUnjail); ok {

			if err := projection.handleValidatorUnjailed(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddr,
			); err != nil {
				return fmt.Errorf("error handleValidatorUnjailed: %v", err)
			}

		}
	}

	// Handle events in EndBlock.
	//
	// Handle ValidatorUpdate events. These events are calculated in each EndBlock.
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.PowerChanged); ok {

			if err := projection.handlePowerChanged(
				rdbTxHandle,
				height,
				blockTime,
				typedEvent.TendermintPubkey,
				typedEvent.Power,
			); err != nil {
				return fmt.Errorf("error handlePowerChanged: %v", err)
			}

		}
	}

	if err := projection.handleMatureUnbondingValidators(
		rdbTxHandle,
		height,
		blockTime,
	); err != nil {
		return fmt.Errorf("error HandleMatureUnbondingValidators(): %v", err)
	}

	if err := projection.handleMatureUBDQueueEntries(
		rdbTxHandle,
		height,
		blockTime,
	); err != nil {
		return fmt.Errorf("error HandleMatureUBDQueueEntries(): %v", err)
	}

	// TODO: check mature Redelegation

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}
