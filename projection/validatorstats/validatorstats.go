package validatorstats

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	entity_projection "github.com/crypto-com/chain-indexing/entity/projection"
	"github.com/crypto-com/chain-indexing/external/json"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/projection/validatorstats/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/golang-migrate/migrate/v4"
)

var _ entity_projection.Projection = &ValidatorStats{}

const TOTAL_REWARD = "total_reward"
const TOTAL_DELEGATE = "total_delegate"

type ValidatorStats struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewValidatorStats(logger applogger.Logger, rdbConn rdb.Conn) *ValidatorStats {
	return &ValidatorStats{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "ValidatorStats"),

		rdbConn,
		logger,
	}
}

func (_ *ValidatorStats) GetEventsToListen() []string {
	return []string{
		event_usecase.GENESIS_VALIDATOR_CREATED,
		event_usecase.MSG_CREATE_VALIDATOR_CREATED,
		event_usecase.BLOCK_PROPOSER_REWARDED,
		event_usecase.BLOCK_REWARDED,
		event_usecase.MSG_DELEGATE_CREATED,
		event_usecase.MSG_UNDELEGATE_CREATED,
	}
}

const (
	MIGRATION_TABLE_NAME = "validator_stats_schema_migrations"
	MIGRATION_GITHUB_TARGET = "github://public:token@crypto-com/chain-indexing/projection/validatorstats/migrations#migration-sharing"
)

func (projection *ValidatorStats) migrationDBConnString() string {
	conn := projection.rdbConn.(*pg.PgxConn)
	connString := conn.ConnString()
	if connString[len(connString)-1:] == "?" {
		return connString + "x-migrations-table=" + MIGRATION_TABLE_NAME
	} else {
		return connString + "&x-migrations-table=" + MIGRATION_TABLE_NAME
	}
}

func (projection *ValidatorStats) OnInit() error {
	m, err := migrate.New(
		MIGRATION_GITHUB_TARGET,
		projection.migrationDBConnString(),
	)
	if err != nil {
		projection.logger.Errorf("failed to init migration: %v", err)
		return err
	}

	if err := m.Up(); err != nil {
		projection.logger.Errorf("failed to run migration: %v", err)
		return err
	}

	return nil
}

func (projection *ValidatorStats) HandleEvents(height int64, events []event_entity.Event) error {
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
	validatorStatsView := view.NewValidatorStats(rdbTxHandle)

	rawTotalReward, err := validatorStatsView.FindBy(TOTAL_REWARD)
	if err != nil {
		return fmt.Errorf("error getting total reward metrics: %v", err)
	}
	var totalReward coin.DecCoins
	if rawTotalReward == "" {
		totalReward = coin.NewEmptyDecCoins()
	} else {
		json.MustUnmarshalFromString(rawTotalReward, &totalReward)
	}

	rawTotalDelegate, err := validatorStatsView.FindBy(TOTAL_DELEGATE)
	if err != nil {
		return fmt.Errorf("error getting total reward metrics: %v", err)
	}
	//totalDelegate := coin.MustParseCoinsNormalized(rawTotalDelegate)
	var totalDelegate coin.Coins
	if rawTotalDelegate == "" {
		totalDelegate = coin.NewEmptyCoins()
	} else {
		json.MustUnmarshalFromString(rawTotalDelegate, &totalDelegate)
	}

	for _, event := range events {
		if createValidatorEvent, ok := event.(*event_usecase.CreateGenesisValidator); ok {
			totalDelegate = totalDelegate.Add(createValidatorEvent.Amount)
		} else if createValidatorEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			totalDelegate = totalDelegate.Add(createValidatorEvent.Amount)
		} else if blockProposerRewardedEvent, ok := event.(*event_usecase.BlockProposerRewarded); ok {
			totalReward = totalReward.Add(blockProposerRewardedEvent.Amount...)
			if err != nil {
				return fmt.Errorf("error adding rewards: %v", err)
			}
		} else if blockRewardedEvent, ok := event.(*event_usecase.BlockRewarded); ok {
			totalReward = totalReward.Add(blockRewardedEvent.Amount...)
			if err != nil {
				return fmt.Errorf("error adding rewards: %v", err)
			}
		} else if delegateEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			totalDelegate = totalDelegate.Add(delegateEvent.Amount)
			if err != nil {
				return fmt.Errorf("error adding delegate: %v", err)
			}
		} else if undelegateEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			totalDelegate = totalDelegate.Add(undelegateEvent.Amount.Neg())
			if err != nil {
				return fmt.Errorf("error subtracting delegate: %v", err)
			}
		}
		// TODO: Handle slashed delegation
	}

	if err = validatorStatsView.Set(TOTAL_REWARD, json.MustMarshalToString(totalReward)); err != nil {
		return fmt.Errorf("error updating total reward")
	}
	if err = validatorStatsView.Set(TOTAL_DELEGATE, json.MustMarshalToString(totalDelegate)); err != nil {
		return fmt.Errorf("error updating total delegate")
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func TrimDecimalPlaces(s string) string {
	parts := strings.Split(s, ".")
	return parts[0]
}
