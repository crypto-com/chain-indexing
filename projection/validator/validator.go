package validator

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/validator/constants"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/projection/validator/view"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &Validator{}

var (
	NewValidators                     = view.NewValidatorsView
	NewValidatorActivities            = view.NewValidatorActivitiesView
	NewValidatorActivitiesTotal       = view.NewValidatorActivitiesTotalView
	NewValidatorBlockCommitments      = view.NewValidatorBlockCommitments
	NewValidatorBlockCommitmentsTotal = view.NewValidatorBlockCommitmentsTotal
	UpdateLastHandledEventHeight      = (*Validator).UpdateLastHandledEventHeight
)

const DO_NOT_MODIFY = "[do-not-modify]"

type Validator struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	conNodeAddressPrefix string

	migrationHelper migrationhelper.MigrationHelper

	config *Config
}

type Config struct {
	AttentionStatusRules AttentionStatusRules `mapstructure:"attention_status_rules"`
}

type AttentionStatusRules struct {
	MaxCommissionRateChange MaxCommissionRateChange `mapstructure:"max_commission_rate_change"`
	MaxEditQuota            MaxEditQuota            `mapstructure:"max_edit_quota"`
}

type MaxCommissionRateChange struct {
	Enable    bool    `mapstructure:"enable"`
	MaxChange float64 `mapstructure:"max_change"`
}
type MaxEditQuota struct {
	Enable   bool           `mapstructure:"enable"`
	Quota    map[string]int `mapstructure:"quota"`
	Interval string         `mapstructure:"interval"`
}

func ConfigFromInterface(data interface{}) (Config, error) {
	config := Config{}

	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &config,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		return config, fmt.Errorf("error creating projection config decoder: %v", decoderErr)
	}

	if err := decoder.Decode(data); err != nil {
		return config, fmt.Errorf("error decoding projection ValidatorAttentionRules config: %v", err)
	}

	return config, nil
}

func NewValidator(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	conNodeAddressPrefix string,
	migrationHelper migrationhelper.MigrationHelper,
	config *Config,
) *Validator {
	return &Validator{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"Validator",
		),

		rdbConn,
		logger,
		conNodeAddressPrefix,
		migrationHelper,
		config,
	}
}

func (_ *Validator) GetEventsToListen() []string {
	return []string{
		event_usecase.GENESIS_CREATED,
		event_usecase.BLOCK_CREATED,
		event_usecase.GENESIS_VALIDATOR_CREATED,
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.MSG_EDIT_VALIDATOR_CREATED,
		event_usecase.MSG_DELEGATE_CREATED,
		event_usecase.MSG_BEGIN_REDELEGATE_CREATED,
		event_usecase.MSG_UNDELEGATE_CREATED,
		event_usecase.MSG_WITHDRAW_DELEGATOR_REWARD_CREATED,
		event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED,
		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.BLOCK_COMMISSIONED,
		event_usecase.VALIDATOR_JAILED,
		event_usecase.VALIDATOR_SLASHED,
		event_usecase.MSG_UNJAIL_CREATED,
		event_usecase.POWER_CHANGED,
		event_usecase.MSG_VOTE_CREATED,
	}
}

// const (
// 	MIGRATION_TABLE_NAME = "validator_schema_migrations"
// 	MIGRATION_DIRECOTRY  = "projection/validator/migrations"
// )

func (projection *Validator) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}

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
	validatorsView := NewValidators(rdbTxHandle)
	validatorBlockCommitmentsView := NewValidatorBlockCommitments(rdbTxHandle)
	validatorBlockCommitmentsTotalView := NewValidatorBlockCommitmentsTotal(rdbTxHandle)
	validatorActivitiesView := NewValidatorActivities(rdbTxHandle)
	validatorActivitiesTotalView := NewValidatorActivitiesTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	var blockProposer string
	for _, event := range events {
		if genesisEvent, ok := event.(*event_usecase.GenesisCreated); ok {
			blockTime = utctime.MustParse(time.RFC3339, genesisEvent.Genesis.GenesisTime)
			blockHash = "genesis"
		} else if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
			blockProposer = blockCreatedEvent.Block.ProposerAddress
		}
	}
	if projectErr := projection.projectValidatorView(validatorsView, &validatorActivitiesView, blockTime, height, events); projectErr != nil {
		return fmt.Errorf("error projecting validator view: %v", projectErr)
	}

	if projectErr := projection.projectValidatorActivitiesView(
		&validatorsView,
		&validatorActivitiesView,
		&validatorActivitiesTotalView,
		blockHash,
		blockTime,
		events,
	); projectErr != nil {
		return fmt.Errorf("error projecting validator activities view: %v", err)
	}

	validatorList, listValidatorErr := validatorsView.ListAll(view.ValidatorsListFilter{
		MaybeStatuses: nil,
	}, view.ValidatorsListOrder{MaybePower: nil})
	if listValidatorErr != nil {
		return fmt.Errorf("error retrieving validator list: %v", listValidatorErr)
	}
	validatorMap := make(map[string]*view.ValidatorRow)
	for i, validator := range validatorList {
		validatorMap[validator.TendermintAddress] = &validatorList[i]
	}

	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			signatureCount := len(blockCreatedEvent.Block.Signatures)
			commitmentRows := make([]view.ValidatorBlockCommitmentRow, 0, signatureCount)

			identities := make([]string, 0, len(blockCreatedEvent.Block.Signatures))
			heightValidatorIdentities := make([]string, 0, len(blockCreatedEvent.Block.Signatures))

			commitmentMap := make(map[string]bool)
			for _, signature := range blockCreatedEvent.Block.Signatures {
				signedValidator, exist := validatorMap[signature.ValidatorAddress]
				if !exist {
					return errors.New("error looking for signing validator details: not found")
				}
				commitmentRows = append(commitmentRows, view.ValidatorBlockCommitmentRow{
					ConsensusNodeAddress: signedValidator.ConsensusNodeAddress,
					BlockHeight:          height,
					IsProposer:           blockProposer == signature.ValidatorAddress,
					Signature:            signature.Signature,
					Timestamp:            signature.Timestamp,
				})

				identities = append(identities, fmt.Sprintf("-:%s", signedValidator.ConsensusNodeAddress))
				heightValidatorIdentities = append(
					heightValidatorIdentities,
					fmt.Sprintf("%s:%s", strconv.FormatInt(height, 10), signedValidator.ConsensusNodeAddress),
				)

				commitmentMap[signedValidator.ConsensusNodeAddress] = true
			}
			if err := validatorBlockCommitmentsView.InsertAll(commitmentRows); err != nil {
				return fmt.Errorf("error incrementing validator block commitment rows: %v", err)
			}
			if err := validatorBlockCommitmentsTotalView.Set(
				fmt.Sprintf("%s:-", strconv.FormatInt(height, 10)), int64(signatureCount),
			); err != nil {
				return fmt.Errorf("error incrementing height validator block commitments total: %v", err)
			}
			if err := validatorBlockCommitmentsTotalView.IncrementAll(
				identities, 1,
			); err != nil {
				return fmt.Errorf("error incrementing validator validator block commitments total: %v", err)
			}
			if err := validatorBlockCommitmentsTotalView.IncrementAll(
				heightValidatorIdentities, 1,
			); err != nil {
				return fmt.Errorf("error incrementing height-valiadtor block commitments total: %v", err)
			}
			if err := validatorBlockCommitmentsTotalView.Increment(
				"-:-", int64(signatureCount),
			); err != nil {
				return fmt.Errorf("error incrementing overall validator block commitments total: %v", err)
			}

			// Update validator up time
			activeValidators, activeValidatorsQueryErr := validatorsView.ListAll(view.ValidatorsListFilter{
				MaybeStatuses: []constants.Status{constants.BONDED},
			}, view.ValidatorsListOrder{})
			if activeValidatorsQueryErr != nil {
				return fmt.Errorf("error querying active validators: %v", activeValidatorsQueryErr)
			}

			var mutActiveValidators []view.ValidatorRow
			for _, activeValidator := range activeValidators {
				mutActiveValidator := activeValidator
				if commitmentMap[mutActiveValidator.ConsensusNodeAddress] {
					mutActiveValidator.TotalSignedBlock += 1
				} else {
					// give 10 blocks buffer on validator first join
					if height-mutActiveValidator.JoinedAtBlockHeight < 10 {
						mutActiveValidator.TotalSignedBlock += 1
					}
				}
				mutActiveValidator.TotalActiveBlock += 1

				mutActiveValidator.ImpreciseUpTime = new(big.Float).Quo(
					new(big.Float).SetInt64(mutActiveValidator.TotalSignedBlock),
					new(big.Float).SetInt64(mutActiveValidator.TotalActiveBlock),
				)

				mutActiveValidators = append(mutActiveValidators, mutActiveValidator)
			}

			if activeValidatorUpdateErr := validatorsView.UpdateAllValidatorUpTime(mutActiveValidators); activeValidatorUpdateErr != nil {
				return fmt.Errorf("error updating active validators up time data: %v", activeValidatorUpdateErr)
			}

		} else if votedEvent, ok := event.(*event_usecase.MsgVote); ok {
			projection.logger.Debug("handling MsgVote event")

			mutVotedValidator, votedValidatorQueryErr := validatorsView.FindBy(view.ValidatorIdentity{
				MaybeInitialDelegatorAddress: &votedEvent.Voter,
			})

			if votedValidatorQueryErr != nil {
				if errors.Is(votedValidatorQueryErr, rdb.ErrNoRows) {
					// the vote belongs to a non-validator account
					break
				}
				return fmt.Errorf("error querying voted validator: %v", votedValidatorQueryErr)
			}

			mutVotedValidator.VotedGovProposal = new(big.Int).Add(mutVotedValidator.VotedGovProposal, big.NewInt(1))
			if votedValidatorUpdateErr := validatorsView.Update(mutVotedValidator); votedValidatorUpdateErr != nil {
				return fmt.Errorf("error updating voted validator: %v", votedValidatorUpdateErr)
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

func (projection *Validator) projectValidatorView(
	validatorsView view.Validators,
	validatorActivities *view.ValidatorActivities,
	blockTime utctime.UTCTime,
	blockHeight int64,
	events []event_entity.Event,
) error {
	// MsgCreateValidator should be handled first
	for _, event := range events {
		if createGenesisValidator, ok := event.(*event_usecase.CreateGenesisValidator); ok {
			projection.logger.Debug("handling CreateGenesisValidator event")
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

			power := "0"
			if createGenesisValidator.Status == constants.BONDED {
				power = createGenesisValidator.Amount.Amount.QuoRaw(1000000).String()
			}

			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:    consensusNodeAddress,
				OperatorAddress:         createGenesisValidator.ValidatorAddress,
				InitialDelegatorAddress: createGenesisValidator.DelegatorAddress,
				TendermintAddress:       tendermintAddress,
				TendermintPubkey:        createGenesisValidator.TendermintPubkey,
				MinSelfDelegation:       createGenesisValidator.MinSelfDelegation,
				Status:                  createGenesisValidator.Status,
				Jailed:                  createGenesisValidator.Jailed,
				JoinedAtBlockHeight:     blockHeight,
				// TODO:  https://github.com/cosmos/cosmos-sdk/pull/8505
				Power:                   power,
				Moniker:                 createGenesisValidator.Description.Moniker,
				Identity:                createGenesisValidator.Description.Identity,
				Website:                 createGenesisValidator.Description.Website,
				SecurityContact:         createGenesisValidator.Description.SecurityContact,
				Details:                 createGenesisValidator.Description.Details,
				CommissionRate:          createGenesisValidator.CommissionRates.Rate,
				CommissionMaxRate:       createGenesisValidator.CommissionRates.MaxRate,
				CommissionMaxChangeRate: createGenesisValidator.CommissionRates.MaxChangeRate,
				TotalSignedBlock:        0,
				TotalActiveBlock:        0,
				ImpreciseUpTime:         big.NewFloat(1),
				VotedGovProposal:        big.NewInt(0),
			}

			// Validator re-joins
			isJoined, joinedAtBlockHeight, err := validatorsView.LastJoinedBlockHeight(
				validatorRow.OperatorAddress, validatorRow.ConsensusNodeAddress,
			)
			if err != nil {
				return fmt.Errorf("error querying validator last joined block height: %v", err)
			}
			if isJoined {
				validatorRow.JoinedAtBlockHeight = joinedAtBlockHeight
			}

			if err := validatorsView.Upsert(&validatorRow); err != nil {
				return fmt.Errorf("error inserting new validator into view: %v", err)
			}
		} else if msgCreateValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			projection.logger.Debug("handling MsgCreateValidator event")
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
			status := constants.UNBONDED
			if blockHeight == 0 {
				// genesis validators are always bonded
				status = constants.BONDED
			}

			validatorRow := view.ValidatorRow{
				ConsensusNodeAddress:    consensusNodeAddress,
				OperatorAddress:         msgCreateValidatorEvent.ValidatorAddress,
				InitialDelegatorAddress: msgCreateValidatorEvent.DelegatorAddress,
				TendermintAddress:       tendermintAddress,
				TendermintPubkey:        msgCreateValidatorEvent.TendermintPubkey,
				MinSelfDelegation:       msgCreateValidatorEvent.MinSelfDelegation,
				Status:                  status,
				Jailed:                  false,
				JoinedAtBlockHeight:     blockHeight,
				// TODO:  https://github.com/cosmos/cosmos-sdk/pull/8505
				Power:                   msgCreateValidatorEvent.Amount.Amount.QuoRaw(1000000).String(),
				Moniker:                 msgCreateValidatorEvent.Description.Moniker,
				Identity:                msgCreateValidatorEvent.Description.Identity,
				Website:                 msgCreateValidatorEvent.Description.Website,
				SecurityContact:         msgCreateValidatorEvent.Description.SecurityContact,
				Details:                 msgCreateValidatorEvent.Description.Details,
				CommissionRate:          msgCreateValidatorEvent.CommissionRates.Rate,
				CommissionMaxRate:       msgCreateValidatorEvent.CommissionRates.MaxRate,
				CommissionMaxChangeRate: msgCreateValidatorEvent.CommissionRates.MaxChangeRate,
				TotalSignedBlock:        0,
				TotalActiveBlock:        0,
				ImpreciseUpTime:         big.NewFloat(1),
				VotedGovProposal:        big.NewInt(0),
			}

			// Validator re-joins
			isJoined, joinedAtBlockHeight, err := validatorsView.LastJoinedBlockHeight(
				validatorRow.OperatorAddress, validatorRow.ConsensusNodeAddress,
			)
			if err != nil {
				return fmt.Errorf("error querying validator last joined block height: %v", err)
			}
			if isJoined {
				validatorRow.JoinedAtBlockHeight = joinedAtBlockHeight
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
					"error getting existing validator %s from view", msgEditValidatorEvent.ValidatorAddress,
				)
			}

			// get checking interval from config
			duration, durationParserErr := time.ParseDuration(projection.config.AttentionStatusRules.MaxEditQuota.Interval)
			if durationParserErr != nil {
				return fmt.Errorf("error parsing config interval %v", durationParserErr)
			}

			afterBlockTime := utctime.FromUnixNano(blockTime.UnixNano() - duration.Nanoseconds())

			mutValidatorActivities, _, err := (*validatorActivities).List(
				view.ValidatorActivitiesListFilter{
					AfterBlockTime:       &afterBlockTime,
					MaybeOperatorAddress: &msgEditValidatorEvent.ValidatorAddress,
				},
				view.ValidatorActivitiesListOrder{},
				&pagination.Pagination{},
			)
			if err != nil {
				return fmt.Errorf(
					"error getting existing validator activities %s from view", msgEditValidatorEvent.ValidatorAddress,
				)
			}

			if msgEditValidatorEvent.Description.Moniker != DO_NOT_MODIFY {
				errAttentionStatus := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "moniker")
				if errAttentionStatus != nil {
					return fmt.Errorf(
						"error checking attention status validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}
				mutValidatorRow.Moniker = msgEditValidatorEvent.Description.Moniker
			}
			if msgEditValidatorEvent.Description.Identity != DO_NOT_MODIFY {
				errAttentionOnNumOfChanges := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "identity")
				if errAttentionOnNumOfChanges != nil {
					return fmt.Errorf(
						"error checking attention status on number of changes validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}
				mutValidatorRow.Identity = msgEditValidatorEvent.Description.Identity
			}
			if msgEditValidatorEvent.Description.Details != DO_NOT_MODIFY {
				errAttentionOnNumOfChanges := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "details")
				if errAttentionOnNumOfChanges != nil {
					return fmt.Errorf(
						"error checking attention status on number of changes validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}
				mutValidatorRow.Details = msgEditValidatorEvent.Description.Details
			}
			if msgEditValidatorEvent.Description.SecurityContact != DO_NOT_MODIFY {
				errAttentionOnNumOfChanges := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "securityContact")
				if errAttentionOnNumOfChanges != nil {
					return fmt.Errorf(
						"error checking attention status on number of changes validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}
				mutValidatorRow.SecurityContact = msgEditValidatorEvent.Description.SecurityContact
			}
			if msgEditValidatorEvent.Description.Website != DO_NOT_MODIFY {
				errAttentionOnNumOfChanges := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "website")
				if errAttentionOnNumOfChanges != nil {
					return fmt.Errorf(
						"error checking attention status on number of changes validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}
				mutValidatorRow.Website = msgEditValidatorEvent.Description.Website
			}

			if msgEditValidatorEvent.MaybeCommissionRate != nil {
				errAttentionOnNumOfChanges := projection.checkAttentionOnNumOfChanges(mutValidatorRow, mutValidatorActivities, &msgEditValidatorEvent.ValidatorAddress, "commissionRate")
				if errAttentionOnNumOfChanges != nil {
					return fmt.Errorf(
						"error checking attention status on number of changes validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}

				errAttentionOnCommission := projection.checkAttentionOnCommission(mutValidatorRow, *msgEditValidatorEvent.MaybeCommissionRate, mutValidatorRow.CommissionRate, msgEditValidatorEvent.ValidatorAddress)
				if errAttentionOnCommission != nil {
					return fmt.Errorf(
						"error checking attention status on commission rate validator %s from view", msgEditValidatorEvent.ValidatorAddress,
					)
				}

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

			// Unjailed validator will become inactive first, if there's voting power changes then it becomes bonded
			mutValidatorRow.Status = constants.INACTIVE
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
				mutValidatorRow.Status = constants.INACTIVE
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

func (projection *Validator) checkAttentionOnCommission(mutValidatorRow *view.ValidatorRow, maybeCommissionRate string, commissionRate string, validatorAddress string) error {
	if projection.config.AttentionStatusRules.MaxCommissionRateChange.Enable {
		newCommission, newCommissionErr := strconv.ParseFloat(maybeCommissionRate, 64)
		if newCommissionErr != nil {
			return fmt.Errorf(
				"error converting new commission rate to float validator %s from view", validatorAddress,
			)
		}
		currentCommission, currentCommissionErr := strconv.ParseFloat(commissionRate, 64)
		if currentCommissionErr != nil {
			return fmt.Errorf(
				"error converting current commission rate to float validator %s from view", validatorAddress,
			)
		}

		if newCommission-currentCommission > projection.config.AttentionStatusRules.MaxCommissionRateChange.MaxChange {
			mutValidatorRow.Status = constants.ATTENTION
		}
	}

	return nil
}

func (projection *Validator) checkAttentionOnNumOfChanges(mutValidatorRow *view.ValidatorRow, mutValidatorActivities []view.ValidatorActivityRow, validatorAddress *string, editField string) error {
	if projection.config.AttentionStatusRules.MaxEditQuota.Enable {
		editQuotaCounter := make(map[string]int)
		for key, quota := range projection.config.AttentionStatusRules.MaxEditQuota.Quota {
			editQuotaCounter[key] = quota
		}

		// skip validator with "Attention" status
		if mutValidatorRow.Status == constants.ATTENTION {
			return nil
		}

		// count the current change
		if _, exists := editQuotaCounter[editField]; exists {
			editQuotaCounter[editField]--
		}

		// count previous changes
		for _, activity := range mutValidatorActivities {
			content := activity.Data.Content.(map[string]interface{})
			pastDescription := content["description"].(map[string]interface{})

			if pastDescription["moniker"] != DO_NOT_MODIFY {
				checkAndUpdateQuota("moniker", &editQuotaCounter)
			}
			if pastDescription["identity"] != DO_NOT_MODIFY {
				checkAndUpdateQuota("identity", &editQuotaCounter)
			}
			if pastDescription["details"] != DO_NOT_MODIFY {
				checkAndUpdateQuota("details", &editQuotaCounter)
			}
			if pastDescription["securityContact"] != DO_NOT_MODIFY {
				checkAndUpdateQuota("securityContact", &editQuotaCounter)
			}
			if pastDescription["website"] != DO_NOT_MODIFY {
				checkAndUpdateQuota("website", &editQuotaCounter)
			}
			if content["commissionRate"] != nil {
				checkAndUpdateQuota("commissionRate", &editQuotaCounter)
			}
		}
		// check counter by each field
		for _, count := range editQuotaCounter {
			if count < 0 {
				mutValidatorRow.Status = constants.ATTENTION
			}
		}

	}

	return nil
}

func checkAndUpdateQuota(key string, counter *map[string]int) {
	if _, exists := (*counter)[key]; exists {
		(*counter)[key]--
	}
}
