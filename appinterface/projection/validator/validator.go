package validator

import (
	"encoding/base64"
	"fmt"

	"github.com/crypto-com/chain-indexing/internal/utctime"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbbase"
	"github.com/crypto-com/chain-indexing/appinterface/projection/validator/constants"
	"github.com/crypto-com/chain-indexing/appinterface/projection/validator/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &Validator{}

const DO_NOT_MODIFY = "[do-not-modify]"

type Validator struct {
	*rdbbase.RDbBase

	rdbConn rdb.Conn
	logger  applogger.Logger

	conNodeAddressPrefix string
}

func NewValidator(logger applogger.Logger, rdbConn rdb.Conn, conNodeAddressPrefix string) *Validator {
	return &Validator{
		rdbbase.NewRDbBase(rdbConn.ToHandle(), "Validator"),

		rdbConn,
		logger,
		conNodeAddressPrefix,
	}
}

func (_ *Validator) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_FAILED,
		event_usecase.MSG_DELEGATE_CREATED,
		event_usecase.MSG_DELEGATE_FAILED,
		event_usecase.MSG_BEGIN_REDELEGATE_CREATED,
		event_usecase.MSG_BEGIN_REDELEGATE_FAILED,
		event_usecase.MSG_UNDELEGATE_CREATED,
		event_usecase.MSG_UNDELEGATE_FAILED,
		event_usecase.MSG_WITHDRAW_DELEGATOR_REWARD_CREATED,
		event_usecase.MSG_WITHDRAW_DELEGATOR_REWARD_FAILED,
		event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED,
		event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED,
		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.BLOCK_COMMISSIONED,
		event_usecase.VALIDATOR_JAILED,
		event_usecase.VALIDATOR_SLASHED,
		event_usecase.MSG_UNJAIL_CREATED,
		event_usecase.MSG_UNJAIL_FAILED,
		event_usecase.POWER_CHANGED,
	}
}

func (projection *Validator) OnInit() error {
	return nil
}

func (projection *Validator) HandleEvents(height int64, events []event_entity.Event) error {
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
	validatorsView := view.NewValidators(rdbTxHandle)
	validatorActivitiesView := view.NewValidatorActivities(rdbTxHandle)
	validatorActivitiesTotalView := view.NewValidatorActivitiesTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		}
	}

	if err := projection.projectValidatorView(validatorsView, height, events); err != nil {
		return fmt.Errorf("error projecting validator view: %v", err)
	}

	if err := projection.projectValidatorActivitiesView(
		validatorsView,
		validatorActivitiesView,
		validatorActivitiesTotalView,
		blockHash,
		blockTime,
		events,
	); err != nil {
		return fmt.Errorf("error projecting validator activities view: %v", err)
	}

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (projection *Validator) projectValidatorView(
	validatorsView *view.Validators,
	blockHeight int64,
	events []event_entity.Event,
) error {
	// MsgCreateValidator should be handled first
	for _, event := range events {
		if msgCreateValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			projection.logger.Debug("handling MsgCreateValidator event")

			consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromPubKey(
				projection.conNodeAddressPrefix, msgCreateValidatorEvent.Pubkey,
			)
			if err != nil {
				return fmt.Errorf("error converting consensus node pubkey to address: %v", err)
			}
			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:         consensusNodeAddress,
				OperatorAddress:              msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress:      msgCreateValidatorEvent.DelegatorAddress,
				MinSelfDelegation:            msgCreateValidatorEvent.MinSelfDelegation,
				Status:                       constants.UNBONDED,
				Jailed:                       false,
				JoinedAtBlockHeight:          blockHeight,
				Power:                        "0",
				MaybeUnbondingHeight:         nil,
				MaybeUnbondingCompletionTime: nil,
				Moniker:                      msgCreateValidatorEvent.Description.Moniker,
				Identity:                     msgCreateValidatorEvent.Description.Identity,
				Website:                      msgCreateValidatorEvent.Description.Website,
				SecurityContact:              msgCreateValidatorEvent.Description.SecurityContact,
				Details:                      msgCreateValidatorEvent.Description.Details,
				CommissionRate:               msgCreateValidatorEvent.CommissionRates.Rate,
				CommissionMaxRate:            msgCreateValidatorEvent.CommissionRates.MaxRate,
				CommissionMaxChangeRate:      msgCreateValidatorEvent.CommissionRates.MaxChangeRate,
			}

			if err := validatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error inserting new validator into view: %v", err)
			}
		}
	}

	for _, event := range events {
		if msgEditValidatorEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
			projection.logger.Debug("handling MsgEditValidator event")

			mutValidatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeOperatorAddress: &msgEditValidatorEvent.ValidatorAddress,
			})
			if err != nil {
				return fmt.Errorf(
					"error getting existing valiadtor %s from view", msgEditValidatorEvent.ValidatorAddress,
				)
			}

			if msgEditValidatorEvent.Description.Moniker != DO_NOT_MODIFY {
				mutValidatorRow.Moniker = msgEditValidatorEvent.Description.Moniker
			}
			if msgEditValidatorEvent.Description.Identity != DO_NOT_MODIFY {
				mutValidatorRow.Identity = msgEditValidatorEvent.Description.Identity
			}
			if msgEditValidatorEvent.Description.Details != DO_NOT_MODIFY {
				mutValidatorRow.Details = msgEditValidatorEvent.Description.Details
			}
			if msgEditValidatorEvent.Description.SecurityContact != DO_NOT_MODIFY {
				mutValidatorRow.SecurityContact = msgEditValidatorEvent.Description.SecurityContact
			}
			if msgEditValidatorEvent.Description.Website != DO_NOT_MODIFY {
				mutValidatorRow.Website = msgEditValidatorEvent.Description.Website
			}

			if msgEditValidatorEvent.MaybeCommissionRate != nil {
				mutValidatorRow.CommissionRate = *msgEditValidatorEvent.MaybeCommissionRate
			}
			if msgEditValidatorEvent.MaybeMinSelfDelegation != nil {
				mutValidatorRow.MinSelfDelegation = *msgEditValidatorEvent.MaybeMinSelfDelegation
			}

			if err := validatorsView.Update(mutValidatorRow); err != nil {
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		} else if validatorJailedEvent, ok := event.(*event_usecase.ValidatorJailed); ok {
			projection.logger.Debug("handling ValidatorJailed event")

			mutValidatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeConsensusNodeAddress: &validatorJailedEvent.ConsensusNodeAddress,
			})
			if err != nil {
				return fmt.Errorf(
					"error getting existing validator `%s` from view", validatorJailedEvent.ConsensusNodeAddress,
				)
			}

			mutValidatorRow.Status = constants.JAILED
			mutValidatorRow.Jailed = true

			if err := validatorsView.Update(mutValidatorRow); err != nil {
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		} else if msgUnjailEvent, ok := event.(*event_usecase.MsgUnjail); ok {
			projection.logger.Debug("handling MsgUnjail event")

			mutValidatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeOperatorAddress: &msgUnjailEvent.ValidatorAddr,
			})
			if err != nil {
				return fmt.Errorf("error getting existing validator `%s` from view", msgUnjailEvent.ValidatorAddr)
			}

			mutValidatorRow.Status = constants.BONDED
			mutValidatorRow.Jailed = false

			if err := validatorsView.Update(mutValidatorRow); err != nil {
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		} else if powerChangedEvent, ok := event.(*event_usecase.PowerChanged); ok {
			projection.logger.Debug("handling PowerChange event")

			pubkey, convErr := base64.StdEncoding.DecodeString(powerChangedEvent.TendermintPubkey)
			if convErr != nil {
				return fmt.Errorf("error base64 decoding tendermint pubkey")
			}
			consensusNodeAddress, convErr := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				projection.conNodeAddressPrefix, pubkey,
			)
			if convErr != nil {
				return fmt.Errorf("error converting tendermint pubkey to consensus pubkey")
			}

			mutValidatorRow, err := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeConsensusNodeAddress: &consensusNodeAddress,
			})
			if err != nil {
				return fmt.Errorf("error getting existing validator `%s` from view", consensusNodeAddress)
			}

			mutValidatorRow.Power = powerChangedEvent.Power
			if powerChangedEvent.Power == "0" && !mutValidatorRow.Jailed {
				mutValidatorRow.Status = constants.UNBONDED
			} else if powerChangedEvent.Power != "0" {
				mutValidatorRow.Status = constants.BONDED
			}

			if err := validatorsView.Update(mutValidatorRow); err != nil {
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		}

	}

	return nil
}
