package crossfire

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

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

	Client                   *HTTPClient
	conNodeAddressPrefix     string
	validatorAddressPrefix   string
	phaseOneStartTime        utctime.UTCTime
	phaseTwoStartTime        utctime.UTCTime
	phaseThreeStartTime      utctime.UTCTime
	competitionEndTime       utctime.UTCTime
	adminAddress             string
	networkUpgradeProposalID string
	participantsListUrl      string
	rdbConn                  rdb.Conn
	logger                   applogger.Logger
}

func NewCrossfire(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	conNodeAddressPrefix string,
	validatorAddressPrefix string,
	unixPhaseOneStartTime int64,
	unixPhaseTwoStartTime int64,
	unixPhaseThreeStartTime int64,
	unixCompetitionEndTime int64,
	adminAddress string,
	networkUpgradeProposalID string,
	participantsListURL string,
) *Crossfire {
	return &Crossfire{
		Base: rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Crossfire"),

		Client:                   NewHTTPClient(participantsListURL),
		conNodeAddressPrefix:     conNodeAddressPrefix,
		validatorAddressPrefix:   validatorAddressPrefix,
		phaseOneStartTime:        utctime.FromUnixNano(unixPhaseOneStartTime),
		phaseTwoStartTime:        utctime.FromUnixNano(unixPhaseTwoStartTime),
		phaseThreeStartTime:      utctime.FromUnixNano(unixPhaseThreeStartTime),
		competitionEndTime:       utctime.FromUnixNano(unixCompetitionEndTime),
		adminAddress:             adminAddress, // TODO: address prefix check
		networkUpgradeProposalID: networkUpgradeProposalID,
		participantsListUrl:      participantsListURL,
		rdbConn:                  rdbConn,
		logger:                   logger,
	}
}

func (_ *Crossfire) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.MSG_VOTE_CREATED,
		event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED,
		event_usecase.TRANSACTION_CREATED,
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
	crossfireChainStatsView := view.NewCrossfireChainStats(rdbTxHandle)
	crossfireValidatorsStatsView := view.NewCrossfireValidatorsStats(rdbTxHandle)

	// TODO: views preparation starts
	// handle block created event
	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			if handleErr := projection.handleBlockCreatedEvent(crossfireChainStatsView, crossfireValidatorsView, crossfireValidatorsStatsView, blockCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling BlockCreatedEvent: %v", handleErr)
			}
		}
	}

	if err := projection.projectCrossfireValidatorView(crossfireValidatorsView, crossfireChainStatsView, crossfireValidatorsStatsView, height, blockTime, events); err != nil {
		return fmt.Errorf("error projecting validator view: %v", err)
	}

	// Recompute TxSentRank
	errTxSentRank := projection.computeTxSentRank(crossfireValidatorsStatsView, crossfireValidatorsView)
	if errTxSentRank != nil {
		return fmt.Errorf("[error] Updating TxSentTask Rank %w", errTxSentRank)
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

func (projection *Crossfire) handleBlockCreatedEvent(
	crossfireChainStatsView *view.CrossfireChainStats,
	crossfireValidatorsView *view.CrossfireValidators,
	crossfireValidatorsStatsView *view.CrossfireValidatorsStats,
	event *event_usecase.BlockCreated,
) error {
	blockTime := event.Block.Time
	blockHeight := event.Block.Height

	// increment current phase block number
	if blockTime.After(projection.phaseOneStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
		err := crossfireChainStatsView.IncrementOne(constants.PHASE1_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("error increment phase1 block count")
		}
	} else if blockTime.After(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
		err := crossfireChainStatsView.IncrementOne(constants.PHASE2_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("error increment phase2 block count")
		}
	} else if blockTime.After(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
		// check the keep active task, throttling with every 10 blocks
		if blockHeight%10 == 0 {
			if err := projection.checkTaskKeepActive(crossfireChainStatsView, crossfireValidatorsView, crossfireValidatorsStatsView); err != nil {
				return fmt.Errorf("error checkTaskKeepActive: %v", err)
			}
		}

		err := crossfireChainStatsView.IncrementOne(constants.PHASE3_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("error increment phase3 block count")
		}
	}

	// increment current block validator commitment number
	for _, signature := range event.Block.Signatures {
		validator, err := crossfireValidatorsView.FindBy(view.CrossfireValidatorIdentity{
			MaybeTendermintAddress: &signature.ValidatorAddress,
		})
		if errors.Is(err, rdb.ErrNoRows) {
			// no validator found by tendermint address
			continue
		}
		if err != nil {
			return fmt.Errorf(
				"error getting existing validator by block commitment of %s from view %s", signature.ValidatorAddress, err,
			)
		}

		if blockTime.After(projection.phaseOneStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE1_COMMIT)
			if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
				return fmt.Errorf(
					"error increment validator commitment count %s", key,
				)
			}
		} else if blockTime.After(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE2_COMMIT)
			if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
				return fmt.Errorf(
					"error increment validator commitment count %s", key,
				)
			}

			//Check task network upgrade
			errUpdatingTask := projection.checkTaskNetworkUpgrade(crossfireValidatorsView, crossfireChainStatsView, validator, blockTime, blockHeight)
			if errUpdatingTask != nil {
				return fmt.Errorf(
					"error Updating crossfire task completion %s", errUpdatingTask,
				)
			}
		} else if blockTime.After(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE3_COMMIT)
			if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
				return fmt.Errorf(
					"error increment validator commitment count %s", key,
				)
			}
		}
	}

	return nil
}

func (projection *Crossfire) projectCrossfireValidatorView(
	crossfireValidatorsView *view.CrossfireValidators,
	crossfireChainStatsView *view.CrossfireChainStats,
	crossfireValidatorStatsView *view.CrossfireValidatorsStats,
	blockHeight int64,
	blockTime utctime.UTCTime,
	events []event_entity.Event,
) error {
	// MsgCreateValidator should be handled first
	for _, event := range events {
		if msgCreateValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			projection.logger.Debug("handling MsgCreateValidator event")

			pubKey, err := base64.StdEncoding.DecodeString(msgCreateValidatorEvent.TendermintPubkey)
			if err != nil {
				return fmt.Errorf("error base64 decoding Tendermint node pubkey: %v", err)
			}
			tendermintAddress := tmcosmosutils.TmAddressFromTmPubKey(pubKey)

			consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				projection.conNodeAddressPrefix, pubKey,
			)
			if err != nil {
				return fmt.Errorf("error converting consensus node pubkey to address: %v", err)
			}

			validatorRow := view.CrossfireValidatorRow{
				ConsensusNodeAddress:            consensusNodeAddress,
				OperatorAddress:                 msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress:         msgCreateValidatorEvent.DelegatorAddress,
				TendermintPubkey:                msgCreateValidatorEvent.TendermintPubkey,
				TendermintAddress:               tendermintAddress,
				Status:                          constants.UNBONDED,
				Jailed:                          false,
				JoinedAtBlockHeight:             blockHeight,
				JoinedAtBlockTime:               blockTime,
				Moniker:                         msgCreateValidatorEvent.Description.Moniker,
				Identity:                        msgCreateValidatorEvent.Description.Identity,
				Website:                         msgCreateValidatorEvent.Description.Website,
				SecurityContact:                 msgCreateValidatorEvent.Description.SecurityContact,
				Details:                         msgCreateValidatorEvent.Description.Details,
				TaskPhase1NodeSetup:             constants.INCOMPLETED,
				TaskPhase2KeepNodeActive:        constants.INCOMPLETED,
				TaskPhase2ProposalVote:          constants.INCOMPLETED,
				TaskPhase2NetworkUpgrade:        constants.INCOMPLETED,
				RankTaskPhase1n2CommitmentCount: 0,
				RankTaskPhase3CommitmentCount:   0,
				RankTaskHighestTxSent:           0,
			}

			isJoined, joinedAtBlockHeight, joinedAtBlockTime, err := crossfireValidatorsView.LastJoinedBlockHeight(
				validatorRow.OperatorAddress, validatorRow.ConsensusNodeAddress,
			)
			if err != nil {
				return fmt.Errorf("error querying validator last joined block height: %v", err)
			}
			if isJoined {
				validatorRow.JoinedAtBlockHeight = joinedAtBlockHeight
				validatorRow.JoinedAtBlockTime = joinedAtBlockTime
			}

			if err := crossfireValidatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error inserting new validator into view: %v", err)
			}

			// checkTaskSetup
			if err := projection.checkTaskSetup(
				crossfireValidatorsView,
				validatorRow.OperatorAddress,
				validatorRow.ConsensusNodeAddress,
				validatorRow.JoinedAtBlockTime,
			); err != nil {
				return fmt.Errorf("error check Setup task for new validator: %v", err)
			}
		}
	}
	for _, event := range events {
		if msgSubmitSoftwareUpgradeProposalEvent, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			projection.logger.Debug("handling MsgSubmitSoftwareUpgradeProposal event")

			// Check if proposed after competition has ended
			if blockTime.After(projection.competitionEndTime) {
				return fmt.Errorf("error Competition has already ended")
			}

			// Check if proposed before OR after Phase 2
			if blockTime.Before(projection.phaseTwoStartTime) || blockTime.After(projection.phaseThreeStartTime) {
				return fmt.Errorf("error This proposal does not occur in Phase 2")
			}

			// Check if proposed by NOT an admin
			if msgSubmitSoftwareUpgradeProposalEvent.ProposerAddress != projection.adminAddress {
				return fmt.Errorf("error checking proposer address equals admin address")
			}

			// Check if proposal ID does not match the required ID
			if msgSubmitSoftwareUpgradeProposalEvent.MaybeProposalId != nil && *msgSubmitSoftwareUpgradeProposalEvent.MaybeProposalId != projection.networkUpgradeProposalID {
				return fmt.Errorf("error checking Proposal ID in proposal")
			}

			networkUpgradeTimestamp := msgSubmitSoftwareUpgradeProposalEvent.MsgSubmitSoftwareUpgradeProposalParams.Content.Plan.Time
			networkUpgradeBlockheight := msgSubmitSoftwareUpgradeProposalEvent.MsgSubmitSoftwareUpgradeProposalParams.Content.Plan.Height

			// Update Network upgrade target Timestamp in DB
			networkUpgradeTimestampDBKey := fmt.Sprintf("%s%s%s", constants.NETWORK_UPGRADE, constants.DB_KEY_SEPARATOR, "timestamp")
			errTSUpdate := crossfireChainStatsView.Set(networkUpgradeTimestampDBKey, networkUpgradeTimestamp.UnixNano())
			if errTSUpdate != nil {
				return fmt.Errorf("error updating network_upgrade timestamp: %v", errTSUpdate)
			}

			// Update Network upgrade target Blockheight in DB
			networkUpgradeBlockheightDBKey := fmt.Sprintf("%s%s%s", constants.NETWORK_UPGRADE, constants.DB_KEY_SEPARATOR, "blockheight")
			errBlockheightUpdate := crossfireChainStatsView.Set(networkUpgradeBlockheightDBKey, networkUpgradeBlockheight)

			if errBlockheightUpdate != nil {
				return fmt.Errorf("error updating network_upgrade blockheight: %v", errBlockheightUpdate)
			}

		} else if msgVoteCreated, ok := event.(*event_usecase.MsgVote); ok {
			projection.logger.Debug("handling MsgVote event")

			// Check if proposed after competition has ended
			if blockTime.After(projection.competitionEndTime) {
				return fmt.Errorf("error Competition has already ended")
			}

			// Check if proposed before OR after Phase 2
			if blockTime.Before(projection.phaseTwoStartTime) || blockTime.After(projection.phaseThreeStartTime) {
				return fmt.Errorf("error Ineligible vote as it does not occur in Phase 2")
			}

			// Check if proposal ID does not match the required ID
			if msgVoteCreated.ProposalId != "" && msgVoteCreated.ProposalId != projection.networkUpgradeProposalID {
				return fmt.Errorf("error checking Proposal ID in Vote not matching")
			}

			// Check if Vote is NOT Yes or Abstain
			// TODO: Whether keep VOTE_OPTION_UNSPECIFIED or not?
			if !(strings.ToUpper(msgVoteCreated.Option) == constants.VOTE_OPTION_YES || strings.ToUpper(msgVoteCreated.Option) == constants.VOTE_OPTION_ABSTAIN || strings.ToUpper(msgVoteCreated.Option) == constants.VOTE_OPTION_UNSPECIFIED) {
				return fmt.Errorf("error Ineligible Vote. Casted vote: %s", msgVoteCreated.Option)
			}

			//TODO: Is this assumption correct? How to fetch operatorAddress / consensusAddress
			operatorAddress, errConverting := tmcosmosutils.ValidatorAddressFromPubAddress(projection.validatorAddressPrefix, msgVoteCreated.Voter)

			if errConverting != nil {
				return fmt.Errorf("error In converting voter address to validator address: %v", errConverting)
			}

			errCheckingTask := projection.checkTaskNetworkProposalVote(crossfireValidatorsView, operatorAddress, blockTime)

			if errCheckingTask != nil {
				return fmt.Errorf("error In checking Network proposal vote: %v", errCheckingTask)
			}
			// Update the proposed ID against the voter in Database
			voted_proposal_id_db_key := constants.VOTED_PROPOSAL_ID + constants.DB_KEY_SEPARATOR + msgVoteCreated.Voter

			proposalIdAsInt64, errConversion := strconv.ParseInt(msgVoteCreated.ProposalId, 10, 64)
			if errConversion != nil {
				return fmt.Errorf("error converting ProposalID to int64: %v", errConversion)
			}
			errUpdateValidatorStats := crossfireValidatorStatsView.Set(voted_proposal_id_db_key, proposalIdAsInt64)

			if errUpdateValidatorStats != nil {
				return fmt.Errorf("error Updating ProposalID for the voter: %v", errUpdateValidatorStats)
			}
		} else if transactionCreatedEvent, ok := event.(*event_usecase.TransactionCreated); ok {
			projection.logger.Debug("handling TransactionCreated event")

			// Check if proposed after competition has ended
			if blockTime.After(projection.competitionEndTime) {
				return fmt.Errorf("error Competition has already ended")
			}

			// Update the Tx Count
			errUpdateTxCount := projection.updateTxSentCount(crossfireValidatorStatsView, blockTime, transactionCreatedEvent)
			if errUpdateTxCount != nil {
				return fmt.Errorf("error Updating tx sent count: %v", errUpdateTxCount)
			}
		}
	}
	return nil
}

// checkTaskSetup checks if node joins before phase 2 and update task_phase_1_node_setup
func (projection *Crossfire) checkTaskSetup(
	crossfireValidatorsView *view.CrossfireValidators,
	operatorAddress string,
	consensusNodeAddress string,
	joinedAtBlockTime utctime.UTCTime,
) error {
	if joinedAtBlockTime.Before(projection.phaseTwoStartTime) {
		if err := crossfireValidatorsView.UpdateTask(
			"task_phase_1_node_setup",
			constants.COMPLETED,
			operatorAddress,
			consensusNodeAddress,
		); err != nil {
			return fmt.Errorf("error updating validator TaskPhase1NodeSetup as completed: %v", err)
		}
	}

	if joinedAtBlockTime.After(projection.phaseTwoStartTime) {
		if err := crossfireValidatorsView.UpdateTask(
			"task_phase_1_node_setup",
			constants.MISSED,
			operatorAddress,
			consensusNodeAddress,
		); err != nil {
			return fmt.Errorf("error updating validator TaskPhase1NodeSetup as missed: %v", err)
		}
	}

	return nil
}

// checkTaskKeepActive checks if a validator has over half of commitments of phase 2
func (projection *Crossfire) checkTaskKeepActive(
	crossfireChainStatsView *view.CrossfireChainStats,
	crossfireValidatorsView *view.CrossfireValidators,
	crossfireValidatorsStatsView *view.CrossfireValidatorsStats,
) error {
	phase2TotalBlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE2_BLOCK_COUNT)
	if err != nil {
		return fmt.Errorf("error getting PHASE2_BLOCK_COUNT: %v", err)
	}

	validators, err := crossfireValidatorsView.List()
	if err != nil {
		return fmt.Errorf("error listing all validators for checkTaskKeepActive: %v", err)
	}

	for _, validator := range validators {
		key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE2_COMMIT)
		validatorPhase2CommitCount, err := crossfireValidatorsStatsView.FindBy(key)
		if err != nil {
			return fmt.Errorf("error find current validator's phase commit count: %v %s", err, validator.OperatorAddress)
		}
		// over 50% of the commitments
		if validatorPhase2CommitCount >= phase2TotalBlockCount/2 {
			if err := crossfireValidatorsView.UpdateTask(
				"task_phase_2_keep_node_active",
				constants.COMPLETED,
				validator.OperatorAddress,
				validator.ConsensusNodeAddress,
			); err != nil {
				return fmt.Errorf("error updating validator task_phase_2_keep_node_active as completed for validator: %v %s", err, validator.OperatorAddress)
			}
		}
	}

	return nil
}

// checkTaskNetworkProposalVote
func (projection *Crossfire) checkTaskNetworkProposalVote(
	crossfireValidatorsView *view.CrossfireValidators,
	voterAddress string,
	txBlockTime utctime.UTCTime,
) error {
	if txBlockTime.Before(projection.phaseThreeStartTime) && txBlockTime.After(projection.phaseTwoStartTime) {
		if err := crossfireValidatorsView.UpdateTaskForOperatorAddress(
			constants.TASK_PHASE_2_PROPOSAL_VOTE_COLUMN_NAME,
			constants.COMPLETED,
			voterAddress,
		); err != nil {
			return fmt.Errorf("error updating validator Phase_2 Task_1 as completed %v", err)
		}
	}

	if txBlockTime.After(projection.phaseThreeStartTime) {
		if err := crossfireValidatorsView.UpdateTaskForOperatorAddress(
			constants.TASK_PHASE_2_PROPOSAL_VOTE_COLUMN_NAME,
			constants.MISSED,
			voterAddress,
		); err != nil {
			return fmt.Errorf("error updating validator Phase_2 Task_1 as missed: %v", err)
		}
	}

	return nil
}

// Update Tx sent count for sender
func (projection *Crossfire) updateTxSentCount(
	crossfireValidatorStatsView *view.CrossfireValidatorsStats,
	blockTime utctime.UTCTime,
	msgTransactionCreated *event_usecase.TransactionCreated,
) error {
	var phaseNumberPrefix string
	if blockTime.After(projection.phaseOneStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
		phaseNumberPrefix = constants.PHASE_1_TX_SENT_PREFIX
	} else if blockTime.After(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
		phaseNumberPrefix = constants.PHASE_2_TX_SENT_PREFIX
	} else if blockTime.After(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
		phaseNumberPrefix = constants.PHASE_3_TX_SENT_PREFIX
	}
	for _, sender := range msgTransactionCreated.Senders {

		// Only considering Pubkey address for now
		if sender.Type == constants.TYPE_URL_PUBKEY && sender.MaybeThreshold == nil {

			// Increment count for address as per PHASE
			phaseAddressCountDbKey := phaseNumberPrefix + constants.DB_KEY_SEPARATOR + sender.Pubkeys[0]
			errIncrementing := crossfireValidatorStatsView.IncrementOne(phaseAddressCountDbKey)
			if errIncrementing != nil {
				return fmt.Errorf("error Phase wise tx sent count increment: %v", errIncrementing)
			}

			// Increment TOTAL count for address
			totalAddressCountDbKey := constants.TOTAL_TX_SENT_PREFIX + constants.DB_KEY_SEPARATOR + sender.Pubkeys[0]
			errIncrementingTotal := crossfireValidatorStatsView.IncrementOne(totalAddressCountDbKey)
			if errIncrementingTotal != nil {
				return fmt.Errorf("error Incrementing tx sent count: %v", errIncrementingTotal)
			}

		}
	}
	return nil
}

// Compute Tx Sent Rank for participating validators
func (projection *Crossfire) computeTxSentRank(
	crossfireValidatorStatsView *view.CrossfireValidatorsStats,
	crossfireValidatorView *view.CrossfireValidators,
) error {
	// Using dummy list for development
	participantsList, errGetRequest := projection.Client.Participants()
	if errGetRequest != nil {
		return fmt.Errorf("[error] Getting Participants list %w", errGetRequest)
	}

	// Get All Tx Count Sorted
	dbParticipantWithCountList, errDbCount := crossfireValidatorStatsView.FindAllLike(constants.TOTAL_TX_SENT_PREFIX)
	if errDbCount != nil {
		return fmt.Errorf("[error] Database Participant With Total Count %w", errDbCount)
	}
	if len(dbParticipantWithCountList) < 1 {
		// No Participants to Rank
		return nil
	}

	// Sort With Most Tx Sent First
	sort.SliceStable(dbParticipantWithCountList, func(i, j int) bool {
		return dbParticipantWithCountList[i].Value > dbParticipantWithCountList[j].Value
	})

	var rank int = 1
	for index, dbParticipant := range dbParticipantWithCountList {
		dbParticipantPrimaryAddress := strings.Split(dbParticipant.Key, constants.DB_KEY_SEPARATOR)[1]
		// Check for each Participant in URL
		for _, participant := range *participantsList {

			// Check if the primary address match
			if participant.PrimaryAddress == dbParticipantPrimaryAddress {

				// Check if Participant is a Registered Crossfire Validator
				registeredCrossfireValidator, errFindingValidator := crossfireValidatorView.FindBy(view.CrossfireValidatorIdentity{
					MaybeOperatorAddress: &participant.OperatorAddress,
				})
				if errFindingValidator != nil || registeredCrossfireValidator == nil {
					break
				}

				// Update Ranks with specific Validator
				errUpdating := crossfireValidatorView.UpdateTxSentRank(rank, participant.PrimaryAddress, participant.OperatorAddress)
				if errUpdating != nil {
					return fmt.Errorf("[error] Updating TX SENT Task Rank %w", errUpdating)
				}
				if index+1 < len(dbParticipantWithCountList) && dbParticipantWithCountList[index].Value != dbParticipantWithCountList[index+1].Value {
					rank++
				}
			}
		}
	}
	return nil
}

// checkTaskNetworkUpgrade
func (projection *Crossfire) checkTaskNetworkUpgrade(
	crossfireValidatorsView *view.CrossfireValidators,
	crossfireChainStatsView *view.CrossfireChainStats,
	validator *view.CrossfireValidatorRow,
	blockTime utctime.UTCTime,
	blockHeight int64,
) error {

	// Check if Validator's upgrade is already successful
	if validator.TaskPhase2NetworkUpgrade == constants.COMPLETED {
		return nil
	}

	targetTimestampDBKey := fmt.Sprintf("%s%s%s", constants.NETWORK_UPGRADE, constants.DB_KEY_SEPARATOR, "timestamp")
	targetBlockHeightDBKey := fmt.Sprintf("%s%s%s", constants.NETWORK_UPGRADE, constants.DB_KEY_SEPARATOR, "blockheight")

	networkUpgradeTimestampNanoSec, errTimestamp := crossfireChainStatsView.FindBy(targetTimestampDBKey)
	if errTimestamp != nil {
		return fmt.Errorf("error getting network Upgrade timestamp: %v", errTimestamp)
	}

	networkUpgradeBlockheight, errBlockheight := crossfireChainStatsView.FindBy(targetBlockHeightDBKey)
	if errBlockheight != nil {
		return fmt.Errorf("error getting network Upgrade Blockheight: %v", errBlockheight)
	}

	// Check if current block is before the network upgrade
	if blockTime.Before(utctime.FromUnixNano(networkUpgradeTimestampNanoSec)) && blockHeight <= networkUpgradeBlockheight {
		return nil
	}
	errUpdatingTaskCompletion := crossfireValidatorsView.UpdateTaskForOperatorAddress("task_phase_2_network_upgrade", constants.COMPLETED, validator.OperatorAddress)
	if errUpdatingTaskCompletion != nil {
		return fmt.Errorf("error Updating Task completion: %v", errUpdatingTaskCompletion)
	}

	return nil
}

type ParticipantDetail struct {
	OperatorAddress string `json:"operatorAddress"`
	PrimaryAddress  string `json:"primaryAddress"`
	Moniker         string `json:"moniker"`
}
type RankedParticipant struct {
	PrimaryAddress string `json:"primaryAddress"`
	Count          int64  `json:"txSentCount"`
	Rank           int    `json:"rank"`
}
