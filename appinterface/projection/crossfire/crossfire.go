package crossfire

import (
	"fmt"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/constants"
	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire/view"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type Crossfire struct {
	*rdbprojectionbase.Base

	rdbConn              rdb.Conn
	logger               applogger.Logger
	conNodeAddressPrefix string
}

func NewCrossfire(logger applogger.Logger, rdbConn rdb.Conn, conNodeAddressPrefix string) *Crossfire {
	return &Crossfire{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Crossfire"),

		rdbConn,
		logger,
		conNodeAddressPrefix,
	}
}

func (_ *Crossfire) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,
	}
}

func (projection *Crossfire) OnInit() error {
	return nil
}

func (projection *Crossfire) HandleEvents(height int64, events []event_entity.Event) error {
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
	crossfireValidatorsView := view.NewCrossfireValidators(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		}
	}
	// TODO: remove print for projectValidatorTx view
	fmt.Println(blockTime, blockHash)

	// TODO: views preparation starts
	if err := projection.projectCrossfireValidatorView(crossfireValidatorsView, height, events); err != nil {
		return fmt.Errorf("error projecting validator view: %v", err)
	}
	// TODO ends: views preparation ends and update current height as handled

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (projection *Crossfire) projectCrossfireValidatorView(
	crossfireValidatorsView *view.CrossfireValidators,
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
				ConsensusNodeAddress:                consensusNodeAddress,
				OperatorAddress:                     msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress:             msgCreateValidatorEvent.DelegatorAddress,
				Status:                              constants.UNBONDED,
				Jailed:                              false,
				JoinedAtBlockHeight:                 blockHeight,
				Moniker:                             msgCreateValidatorEvent.Description.Moniker,
				Identity:                            msgCreateValidatorEvent.Description.Identity,
				Website:                             msgCreateValidatorEvent.Description.Website,
				SecurityContact:                     msgCreateValidatorEvent.Description.SecurityContact,
				Details:                             msgCreateValidatorEvent.Description.Details,
				Phase0TaskRegistration:              constants.INCOMPLETED,
				Phase1TaskNodeSetup:                 constants.INCOMPLETED,
				Phase1TaskBlockValidCommit:          constants.INCOMPLETED,
				Phase2TaskKeepNodeActive:            constants.INCOMPLETED,
				Phase2TaskProposalVote:              constants.INCOMPLETED,
				Phase2TaskNetworkUpgrade:            constants.INCOMPLETED,
				Phase2TaskNetworkUpgradeBlockCommit: constants.INCOMPLETED,
				Phase1n2TaskCommitmentCountRank:     0,
				Phase3TaskCommitmentCountRank:       0,
				TaskHighestSequenceRank:             0,
			}

			isJoined, joinedAtBlockHeight, err := crossfireValidatorsView.LastJoinedBlockHeight(
				validatorRow.OperatorAddress, validatorRow.ConsensusNodeAddress,
			)
			if err != nil {
				return fmt.Errorf("error querying validator last joined block height: %v", err)
			}
			if isJoined {
				validatorRow.JoinedAtBlockHeight = joinedAtBlockHeight
			}

			if err := crossfireValidatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error inserting new validator into view: %v", err)
			}
		}
	}

	return nil
}
