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
	primaryAddressPrefix     string
	phaseOneStartTime        utctime.UTCTime
	phaseTwoStartTime        utctime.UTCTime
	phaseThreeStartTime      utctime.UTCTime
	competitionEndTime       utctime.UTCTime
	JackpotOneStartTime      utctime.UTCTime
	JackpotTwoStartTime      utctime.UTCTime
	JackpotThreeStartTime    utctime.UTCTime
	JackpotFourStartTime     utctime.UTCTime
	JackpotFourEndTime       utctime.UTCTime
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
	primaryAddressPrefix string,
	unixPhaseOneStartTime int64,
	unixPhaseTwoStartTime int64,
	unixPhaseThreeStartTime int64,
	unixCompetitionEndTime int64,
	adminAddress string,
	networkUpgradeProposalID string,
	participantsListURL string,
	JackpotOneStartTime int64,
	JackpotTwoStartTime int64,
	JackpotThreeStartTime int64,
	JackpotFourStartTime int64,
	JackpotFourEndTime int64,
) *Crossfire {
	return &Crossfire{
		Base: rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Crossfire"),

		Client:                 NewHTTPClient(participantsListURL),
		conNodeAddressPrefix:   conNodeAddressPrefix,
		validatorAddressPrefix: validatorAddressPrefix,
		primaryAddressPrefix:   primaryAddressPrefix,
		phaseOneStartTime:      utctime.FromUnixNano(unixPhaseOneStartTime),
		phaseTwoStartTime:      utctime.FromUnixNano(unixPhaseTwoStartTime),
		phaseThreeStartTime:    utctime.FromUnixNano(unixPhaseThreeStartTime),
		competitionEndTime:     utctime.FromUnixNano(unixCompetitionEndTime),

		JackpotOneStartTime:   utctime.FromUnixNano(JackpotOneStartTime),
		JackpotTwoStartTime:   utctime.FromUnixNano(JackpotTwoStartTime),
		JackpotThreeStartTime: utctime.FromUnixNano(JackpotThreeStartTime),
		JackpotFourStartTime:  utctime.FromUnixNano(JackpotFourStartTime),
		JackpotFourEndTime:    utctime.FromUnixNano(JackpotFourEndTime),

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
		return fmt.Errorf("[Crossfire] error beginning transaction: %v", err)
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

	// Event projection
	projection.profile("begin handling event")
	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			projection.profile("begin handling block created event")
			if handleErr := projection.handleBlockCreatedEvent(crossfireChainStatsView, crossfireValidatorsView, crossfireValidatorsStatsView, blockCreatedEvent); handleErr != nil {
				return fmt.Errorf("[Crossfire] error handling BlockCreatedEvent: %v", handleErr)
			}
			projection.profile("end handling block created event")
		}
	}

	projection.profile("begin projecting validator view")
	if err := projection.projectCrossfireValidatorView(crossfireValidatorsView, crossfireChainStatsView, crossfireValidatorsStatsView, height, blockTime, events); err != nil {
		return fmt.Errorf("[Crossfire] error projecting validator view: %v", err)
	}
	projection.profile("end projecting validator view")

	// Ranks computation every 100 blocks and/or at the end of event projection
	if height%100 == 0 {
		projection.profile("begin computing transaction sent rank")
		if err := projection.computeTxSentRank(crossfireValidatorsStatsView, crossfireValidatorsView); err != nil {
			return fmt.Errorf("[Crossfire] error Updating TxSentTask Rank %v", err)
		}
		projection.profile("end computing transaction sent rank")

		projection.profile("begin phase 1 and 2 computing commitment rank")
		if err := projection.computeCommitmentRank(
			"rank_task_phase_1_2_commitment_count",
			constants.PHASE_1N2_COMMIT_PREFIX,
			crossfireValidatorsStatsView,
			crossfireValidatorsView,
		); err != nil {
			return fmt.Errorf("[Crossfire] error compute rank_task_phase_1_2_commitment_count %v", err)
		}
		projection.profile("end phase 1 and 2 computing commitment rank")

		if blockTime.AfterOrEqual(projection.phaseThreeStartTime) {
			projection.profile("begin phase 3 computing commitment rank")
			if err := projection.computeCommitmentRank(
				"rank_task_phase_3_commitment_count",
				constants.PHASE_3_COMMIT_PREFIX,
				crossfireValidatorsStatsView,
				crossfireValidatorsView,
			); err != nil {
				return fmt.Errorf("[Crossfire] error compute rank_task_phase_3_commitment_count %v", err)
			}
			projection.profile("end phase 3 computing commitment rank")
		}
	}

	// Done current height projection and ranking computation
	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("[Crossfire] error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("[Crossfire] error committing changes: %v", err)
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
	if blockTime.AfterOrEqual(projection.phaseOneStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
		err := crossfireChainStatsView.IncrementOne(constants.PHASE_1_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("[Crossfire] error increment phase1 block count")
		}
	} else if blockTime.AfterOrEqual(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
		err := crossfireChainStatsView.IncrementOne(constants.PHASE_2_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("[Crossfire] error increment phase2 block count")
		}
	} else if blockTime.AfterOrEqual(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
		// check the keep active task, throttling with every 100 blocks
		// keep active task is only counting commitments during phase 2, so checking it in phase 3
		// guarantees correctness
		if blockHeight%100 == 0 {
			projection.profile("begin checking keep alive task")
			if err := projection.checkTaskKeepActive(crossfireChainStatsView, crossfireValidatorsView, crossfireValidatorsStatsView); err != nil {
				return fmt.Errorf("[Crossfire] error checkTaskKeepActive: %v", err)
			}
			projection.profile("end checking keep alive task")
		}

		err := crossfireChainStatsView.IncrementOne(constants.PHASE_3_BLOCK_COUNT)
		if err != nil {
			return fmt.Errorf("[Crossfire] error increment phase3 block count")
		}
	}

	// increment current block validator commitment number
	projection.profile("begin looping block validator commitment")
	keysToIncrementByOne := make([]string, 0, len(event.Block.Signatures))
	identityType := "tendermint_address"
	identities := make([]view.CrossfireValidatorIdentity, 0, len(event.Block.Signatures))
	for i := range event.Block.Signatures {
		identities = append(identities, view.CrossfireValidatorIdentity{
			MaybeTendermintAddress: &event.Block.Signatures[i].ValidatorAddress,
		})
	}

	if len(identities) == 0 {
		// No signatures in block
		projection.logger.Debugf("[Crossfire] non-critical error: no signatures in block: %s", event.BlockHeight)
		return nil
	}
	projection.profile("begin finding all block validators")
	validators, err := crossfireValidatorsView.FindAllBy(identityType, identities)
	if err != nil {
		return fmt.Errorf(
			"[Crossfire] error getting existing validators by block commitments: %v", err,
		)
	}
	projection.profile("end finding all block validators")

	validatorsMap := make(map[string]*view.CrossfireValidatorRow)
	for i, validator := range validators {
		validatorsMap[validator.TendermintAddress] = &validators[i]
	}

	for _, signature := range event.Block.Signatures {
		validator, exist := validatorsMap[signature.ValidatorAddress]
		if !exist {
			projection.logger.Errorf("[Crossfire] non-fatal error updating validator commitment number: not found, %s", signature.ValidatorAddress)
			continue
		}
		if blockTime.AfterOrEqual(projection.phaseOneStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE_1N2_COMMIT_PREFIX)
			keysToIncrementByOne = append(keysToIncrementByOne, key)
			//if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
			//	return fmt.Errorf(
			//		"[Crossfire] error increment validator commitment count %s", key,
			//	)
			//}
		}

		if blockTime.AfterOrEqual(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE_2_COMMIT_PREFIX)
			keysToIncrementByOne = append(keysToIncrementByOne, key)
			//if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
			//	return fmt.Errorf(
			//		"[Crossfire] error increment validator commitment count %s", key,
			//	)
			//}

			////Check task network upgrade
			//errUpdatingTask := projection.checkTaskNetworkUpgrade(crossfireValidatorsView, crossfireChainStatsView, validator, blockTime)
			//if errUpdatingTask != nil {
			//	return fmt.Errorf(
			//		"[Crossfire] error Updating crossfire task completion %s", errUpdatingTask,
			//	)
			//}
		}
		if blockTime.AfterOrEqual(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
			key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE_3_COMMIT_PREFIX)
			keysToIncrementByOne = append(keysToIncrementByOne, key)
			//if err := crossfireValidatorsStatsView.IncrementOne(key); err != nil {
			//	return fmt.Errorf(
			//		"[Crossfire] error increment validator commitment count %s", key,
			//	)
			//}
		}
	}
	if len(keysToIncrementByOne) == 0 {
		// No signatures in block
		projection.logger.Debugf("[Crossfire] non-critical error: no validators need to increment commitment count in block height, perhaps competition not started?: %s", event.BlockHeight)
		return nil
	}
	projection.profile("begin persisting validator commitments")
	if err := crossfireValidatorsStatsView.IncrementAllByOne(keysToIncrementByOne); err != nil {
		return fmt.Errorf(
			"[Crossfire] error increment validator commitment counts: %v", err,
		)
	}
	projection.profile("end persisting validator commitments")

	for _, signature := range event.Block.Signatures {
		validator, exist := validatorsMap[signature.ValidatorAddress]
		if !exist {
			projection.logger.Errorf("[Crossfire] non-fatal error preparing to checkTaskNetworkUpgrade: not found, %s", signature.ValidatorAddress)
			continue
		}

		if blockTime.AfterOrEqual(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
			//Check task network upgrade
			errUpdatingTask := projection.checkTaskNetworkUpgrade(crossfireValidatorsView, crossfireChainStatsView, validator, blockTime)
			if errUpdatingTask != nil {
				return fmt.Errorf(
					"[Crossfire] error Updating crossfire task completion %s", errUpdatingTask,
				)
			}
		}
	}

	projection.profile("end looping block validator commitments")

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
			projection.logger.Info("[Crossfire] handling MsgCreateValidator")

			pubKey, err := base64.StdEncoding.DecodeString(msgCreateValidatorEvent.TendermintPubkey)
			if err != nil {
				return fmt.Errorf("[Crossfire] error base64 decoding Tendermint node pubkey: %v", err)
			}
			tendermintAddress := tmcosmosutils.TmAddressFromTmPubKey(pubKey)

			consensusNodeAddress, err := tmcosmosutils.ConsensusNodeAddressFromTmPubKey(
				projection.conNodeAddressPrefix, pubKey,
			)
			if err != nil {
				return fmt.Errorf("[Crossfire] error converting consensus node pubkey to address: %v", err)
			}

			if markErr := crossfireValidatorsView.MarkOldValidatorSecondary(msgCreateValidatorEvent.ValidatorAddress); markErr != nil {
				return fmt.Errorf("[Crossfire] error marking old validator as secondary into view: %v", err)
			}

			validatorRow := view.CrossfireValidatorRow{
				ConsensusNodeAddress:            consensusNodeAddress,
				OperatorAddress:                 msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress:         msgCreateValidatorEvent.DelegatorAddress,
				TendermintPubkey:                msgCreateValidatorEvent.TendermintPubkey,
				TendermintAddress:               tendermintAddress,
				Status:                          constants.PRIMARY,
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
				return fmt.Errorf("[Crossfire] error querying validator last joined block height: %v", err)
			}
			if isJoined {
				validatorRow.JoinedAtBlockHeight = joinedAtBlockHeight
				validatorRow.JoinedAtBlockTime = joinedAtBlockTime
			}

			if err := crossfireValidatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("[Crossfire] error inserting new validator into view: %v", err)
			}

			// checkTaskSetup
			if err := projection.checkTaskSetup(
				crossfireValidatorsView,
				validatorRow.OperatorAddress,
				validatorRow.ConsensusNodeAddress,
				validatorRow.JoinedAtBlockTime,
			); err != nil {
				return fmt.Errorf("[Crossfire] error check Setup task for new validator: %v", err)
			}
		}
	}
	for _, event := range events {
		if msgSubmitSoftwareUpgradeProposalEvent, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			projection.logger.Debug("[Crossfire] handling MsgSubmitSoftwareUpgradeProposal event")

			// Check if proposed after competition has ended
			if blockTime.AfterOrEqual(projection.competitionEndTime) {
				projection.logger.Debug("[Crossfire] error Competition has already ended")
				continue
			}

			// Check if proposed before OR after Phase 2
			if blockTime.Before(projection.phaseTwoStartTime) || blockTime.AfterOrEqual(projection.phaseThreeStartTime) {
				projection.logger.Debug("[Crossfire] error This proposal does not occur in Phase 2")
				continue
			}

			// Check if proposed by NOT an admin
			if msgSubmitSoftwareUpgradeProposalEvent.ProposerAddress != projection.adminAddress {
				projection.logger.Debug("[Crossfire] error checking proposer address equals admin address")
				continue
			}

			// Check if proposal ID does not match the required ID
			if msgSubmitSoftwareUpgradeProposalEvent.MaybeProposalId != nil && *msgSubmitSoftwareUpgradeProposalEvent.MaybeProposalId != projection.networkUpgradeProposalID {
				projection.logger.Debug("[Crossfire] error checking Proposal ID in proposal")
				continue
			}

			networkUpgradeTimestamp := msgSubmitSoftwareUpgradeProposalEvent.MsgSubmitSoftwareUpgradeProposalParams.Content.Plan.Time
			networkUpgradeBlockheight := msgSubmitSoftwareUpgradeProposalEvent.MsgSubmitSoftwareUpgradeProposalParams.Content.Plan.Height

			// Update Network upgrade target Timestamp in DB
			networkUpgradeTimestampDBKey := constants.NETWORK_UPGRADE_TARGET_TIMESTAMP_KEY()
			errTSUpdate := crossfireChainStatsView.Set(networkUpgradeTimestampDBKey, networkUpgradeTimestamp.UnixNano())
			if errTSUpdate != nil {
				return fmt.Errorf("[Crossfire] error updating network_upgrade timestamp: %v", errTSUpdate)
			}

			// Update Network upgrade target Blockheight in DB
			networkUpgradeBlockheightDBKey := constants.NETWORK_UPGRADE_TARGET_BLOCKHEIGHT_KEY()
			errBlockheightUpdate := crossfireChainStatsView.Set(networkUpgradeBlockheightDBKey, networkUpgradeBlockheight)

			if errBlockheightUpdate != nil {
				return fmt.Errorf("[Crossfire] error updating network_upgrade blockheight: %v", errBlockheightUpdate)
			}

			projection.logger.Infof("[Crossfire] updated network upgrade proposal to id: %s, time: %s", *msgSubmitSoftwareUpgradeProposalEvent.MaybeProposalId, networkUpgradeTimestamp)
		} else if msgVoteCreatedEvent, ok := event.(*event_usecase.MsgVote); ok {
			projection.logger.Debug("[Crossfire] handling MsgVote event")

			// Check if proposed after competition has ended
			if blockTime.AfterOrEqual(projection.competitionEndTime) {
				projection.logger.Debug("[Crossfire] error Competition has already ended")
				continue
			}

			// Check if proposed before OR after Phase 2
			if blockTime.Before(projection.phaseTwoStartTime) || blockTime.AfterOrEqual(projection.phaseThreeStartTime) {
				projection.logger.Debug("[Crossfire] error Ineligible vote as it does not occur in Phase 2")
				continue
			}

			// Check if proposal ID does not match the required ID
			if msgVoteCreatedEvent.ProposalId != "" && msgVoteCreatedEvent.ProposalId != projection.networkUpgradeProposalID {
				projection.logger.Debug("[Crossfire] error checking Proposal ID in Vote not matching")
				continue
			}

			// Update the proposed ID against the voter in Database
			votedProposalIdDbKey := constants.VOTED_PROPOSAL_ID + constants.DB_KEY_SEPARATOR + msgVoteCreatedEvent.Voter

			proposalIdAsInt64, errConversion := strconv.ParseInt(msgVoteCreatedEvent.ProposalId, 10, 64)
			if errConversion != nil {
				return fmt.Errorf("[Crossfire] error converting ProposalID to int64: %v", errConversion)
			}
			errUpdateValidatorStats := crossfireValidatorStatsView.Set(votedProposalIdDbKey, proposalIdAsInt64)

			if errUpdateValidatorStats != nil {
				return fmt.Errorf("[Crossfire] error Updating ProposalID for the voter: %v", errUpdateValidatorStats)
			}
		} else if transactionCreatedEvent, ok := event.(*event_usecase.TransactionCreated); ok {
			// Check if proposed after competition has ended
			if blockTime.AfterOrEqual(projection.competitionEndTime) {
				return fmt.Errorf("[Crossfire] error Competition has already ended")
			}

			// Update the Tx Count
			errUpdateTxCount := projection.updateTxSentCount(crossfireValidatorStatsView, blockTime, transactionCreatedEvent)
			if errUpdateTxCount != nil {
				return fmt.Errorf("[Crossfire] error Updating tx sent count: %v", errUpdateTxCount)
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
	if joinedAtBlockTime.AfterOrEqual(projection.phaseOneStartTime) && joinedAtBlockTime.Before(projection.phaseTwoStartTime) {
		if err := crossfireValidatorsView.UpdateTask(
			"task_phase_1_node_setup",
			constants.COMPLETED,
			operatorAddress,
			consensusNodeAddress,
		); err != nil {
			return fmt.Errorf("error updating validator TaskPhase1NodeSetup as completed: %v", err)
		}
	}

	if joinedAtBlockTime.AfterOrEqual(projection.phaseTwoStartTime) {
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
	// This should only be called during phase 3, otherwise the default value of FindBy is 0 and
	// will make all nodes be treated as completed.
	phase2TotalBlockCount, err := crossfireChainStatsView.FindBy(constants.PHASE_2_BLOCK_COUNT)
	if err != nil {
		return fmt.Errorf("error getting phase 2 block count: %v", err)
	}

	validators, err := crossfireValidatorsView.List()
	if err != nil {
		return fmt.Errorf("error listing all validators for checkTaskKeepActive: %v", err)
	}

	projection.profile("begin checking alive task - looping validators")
	for _, validator := range validators {
		// If validator missed the task 1, it cannot complete the active task
		if validator.TaskPhase1NodeSetup == constants.MISSED {
			projection.logger.Debugf("[Crossfire] Won't handle task keepActive for %s because it missed taskSetup", validator.OperatorAddress)
			continue
		}

		key := constants.ValidatorCommitmentKey(validator.OperatorAddress, constants.PHASE_2_COMMIT_PREFIX)
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
	projection.profile("end checking alive task - looping validators")

	return nil
}

/*
// checkTaskNetworkProposalVote
func (projection *Crossfire) checkTaskNetworkProposalVote(
	crossfireValidatorsView *view.CrossfireValidators,
	voterAddress string,
	txBlockTime utctime.UTCTime,
) error {
	if txBlockTime.Before(projection.phaseThreeStartTime) && txBlockTime.AfterOrEqual(projection.phaseTwoStartTime) {
		if err := crossfireValidatorsView.UpdateTaskForOperatorAddress(
			constants.TASK_PHASE_2_PROPOSAL_VOTE_COLUMN_NAME,
			constants.COMPLETED,
			voterAddress,
		); err != nil {
			return fmt.Errorf("error updating validator Phase_2 Task_1 as completed %v", err)
		}
	}

	if txBlockTime.AfterOrEqual(projection.phaseThreeStartTime) {
		if err := crossfireValidatorsView.UpdateTaskForOperatorAddress(
			constants.TASK_PHASE_2_PROPOSAL_VOTE_COLUMN_NAME,
			constants.MISSED,
			voterAddress,
		); err != nil {
			return fmt.Errorf("error updating validator Phase_2 Task_1 as missed: %v", err)
		}
	}

	return nil
}*/

// Update Tx sent count for sender
func (projection *Crossfire) updateTxSentCount(
	crossfireValidatorStatsView *view.CrossfireValidatorsStats,
	blockTime utctime.UTCTime,
	msgTransactionCreated *event_usecase.TransactionCreated,
) error {
	var phaseNumberPrefix string
	var weeklyJackpotDBKey string
	if blockTime.AfterOrEqual(projection.phaseOneStartTime) && blockTime.Before(projection.phaseTwoStartTime) {
		phaseNumberPrefix = constants.PHASE_1_TX_SENT_PREFIX
	} else if blockTime.AfterOrEqual(projection.phaseTwoStartTime) && blockTime.Before(projection.phaseThreeStartTime) {
		phaseNumberPrefix = constants.PHASE_2_TX_SENT_PREFIX
	} else if blockTime.AfterOrEqual(projection.phaseThreeStartTime) && blockTime.Before(projection.competitionEndTime) {
		phaseNumberPrefix = constants.PHASE_3_TX_SENT_PREFIX
	}

	if blockTime.AfterOrEqual(projection.JackpotOneStartTime) && blockTime.Before(projection.JackpotTwoStartTime) {
		//Jackpot Week 1
		weeklyJackpotDBKey = constants.JACKPOT_1_TX_SENT_PREFIX
	} else if blockTime.AfterOrEqual(projection.JackpotTwoStartTime) && blockTime.Before(projection.JackpotThreeStartTime) {
		//Jackpot Week 2
		weeklyJackpotDBKey = constants.JACKPOT_2_TX_SENT_PREFIX

	} else if blockTime.AfterOrEqual(projection.JackpotThreeStartTime) && blockTime.Before(projection.JackpotFourStartTime) {
		//Jackpot Week 3
		weeklyJackpotDBKey = constants.JACKPOT_3_TX_SENT_PREFIX

	} else if blockTime.AfterOrEqual(projection.JackpotFourStartTime) && blockTime.Before(projection.JackpotFourEndTime) {
		//Jackpot Week 4
		weeklyJackpotDBKey = constants.JACKPOT_4_TX_SENT_PREFIX
	}

	if phaseNumberPrefix == "" {
		// competition not start yet
		return nil
	}

	for _, sender := range msgTransactionCreated.Senders {
		// Only considering Pubkey address for now
		if sender.Type == constants.TYPE_URL_PUBKEY && sender.MaybeThreshold == nil {
			pubKey, _ := base64.StdEncoding.DecodeString(sender.Pubkeys[0])
			// TODO: Support MultiSig address
			primaryAddress, err := tmcosmosutils.AccountAddressFromPubKey(projection.primaryAddressPrefix, pubKey)
			if err != nil {
				projection.logger.Errorf("[Crossfire] error getting account address from transaction sender pubkey: %v", err)
				continue
			}

			if phaseNumberPrefix != "" {
				// Increment count for address as per PHASE
				phaseAddressCountDbKey := phaseNumberPrefix + constants.DB_KEY_SEPARATOR + primaryAddress
				errIncrementing := crossfireValidatorStatsView.IncrementOne(phaseAddressCountDbKey)
				if errIncrementing != nil {
					return fmt.Errorf("error Phase wise tx sent count increment: %v", errIncrementing)
				}

				// Increment TOTAL count for address
				totalAddressCountDbKey := constants.TOTAL_TX_SENT_PREFIX + constants.DB_KEY_SEPARATOR + primaryAddress
				errIncrementingTotal := crossfireValidatorStatsView.IncrementOne(totalAddressCountDbKey)
				if errIncrementingTotal != nil {
					return fmt.Errorf("[Crossfire] error Incrementing tx sent count: %v", errIncrementingTotal)
				}
			}

			//Increment count for Weekly Jackpot (If Applicable)
			if weeklyJackpotDBKey != "" {
				key := fmt.Sprintf("%s%s%s", weeklyJackpotDBKey, constants.DB_KEY_SEPARATOR, primaryAddress)
				errIncrementingJackpotTxSentCount := crossfireValidatorStatsView.IncrementOne(key)
				if errIncrementingJackpotTxSentCount != nil {
					return fmt.Errorf("[Crossfire] error Incrementing Weekly tx sent count: %v", errIncrementingJackpotTxSentCount)
				}
			}
		}
	}
	return nil
}

func (projection *Crossfire) computeCommitmentRank(
	rankKey string,
	crossfireValidatorStatsPrefix string,
	crossfireValidatorStatsView *view.CrossfireValidatorsStats,
	crossfireValidatorView *view.CrossfireValidators,
) error {
	participants, err := projection.Client.Participants()
	if err != nil {
		return fmt.Errorf("error getting participants list %v", err)
	}

	participantsMap := make(map[string]bool)
	for _, participant := range *participants {
		participantsMap[participant.OperatorAddress] = true
	}

	projection.profile("begin filtering participants commitment records")
	validatorCommits, err := crossfireValidatorStatsView.FindAllLike(crossfireValidatorStatsPrefix)
	if err != nil {
		return fmt.Errorf("error getting validators' %s number from stats %v", crossfireValidatorStatsPrefix, err)
	}

	// filter out validators that not on list and missed taskSetup
	var participatedValidatorCommits []view.CrossfireValidatorsStatsRow
	for _, validatorCommit := range validatorCommits {
		statsOperatorAddress := strings.Split(validatorCommit.Key, constants.DB_KEY_SEPARATOR)[1]

		// check whether validator missed taskSetup or not
		validatorRow, err := crossfireValidatorView.FindBy(view.CrossfireValidatorIdentity{
			MaybeOperatorAddress: &statsOperatorAddress,
		})
		if errors.Is(err, rdb.ErrNoRows) {
			// no validator found by tendermint address
			projection.logger.Error("[Crossfire] error finding participant when computing commitment rank: record not found")
			continue
		}
		if err != nil {
			return fmt.Errorf("[Crossfire] error getting validator with operator addr %s %v", statsOperatorAddress, err)
		}
		if validatorRow.TaskPhase1NodeSetup == constants.MISSED || validatorRow.TaskPhase1NodeSetup == constants.INCOMPLETED {
			projection.logger.Debugf("[Crossfire] Won't handle commit rank for %s because it missed taskSetup", statsOperatorAddress)
			continue
		}

		// check whether validator is on list or not
		if participantsMap[statsOperatorAddress] {
			participatedValidatorCommits = append(participatedValidatorCommits, validatorCommit)
		}
	}
	projection.profile("end filtering participants commitment records")

	projection.profile("begin sorting participants commitment records")
	sort.SliceStable(participatedValidatorCommits, func(i, j int) bool {
		return participatedValidatorCommits[i].Value > participatedValidatorCommits[j].Value
	})

	rank := 1
	for index, validatorCommit := range participatedValidatorCommits {
		// current value smaller than previous value, increment rank
		if index > 0 && validatorCommit.Value != participatedValidatorCommits[index-1].Value {
			rank++
		}

		statsOperatorAddress := strings.Split(validatorCommit.Key, constants.DB_KEY_SEPARATOR)[1]
		if err := crossfireValidatorView.UpdateRank(
			rankKey,
			int64(rank),
			statsOperatorAddress,
		); err != nil {
			return fmt.Errorf("error UpdateRank for %s %v", rankKey, err)
		}
	}
	projection.profile("end sorting participants commitment records")

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
		return fmt.Errorf("[Crossfire] error getting Participants list %w", errGetRequest)
	}

	participantsMap := make(map[string]Participant)
	for _, participant := range *participantsList {
		participantsMap[participant.PrimaryAddress] = participant
	}

	projection.profile("begin getting all transaction sent records")
	// Get All Tx Count Sorted
	dbTxSendersWithCountList, errDbCount := crossfireValidatorStatsView.FindAllLike(constants.TOTAL_TX_SENT_PREFIX)
	if errDbCount != nil {
		return fmt.Errorf("[Crossfire] error Database Participant With Total Count %w", errDbCount)
	}
	if len(dbTxSendersWithCountList) < 1 {
		// No Participants to Rank
		return nil
	}
	projection.profile("end getting all transaction sent records")

	projection.profile("begin filtering participants transaction records")
	// Filter only participants from tx senders list
	dbParticipantWithCountList := make([]view.CrossfireValidatorsStatsRow, 0)
	for _, dbTxSender := range dbTxSendersWithCountList {
		dbParticipantPrimaryAddress := strings.Split(dbTxSender.Key, constants.DB_KEY_SEPARATOR)[1]
		if _, ok := participantsMap[dbParticipantPrimaryAddress]; ok {
			dbParticipantWithCountList = append(dbParticipantWithCountList, dbTxSender)
		}
	}
	projection.profile("end filtering participants transaction records")

	projection.profile("begin sorting participants transaction records")
	// Sort With Most Tx Sent First
	sort.SliceStable(dbParticipantWithCountList, func(i, j int) bool {
		return dbParticipantWithCountList[i].Value > dbParticipantWithCountList[j].Value
	})

	var rank int = 1
	for index, dbParticipant := range dbParticipantWithCountList {
		dbParticipantPrimaryAddress := strings.Split(dbParticipant.Key, constants.DB_KEY_SEPARATOR)[1]

		participant := participantsMap[dbParticipantPrimaryAddress]
		// Check if Participant is a Registered Crossfire Validator
		_, errFindingValidator := crossfireValidatorView.FindBy(view.CrossfireValidatorIdentity{
			MaybeOperatorAddress: &participant.OperatorAddress,
		})
		if errors.Is(errFindingValidator, rdb.ErrNoRows) {
			// no validator found by tendermint address
			projection.logger.Errorf("[Crossfire] error finding participant when computing tx rank: record not found: %s", participant.OperatorAddress)
			continue
		}
		if errFindingValidator != nil {
			return fmt.Errorf("[Crossfire] error getting validator when computing tx rank with operator addr %s %v", participant.OperatorAddress, errFindingValidator)
		}

		// Update Ranks with specific Validator
		errUpdating := crossfireValidatorView.UpdateTxSentRank(rank, participant.OperatorAddress)
		if errUpdating != nil {
			return fmt.Errorf("[Crossfire] error updating transaction sent Task Rank %w", errUpdating)
		}
		if index+1 < len(dbParticipantWithCountList) && dbParticipantWithCountList[index].Value != dbParticipantWithCountList[index+1].Value {
			rank++
		}
	}
	projection.profile("end sorting participants transaction records")
	return nil
}

// checkTaskNetworkUpgrade
func (projection *Crossfire) checkTaskNetworkUpgrade(
	crossfireValidatorsView *view.CrossfireValidators,
	crossfireChainStatsView *view.CrossfireChainStats,
	validator *view.CrossfireValidatorRow,
	blockTime utctime.UTCTime,
) error {

	// Check if Validator's upgrade is already successful
	if validator.TaskPhase2NetworkUpgrade == constants.COMPLETED {
		return nil
	}

	targetTimestampDBKey := constants.NETWORK_UPGRADE_TARGET_TIMESTAMP_KEY()

	networkUpgradeTimestampNanoSec, errTimestamp := crossfireChainStatsView.FindBy(targetTimestampDBKey)
	if errTimestamp != nil {
		return fmt.Errorf("[Crossfire] error getting network Upgrade timestamp: %v", errTimestamp)
	}
	if networkUpgradeTimestampNanoSec == 0 {
		return nil
	}

	// Check if current block is before the network upgrade
	if blockTime.Before(utctime.FromUnixNano(networkUpgradeTimestampNanoSec)) {
		return nil
	}
	errUpdatingTaskCompletion := crossfireValidatorsView.UpdateTaskForOperatorAddress("task_phase_2_network_upgrade", constants.COMPLETED, validator.OperatorAddress)
	if errUpdatingTaskCompletion != nil {
		return fmt.Errorf("[Crossfire] error Updating Task completion: %v", errUpdatingTaskCompletion)
	}

	return nil
}

func (projection *Crossfire) profile(label string) {
	projection.logger.Debugf("[Crossfire] [Profile] %s: %s", label, utctime.Now())
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
