package validator_delegation

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/types"
	"github.com/crypto-com/chain-indexing/projection/validator_delegation/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &ValidatorDelegation{}

type Config struct {
	AccountAddressPrefix   string
	ValidatorAddressPrefix string
	ConNodeAddressPrefix   string
	// set in genesis `unbonding_time`
	UnbondingTime time.Duration
	// set in genesis `slash_fraction_double_sign`
	SlashFractionDoubleSign coin.Dec
	// set in genesis `slash_fraction_downtime`
	SlashFractionDowntime coin.Dec
	// PowerReduction - is the amount of staking tokens required for 1 unit of consensus-engine power.
	// Currently, this returns a global variable that the app developer can tweak.
	// https://github.com/cosmos/cosmos-sdk/blob/0cb7fd081e05317ed7a2f13b0e142349a163fe4d/x/staking/keeper/params.go#L46
	DefaultPowerReduction coin.Int
}

type CustomConfig struct {
	UnbondingTime           time.Duration `mapstructure:"unbonding_time"`
	SlashFractionDoubleSign string        `mapstructure:"slash_fraction_double_sign"`
	SlashFractionDowntime   string        `mapstructure:"slash_fraction_downtime"`
	DefaultPowerReduction   string        `mapstructure:"default_power_reduction"`
}

func CustomConfigFromInterface(data interface{}) (CustomConfig, error) {
	customConfig := CustomConfig{}

	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
		),
		Result: &customConfig,
	}
	decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
	if decoderErr != nil {
		return customConfig, fmt.Errorf("error creating projection config decoder: %v", decoderErr)
	}

	if err := decoder.Decode(data); err != nil {
		return customConfig, fmt.Errorf("error decoding projection ValidatorDelegation config: %v", err)
	}

	return customConfig, nil
}

func PrepareConfig(
	customConfig CustomConfig,
	accountAddressPrefix string,
	validatorAddressPrefix string,
	conNodeAddressPrefix string,
) (Config, error) {

	config := Config{}
	config.AccountAddressPrefix = accountAddressPrefix
	config.ValidatorAddressPrefix = validatorAddressPrefix
	config.ConNodeAddressPrefix = conNodeAddressPrefix
	config.UnbondingTime = customConfig.UnbondingTime

	var err error
	var ok bool

	config.SlashFractionDoubleSign, err = coin.NewDecFromStr(customConfig.SlashFractionDoubleSign)
	if err != nil {
		return config, fmt.Errorf("error parsing slashFractionDoubleSign from RawConfig: %v", err)
	}
	config.SlashFractionDowntime, err = coin.NewDecFromStr(customConfig.SlashFractionDowntime)
	if err != nil {
		return config, fmt.Errorf("error parsing slashFractionDowntime from RawConfig: %v", err)
	}
	config.DefaultPowerReduction, ok = coin.NewIntFromString(customConfig.DefaultPowerReduction)
	if !ok {
		return config, errors.New("error parsing defaultPowerReduction from RawConfig")
	}

	return config, nil
}

var (
	NewValidators                = view.NewValidatorsView
	NewUnbondingValidators       = view.NewUnbondingValidatorsView
	NewDelegations               = view.NewDelegationsView
	NewUnbondingDelegations      = view.NewUnbondingDelegationsView
	NewUnbondingDelegationQueue  = view.NewUnbondingDelegationQueueView
	NewRedelegations             = view.NewRedelegationsView
	NewRedelegationQueue         = view.NewRedelegationQueueView
	NewEvidences                 = view.NewEvidencesView
	UpdateLastHandledEventHeight = (*ValidatorDelegation).UpdateLastHandledEventHeight

	// DEBUG
	NewOperationPerformanceLogs = view.NewOperationPerformanceLogsView
)

type ValidatorDelegation struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	config Config

	migrationHelper migrationhelper.MigrationHelper
}

func NewValidatorDelegation(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config Config,
	migrationHelper migrationhelper.MigrationHelper,
) *ValidatorDelegation {
	return &ValidatorDelegation{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"ValidatorDelegation",
		),

		rdbConn,
		logger,

		config,

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
		event_usecase.EVIDENCE_SUBMITTED,
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

	// Handle events in Genesis block
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.CreateGenesisValidator); ok {

			if err := projection.handleGenesisCreateNewValidator(
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

	// DEBUG
	start := time.Now()

	// Handle evidence.
	//
	// NOTES: Although Evidences are handled by CosmosSDK in BeginBlock,
	//        here, we use a separate loop to handle Evidences.
	//        As later, `slash` event will need the information from Evidences.
	//        Therefore Evidence must be written to DB before we handle `slash` event in BeginBlock.
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.EvidenceSubmitted); ok {

			if err := projection.handleEvidence(
				rdbTxHandle,
				height,
				typedEvent.TendermintAddress,
				typedEvent.InfractionHeight,
				typedEvent.RawEvidence,
			); err != nil {
				return fmt.Errorf("error handling EvidenceSubmitted event: %v", err)
			}

		}
	}
	// DEBUG
	end := time.Now()
	performanceView := NewOperationPerformanceLogs(rdbTxHandle)
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "Evidence",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

	// Handle events in BeginBlock.
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.ValidatorJailed); ok {

			if err := projection.handleValidatorJailed(
				rdbTxHandle,
				height,
				typedEvent.ConsensusNodeAddress,
			); err != nil {
				return fmt.Errorf("error handling ValidatorJailed event: %v", err)
			}

		} else if typedEvent, ok := event.(*event_usecase.ValidatorSlashed); ok {

			// Important NOTES:
			//
			// There are two cases when a slash will happen:
			// - Missing Validator Signature
			// - Double Signing
			//
			// One important information for slashing is `distributionHeight`.
			// This field is used to retrieve the affected UnbondingDelegations and Redelegations.
			// Unluckily, it is not in a `slash` event.
			//
			// In the case of `Missing Validator Signature`, we can always calculate it by ourselves:
			// - `distributionHeight = currentBlockHeight - 2`.
			// https://github.com/cosmos/cosmos-sdk/blob/b5477dfee9e0785fe651fc603ca53f72a34d9bfd/x/slashing/keeper/infractions.go#L85
			//
			// In the case of `Double Signing`, we can retrieve the `infractionHeight` through Evidence.
			// - `distributionHeight = evidence.InfractionHeight - 1`
			// https://github.com/cosmos/cosmos-sdk/blob/b5477dfee9e0785fe651fc603ca53f72a34d9bfd/x/evidence/keeper/infraction.go#L101

			if typedEvent.Reason == string(types.MISSING_SIGNATURE) {

				distributionHeight := height - 2

				power, err := strconv.ParseInt(typedEvent.SlashedPower, 10, 64)
				if err != nil {
					return fmt.Errorf("error in parsing event.SlashedPower to int64: %v", err)
				}

				if err := projection.handleSlash(
					rdbTxHandle,
					height,
					blockTime,
					typedEvent.ConsensusNodeAddress,
					distributionHeight,
					power,
					projection.config.SlashFractionDowntime,
				); err != nil {
					return fmt.Errorf("error in handling slash event with missing_signature: %v", err)
				}

			} else if typedEvent.Reason == string(types.DOUBLE_SIGN) {

				evidencesView := NewEvidences(rdbTxHandle)
				validatorsView := NewValidators(rdbTxHandle)

				// Get the validator, to retrieve the TendermintAddress
				validator, found, err := validatorsView.FindByConsensusNodeAddr(typedEvent.ConsensusNodeAddress, height)
				if err != nil {
					return fmt.Errorf("error finding validator by consensusNodeAddr: %v", err)
				}
				if !found {
					return fmt.Errorf("error validator not found, conNodeAddr: %v", typedEvent.ConsensusNodeAddress)
				}

				// Get evidence related to this slash, to retrieve the infractionHeight
				evidence, err := evidencesView.FindBy(height, validator.TendermintAddress)
				if err != nil {
					return fmt.Errorf("error in finding the evidence: %v", err)
				}

				distributionHeight := evidence.InfractionHeight - 1

				power, err := strconv.ParseInt(typedEvent.SlashedPower, 10, 64)
				if err != nil {
					return fmt.Errorf("error in parsing event.SlashedPower to int64: %v", err)
				}

				if err := projection.handleSlash(
					rdbTxHandle,
					height,
					blockTime,
					typedEvent.ConsensusNodeAddress,
					distributionHeight,
					power,
					projection.config.SlashFractionDoubleSign,
				); err != nil {
					return fmt.Errorf("error in handling slash event with double_sign: %v", err)
				}

			}

		}
	}
	// DEBUG
	end = time.Now()
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "BeginBlock",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

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
			// DEBUG
			end = time.Now()
			if err := performanceView.Insert(
				view.OperationPerformanceRow{
					Height:   height,
					Action:   "MsgCreateValidator",
					Duration: end.Sub(start),
				},
			); err != nil {
				return fmt.Errorf("error in inserting performance log: %v", err)
			}
			start = end

		} else if typedEvent, ok := event.(*event_usecase.MsgEditValidator); ok {

			if typedEvent.MaybeMinSelfDelegation != nil {

				if err := projection.handleEditValidator(
					rdbTxHandle,
					height,
					typedEvent.ValidatorAddress,
					*typedEvent.MaybeMinSelfDelegation,
				); err != nil {
					return fmt.Errorf("error handling MsgEditValidator: %v", err)
				}
				// DEBUG
				end = time.Now()
				if err := performanceView.Insert(
					view.OperationPerformanceRow{
						Height:   height,
						Action:   "MsgEditValidator",
						Duration: end.Sub(start),
					},
				); err != nil {
					return fmt.Errorf("error in inserting performance log: %v", err)
				}
				start = end

			}

		} else if typedEvent, ok := event.(*event_usecase.MsgDelegate); ok {

			if _, err := projection.handleDelegate(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddress,
				typedEvent.DelegatorAddress,
				typedEvent.Amount.Amount,
			); err != nil {
				return fmt.Errorf("error handling MsgDelegate: %v", err)
			}
			// DEBUG
			end = time.Now()
			if err := performanceView.Insert(
				view.OperationPerformanceRow{
					Height:   height,
					Action:   "MsgDelegate",
					Duration: end.Sub(start),
				},
			); err != nil {
				return fmt.Errorf("error in inserting performance log: %v", err)
			}
			start = end

		} else if typedEvent, ok := event.(*event_usecase.MsgUndelegate); ok {

			// Successful tx with MsgUndelegate always has unbonding completion time in the Msg.
			if typedEvent.MaybeUnbondCompleteAt != nil {

				if err := projection.handleUndelegate(
					rdbTxHandle,
					height,
					typedEvent.ValidatorAddress,
					typedEvent.DelegatorAddress,
					typedEvent.Amount.Amount,
					*typedEvent.MaybeUnbondCompleteAt,
				); err != nil {
					return fmt.Errorf("error handling MsgUndelegate: %v", err)
				}
				// DEBUG
				end = time.Now()
				if err := performanceView.Insert(
					view.OperationPerformanceRow{
						Height:   height,
						Action:   "MsgUndelegate",
						Duration: end.Sub(start),
					},
				); err != nil {
					return fmt.Errorf("error in inserting performance log: %v", err)
				}
				start = end

			}

		} else if typedEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {

			if err := projection.handleRedelegate(
				rdbTxHandle,
				height,
				blockTime,
				typedEvent.DelegatorAddress,
				typedEvent.ValidatorSrcAddress,
				typedEvent.ValidatorDstAddress,
				typedEvent.Amount.Amount,
			); err != nil {
				return fmt.Errorf("error handling MsgBeginRedelegate: %v", err)
			}
			// DEBUG
			end = time.Now()
			if err := performanceView.Insert(
				view.OperationPerformanceRow{
					Height:   height,
					Action:   "MsgBeginRedelegate",
					Duration: end.Sub(start),
				},
			); err != nil {
				return fmt.Errorf("error in inserting performance log: %v", err)
			}
			start = end

		} else if typedEvent, ok := event.(*event_usecase.MsgUnjail); ok {

			if err := projection.handleValidatorUnjailed(
				rdbTxHandle,
				height,
				typedEvent.ValidatorAddr,
			); err != nil {
				return fmt.Errorf("error handling MsgUnjail: %v", err)
			}
			// DEBUG
			end = time.Now()
			if err := performanceView.Insert(
				view.OperationPerformanceRow{
					Height:   height,
					Action:   "MsgUnjail",
					Duration: end.Sub(start),
				},
			); err != nil {
				return fmt.Errorf("error in inserting performance log: %v", err)
			}
			start = end

		}
	}

	// Handle events in EndBlock.
	//
	// Handle ValidatorUpdate events. These events are emitted in each EndBlock.
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.PowerChanged); ok {

			if err := projection.handlePowerChanged(
				rdbTxHandle,
				height,
				blockTime,
				typedEvent.TendermintPubkey,
				typedEvent.Power,
			); err != nil {
				return fmt.Errorf("error handling PowerChanged event: %v", err)
			}

		}
	}
	// DEBUG
	end = time.Now()
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "EndBlock",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

	if err := projection.handleMatureUnbondingValidators(
		rdbTxHandle,
		height,
		blockTime,
	); err != nil {
		return fmt.Errorf("error handling mature unbonding validators: %v", err)
	}
	// DEBUG
	end = time.Now()
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "MatureUnbondingValidators",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

	if err := projection.handleMatureUnbondingDelegationQueueEntries(
		rdbTxHandle,
		height,
		blockTime,
	); err != nil {
		return fmt.Errorf("error handling mature unbonding delegation queue entries: %v", err)
	}
	// DEBUG
	end = time.Now()
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "MatureUnbondingDelegations",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

	if err := projection.handleMatureRedelegationQueueEntries(
		rdbTxHandle,
		height,
		blockTime,
	); err != nil {
		return fmt.Errorf("error handling mature redelegation queue entries(): %v", err)
	}
	// DEBUG
	end = time.Now()
	if err := performanceView.Insert(
		view.OperationPerformanceRow{
			Height:   height,
			Action:   "MatureRedelegations",
			Duration: end.Sub(start),
		},
	); err != nil {
		return fmt.Errorf("error in inserting performance log: %v", err)
	}
	start = end

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}
