package rdbvalidatorbase

import (
	"encoding/base64"
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbvalidatorbase/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

const DO_NOT_MODIFY = "[do-not-modify]"

var (
	NewValidators = view.NewValidatorsView
)

// a generic Validator projection. For table schema refer to view/validators.go
type Base struct {
	tableName string

	conNodeAddressPrefix string
}

func NewBase(tableName string, conNodeAddressPrefix string) *Base {
	return &Base{
		tableName,

		conNodeAddressPrefix,
	}
}

func (projection *Base) GetEventsToListen() []string {
	return []string{
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.GENESIS_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_CREATED,
	}
}

// Handle
func (projection *Base) HandleEvents(conn *rdb.Handle, logger logger.Logger, events []event_entity.Event) error {
	validatorsView := NewValidators(conn, projection.tableName)
	for _, event := range events {
		if createGenesisValidator, ok := event.(*event_usecase.CreateGenesisValidator); ok {
			logger.Debug("handling CreateGenesisValidator event")
			tendermintPubkey, err := base64.StdEncoding.DecodeString(createGenesisValidator.TendermintPubkey)
			if err != nil {
				return fmt.Errorf("error base64 decoding Tendermint node pubkey: %v", err)
			}
			consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				projection.conNodeAddressPrefix, tendermintPubkey,
			)
			if err != nil {
				return fmt.Errorf("error converting Tendermint node pubkey to address: %v", err)
			}
			tendermintAddress := tmcosmosutils.TmAddressFromTmPubKey(tendermintPubkey)
			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:    consensusNodeAddress,
				OperatorAddress:         createGenesisValidator.ValidatorAddress,
				InitialDelegatorAddress: createGenesisValidator.DelegatorAddress,
				TendermintAddress:       tendermintAddress,
				TendermintPubkey:        createGenesisValidator.TendermintPubkey,
				Moniker:                 createGenesisValidator.Description.Moniker,
			}

			if err := validatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error upserting new validator into view: %v", err)
			}
		} else if msgCreateValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			logger.Debug("handling MsgCreateValidator event")
			tendermintPubkey, err := base64.StdEncoding.DecodeString(msgCreateValidatorEvent.TendermintPubkey)
			if err != nil {
				return fmt.Errorf("error base64 decoding Tendermint node pubkey: %v", err)
			}
			consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				projection.conNodeAddressPrefix, tendermintPubkey,
			)
			if err != nil {
				return fmt.Errorf("error converting Tendermint node pubkey to address: %v", err)
			}
			tendermintAddress := tmcosmosutils.TmAddressFromTmPubKey(tendermintPubkey)
			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:    consensusNodeAddress,
				OperatorAddress:         msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress: msgCreateValidatorEvent.DelegatorAddress,
				TendermintAddress:       tendermintAddress,
				TendermintPubkey:        msgCreateValidatorEvent.TendermintPubkey,
				Moniker:                 msgCreateValidatorEvent.Description.Moniker,
			}

			if err := validatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error upserting new validator into view: %v", err)
			}
		} else if msgEditValidatorEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
			logger.Debug("handling MsgEditValidator event")

			mutValidatorRow, err := validatorsView.FindLastBy(view.ValidatorIdentity{
				MaybeOperatorAddress: &msgEditValidatorEvent.ValidatorAddress,
			})
			if err != nil {
				return fmt.Errorf(
					"error getting existing validator %s from view", msgEditValidatorEvent.ValidatorAddress,
				)
			}

			if msgEditValidatorEvent.Description.Moniker != DO_NOT_MODIFY {
				mutValidatorRow.Moniker = msgEditValidatorEvent.Description.Moniker
			}

			if err := validatorsView.Update(mutValidatorRow); err != nil {
				return fmt.Errorf("error updating validator into view: %v", err)
			}
		}
	}

	return nil
}

func (projection *Base) GetView(conn *rdb.Handle) view.Validators {
	return NewValidators(conn, projection.tableName)
}
