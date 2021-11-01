package proposal

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/projection/proposal/types"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase"
	rdbparambase_types "github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/types"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbvalidatorbase"
	validatorbase_view "github.com/crypto-com/chain-indexing/appinterface/projection/rdbvalidatorbase/view"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/projection/proposal/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &Proposal{}

var (
	NewProposals       = view.NewProposalsView
	NewVotes           = view.NewVotesView
	NewVotesTotal      = view.NewVotesTotalView
	NewDepositors      = view.NewDepositorsView
	NewDepositorsTotal = view.NewDepositorsTotalView

	UpdateLastHandledEventHeight = (*Proposal).UpdateLastHandledEventHeight

	ParamBaseHandleEvents     = (*rdbparambase.Base).HandleEvents
	ValidatorBaseHandleEvents = (*rdbvalidatorbase.Base).HandleEvents
	ParamBaseGetView          = (*rdbparambase.Base).GetView
	ValidatorBaseGetView      = (*rdbvalidatorbase.Base).GetView
)

type Proposal struct {
	*rdbprojectionbase.Base
	paramBase     *rdbparambase.Base
	validatorBase *rdbvalidatorbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	conNodeAddressPrefix string
}

func NewProposal(logger applogger.Logger, rdbConn rdb.Conn, conNodeAddressPrefix string) *Proposal {
	return &Proposal{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Proposal"),

		rdbparambase.NewBase(view.PARAMS_TABLE_NAME, []rdbparambase_types.ParamAccessor{{
			Module: "gov",
			Key:    "max_deposit_period",
		}, {
			Module: "gov",
			Key:    "voting_period",
		}, {
			Module: "gov",
			Key:    "quorum",
		}, {
			Module: "gov",
			Key:    "threshold",
		}, {
			Module: "gov",
			Key:    "veto_threshold",
		}}),
		rdbvalidatorbase.NewBase(view.VALIDATORS_TABLE_NAME, conNodeAddressPrefix),

		rdbConn,
		logger,
		conNodeAddressPrefix,
	}
}

func (proposal *Proposal) GetEventsToListen() []string {
	return append(
		append(
			[]string{
				event_usecase.BLOCK_CREATED,
				event_usecase.MSG_SUBMIT_TEXT_PROPOSAL_CREATED,
				event_usecase.MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED,
				event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED,
				event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED,
				event_usecase.MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED,
				event_usecase.PROPOSAL_VOTING_PERIOD_STARTED,
				event_usecase.PROPOSAL_INACTIVED,
				event_usecase.PROPOSAL_ENDED,
				event_usecase.MSG_DEPOSIT_CREATED,
				event_usecase.MSG_VOTE_CREATED,
			},
			proposal.paramBase.GetEventsToListen()...,
		),
		proposal.validatorBase.GetEventsToListen()...,
	)
}

func (_ *Proposal) OnInit() error {
	return nil
}

func (projection *Proposal) HandleEvents(height int64, events []event_entity.Event) error {
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

	if err := ParamBaseHandleEvents(projection.paramBase, rdbTxHandle, projection.logger, events); err != nil {
		return fmt.Errorf("error handling event in param base: %v", err)
	}
	if err := ValidatorBaseHandleEvents(projection.validatorBase, rdbTxHandle, projection.logger, events); err != nil {
		return fmt.Errorf("error handling event in validator base: %v", err)
	}

	proposalsView := NewProposals(rdbTxHandle)

	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	for _, event := range events {
		if msgSubmitProposal, ok := event.(*event_usecase.MsgSubmitTextProposal); ok {
			context, err := projection.prepareNewProposalSubmissionContext(rdbTxHandle, msgSubmitProposal.ProposerAddress)
			if err != nil {
				return err
			}

			depositEndTime := blockTime.Add(context.maxDepositPeriod)
			row := view.ProposalRow{
				ProposalId:                   *msgSubmitProposal.MaybeProposalId,
				Title:                        msgSubmitProposal.Content.Title,
				Description:                  msgSubmitProposal.Content.Description,
				Type:                         msgSubmitProposal.Content.Type,
				Status:                       view.PROPOSAL_STATUS_DEPOSIT_PERIOD,
				ProposerAddress:              msgSubmitProposal.ProposerAddress,
				MaybeProposerOperatorAddress: context.maybeProposerValidatorAddress,
				Data:                         nil,
				InitialDeposit:               msgSubmitProposal.InitialDeposit,
				TotalDeposit:                 msgSubmitProposal.InitialDeposit,
				TotalVote:                    big.NewInt(0),
				TransactionHash:              msgSubmitProposal.TxHash(),
				SubmitBlockHeight:            height,
				SubmitTime:                   blockTime,
				DepositEndTime:               depositEndTime,
				MaybeVotingStartTime:         nil,
				MaybeVotingEndTime:           nil,
				MaybeVotingEndBlockHeight:    nil,
			}

			if insertProposalErr := proposalsView.Insert(&row); insertProposalErr != nil {
				return fmt.Errorf("error inserting  text proposal proposal into view: %v", insertProposalErr)
			}

			maybeDepositorValidatorAddress := context.maybeProposerValidatorAddress

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    *msgSubmitProposal.MaybeProposalId,
				DepositorAddress:              msgSubmitProposal.ProposerAddress,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               msgSubmitProposal.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        msgSubmitProposal.InitialDeposit,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting proposer deposit record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				*msgSubmitProposal.MaybeProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting proposer deposit total record into view: %v", updateDepositorTotalErr)
			}

		} else if msgSubmitProposal, ok := event.(*event_usecase.MsgSubmitParamChangeProposal); ok {
			context, err := projection.prepareNewProposalSubmissionContext(rdbTxHandle, msgSubmitProposal.ProposerAddress)
			if err != nil {
				return err
			}

			depositEndTime := blockTime.Add(context.maxDepositPeriod)
			row := view.ProposalRow{
				ProposalId:                   *msgSubmitProposal.MaybeProposalId,
				Title:                        msgSubmitProposal.Content.Title,
				Description:                  msgSubmitProposal.Content.Description,
				Type:                         msgSubmitProposal.Content.Type,
				Status:                       view.PROPOSAL_STATUS_DEPOSIT_PERIOD,
				ProposerAddress:              msgSubmitProposal.ProposerAddress,
				MaybeProposerOperatorAddress: context.maybeProposerValidatorAddress,
				Data:                         msgSubmitProposal.Content.Changes,
				InitialDeposit:               msgSubmitProposal.InitialDeposit,
				TotalDeposit:                 msgSubmitProposal.InitialDeposit,
				TotalVote:                    big.NewInt(0),
				TransactionHash:              msgSubmitProposal.TxHash(),
				SubmitBlockHeight:            height,
				SubmitTime:                   blockTime,
				DepositEndTime:               depositEndTime,
				MaybeVotingStartTime:         nil,
				MaybeVotingEndTime:           nil,
				MaybeVotingEndBlockHeight:    nil,
			}

			if insertProposalErr := proposalsView.Insert(&row); insertProposalErr != nil {
				return fmt.Errorf("error inserting param change proposal into view: %v", insertProposalErr)
			}

			maybeDepositorValidatorAddress := context.maybeProposerValidatorAddress

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    *msgSubmitProposal.MaybeProposalId,
				DepositorAddress:              msgSubmitProposal.ProposerAddress,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               msgSubmitProposal.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        msgSubmitProposal.InitialDeposit,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting proposer deposit record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				*msgSubmitProposal.MaybeProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting proposer deposit total record into view: %v", updateDepositorTotalErr)
			}

		} else if msgSubmitProposal, ok := event.(*event_usecase.MsgSubmitCommunityPoolSpendProposal); ok {
			context, err := projection.prepareNewProposalSubmissionContext(rdbTxHandle, msgSubmitProposal.ProposerAddress)
			if err != nil {
				return err
			}

			depositEndTime := blockTime.Add(context.maxDepositPeriod)
			row := view.ProposalRow{
				ProposalId:                   *msgSubmitProposal.MaybeProposalId,
				Title:                        msgSubmitProposal.Content.Title,
				Description:                  msgSubmitProposal.Content.Description,
				Type:                         msgSubmitProposal.Content.Type,
				Status:                       view.PROPOSAL_STATUS_DEPOSIT_PERIOD,
				ProposerAddress:              msgSubmitProposal.ProposerAddress,
				MaybeProposerOperatorAddress: context.maybeProposerValidatorAddress,
				Data: types.CommunityPoolSpendData{
					RecipientAddress: msgSubmitProposal.Content.RecipientAddress,
					Amount:           msgSubmitProposal.Content.Amount,
				},
				InitialDeposit:            msgSubmitProposal.InitialDeposit,
				TotalDeposit:              msgSubmitProposal.InitialDeposit,
				TotalVote:                 big.NewInt(0),
				TransactionHash:           msgSubmitProposal.TxHash(),
				SubmitBlockHeight:         height,
				SubmitTime:                blockTime,
				DepositEndTime:            depositEndTime,
				MaybeVotingStartTime:      nil,
				MaybeVotingEndTime:        nil,
				MaybeVotingEndBlockHeight: nil,
			}

			if insertProposalErr := proposalsView.Insert(&row); insertProposalErr != nil {
				return fmt.Errorf("error inserting param change proposal into view: %v", insertProposalErr)
			}

			maybeDepositorValidatorAddress := context.maybeProposerValidatorAddress

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    *msgSubmitProposal.MaybeProposalId,
				DepositorAddress:              msgSubmitProposal.ProposerAddress,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               msgSubmitProposal.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        msgSubmitProposal.InitialDeposit,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting proposer deposit record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				*msgSubmitProposal.MaybeProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting proposer deposit total record into view: %v", updateDepositorTotalErr)
			}

		} else if msgSubmitProposal, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			context, err := projection.prepareNewProposalSubmissionContext(rdbTxHandle, msgSubmitProposal.ProposerAddress)
			if err != nil {
				return err
			}

			depositEndTime := blockTime.Add(context.maxDepositPeriod)
			row := view.ProposalRow{
				ProposalId:                   *msgSubmitProposal.MaybeProposalId,
				Title:                        msgSubmitProposal.Content.Title,
				Description:                  msgSubmitProposal.Content.Description,
				Type:                         msgSubmitProposal.Content.Type,
				Status:                       view.PROPOSAL_STATUS_DEPOSIT_PERIOD,
				ProposerAddress:              msgSubmitProposal.ProposerAddress,
				MaybeProposerOperatorAddress: context.maybeProposerValidatorAddress,
				Data:                         msgSubmitProposal.Content.Plan,
				InitialDeposit:               msgSubmitProposal.InitialDeposit,
				TotalDeposit:                 msgSubmitProposal.InitialDeposit,
				TotalVote:                    big.NewInt(0),
				TransactionHash:              msgSubmitProposal.TxHash(),
				SubmitBlockHeight:            height,
				SubmitTime:                   blockTime,
				DepositEndTime:               depositEndTime,
				MaybeVotingStartTime:         nil,
				MaybeVotingEndTime:           nil,
				MaybeVotingEndBlockHeight:    nil,
			}

			if insertProposalErr := proposalsView.Insert(&row); insertProposalErr != nil {
				return fmt.Errorf("error inserting param software upgrade proposal into view: %v", insertProposalErr)
			}

			maybeDepositorValidatorAddress := context.maybeProposerValidatorAddress

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    *msgSubmitProposal.MaybeProposalId,
				DepositorAddress:              msgSubmitProposal.ProposerAddress,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               msgSubmitProposal.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        msgSubmitProposal.InitialDeposit,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting proposer deposit record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				*msgSubmitProposal.MaybeProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting proposer deposit total record into view: %v", updateDepositorTotalErr)
			}

		} else if msgSubmitProposal, ok := event.(*event_usecase.MsgSubmitCancelSoftwareUpgradeProposal); ok {
			context, err := projection.prepareNewProposalSubmissionContext(rdbTxHandle, msgSubmitProposal.ProposerAddress)
			if err != nil {
				return err
			}

			depositEndTime := blockTime.Add(context.maxDepositPeriod)
			row := view.ProposalRow{
				ProposalId:                   *msgSubmitProposal.MaybeProposalId,
				Title:                        msgSubmitProposal.Content.Title,
				Description:                  msgSubmitProposal.Content.Description,
				Type:                         msgSubmitProposal.Content.Type,
				Status:                       view.PROPOSAL_STATUS_DEPOSIT_PERIOD,
				ProposerAddress:              msgSubmitProposal.ProposerAddress,
				MaybeProposerOperatorAddress: context.maybeProposerValidatorAddress,
				Data:                         nil,
				InitialDeposit:               msgSubmitProposal.InitialDeposit,
				TotalDeposit:                 msgSubmitProposal.InitialDeposit,
				TotalVote:                    big.NewInt(0),
				TransactionHash:              msgSubmitProposal.TxHash(),
				SubmitBlockHeight:            height,
				SubmitTime:                   blockTime,
				DepositEndTime:               depositEndTime,
				MaybeVotingStartTime:         nil,
				MaybeVotingEndTime:           nil,
				MaybeVotingEndBlockHeight:    nil,
			}

			if insertProposalErr := proposalsView.Insert(&row); insertProposalErr != nil {
				return fmt.Errorf("error inserting cancel software upgrade proposal proposal into view: %v", insertProposalErr)
			}

			maybeDepositorValidatorAddress := context.maybeProposerValidatorAddress

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    *msgSubmitProposal.MaybeProposalId,
				DepositorAddress:              msgSubmitProposal.ProposerAddress,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               msgSubmitProposal.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        msgSubmitProposal.InitialDeposit,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting proposer deposit record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				*msgSubmitProposal.MaybeProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting proposer deposit total record into view: %v", updateDepositorTotalErr)
			}

		} else if proposalVotingPeriodStarted, ok := event.(*event_usecase.ProposalVotingPeriodStarted); ok {
			mutProposal, err := proposalsView.FindById(proposalVotingPeriodStarted.ProposalId)
			if err != nil {
				return fmt.Errorf("error finding proposal which voting period has started: %v", err)
			}

			paramsView := ParamBaseGetView(projection.paramBase, rdbTxHandle)
			votingPeriod, err := paramsView.FindDurationBy(rdbparambase_types.ParamAccessor{
				Module: "gov",
				Key:    "voting_period",
			})
			if err != nil {
				return fmt.Errorf("error retrieving voting_period param: %v", err)
			}

			votingEndTime := blockTime.Add(votingPeriod)
			mutProposal.Status = view.PROPOSAL_STATUS_VOTING_PERIOD
			mutProposal.MaybeVotingStartTime = &blockTime
			mutProposal.MaybeVotingEndTime = &votingEndTime
			if err := proposalsView.Update(&mutProposal.ProposalRow); err != nil {
				return fmt.Errorf("error updating proposal which voting period has started: %v", err)
			}

		} else if proposalInactived, ok := event.(*event_usecase.ProposalInactived); ok {
			mutProposal, err := proposalsView.FindById(proposalInactived.ProposalId)
			if err != nil {
				return fmt.Errorf("error finding proposal which becomes inactive: %v", err)
			}

			mutProposal.Status = view.PROPOSAL_STATUS_INACTIVE
			if err := proposalsView.Update(&mutProposal.ProposalRow); err != nil {
				return fmt.Errorf("error updating proposal which becomes inactive: %v", err)
			}

		} else if proposalEnded, ok := event.(*event_usecase.ProposalEnded); ok {
			mutProposal, err := proposalsView.FindById(proposalEnded.ProposalId)
			if err != nil {
				return fmt.Errorf("error finding proposal which has ended: %v", err)
			}

			if proposalEnded.Result == "proposal_passed" {
				mutProposal.Status = view.PROPOSAL_STATUS_PASSED
			} else if proposalEnded.Result == "proposal_failed" {
				mutProposal.Status = view.PROPOSAL_STATUS_FAILED
			} else if proposalEnded.Result == "proposal_rejected" {
				mutProposal.Status = view.PROPOSAL_STATUS_REJECTED
			} else {
				return fmt.Errorf("unrecognized proposal end status: %s", proposalEnded.Result)
			}
			mutProposal.MaybeVotingEndBlockHeight = &height
			if err := proposalsView.Update(&mutProposal.ProposalRow); err != nil {
				return fmt.Errorf("error updating proposal which has ended: %v", err)
			}

		} else if deposit, ok := event.(*event_usecase.MsgDeposit); ok {
			mutProposal, queryProposalErr := proposalsView.FindById(deposit.ProposalId)
			if queryProposalErr != nil {
				return fmt.Errorf("error querying proposal which has deposit: %v", queryProposalErr)
			}

			mutProposal.TotalDeposit = mutProposal.TotalDeposit.Add(deposit.Amount...)
			if updateProposalErr := proposalsView.Update(&mutProposal.ProposalRow); updateProposalErr != nil {
				return fmt.Errorf("error updating proposal which has deposit: %v", updateProposalErr)
			}

			validatorsView := ValidatorBaseGetView(projection.validatorBase, rdbTxHandle)
			maybeDepositorValidatorRow, err := validatorsView.FindLastBy(validatorbase_view.ValidatorIdentity{
				MaybeInititalDelegatorAddress: &deposit.Depositor,
			})
			var maybeDepositorValidatorAddress *string
			if err != nil {
				if !errors.Is(err, rdb.ErrNoRows) {
					return fmt.Errorf("error querying depositor validator address: %v", err)
				}
			} else {
				maybeDepositorValidatorAddress = &maybeDepositorValidatorRow.OperatorAddress
			}

			depositorsView := NewDepositors(rdbTxHandle)
			depositorsTotalView := NewDepositorsTotal(rdbTxHandle)
			if insertDepositorErr := depositorsView.Insert(&view.DepositorRow{
				ProposalId:                    deposit.ProposalId,
				DepositorAddress:              deposit.Depositor,
				MaybeDepositorOperatorAddress: maybeDepositorValidatorAddress,
				TransactionHash:               deposit.TxHash(),
				DepositAtBlockHeight:          height,
				DepositAtBlockTime:            blockTime,
				Amount:                        deposit.Amount,
			}); insertDepositorErr != nil {
				return fmt.Errorf("error inserting depositor record into view: %v", insertDepositorErr)
			}
			if updateDepositorTotalErr := depositorsTotalView.Increment(
				deposit.ProposalId, 1,
			); updateDepositorTotalErr != nil {
				return fmt.Errorf("error inserting depositor total record into view: %v", updateDepositorTotalErr)
			}

		} else if vote, ok := event.(*event_usecase.MsgVote); ok {
			validatorsView := ValidatorBaseGetView(projection.validatorBase, rdbTxHandle)
			var maybeVoterOperatorAddress *string
			maybeVoterValidatorRow, err := validatorsView.FindLastBy(validatorbase_view.ValidatorIdentity{
				MaybeInititalDelegatorAddress: &vote.Voter,
			})
			if err != nil {
				if !errors.Is(err, rdb.ErrNoRows) {
					return fmt.Errorf("error querying voter validator address: %v", err)
				}
			} else {
				maybeVoterOperatorAddress = &maybeVoterValidatorRow.OperatorAddress
			}

			votesView := NewVotes(rdbTxHandle)

			mutVoteRow, queryExistingVoteRowErr := votesView.FindByProposalIdVoter(vote.ProposalId, vote.Voter)
			if queryExistingVoteRowErr != nil {
				if !errors.Is(queryExistingVoteRowErr, rdb.ErrNoRows) {
					return fmt.Errorf(
						"error finding voter record with same proposal id and voter: %v",
						queryExistingVoteRowErr,
					)
				}

				// vote record does not exists
				if insertVoteErr := votesView.Insert(&view.VoteRow{
					ProposalId:                vote.ProposalId,
					VoterAddress:              vote.Voter,
					MaybeVoterOperatorAddress: maybeVoterOperatorAddress,
					TransactionHash:           vote.TxHash(),
					VoteAtBlockHeight:         height,
					VoteAtBlockTime:           blockTime,
					Answer:                    vote.Option,
					Histories:                 make([]view.VoteHistory, 0),
				}); insertVoteErr != nil {
					return fmt.Errorf("error inserting vote record to view: %v", insertVoteErr)
				}

				mutProposal, queryVotedProposalErr := proposalsView.FindById(vote.ProposalId)
				if queryVotedProposalErr != nil {
					return fmt.Errorf("error querying proposal which has new vote: %v", queryVotedProposalErr)
				}

				mutProposal.TotalVote = new(big.Int).Add(mutProposal.TotalVote, new(big.Int).SetInt64(int64(1)))
				if updateProposalErr := proposalsView.Update(&mutProposal.ProposalRow); updateProposalErr != nil {
					return fmt.Errorf("error updating proposal which has new vote: %v", updateProposalErr)
				}

				votesTotalView := NewVotesTotal(rdbTxHandle)
				if updateVoteTotalErr := votesTotalView.Increment(vote.ProposalId, 1); updateVoteTotalErr != nil {
					return fmt.Errorf("error updating votes total record: %v", updateVoteTotalErr)
				}

			} else {
				// vote record already exists
				mutVoteRow.Histories = append(mutVoteRow.Histories, view.VoteHistory{
					TransactionHash:   mutVoteRow.TransactionHash,
					VoteAtBlockHeight: mutVoteRow.VoteAtBlockHeight,
					VoteAtBlockTime:   mutVoteRow.VoteAtBlockTime,
					Answer:            mutVoteRow.Answer,
				})
				mutVoteRow.TransactionHash = vote.TxHash()
				mutVoteRow.VoteAtBlockHeight = height
				mutVoteRow.VoteAtBlockTime = blockTime
				mutVoteRow.Answer = vote.Option

				if updateVoteErr := votesView.Update(&mutVoteRow.VoteRow); updateVoteErr != nil {
					return fmt.Errorf("error updating existing vote record: %v", updateVoteErr)
				}
			}
		}
	}

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (projection *Proposal) prepareNewProposalSubmissionContext(
	rdbTxHandle *rdb.Handle,
	proposerAddress string,
) (newProposalSubmissionContext, error) {
	validatorsView := ValidatorBaseGetView(projection.validatorBase, rdbTxHandle)
	maybeProposerValidatorRow, err := validatorsView.FindLastBy(validatorbase_view.ValidatorIdentity{
		MaybeInititalDelegatorAddress: &proposerAddress,
	})
	var maybeProposerValidatorAddress *string
	if err != nil {
		if !errors.Is(err, rdb.ErrNoRows) {
			return newProposalSubmissionContext{}, fmt.Errorf("error querying proposer validator address: %v", err)
		}
	} else {
		maybeProposerValidatorAddress = &maybeProposerValidatorRow.OperatorAddress
	}

	paramsView := ParamBaseGetView(projection.paramBase, rdbTxHandle)
	maxDepositPeriod, err := paramsView.FindDurationBy(rdbparambase_types.ParamAccessor{
		Module: "gov",
		Key:    "max_deposit_period",
	})
	if err != nil {
		return newProposalSubmissionContext{}, fmt.Errorf("error retrieving max_deposit_period param: %v", err)
	}

	return newProposalSubmissionContext{
		maybeProposerValidatorAddress,
		maxDepositPeriod,
	}, nil
}

type newProposalSubmissionContext struct {
	maybeProposerValidatorAddress *string
	maxDepositPeriod              time.Duration
}
