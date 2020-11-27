package projection

import (
	"encoding/base64"
	"fmt"

	"github.com/crypto-com/chainindex/appinterface/pagination"

	"github.com/crypto-com/chainindex/appinterface/projection/rdbbase"
	"github.com/crypto-com/chainindex/appinterface/projection/view"
	"github.com/crypto-com/chainindex/appinterface/rdb"
	event_entity "github.com/crypto-com/chainindex/entity/event"
	projection_entity "github.com/crypto-com/chainindex/entity/projection"
	applogger "github.com/crypto-com/chainindex/internal/logger"
	"github.com/crypto-com/chainindex/internal/tmcosmosutils"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ projection_entity.Projection = &Validator{}

const JAILED = "Jailed"
const BONDED = "Bonded"
const UNBONDED = "Unbonded"
const UNBONDING = "Unbonding"

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
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_CREATED,
		event_usecase.VALIDATOR_JAILED,
		event_usecase.MSG_UNJAIL_CREATED,
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
		fmt.Println("rrrrrrrrrrrrrrrrrun here")
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()
	validatorsView := view.NewValidators(rdbTxHandle)

	for _, event := range events {
		rows, _, err := validatorsView.List(pagination.NewOffsetPagination(1, 20))
		fmt.Println(rows, err)
		if msgCreateValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			projection.logger.Debug("handling MsgCreateValidator event")

			consensusNodeAddress, convErr := tmcosmosutils.ConsensusNodeAddressFromPubKey(
				projection.conNodeAddressPrefix, msgCreateValidatorEvent.Pubkey,
			)
			if convErr != nil {
				return fmt.Errorf("error converting consensus node pubkey to address: %v", err)
			}
			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:         consensusNodeAddress,
				OperatorAddress:              msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress:      msgCreateValidatorEvent.DelegatorAddress,
				MinSelfDelegation:            msgCreateValidatorEvent.MinSelfDelegation,
				Status:                       BONDED,
				Jailed:                       false,
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

			if err := validatorsView.Insert(&validatorRow); err != nil {
				return fmt.Errorf("error inserting new validator into view: %v", err)
			}
		} else if msgEditValidatorEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
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

			mutValidatorRow.Status = JAILED
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

			mutValidatorRow.Status = BONDED
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
				mutValidatorRow.Status = UNBONDED
			}

			//execResult, err := rdbTx.Exec("UPDATE view_validators SET power = $1 WHERE consensus_node_address = $2", "10000000", mutValidatorRow.ConsensusNodeAddress)
			//fmt.Println(execResult.IsUpdate(), execResult.RowsAffected(), err)
			if err := validatorsView.Update(mutValidatorRow); err != nil {
				rows, _, err := validatorsView.List(pagination.NewOffsetPagination(1, 20))
				fmt.Println(rows, err)
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		}
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
